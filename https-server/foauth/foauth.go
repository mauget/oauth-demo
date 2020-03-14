package foauth

import (
	"context"
	"fly-world/domain/randstr"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type AccessToken struct {
	Token  string
	Expiry int64
}

type FUser struct {
	Id            string
	Email         string
	VerifiedEmail bool
	Name          string
	GivenName     string
	FamilyName    string
	Picture       string
	Locale        string
}

var endpoint = oauth2.Endpoint{
	AuthURL:   "https://www.facebook.com/v3.2/dialog/oauth",
	TokenURL:  "https://graph.facebook.com/v3.2/oauth/access_token",
	AuthStyle: oauth2.AuthStyleAutoDetect,
}

var (
	facebookOauthConf = &oauth2.Config{
		RedirectURL:  os.Getenv("FLYWORLD_FB_CALLBACK"),
		ClientID:     os.Getenv("FACEBOOK_ID"),
		ClientSecret: os.Getenv("FACEBOOK_SECRET"),
		Scopes:       []string{"public_profile", "email"},
		Endpoint:     endpoint,
	}

	// A random string, generated for each user
	oauthStateString = randstr.RandStringBytesMaskImprSrcUnsafe(16)


	// SPA redirect target
	//clientRedirectTail = "/#/oauth/facebook"
)

func HandleFacebookLogin(w http.ResponseWriter, r *http.Request) {

	loginUrl := facebookOauthConf.AuthCodeURL(oauthStateString)

	log.Println("FLYWORLD FB login URL: " + loginUrl)

	http.Redirect(w, r, loginUrl, http.StatusTemporaryRedirect)
}

func HandleFacebookCallback(w http.ResponseWriter, r *http.Request) {

	// Get a session. Get() always returns a session, even if empty.
	//store := utils.SessionStore
	//session, err := store.New(r, user.FlyworldSession)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	state := r.FormValue("state")

	if state != oauthStateString {
		log.Printf("FLYWORLD: invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	log.Printf("FB code: '%s'", code)

	token, err := facebookOauthConf.Exchange(context.TODO(), code)

	if err != nil {
		log.Printf("code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	log.Printf("FB token: '%v'", token)

	urlStr := fmt.Sprintf("%s?client_id=%s&client_secret=%s&code=%s&grant_type=%s&redirect_uri=%s",
		endpoint.TokenURL,
		facebookOauthConf.ClientID,
		facebookOauthConf.ClientSecret,
		code,
		"client_credentials",
		facebookOauthConf.RedirectURL)

	log.Printf("FB urlStr: '%s'", urlStr)

	response, err := http.Get(urlStr)

	if err != nil {

		emsg := fmt.Sprintf("Oauth2 error: %s\n", err)
		log.Println(emsg)
		_, _ = fmt.Fprintf(w, emsg)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

	} else if response != nil {
		//defer response.Body.Close()

		log.Printf("endpoint response %s\n", response)

		var content []byte
		content, err := ioutil.ReadAll(response.Body)

		if err == nil {

			log.Printf("#### %s\n", content, token.AccessToken)

			//var u FUser
			//u, err = extractUserFromAuthContent(content)
			//
			//if err == nil {
			//
			//	//var token = token.AccessToken
			//	userid := createUserID(u.Name, u.Id)
			//
			//	// temp - remove THIS FAKES out the issing user info
			//	//userid = "facebook@temp.org"
			//
			//	log.Println("### FB u.Name:" + u.Name + " u.Id:" + u.Id)
			//
			//	// Set the userid to in-session in user collection.
			//	// userid persists across sessions within user collection.
			//	// The role defaults to "user".
			//	// The userid and role persists across sessions.
			//	// An external administrator action could  alter the user doc to role "admin".
			//	_, err = user.UpsertOAuthUser(userid, token.AccessToken)
			//
			//	if err != nil {
			//		log.Printf("UpsertOAuther error %s", err)
			//		http.Error(w, err.Error(), http.StatusInternalServerError)
			//	}
			//
			//	u2, err := user.GetForUserID(userid)
			//	if err != nil {
			//		log.Printf("UpsertOAuther has no role %s", err)
			//		http.Error(w, err.Error(), http.StatusInternalServerError)
			//		return
			//	}
			//
			//	// Save the token and role in a session cookie
			//	session.Values["role"] = u2.Role
			//	session.Values["token"] = token.AccessToken
			//
			//	err = session.Save(r, w)
			//
			//	if err != nil {
			//		log.Printf("Session save error %s", err)
			//		http.Error(w, err.Error(), http.StatusInternalServerError)
			//		return
			//	}
			//
			//	// Pass the user ID to the SPA client via a 307 redirect
			//	userIdTail := "/" + userid
			//
			//	http.Redirect(w, r, clientRedirectTail+userIdTail, http.StatusTemporaryRedirect)*/

			} else {

				log.Println("Decode OAuth content failed: %s\n", err.Error())
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			}
		}

	//} else {
	//	log.Println("Panic")
	//}

}


//func createUserID(email string, _ string) string {
//	// Commented this out don't think we need the id field
//	//return email + ";" + id
//	return email
//}

/*func extractUserFromAuthContent(content []byte) (FUser, error) {

	log.Printf("Raw content: %s", string(content))

	var u FUser
	err := json.Unmarshal(content, &u)

	log.Printf("AuthContent: %v", u)

	log.Println("User: " + u.Name)
	log.Println("EMail: " + u.Email)
	log.Println("ID: " + u.Id)
	return u, err
}*/
