package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Danr17/GO_coding/Scripts/Microservice/homepage"
	"github.com/Danr17/GO_coding/Scripts/Microservice/server"
)

var (
	GcukCertFile    = os.Getenv("GCUK_CERT_FILE")
	GcukKeyFile     = os.Getenv("GCUK_KEY_FILE")
	GcukServiceAddr = os.Getenv("GCUK_SERVICE_ADDR")
)

func main() {
	logger := log.New(os.Stdout, "gcuk ", log.LstdFlags|log.Lshortfile)

	h := homepage.NewHandlers(logger)

	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	srv := server.New(mux, GcukServiceAddr)

	logger.Println("server starting")
	err := srv.ListenAndServeTLS(GcukCertFile, GcukKeyFile)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
