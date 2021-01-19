package cliping

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	icmpPing "github.com/sparrc/go-ping"
)

//StartICMP starts an icmp ping
func StartICMP(args []string, count int, timeoutDuration time.Duration, intervalDuration time.Duration, privileged bool) chan bool {
	pinger, err := icmpPing.NewPinger(args[0])
	if err != nil {
		log.Fatalf("ERROR: %s\n", err.Error())
	}

	//notify the main function that it is done
	done := make(chan bool)

	// listen for ctrl-C signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			pinger.Stop()
		}
	}()

	pinger.OnRecv = func(pkt *icmpPing.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
	}
	pinger.OnFinish = func(stats *icmpPing.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	pinger.Count = count
	pinger.Interval = intervalDuration
	//pinger.Timeout = timeoutDuration

	pinger.SetPrivileged(privileged)

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	pinger.Run()
	close(done)
	return done
}
