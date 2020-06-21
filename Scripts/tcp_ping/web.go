package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

type website struct {
	target *Target
	result *Result
}

//WebPing data
type WebPing struct {
	sites []*website
}

// NewWebPing return a new WebPing
func NewWebPing(targets []*Target) *WebPing {
	sites := []*website{}
	for _, target := range targets {
		site := website{
			target: target,
		}
		if site.result == nil {
			site.result = &Result{Target: target}
		}
		sites = append(sites, &site)
	}
	return &WebPing{
		sites: sites,
	}
}

// Start a webping
func (webping WebPing) Start(done chan struct{}) {
	var wg sync.WaitGroup
startLoop:
	for {
		select {
		case <-done:
			break startLoop
		default:
			wg.Add(len(webping.sites))
			for _, site := range webping.sites {
				go pingHost(site, &wg)
			}
			wg.Wait()

		}
		time.Sleep(5 * time.Second)
	}
}

func pingHost(site *website, wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.NewTicker(site.target.Interval)
	defer t.Stop()
pingLoop:
	for {
		select {
		case <-t.C:
			duration, remoteAddr, err := site.ping()
			site.result.Counter++

			if err != nil {
				fmt.Printf("Ping %s - failed: %s\n", site.target, err)
			} else {
				fmt.Printf("Ping %s(%s) - Connected - time=%s\n", site.target, remoteAddr, duration)

				if site.result.Counter == 1 {
					site.result.MinDuration = duration
					site.result.MaxDuration = duration
				}

				switch {
				case duration > site.result.MaxDuration:
					site.result.MaxDuration = duration
				case duration < site.result.MinDuration:
					site.result.MinDuration = duration
				}

				site.result.SuccessCounter++
				site.result.TotalDuration += duration
			}
			if site.result.Counter >= site.target.Counter && site.target.Counter != 0 {
				log.Printf("site %s and counter %d", site.target.Host, site.result.Counter)
				site.result.Counter = 0
				break pingLoop
			}
		}
	}
}

func (site *website) ping() (time.Duration, net.Addr, error) {
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
