package main

import (
	"fly-world/utils"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
	"oauth-demo/domain/randstr"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	port, isPresent := os.LookupEnv("PORT")
	if !isPresent {
		port = "8000"
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/", handleMain).Methods("GET")
	router.HandleFunc("/api/GoogleLogin", handleGoogleLogin).Methods("GET")
	router.HandleFunc("/api/GoogleCallback", handleGoogleCallback).Methods("GET")

	// Serve the UI from the build directory. Assumes we carried out a yarn build.
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./build")))

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v", sig)
		fmt.Println("Wait for 2 second to finish processing")
		utils.Close()
		os.Exit(0)
	}()

	log.Println("Starting  and listening on Port ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}


var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8000/api/GoogleCallback",
		ClientID:     os.Getenv("googlekey"),
		ClientSecret: os.Getenv("googlesecret"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}
	// A random string, generated for each user
	oauthStateString = randstr.RandStringBytesMaskImprSrcUnsafe(16)

	// SPA redirect target
	clientRedirectTail = "/#/oauth/"
)


func handleMain(w http.ResponseWriter, _ *http.Request) {
	htmlIndex := "/"
	_, _ = fmt.Fprintf(w, htmlIndex)
}


func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}


func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")

	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', received '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)

	if err != nil {
		fmt.Println("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}


	// TODO save the token per-user for comparision to token sent with subsequent requests.

	// Pass the token to the APP via a 303 redirect
	session["test"] = token.AccessToken;
	fmt.Printf("token '%s'\n", token.AccessToken)
	http.Redirect(w, r, clientRedirectTail + token.AccessToken, http.StatusSeeOther)

	/*response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	if err != nil {
		_, _ = fmt.Fprintf(w, "Error: %s\n", err)

	} else if response != nil {
		defer response.Body.Close()

		contents, _ := ioutil.ReadAll(response.Body)
		_, _ = fmt.Fprintf(w, "Content: %s\n", contents)

	} else {
		fmt.Println("Panic")
	}*/
}
