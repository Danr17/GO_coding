package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"syscall"

	"github.com/Danr17/GO_coding/ssh_login/pkg/parsefile"
	"github.com/Danr17/GO_coding/ssh_login/pkg/sshinit"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func connect(init *sshinit.SSHInit, addressBatch, results chan string) {
	for address := range addressBatch {
		fmt.Printf("Trying to connect to %s \n", address)
		client, err := ssh.Dial("tcp", address+":22", init.Config)
		if err != nil {
			//fmt.Printf("%#v\n", err)
			if _, ok := err.(*net.OpError); ok {
				results <- "Unreachable" + " " + address
				continue
			}
			results <- "AuthFail" + " " + address
			continue
		}
		// Each ClientConn can support multiple interactive sessions, represented by a Session.
		session, err := client.NewSession()
		if err != nil {
			log.Printf("Failed to create session: %v", err)
		}
		defer session.Close()

		results <- "Done"

		//log.Printf("Successfully connected to host %s", address)
		// Once a Session is created, you can execute a single command on the remote side using the Run method.
		/*
			var b bytes.Buffer
			session.Stdout = &b
			if err := session.Run("show users"); err != nil {
				log.Fatal("Failed to run: " + err.Error())
			}
			fmt.Println(b.String())
		*/
	}
}

func main() {
	username := flag.String("user", "", "provide an username")
	password := flag.String("pass", "", "provide a password")
	filename := flag.String("file", "test.txt", "provide a filename like test.txt")
	interactive := flag.Bool("interactive", false, "interactive mode")

	flag.Parse()

	if *username == "" || *filename == "" {
		log.Fatalln("Either username, password or filename were not provided")
	}

	if *password == "" {
		fmt.Print("Enter Password: \n")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatalf("Could not read the password: %s", err)
		}
		pass := strings.TrimSpace(string(bytePassword))
		*password = pass
	}

	file := fmt.Sprintf("%s.txt", *filename)

	addresses, err := parsefile.Parse(file)
	if err != nil {
		log.Fatalf("could parse the file %s: %v", file, err)
	}

	conninit := sshinit.NewSSHInit(*username, *password, *interactive)

	var notWorking []string
	addressBatch := make(chan string, 30)
	results := make(chan string)

	for i := 0; i < cap(addressBatch); i++ {
		go connect(conninit, addressBatch, results)
	}

	go func() {
		for _, address := range addresses {
			addressBatch <- strings.TrimSpace(address)
		}
	}()

	for i := 0; i < len(addresses); i++ {
		status := <-results
		if status == "Done" {
			continue
		}
		notWorking = append(notWorking, status)
	}

	close(addressBatch)
	close(results)

	f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, notWorkAdd := range notWorking {
		if _, err = f.WriteString(notWorkAdd + "\n"); err != nil {
			panic(err)
		}
	}
}
