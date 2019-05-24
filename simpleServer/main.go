package main

import (
	"log"
	"net/http"
)

const certs = "."

// Ref: https://github.com/denji/golang-tls
func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	_, _ = w.Write([]byte("<h1>This is an example server.</h1>"))
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServeTLS(":443", certs+"/server.crt", certs+"/server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
