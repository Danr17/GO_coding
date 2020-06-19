package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const usage = `1. ping over tcp
    > tcping google.com
2. ping over tcp with custom port
    > tcping google.com -port=443
3. Web Ping over provided sources
    > tcping -web=true -file=devices.txt`

var (
	isWeb     = flag.Bool("web", false, "enable this if you want to see it on Web")
	inWebFile = flag.String("file", "", "specify the filename")
	port      = flag.Int("port", 443, "enable this if you want to see it on Web")
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

	host := args[0]
	parseHost := FormatIP(host)

	timeoutDuration, err := convertTime(*timeout)
	if err != nil {
		log.Fatalln("The value provided for **timeout** is wrong")
	}
	intervalDuration, err := convertTime(*interval)
	if err != nil {
		log.Fatalln("The value provided for **interval** is wrong")
	}

	target := Target{
		Timeout:  timeoutDuration,
		Interval: intervalDuration,
		Host:     parseHost,
		Port:     *port,
		Counter:  *counter,
	}

	if !*isWeb {
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
		return
	}

	if *inWebFile == "" {
		fmt.Println(`The filename should be specified (port is 443 default if not changed):
Example usage: tcp_ping web=true file="filename" port=443`)
	}
	hosts, err := parse(*inWebFile)
	if err != nil {
		log.Fatalf("could parse the file %s: %v", *inWebFile, err)
	}
	pinger := NewWebPing()
	targets := []*Target{}
	for _, host := range hosts {
		target = Target{
			Timeout:  timeoutDuration,
			Interval: intervalDuration,
			Host:     host.ip,
			Port:     *port,
			Counter:  *counter,
		}
		targets = append(targets, &target)
	}
	pinger.SetTarget(targets)
	pingerDone := pinger.Start()
	select {
	case <-pingerDone:
		break
	case <-sigs:
		break
	}

	// fmt.Println(pinger.Result())
	return
}
