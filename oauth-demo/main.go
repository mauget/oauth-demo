package main

import (
	"fly-world/utils"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var port string
	port, isPresent := os.LookupEnv("PORT")
	if !isPresent {
		port = "8000"
	}

	router := mux.NewRouter()

	// Serve the UI from the build directory. Assumes we carried out an npm build (or yarn build).
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
