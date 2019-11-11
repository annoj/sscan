//
// sscann subdomain scanner
// ========================
//
// Scans a given domain for subdomains using DNS service.
// Version 0.0.1
//
// Auhor: Jona Heitzer
//

package sscanner

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func readSubdomainsFile(fp string) *os.File {
	f, err := os.Open(fp)
	if err != nil {
		panic("Subdomainsfile could not be opened")
	}
	return f
}

type Scanner struct {
	domain         string
	subdomainsFile *os.File
	resolver       string
}

func (s *Scanner) Init(
	domain string,
	subdomainsFilePath string,
	resolver string,
) {
	s.domain = domain
	s.subdomainsFile = readSubdomainsFile(subdomainsFilePath)
	s.resolver = resolver
}

func (s *Scanner) Scan() {
	scanner := bufio.NewScanner(s.subdomainsFile)
	for scanner.Scan() {
		subdomain := fmt.Sprintf("%s.%s", scanner.Text(), s.domain)
		ips, err := net.LookupHost(subdomain)
		if err == nil {
			for _, ip := range ips {
				log.Printf("Found %s at %s", subdomain, ip)
			}
		}
	}
}
