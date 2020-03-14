package utils

import (
	"log"
	"os"

	"github.com/gorilla/sessions"

	"gopkg.in/mgo.v2"
)

var tsession *mgo.Session
var db string

// SessionStore - Cookie Session
//var SessionStore *sessions.CookieStore

var SessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func init() {

	skey, found := os.LookupEnv("SESSION_KEY")
	if found {

		SessionStore = sessions.NewCookieStore([]byte(skey))

	} else {

		SessionStore = sessions.NewCookieStore([]byte("localhost"))

	}

	dburi, isPresent := os.LookupEnv("MONGODB_URI")
	if !isPresent {
		dburi = "mongodb://localhost:27017"

	}
	db, isPresent = os.LookupEnv("MONGO_DB")
	if !isPresent {
		db = "flyworld"
	}

	s, err := mgo.Dial(dburi)

	log.Println("Connected to :", dburi)
	if err != nil {
		panic(err)
	}
	tsession = s

}

// Dbsession -  Mongodb session
func Dbsession() *mgo.Session {

	return tsession
}

// Close -  Mongodb session
func Close() {

	tsession.Close()
	log.Println("DB Session closed")
}

// CreateIndex - create search index
func CreateIndex() {

	c := tsession.DB(db).C("categories")

	index := mgo.Index{
		Key:  []string{"$text:description", "$text:type"},
		Name: "search",
	}
	_ = c.EnsureIndex(index)

	log.Println("Category each Index Created")

	c = tsession.DB(db).C("videos")

	index = mgo.Index{
		Key:  []string{"$text:description", "$text:title"},
		Name: "search",
	}
	_ = c.EnsureIndex(index)

	log.Println("Video  search Index Created")

}

/*func ForceSsl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("GO_ENV") == "production" {
			if r.Header.Get("x-forwarded-proto") != "https" {
				sslUrl := "https://" + r.Host + r.RequestURI
				http.Redirect(w, r, sslUrl, http.StatusTemporaryRedirect)
				return
			}
		}

		next.ServeHTTP(w, r)
	})

}*/
