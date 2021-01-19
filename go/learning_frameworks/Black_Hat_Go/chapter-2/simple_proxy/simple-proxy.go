package main

import (
	"io"
	"log"
	"net"
)

func handle(conn net.Conn) {
	dst, err := net.Dial("tcp", "www.google.com:80")
	if err != nil {
		log.Fatalln("Unable to connect to the remote host")
	}
	defer dst.Close()

	//Copy source to destination, run within a goroutine as io.Copy is a blocking function
	go func() {
		if _, err := io.Copy(dst, conn); err != nil {
			log.Fatalln(err)
		}

	}()

	//Copy the destination back to our source
	if _, err := io.Copy(conn, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("could not listen on the port")
	}
	log.Println("listening on 0.0.0.0:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("could not accept the connection")
		}
		log.Printf("new connection from %s", conn.RemoteAddr())

		go handle(conn)
	}
}
