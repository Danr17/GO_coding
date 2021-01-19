package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	cliping "github.com/Danr17/GO_scripts/tree/master/ping_util/pkg/cli_ping"
	"github.com/Danr17/GO_scripts/tree/master/ping_util/pkg/utils"
	webping "github.com/Danr17/GO_scripts/tree/master/ping_util/pkg/web_ping"
)

const usage = `WEB:
1. Web Ping with defauts (defaults: port=443, interval=1s, counter=4)
    > ping_util-linux-amd64 -p web -file example.txt
2. Web Ping with count 10, interval 3s, timeout 3s
    > ping_util-linux-amd64 -p web -file example.txt -c 10 -i 3s -t 3s

CLI:
1. ping over TCP  with defaults (defaults: port=443, interval=1s, counter=4)
    > ping_util-linux-amd64 example.com
2. ping over TCP over with custom port, counter and interval
    > ping_util-linux-amd64 -p tcp -port 80 -c 3 -i 3s example.com

3. ping over UDP with defaults
	> ping_util-linux-amd64 -p udp example.com

4. ping over ICMP in Privilege (run as super-user) mode !!!
	> sudo ./ping_util -p icmp -privileged example.com
5. ping over ICMP without Privilege mode, is actually over UDP
	> sudo ./ping_util -p icmp example.com 
`

var (
	proto      = flag.String("p", "tcp", "protocol, units web, tcp, udp, icmp")
	port       = flag.Int("port", 443, "port, default is 443")
	inWebFile  = flag.String("file", "", "specify the filename")
	counter    = flag.Int("c", 4, "ping counter")
	timeout    = flag.String("t", "1s", `connect timeout, units are "ns", "us" (or "µs"), "ms", "s", "m", "h"`)
	interval   = flag.String("i", "1s", `ping interval, units are "ns", "us" (or "µs"), "ms", "s", "m", "h"`)
	privileged = flag.Bool("privileged", false, "required for ICMP Ping, run as super-user")
)

func main() {
	flag.Parse()
	args := flag.Args()

	//catch CTRL-C
	sigs := make(chan os.Signal, 1)
	//notify web to stop
	done := make(chan bool, 1)
	//channel to listen for errors coming from the listener.
	serverErrors := make(chan error, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	timeoutDuration, err := utils.ConvertTime(*timeout)
	if err != nil {
		log.Fatalln("The value provided for **timeout** is wrong")
	}
	intervalDuration, err := utils.ConvertTime(*interval)
	if err != nil {
		log.Fatalln("The value provided for **interval** is wrong")
	}

	var server *http.Server
	var WEBpinger *webping.WebPing
	var CLIpinger *cliping.CLIping

	switch *proto {
	case "tcp", "udp":
		CLIpinger = cliping.NewCLIping()
		done = startCLI(CLIpinger, args, timeoutDuration, intervalDuration)
	case "icmp":
		done = cliping.StartICMP(args, *counter, intervalDuration, timeoutDuration, *privileged)

	case "web":
		if *inWebFile == "" {
			fmt.Println(usage)
		}

		server, WEBpinger = startWeb(args, timeoutDuration, intervalDuration)

		go func() {
			log.Println("API listening on port 8080. Open browser: http://localhost:8080")
			serverErrors <- server.ListenAndServe()
		}()

		go WEBpinger.Start(done)
	default:
		log.Panicln("The value provided for protocol is not valid, should be --p web, --p tcp, --p udp or --p icmp")

	}

	select {
	case err := <-serverErrors:
		log.Fatalf("error: starting server: %s", err)
	case <-sigs:
		if *proto == "web" {
			log.Println("Closing HTTP server")
			server.Close()
			close(done)
			return
		}
		CLIpinger.Stop()
		return
	case <-done:
		return
	}

}
