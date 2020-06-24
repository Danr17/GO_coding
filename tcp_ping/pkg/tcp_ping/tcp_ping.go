package tcping

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Danr17/GO_scripts/tree/master/Scripts/tcp_ping/pkg/ping"
	"github.com/Danr17/GO_scripts/tree/master/Scripts/tcp_ping/pkg/utils"
)

// TCPing ...
type TCPing struct {
	target *ping.Target
	Done   chan struct{}
	result *ping.Result
}

// NewTCPing return a new TCPing
func NewTCPing() *TCPing {
	tcping := TCPing{
		Done: make(chan struct{}),
	}
	return &tcping
}

// SetTarget set target for TCPing
func (tcping *TCPing) SetTarget(target *ping.Target) {
	tcping.target = target
	if tcping.result == nil {
		tcping.result = &ping.Result{Target: target}
	}
}

// Result return the result
func (tcping TCPing) Result() *ping.Result {
	return tcping.result
}

// Start a tcping
func (tcping TCPing) Start() {
	go func() {
		t := time.NewTicker(tcping.target.Interval)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				duration, remoteAddr, err := tcping.ping()
				tcping.result.Counter++

				if err != nil {
					fmt.Printf("Ping %s - failed: %s\n", tcping.target, err)
				} else {
					fmt.Printf("Ping %s(%s) - Connected - time=%s\n", tcping.target, remoteAddr, duration)

					if tcping.result.Counter == 1 {
						tcping.result.MinDuration = duration
						tcping.result.MaxDuration = duration
					}

					switch {
					case duration > tcping.result.MaxDuration:
						tcping.result.MaxDuration = duration
					case duration < tcping.result.MinDuration:
						tcping.result.MinDuration = duration
					}

					tcping.result.SuccessCounter++
					tcping.result.TotalDuration += duration
				}
				if tcping.result.Counter >= tcping.target.Counter && tcping.target.Counter != 0 {
					log.Println("ping done for site", tcping.target.Host)
					tcping.Stop()
					break
				}
			}
		}
	}()
}

// Stop the tcping
func (tcping *TCPing) Stop() {
	tcping.Done <- struct{}{}
}

func (tcping TCPing) ping() (time.Duration, net.Addr, error) {
	var remoteAddr net.Addr
	duration, errIfce := utils.TimeIt(func() interface{} {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", tcping.target.Host, tcping.target.Port), tcping.target.Timeout)
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
