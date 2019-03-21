package main

import (
	"log"
	"net/http"

	"./handlers"
	"./imgproc"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./ui")))
	http.HandleFunc("/upload", handlers.UploadFile)
	http.HandleFunc("/gallery", handlers.ListFiles)
	log.Println("imgn server running")
	go imgproc.Extract()
	http.ListenAndServe(":8080", nil)
}
