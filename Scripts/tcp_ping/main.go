package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	//	"time"
)

const usage = `1. ping over tcp
    > tcping google.com
2. ping over tcp with custom port
    > tcping google.com 443
3. ping over http
    > tcping -H google.com
4. ping with URI schema
    > tcping http://hui.lu`

var (
	isWeb     = flag.Bool("web", false, "enable this if you want to see it on Web")
	inWebPort = flag.Int("port", 443, "enable this if you want to see it on Web")
	inWebFile = flag.String("file", "", "specify the filename")
	counter   = flag.Int("counter", 4, "ping counter")
	timeout   = flag.String("timeout", "1s", `connect timeout, units are "ns", "us" (or "µs"), "ms", "s", "m", "h"`)
	interval  = flag.String("interval", "1s", `ping interval, units are "ns", "us" (or "µs"), "ms", "s", "m", "h"`)
)

func main() {
	flag.Parse()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println(usage)
		os.Exit(1)
	}

	var port = 443
	if len(args) == 2 {
		port, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("port should be integer")
			return
		}
	}

	host := args[1]
	parseHost := FormatIP(host)
	target := Target{
		//	Timeout:  timeoutDuration,
		//	Interval: intervalDuration,
		Host:    parseHost,
		Port:    port,
		Counter: *counter,
	}

	if *isWeb == false {
		pinger := NewTCPing()
		pinger.SetTarget(&target)
		pingerDone := pinger.Start()
		select {
		case <-pingerDone:
			break
		case <-sigs:
			break
		}

		fmt.Println(pinger.Result())
	}

	if *inWebFile == "" {
		fmt.Println(`Filename and port number should be specified:
Example usage: tcp_ping web=true file="filename" port=443`)
	}
}
