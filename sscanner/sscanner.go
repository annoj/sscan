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
	"sync"
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

func lookup(domain string, wg *sync.WaitGroup) {
	defer wg.Done()
	ips, err := net.LookupHost(domain)
	if err == nil {
		for _, ip := range ips {
			log.Printf("Found %s at %s", domain, ip)
		}
	}
}

func (s *Scanner) Scan() {
	var wg sync.WaitGroup
	scanner := bufio.NewScanner(s.subdomainsFile)
	for scanner.Scan() {
		subdomain := fmt.Sprintf("%s.%s", scanner.Text(), s.domain)
		go lookup(subdomain, &wg)
		wg.Add(1)
	}
	log.Printf("All subdomains parsed.")
	wg.Wait()
}
