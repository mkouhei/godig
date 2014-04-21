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


func Query(msg *dns.Msg, t uint16, c *dns.Client, server string, domain *string) {
	msg.SetQuestion(*domain, t)

	r, _, err := c.Exchange(msg, server)
	if err != nil {
		return
	}
	if r.Rcode != dns.RcodeSuccess {
		return
	}
	
	for _, a := range r.Answer {
		if (t == 1) {
			if res, ok := a.(*dns.A); ok {
				fmt.Printf("%s\n", res.String())
			}
		} else if (t == 28) {
			if res, ok := a.(*dns.AAAA); ok {
				fmt.Printf("%s\n", res.String())
			}
		}
	}

}


func main() {
	var domain = flag.String("d", "example.org.", "specify domain name (example.org.)")
	flag.Parse()

	config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")
	c := new(dns.Client)
	m := new(dns.Msg)
	m.RecursionDesired = true
	server := config.Servers[0]+":"+config.Port

	Query(m, dns.TypeA, c, server, domain)
	Query(m, dns.TypeAAAA, c, server, domain)

}
