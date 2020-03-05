package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	//create buffer to store received data
	b := make([]byte, 512)

	for {
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client Disconnected")
			break
		}
		if err != nil {
			log.Println("undexpected error")
			break
		}
		log.Printf("Received %d bytes: %s", size, string(b))

		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unabel to write data")
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatalln("Unable to bind the port")
	}
	log.Println("Listening on 0.0.0.0:2000")

	for {
		conn, err := listener.Accept()
		log.Println("Connection received")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		go echo(conn)
	}

}
