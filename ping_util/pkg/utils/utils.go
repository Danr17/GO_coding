package utils

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

//TimeIt a wrapper around function t measure the time it took to execute
func TimeIt(f func() interface{}) (int64, interface{}) {
	startAt := time.Now()
	res := f()
	endAt := time.Now()
	return endAt.UnixNano() - startAt.UnixNano(), res
}

// FormatIP - trim spaces and format IP
//
// IP - the provided IP
//
// string - return "" if the input is neither valid IPv4 nor valid IPv6
//          return IPv4 in format like "192.168.9.1"
//          return IPv6 in format like "[2002:ac1f:91c5:1::bd59]"
func FormatIP(IP string) string {

	host := strings.Trim(IP, "[ ]")
	if parseIP := net.ParseIP(host); parseIP != nil && parseIP.To4() == nil {
		host = fmt.Sprintf("[%s]", host)
	}

	return host
}

//ConvertTime converts a string into time.Duration
func ConvertTime(timeout string) (time.Duration, error) {
	var result time.Duration
	if res, err := strconv.Atoi(timeout); err == nil {
		result = time.Duration(res) * time.Millisecond
	} else {
		result, err = time.ParseDuration(timeout)
		if err != nil {
			fmt.Println("parse timeout failed", err)
			return 1 * time.Second, err
		}
	}
	return result, nil
}
