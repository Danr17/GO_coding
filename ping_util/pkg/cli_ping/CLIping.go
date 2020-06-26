package cliping

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Danr17/GO_scripts/tree/master/ping_util/pkg/ping"
	"github.com/Danr17/GO_scripts/tree/master/ping_util/pkg/utils"
)

// CLIping ...
type CLIping struct {
	target *ping.Target
	Done   chan bool
	result *ping.Result
}

// NewCLIping return a new TCPing
func NewCLIping() *CLIping {
	cliping := CLIping{
		Done: make(chan bool),
	}
	return &cliping
}

// SetTarget set target for TCPing
func (cliping *CLIping) SetTarget(target *ping.Target) {
	cliping.target = target
	if cliping.result == nil {
		cliping.result = &ping.Result{Target: target}
	}
}

// Result return the result
func (cliping CLIping) Result() *ping.Result {
	return cliping.result
}

// Start a tcping
func (cliping CLIping) Start() {
	go func() {
		t := time.NewTicker(cliping.target.Interval)
		defer t.Stop()
		for {
			select {
			case <-cliping.Done:
				return
			case <-t.C:
				duration, remoteAddr, err := cliping.ping()
				cliping.result.Counter++

				if err != nil {
					fmt.Printf("Ping %s - failed: %s\n", cliping.target, err)
				} else {
					fmt.Printf("Ping %s(%s) - Connected - time=%s\n", cliping.target, remoteAddr, duration)

					if cliping.result.Counter == 1 {
						cliping.result.MinDuration = duration
						cliping.result.MaxDuration = duration
					}

					switch {
					case duration > cliping.result.MaxDuration:
						cliping.result.MaxDuration = duration
					case duration < cliping.result.MinDuration:
						cliping.result.MinDuration = duration
					}

					cliping.result.SuccessCounter++
					cliping.result.TotalDuration += duration
				}
				if cliping.result.Counter >= cliping.target.Counter && cliping.target.Counter != 0 {
					log.Println("ping done for site", cliping.target.Host)
					cliping.Stop()
					break
				}
			}
		}
	}()
}

// Stop the tcping
func (cliping *CLIping) Stop() {
	close(cliping.Done)
}

func (cliping CLIping) ping() (time.Duration, net.Addr, error) {
	var remoteAddr net.Addr
	var duration int64
	var errIfce interface{}

	switch cliping.target.Proto {
	case "tcp":
		duration, errIfce = utils.TimeIt(func() interface{} {
			conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", cliping.target.Host, cliping.target.Port), cliping.target.Timeout)
			if err != nil {
				return err
			}
			remoteAddr = conn.RemoteAddr()
			conn.Close()
			return nil
		})

	//	return time.Duration(duration), remoteAddr, nil
	case "udp":
		duration, errIfce = utils.TimeIt(func() interface{} {
			conn, err := net.DialTimeout("udp", fmt.Sprintf("%s:%d", cliping.target.Host, cliping.target.Port), cliping.target.Timeout)
			if err != nil {
				return err
			}
			remoteAddr = conn.RemoteAddr()
			conn.Close()
			return nil
		})
	default:
		log.Panicln("The value provided for protocol is not valid, should be --p web, --p tcp, --p udp or --p icmp")
	}
	if errIfce != nil {
		err := errIfce.(error)
		return 0, remoteAddr, err
	}
	return time.Duration(duration), remoteAddr, nil
}
