package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	cliping "github.com/Danr17/GO_scripts/tree/master/ping_util/pkg/cli_ping"
	parse "github.com/Danr17/GO_scripts/tree/master/ping_util/pkg/parsefile"
	"github.com/Danr17/GO_scripts/tree/master/ping_util/pkg/ping"
	"github.com/Danr17/GO_scripts/tree/master/ping_util/pkg/utils"
	webping "github.com/Danr17/GO_scripts/tree/master/ping_util/pkg/web_ping"
)

func startCLI(pinger *cliping.CLIping, args []string, timeoutDuration time.Duration, intervalDuration time.Duration) chan bool {
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
		Proto:    *proto,
		Counter:  *counter,
	}

	pinger.SetTarget(&target)
	pinger.Start()
	<-pinger.Done
	fmt.Println(pinger.Result())
	return (pinger.Done)

}

func startWeb(args []string, timeoutDuration time.Duration, intervalDuration time.Duration) (server *http.Server, pinger *webping.WebPing) {
	hosts, err := parse.File(*inWebFile)
	if err != nil {
		log.Fatalf("could parse the file %s: %v", *inWebFile, err)
	}
	targets := []*ping.Target{}
	for _, host := range hosts {
		webtarget := ping.Target{
			Timeout:    timeoutDuration,
			Interval:   intervalDuration,
			Host:       host.IP,
			HostName:   host.Name,
			Port:       host.Port,
			Proto:      host.Protocol,
			Counter:    *counter,
			Privileged: *privileged,
		}
		targets = append(targets, &webtarget)
	}

	pinger = webping.NewWebPing(targets)

	mux := http.NewServeMux()
	mux.HandleFunc("/", webping.HTMLPage(pinger))

	server = &http.Server{
		Addr:         "localhost:8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	return server, pinger

}
