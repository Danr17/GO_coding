package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const usage = `1. ping over tcp
    > tcp_ping-linux-amd64 google.com
2. ping over tcp with custom port
    > tcp_ping-linux-amd64 google.com -port=443
3. ping over tcp using counter and interval
    > tcp_ping-linux-amd64 google.com -counter=10 -interval="3s"
4. Web Ping over provided sources
    > tcp_ping-linux-amd64 -web -file=example.txt`

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
	args := flag.Args()

	//catch CTRL-C
	sigs := make(chan os.Signal, 1)
	//notify web to stop
	done := make(chan struct{}, 1)
	//channel to listen for errors coming from the listener.
	serverErrors := make(chan error, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		close(done)
	}()

	timeoutDuration, err := convertTime(*timeout)
	if err != nil {
		log.Fatalln("The value provided for **timeout** is wrong")
	}
	intervalDuration, err := convertTime(*interval)
	if err != nil {
		log.Fatalln("The value provided for **interval** is wrong")
	}

	if !*isWeb {
		if len(args) < 1 {
			fmt.Println(usage)
			os.Exit(1)
		}

		host := args[0]
		parseHost := FormatIP(host)

		target := Target{
			Timeout:  timeoutDuration,
			Interval: intervalDuration,
			Host:     parseHost,
			Port:     *port,
			Counter:  *counter,
		}

		pinger := NewTCPing()
		pinger.SetTarget(&target)
		pinger.Start()
		<-pinger.done

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
	targets := []*Target{}
	for _, host := range hosts {
		webtarget := Target{
			Timeout:  timeoutDuration,
			Interval: intervalDuration,
			Host:     host.ip,
			HostName: host.name,
			Port:     *port,
			Counter:  *counter,
		}
		targets = append(targets, &webtarget)
	}

	pinger := NewWebPing(targets)

	mux := http.NewServeMux()
	mux.HandleFunc("/", htmlPage(pinger))

	server := http.Server{
		Addr:         "localhost:8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Println("API listening on port 8080. Open browser: http://localhost:8080")
		serverErrors <- server.ListenAndServe()
	}()

	go pinger.Start(done)

	select {
	case err := <-serverErrors:
		log.Fatalf("error: starting server: %s", err)
	case <-sigs:
		server.Close()
		return
	}

	// fmt.Println(pinger.Result())
	return
}
