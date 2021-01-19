package main

import (
	"bufio"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	//using simple the builtin copy function, which copy the source to destination
	/* if _, err := copy(conn, conn); err != nil {
		log.Fatalln("Unable to read or write data")
	}
	*/

	//or using bufio Reader and Writer

	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("could not read data")
	}

	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("Unable to write data")
	}

	writer.Flush()
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
