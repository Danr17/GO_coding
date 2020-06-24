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

	parse "github.com/Danr17/GO_scripts/tree/master/Scripts/tcp_ping/pkg/parsefile"
	"github.com/Danr17/GO_scripts/tree/master/Scripts/tcp_ping/pkg/ping"
	tcping "github.com/Danr17/GO_scripts/tree/master/Scripts/tcp_ping/pkg/tcp_ping"
	"github.com/Danr17/GO_scripts/tree/master/Scripts/tcp_ping/pkg/utils"
	"github.com/Danr17/GO_scripts/tree/master/Scripts/tcp_ping/pkg/web"
)

const usage = `WEB:
1. Web Ping with defauts (defaults: port=443, interval=1s, counter=4)
    > tcp_ping-linux-amd64 -web -file=example.txt
2. Web Ping with over port=80, interval=3s, counter=10
    > tcp_ping-linux-amd64 -web -file=example.txt -port 80 -interval 3s -counter 10

CLI:
1. ping over tcp  with defaults (defaults: port=443, interval=1s, counter=4)
    > tcp_ping-linux-amd64 example.com
2. ping over tcp over custom port
    > tcp_ping-linux-amd64 -port 22 example.com 
3. ping over tcp using counter and interval
    > tcp_ping-linux-amd64 -port 80 -counter 3 -interval 3s example.com
`

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

	timeoutDuration, err := utils.ConvertTime(*timeout)
	if err != nil {
		log.Fatalln("The value provided for **timeout** is wrong")
	}
	intervalDuration, err := utils.ConvertTime(*interval)
	if err != nil {
		log.Fatalln("The value provided for **interval** is wrong")
	}

	if !*isWeb {
		if len(args) < 1 {
			fmt.Println(usage)
			os.Exit(1)
		}

		host := args[0]
		parseHost := utils.FormatIP(host)

		target := ping.Target{
			Timeout:  timeoutDuration,
			Interval: intervalDuration,
			Host:     parseHost,
			Port:     *port,
			Counter:  *counter,
		}

		pinger := tcping.NewTCPing()
		pinger.SetTarget(&target)
		pinger.Start()
		<-pinger.done

		fmt.Println(pinger.Result())
		return
	}
	if *inWebFile == "" {
		fmt.Println(usage)
	}
	hosts, err := parse.File(*inWebFile)
	if err != nil {
		log.Fatalf("could parse the file %s: %v", *inWebFile, err)
	}
	targets := []*ping.Target{}
	for _, host := range hosts {
		webtarget := ping.Target{
			Timeout:  timeoutDuration,
			Interval: intervalDuration,
			Host:     host.ip,
			HostName: host.name,
			Port:     *port,
			Counter:  *counter,
		}
		targets = append(targets, &webtarget)
	}

	pinger := web.NewWebPing(targets)

	mux := http.NewServeMux()
	mux.HandleFunc("/", web.HTMLPage(pinger))

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
