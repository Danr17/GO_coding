package main

import (
	"fmt"
	"net"
	"time"
)

type website struct {
	target *Target
	result *Result
}

//WebPing data
type WebPing struct {
	sites []*website
	done  chan struct{}
}

// NewWebPing return a new WebPing
func NewWebPing() *WebPing {

	webping := WebPing{
		done: make(chan struct{}),
	}
	return &webping
}

// SetTarget set target for WebPing
func (webping *WebPing) SetTarget(targets []*Target) {
	for i, target := range targets {
		webping.sites[i].target = target
		if webping.sites[i].result == nil {
			webping.sites[i].result = &Result{Target: target}
		}
	}
}

// Start a webping
func (webping WebPing) Start() <-chan struct{} {
	for _, site := range webping.sites {
		go func(site *website) {
			t := time.NewTicker(site.target.Interval)
			defer t.Stop()
			for {
				select {
				case <-t.C:
					if site.result.Counter >= site.target.Counter && site.target.Counter != 0 {
						webping.Stop()
						return
					}
					duration, remoteAddr, err := site.ping()
					site.result.Counter++

					if err != nil {
						fmt.Printf("Ping %s - failed: %s\n", site.target, err)
					} else {
						fmt.Printf("Ping %s(%s) - Connected - time=%s\n", site.target, remoteAddr, duration)

						if site.result.MinDuration == 0 {
							site.result.MinDuration = duration
						}
						if site.result.MaxDuration == 0 {
							site.result.MaxDuration = duration
						}
						site.result.SuccessCounter++
						if duration > site.result.MaxDuration {
							site.result.MaxDuration = duration
						} else if duration < site.result.MinDuration {
							site.result.MinDuration = duration
						}
						site.result.TotalDuration += duration
					}
				case <-webping.done:
					return
				}
			}
		}(site)
	}
	return webping.done
}

// Stop the webping
func (webping *WebPing) Stop() {
	webping.done <- struct{}{}
}

func (site website) ping() (time.Duration, net.Addr, error) {
	var remoteAddr net.Addr
	duration, errIfce := timeIt(func() interface{} {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", site.target.Host, site.target.Port), site.target.Timeout)
		if err != nil {
			return err
		}
		remoteAddr = conn.RemoteAddr()
		conn.Close()
		return nil
	})
	if errIfce != nil {
		err := errIfce.(error)
		return 0, remoteAddr, err
	}
	return time.Duration(duration), remoteAddr, nil
}
