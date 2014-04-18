package main

/*
This code is my learning of Golang.

see also bellow blog.
http://archive.miek.nl/blog/archives/2012/12/07/printing_mx_records_with_go_dns_take_3/index.html
*/

import (
	"github.com/miekg/dns"
	"fmt"
	"flag"
)

func main() {
	var domain = flag.String("d", "example.org.", "specify domain name (example.org.)")
	flag.Parse()

	config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")
	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(*domain, dns.TypeSOA)
	m.RecursionDesired = true
	r, _, err := c.Exchange(m, config.Servers[0]+":"+config.Port)
	if err != nil {
		return
	}
	if r.Rcode != dns.RcodeSuccess {
		return
	}
	for _, a := range r.Answer {
		if soa, ok := a.(*dns.SOA); ok {
			fmt.Printf("%s\n", soa.String())
		}
	}
}
