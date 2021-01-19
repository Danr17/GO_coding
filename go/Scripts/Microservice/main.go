package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/Danr17/GO_coding/Scripts/Microservice/homepage"
	"github.com/Danr17/GO_coding/Scripts/Microservice/server"
)

var (
	GcukCertFile    = flag.String("GCUK_CERT_FILE", "certs/local.localhost.cert", "The TLS cert file")
	GcukKeyFile     = flag.String("GCUK_KEY_FILE", "certs/local.localhost.key", "The TLS key file")
	GcukServiceAddr = flag.String("GCUK_SERVICE_ADDR", ":8080", "The server port")
)

func main() {
	flag.Parse()
	logger := log.New(os.Stdout, "gcuk ", log.LstdFlags|log.Lshortfile)

	h := homepage.NewHandlers(logger)

	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	srv := server.New(mux, *GcukServiceAddr)

	logger.Println("server starting")
	err := srv.ListenAndServeTLS(*GcukCertFile, *GcukKeyFile)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
