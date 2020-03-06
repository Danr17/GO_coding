package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Danr17/GO_coding/ssh_login/pkg/sshinit"
	"golang.org/x/crypto/ssh"
)

func parse(filename string) ([]string, error) {
	addressList := []string{}
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open the file %s", filename)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		address, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return addressList, err
		}
		addressList = append(addressList, address)
	}
	return addressList, nil
}

func connect(init *sshinit.SSHInit, hostname string) {
	client, err := ssh.Dial("tcp", hostname+":22", init.Config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("show users"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}

func main() {
	username := flag.String("user", "", "provide an username")
	password := flag.String("pass", "", "provide a password")
	filename := flag.String("file", "test.txt", "provide a filename like test.txt")

	if *username == "" || *password == "" || *filename == "" {
		log.Fatalln("Either username, password or filename were not provided")
	}

	addresses, err := parse(*filename)
	if err != nil {
		log.Fatalf("could parse the file %s: %v", *filename, err)
	}

	conninit := sshinit.NewSSHInit(*username, *password)

	for _, address := range addresses {
		log.Printf("Connecting to: %s", address)
		connect(conninit, address)
		log.Printf("Done: %s", address)
	}
}
