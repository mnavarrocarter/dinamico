package main

import (
	"flag"
	"fmt"
	"github.com/mnavarrocarter/dinamico/dns"
	"github.com/mnavarrocarter/dinamico/ip"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	path := flag.String("state", "domains.txt", "The file containing the list of the domains to check")
	duration := flag.Duration("interval", time.Second * 60, "The interval to check for ip")
	timeout := flag.Duration("timeout", time.Second * 10, "Timeout for http requests")
	server := flag.String("server", "", "The address to enable the http server on")
	verbose := flag.Bool("log", false, "Enable logs")

	flag.Parse()

	if timeout != nil {
		http.DefaultClient.Timeout = *timeout
	}

	if verbose != nil && *verbose == true {
		dns.WithLogging(nil)
		ip.WithLogging(nil)
	}

	if path == nil {
		fmt.Println("You must provide a config file path")
		os.Exit(1)
	}

	file, err := os.Open(*path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	records, err := dns.ParseRecords(file)

	if server != nil && *server != "" {
		// Run a TCP server in a go routine
		// ListenAndServe
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file.Close()

	ipUpdates := ip.GetIpChanges(*duration)

	for {
		i := <- ipUpdates
		for _, u := range records.All() {
			go func(u *url.URL) {
				_ = dns.DefaultUpdater.Update(u, i)
			}(u)
		}
	}
}