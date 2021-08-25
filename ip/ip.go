package ip

import (
	"errors"
	"net"
	"time"
)

var ErrFecthingIp = errors.New("error fetching current ip")

type Updates <-chan net.IP

// A Provider obtains the current ip address from somewhere
// If Provider fails to obtain an ip, then it returns an error
type Provider interface {
	ObtainIp() (net.IP, error)
}

// DefaultProvider is an instance of Provider that acts as the default
var DefaultProvider Provider = NewGoogleProvider()

// GetIpChanges returns a channel called Updates that receives a net.IP address every time the ip changes.
// Internally, it keeps track of the current ip in memory even though that could change between implementors
var GetIpChanges func(interval time.Duration) Updates = getIpChangesFromDefaultProvider

func getIpChangesFromDefaultProvider(interval time.Duration) Updates {
	ch := make(chan net.IP)
	go func() {
		var ip net.IP
		for {
			ip2, err := DefaultProvider.ObtainIp()
			if err != nil {
				time.Sleep(interval)
				continue
			}
			if !ip2.Equal(ip) {
				ch <- ip2
				ip = ip2
 			}
			time.Sleep(interval)
		}
	}()
	return ch
}