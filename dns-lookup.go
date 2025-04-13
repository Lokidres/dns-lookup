package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func lookupA(domain string) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Printf("Error looking up A records: %v\n", err)
		return
	}
	fmt.Println("=== A Records ===")
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			fmt.Printf("%s\n", ipv4)
		}
	}
}

func lookupMX(domain string) {
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		fmt.Printf("Error looking up MX records: %v\n", err)
		return
	}
	fmt.Println("=== MX Records ===")
	for _, mx := range mxRecords {
		fmt.Printf("%s\tpreference: %d\n", mx.Host, mx.Pref)
	}
}

func lookupNS(domain string) {
	nsRecords, err := net.LookupNS(domain)
	if err != nil {
		fmt.Printf("Error looking up NS records: %v\n", err)
		return
	}
	fmt.Println("=== NS Records ===")
	for _, ns := range nsRecords {
		fmt.Printf("%s\n", ns.Host)
	}
}

func lookupTXT(domain string) {
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Printf("Error looking up TXT records: %v\n", err)
		return
	}
	fmt.Println("=== TXT Records ===")
	for _, txt := range txtRecords {
		fmt.Printf("%s\n", txt)
	}
}

func lookupZone(domain string) {
	fmt.Println("=== Zone Records ===")
	lookupA(domain)
	fmt.Println()
	lookupMX(domain)
	fmt.Println()
	lookupNS(domain)
	fmt.Println()
	lookupTXT(domain)
}

func printHelp() {
	fmt.Println(`=== SUPER DNS LOOKUP TOOL v1.0 ===
A fast and simple DNS record lookup tool.

USAGE:
    dns-lookup [OPTIONS] <domain>

OPTIONS:
    --help     Show this help message
    --a        Look up A records (IPv4 addresses)
    --mx       Look up MX records (mail servers)
    --ns       Look up NS records (name servers)
    --txt      Look up TXT records (text records)
    --zone     Look up all available records

EXAMPLES:
    dns-lookup --a google.com
    dns-lookup --mx --ns example.com
    dns-lookup --zone github.com
    
NOTE:
    - Multiple record types can be queried simultaneously
    - URLs are automatically cleaned (http:// and https:// are removed)
    - Use --zone to get all available records at once
    
Report issues: https://github.com/yourusername/dns-lookup`)
}

func main() {
	help := flag.Bool("help", false, "Show help message")
	a := flag.Bool("a", false, "Look up A records")
	mx := flag.Bool("mx", false, "Look up MX records")
	ns := flag.Bool("ns", false, "Look up NS records")
	txt := flag.Bool("txt", false, "Look up TXT records")
	zone := flag.Bool("zone", false, "Look up all records (zone)")

	flag.Parse()

	if *help {
		printHelp()
		os.Exit(0)
	}

	if len(flag.Args()) != 1 {
		printHelp()
		os.Exit(1)
	}

	domain := flag.Arg(0)
	if strings.HasPrefix(domain, "http://") || strings.HasPrefix(domain, "https://") {
		domain = strings.Split(strings.Split(domain, "://")[1], "/")[0]
	}

	if *zone {
		lookupZone(domain)
		os.Exit(0)
	}

	if *a {
		lookupA(domain)
		fmt.Println()
	}
	if *mx {
		lookupMX(domain)
		fmt.Println()
	}
	if *ns {
		lookupNS(domain)
		fmt.Println()
	}
	if *txt {
		lookupTXT(domain)
	}

	if !*a && !*mx && !*ns && !*txt && !*zone {
		fmt.Println("Please specify at least one record type to look up")
		fmt.Println("Use --help for usage information")
		os.Exit(1)
	}
}
