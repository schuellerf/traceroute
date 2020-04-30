// +build windows

// Package traceroute not yet implemented in windows
package traceroute

import (
	"time"
	"fmt"
)
const DEFAULT_MAX_HOPS = 64

// TracerouteHop type
type TracerouteHop struct {
	Success     bool
	Address     [4]byte
	Host        string
	N           int
	ElapsedTime time.Duration
	TTL         int
}

// TracrouteOptions type
type TracerouteOptions struct {
	port       int
	maxHops    int
	firstHop   int
	timeoutMs  int
	retries    int
	packetSize int
}

// TracerouteResult type
type TracerouteResult struct {
	DestinationAddress [4]byte
	Hops               []TracerouteHop
}

func (options *TracerouteOptions) MaxHops() int {
	if options.maxHops == 0 {
		options.maxHops = DEFAULT_MAX_HOPS
	}
	return options.maxHops
}
func (options *TracerouteOptions) SetMaxHops(maxHops int) {
	options.maxHops = maxHops
}

func (hop *TracerouteHop) AddressString() string {
	return fmt.Sprintf("%v.%v.%v.%v", hop.Address[0], hop.Address[1], hop.Address[2], hop.Address[3])
}

func Traceroute(dest string, options *TracerouteOptions, c ...chan TracerouteHop) (result TracerouteResult, err error) {
	return
}
