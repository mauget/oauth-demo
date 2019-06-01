package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
	"unsafe"
)

var sessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

const goSecureSession = "GO_SECURE_SESSION"

type UserStore struct {
	UserID   string
	Password string
	Token    string
}

// This could be a MongoDB collections
var userMap = make(map[string]UserStore)

func main() {

	const tlsPort = "443"
	const certsLoc = "certs"

	userMap["Lou"] = UserStore{"Lou", "123456", ""}
	userMap["Dave"] = UserStore{"Dave", "123456", ""}
	userMap["Test"] = UserStore{"Test", "123456", ""}

	log.Println("Initailzed userMap")

	//------ router --->
	router := mux.NewRouter()
	router.HandleFunc("/api/testmsg", getTestMsg).Methods("GET")
	router.HandleFunc("/api/login", handleLogin).Methods("POST")
	router.HandleFunc("/api/logout", handleLogout).Methods("POST")
	router.HandleFunc("/api/session", getSession).Methods("GET")

	// Following enables serving files under https://localhost:443/<filename>, where "/" mapped to "./view" dir
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./view"))))
	//<---------------

	//----- Listen for bare http tlsPort 80 request, redirect to https
	go httpsRedirectToTLS(tlsPort)

	//---- Run the secured server
	log.Printf("Starting  and listening on Port %s\n", tlsPort)
	log.Fatal(http.ListenAndServeTLS(":"+tlsPort, certsLoc+"/server.crt", certsLoc+"/server.key", router))

}

// Redirect to something like https://localhost:443/
func httpsRedirectToTLS(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		httpsURL := fmt.Sprintf("https://%s:%s/", "localhost", port)

		log.Printf("Redirecting to %s (TLS)\n", httpsURL)
		http.Redirect(w, r, httpsURL, http.StatusSeeOther)
	})

	_ = http.ListenAndServe(":80", nil)

}

func getTestMsg(w http.ResponseWriter, _ *http.Request) {

	now := time.Now()

	t := now.Format("Mon Jan _2 15:04:05 MST 2006")

	err := json.NewEncoder(w).Encode(t)
	if nil != err {
		log.Fatalf("getTestMsg %s\n", err)
	}
}

func authenticate(userID string, password string) (UserStore, bool) {
	userEntry := userMap[userID]
	return userEntry, (userEntry.UserID == userID) && (userEntry.Password == password)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {

	// Start a new session
	store := sessionStore
	session, err := store.New(r, goSecureSession)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(r.Body)

	// User - user struct
	type User struct {
		UserID   string
		Password string
	}

	var t User
	if decoder.Decode(&t) != nil {
		panic(err)
	}
	log.Println(t.UserID)

	u, isMatch := authenticate(t.UserID, t.Password)

	if !isMatch {
		_ = json.NewEncoder(w).Encode("Error authenticating User")

	} else {

		// Set the existing or new user to in-session
		token := randstr(32)
		u.Token = token

		userMap[t.UserID] = u

		// Save the token and user ID  in a session cookie
		session.Values["token"] = u.Token
		session.Values["userID"] = u.UserID

		// Abandons any existing session for the requesting host
		err = session.Save(r, w)
		if err != nil {
			log.Printf("Session save error %s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		log.Println("userMap", userMap)
		_ = json.NewEncoder(w).Encode(u)

	}

}

// Clear session
func handleLogout(w http.ResponseWriter, r *http.Request) {

	// Get a session. Get() always returns a session, even if empty.
	store := sessionStore
	session, _ := store.Get(r, goSecureSession)

	// Delete the session
	session.Options.MaxAge = -1
	_ = session.Save(r, w)

	// SPA redirect target
	clientRedirectTail := "/"

	http.Redirect(w, r, clientRedirectTail, http.StatusSeeOther)
}

// Return truth of session token match for current user
func getSession(w http.ResponseWriter, r *http.Request) {

	// Get a session. Get() always returns a session, even if empty.
	store := sessionStore
	session, _ := store.Get(r, goSecureSession)

	// Get the token and user ID from the cookie
	val := session.Values["token"]
	token, okToken := val.(string)

	val = session.Values["userID"]
	userID, okUserID := val.(string)

	isSession := false

	if okToken && okUserID {
		userEntry := userMap[userID]
		isSession = userEntry.Token == token
	}
	_ = json.NewEncoder(w).Encode(isSession)
}

/**
 * Good-performing, low garbage, skinny-memory, random string generator.
 *
 * e.g.  randstr.RandStringBytesMaskImprSrcUnsafe(32)
 *
 * Ref: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
 */

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func randstr(n int) string {
	b := make([]byte, n)

	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
