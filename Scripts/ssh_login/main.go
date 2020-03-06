package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/Danr17/GO_coding/ssh_login/pkg/parsefile"
	"github.com/Danr17/GO_coding/ssh_login/pkg/sshinit"
	"golang.org/x/crypto/ssh"
)

func connect(init *sshinit.SSHInit, addressBatch, results chan string) {
	for address := range addressBatch {
		fmt.Printf("Trying to connect to %s \n", address)
		client, err := ssh.Dial("tcp", address+":22", init.Config)
		if err != nil {
			results <- address
			continue
		}
		// Each ClientConn can support multiple interactive sessions, represented by a Session.
		session, err := client.NewSession()
		if err != nil {
			log.Printf("Failed to create session: %v", err)
		}
		defer session.Close()

		fmt.Println("Done")
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

	flag.Parse()

	if *username == "" || *password == "" || *filename == "" {
		log.Fatalln("Either username, password or filename were not provided")
	}
	file := fmt.Sprintf("%s.txt", *filename)

	addresses, err := parsefile.Parse(file)
	if err != nil {
		log.Fatalf("could parse the file %s: %v", file, err)
	}

	conninit := sshinit.NewSSHInit(*username, *password)

	var notWorking []string
	addressBatch := make(chan string, 3)
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

	fmt.Print("\n")
	fmt.Println("We couldn't connect to following devices:")
	for _, notWorkAdd := range notWorking {
		fmt.Printf("To %s can't connect \n", notWorkAdd)
	}
}
