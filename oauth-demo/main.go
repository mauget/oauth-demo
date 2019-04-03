package main

import (
	"fmt"
	"log"
	"net/http"
	"oauth-demo/server"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

func main() {

	var port string
	port, isPresent := os.LookupEnv("PORT")
	if !isPresent {
		port = "8000"
	}

	router := mux.NewRouter()

	// Route to production build of React UI:
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./build")))

	var gracefulStop = make(chan os.Signal)

	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v", sig)
		fmt.Println("Wait for 2 seconds to finish processing")

		server.Close()

		os.Exit(0)
	}()

	log.Println("Starting and listening on Port ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
