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

func Query(msg *dns.Msg, recordType uint16, client *dns.Client, server string, domain *string) {
	msg.SetQuestion(*domain, recordType)

	r, _, err := client.Exchange(msg, server)
	if err != nil {
		return
	}
	if r.Rcode != dns.RcodeSuccess {
		return
	}

	for _, a := range r.Answer {
		fmt.Printf("%s\n", a.String())
	}

}


func main() {
	var domain = flag.String("d", "example.org.", "specify domain name (example.org.)")
	flag.Parse()

	config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")
	client := new(dns.Client)
	msg := new(dns.Msg)
	msg.RecursionDesired = true
	server := config.Servers[0]+":"+config.Port

	Query(msg, dns.TypeA, client, server, domain)
	Query(msg, dns.TypeAAAA, client, server, domain)
}
