package server

import (
	"gopkg.in/mgo.v2"
	"log"
)

var tsession *mgo.Session

func init() {

	/*	dburi, isPresent := os.LookupEnv("MONGODB_URI")
		if !isPresent {
			dburi = "mongodb://localhost:27017"

		}

		s, err := mgo.Dial(dburi)

		log.Println("Connected to :", dburi)
		if err != nil {
			panic(err)
		}
		tsession = s*/

}

// Dbsession -  Mongodb session
/*func Dbsession() *mgo.Session {

	return tsession

}*/

// Close -  Mongodb session
func Close() {

	if tsession != nil {
		tsession.Close()
	}
	log.Println("DB Session closed")
}
