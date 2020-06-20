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
	done   chan struct{}
}

//WebPing data
type WebPing struct {
	sites   []*website
	bigdone chan struct{}
}

// NewWebPing return a new WebPing
func NewWebPing(targets []*Target) *WebPing {
	sites := []*website{}
	for _, target := range targets {
		site := website{
			target: target,
			done:   make(chan struct{}),
		}
		if site.result == nil {
			site.result = &Result{Target: target}
		}
		sites = append(sites, &site)
	}
	return &WebPing{
		sites:   sites,
		bigdone: make(chan struct{}),
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
				go func(site *website) {
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
							if site.result.Counter >= site.target.Counter && site.target.Counter != 0 {
								log.Printf("site %s and counter %d", site.target.Host, site.result.Counter)
								site.result.Counter = 0
								//site.Stop()
								//return
								//site.done <- struct{}{}
								break pingLoop
							}
							// case <-done:
							// 	log.Printf("site %s is done", site.target.Host)
							// 	break startLoop
						}
					}
				}(site)
			}
			wg.Wait()
			time.Sleep(10 * time.Second)
		}
	}
}

// Stop the webping
func (site *website) Stop() {
	site.done <- struct{}{}
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
