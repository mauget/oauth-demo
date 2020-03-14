package main

import (
	"github.com/gorilla/mux"
	"https-server/foauth"
	"log"
	"net/http"
)

func HandlerOne(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("<p><a href=\"https://localhost//api/FacebookLogin\">Facebook login</a>" +
		"</p>\n<p><a href=\"https://locahost/api/FacebookCallback\">Facebook login redirect</a></p>\n"))
}

func HandlerTwo(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("<h1>Gorilla Two</h1>\n"))
}
func HandlerThree(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("<h1>Gorilla Three</h1>\n"))
}

func main() {
	router := mux.NewRouter()
	// Routes consist of a path and a handler function.
	router.HandleFunc("/", HandlerOne)
	router.HandleFunc("/one", HandlerOne)
	router.HandleFunc("/two", HandlerTwo)
	router.HandleFunc("/three", HandlerThree)

	//Facebook
	router.HandleFunc("/api/FacebookLogin", foauth.HandleFacebookLogin)
	//router.HandleFunc("/api/FacebookCallback", foauth.HandleFacebookCallback).Methods("GET")
	router.HandleFunc("/api/FacebookCallback", foauth.HandleFacebookCallback)
	//

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServeTLS(":443", "localhost.crt", "localhost.key", router))
}

