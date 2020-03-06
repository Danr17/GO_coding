package parsefile

import (
	"bufio"
	"io"
	"log"
	"os"
)

//Parse the provided txt file
func Parse(filename string) ([]string, error) {
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
