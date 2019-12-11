package main

import (
	"flag"
	"fmt"

	"github.com/sparrc/go-ping"
)

var (
	// Declare the variables of flags here
	count int
)

type Host string
type hostFlag struct{ Host }

func (h *hostFlag) Set(s string) error {
	// Implement the interface and parse the host name
}

func HostFlag(name string, host Host, usage string) *Host {
	// Implement by yourself
}

func Init() {
	flag.IntVar(&count, "c", 4, "Count")
	// Define the other flags here
}

func main() {
	Init()
	flag.Parse()

	pinger, err := ping.NewPinger(string(*host))

	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	pinger.SetPrivileged(true)
	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}
	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}
	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())

	// You can set the flag value to the pinger as the following
	pinger.Count = count
	// pinger.Count = 4
	// pinger.Interval = time.Second
	// pinger.Timeout = time.Second
	// pinger.Size = 100

	pinger.Run()
}
