package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type Host struct {
	hostname string
	ip       string
}

//Parse the provided txt file
func parse(filename string) ([]Host, error) {
	hosts := []Host{}
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open the file %s", filename)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return hosts, err
		}
		result := strings.Split(line, " ")
		host := strings.TrimSpace(result[0])
		ip := strings.TrimSpace(result[1])
		hosts = append(hosts, Host{hostname: host, ip: ip})
	}
	return hosts, nil
}
