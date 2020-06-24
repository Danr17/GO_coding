package parse

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

//Host holds host informations
type Host struct {
	name string
	ip   string
}

//ParseFile parse the provided txt file
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
		hosts = append(hosts, Host{name: host, ip: ip})
	}
	return hosts, nil
}
