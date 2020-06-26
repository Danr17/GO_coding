package webping

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	icmpPing "github.com/sparrc/go-ping"

	"github.com/Danr17/GO_scripts/tree/master/ping_util/pkg/ping"
	"github.com/Danr17/GO_scripts/tree/master/ping_util/pkg/utils"
)

type website struct {
	Target *ping.Target
	Result *ping.Result
}

//WebPing data
type WebPing struct {
	sites []*website
}

// NewWebPing return a new WebPing
func NewWebPing(targets []*ping.Target) *WebPing {
	sites := []*website{}
	for _, target := range targets {
		site := website{
			Target: target,
		}
		if site.Result == nil {
			site.Result = &ping.Result{Target: target}
		}
		sites = append(sites, &site)
	}
	return &WebPing{
		sites: sites,
	}
}

// Start a webping
func (webping WebPing) Start(done chan bool) {
	var wg sync.WaitGroup
startLoop:
	for {
		select {
		case <-done:
			break startLoop
		default:
			wg.Add(len(webping.sites))
			for _, site := range webping.sites {
				if site.Target.Proto == "icmp" {
					go pingIcmp(site, &wg)
				} else {
					go pingHost(site, &wg)
				}
			}
			wg.Wait()

		}
		time.Sleep(5 * time.Second)
	}
}

func pingHost(site *website, wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.NewTicker(site.Target.Interval)
	defer t.Stop()
pingLoop:
	for {
		select {
		case <-t.C:
			duration, _, err := site.ping()
			site.Result.Counter++

			if err != nil {
				break
			} else {
				if site.Result.Counter == 1 {
					site.Result.MinDuration = duration
					site.Result.MaxDuration = duration
				}

				switch {
				case duration > site.Result.MaxDuration:
					site.Result.MaxDuration = duration
				case duration < site.Result.MinDuration:
					site.Result.MinDuration = duration
				}

				site.Result.SuccessCounter++
				site.Result.TotalDuration += duration
			}
			if site.Result.Counter >= site.Target.Counter && site.Target.Counter != 0 {
				//log.Printf("site %s and counter %d", site.Target.Host, site.Result.Counter)
				site.Result.Counter = 0
				site.Result.Status = true
				hour, min, sec := time.Now().Local().Clock()
				site.Result.LastSeen = strconv.Itoa(hour) + ":" + strconv.Itoa(min) + ":" + strconv.Itoa(sec)
				break pingLoop
			}
		}
	}
}

func pingIcmp(site *website, wg *sync.WaitGroup) {
	defer wg.Done()
	pinger, err := icmpPing.NewPinger(site.Target.Host)
	if err != nil {
		panic(err)
	}
	pinger.Count = site.Target.Counter
	pinger.Interval = site.Target.Interval

	pinger.SetPrivileged(site.Target.Privileged)

	pinger.Run() // blocks until finished

	stats := pinger.Statistics()
	hour, min, sec := time.Now().Local().Clock()
	site.Result.LastSeen = strconv.Itoa(hour) + ":" + strconv.Itoa(min) + ":" + strconv.Itoa(sec)
	site.Result.MinDuration = stats.MinRtt
	site.Result.MaxDuration = stats.MaxRtt
	site.Result.Status = true
	if stats.PacketLoss > 0 {
		site.Result.Status = false
	}
}

func (site *website) ping() (time.Duration, net.Addr, error) {
	var remoteAddr net.Addr
	duration, errIfce := utils.TimeIt(func() interface{} {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", site.Target.Host, site.Target.Port), site.Target.Timeout)
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
