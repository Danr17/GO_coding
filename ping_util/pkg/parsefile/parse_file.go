package parse

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

//Host holds host informations
type Host struct {
	Name     string
	IP       string
	Protocol string
	Port     int
}

//File parse the provided txt file
func File(filename string) ([]Host, error) {
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
		protocol := strings.TrimSpace(result[2])
		portRaw := strings.TrimSpace(result[3])
		port, err := strconv.Atoi(portRaw)
		if err != nil {
			log.Fatalln("The port number provided is not a Number")
		}
		hosts = append(hosts,
			Host{
				Name:     host,
				IP:       ip,
				Protocol: protocol,
				Port:     port,
			})
	}
	return hosts, nil
}
