package main

import (
	"log"
	"net/http"
	"os"

	hello "github.com/Danr17/GO_coding/Scripts/Hello_Trace"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/hello", hello.Hello)
	log.Printf("Listening on localhost:%v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
