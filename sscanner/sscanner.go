
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

type Scanner struct {
	domain			string
	subdomainsFile	*os.File
	resolver		string
	outfilePath		string
}

func readSubdomainsFile(fp string) *os.File {
	f, err := os.Open(fp)
	if err != nil {
		log.Panic("Subdomainsfile could not be opened")
	}
	return f
}

func (s *Scanner) Init(
	domain string,
	subdomainsFilePath string,
	resolver string,
	outfilePath string,
) {
	log.Println("Initializing scanner")
	s.domain = domain
	s.subdomainsFile = readSubdomainsFile(subdomainsFilePath)
	s.resolver = resolver
	s.outfilePath = outfilePath
}

func lookup(
	domain string,
	wg *sync.WaitGroup,
	outChan chan string,
	done chan bool,
) {
	defer wg.Done()
	ips, err := net.LookupHost(domain)
	if err == nil {
		for _, ip := range ips {
			log.Printf("Found %s at %s", domain, ip)
			outChan <- fmt.Sprintf("%s,%s", domain, ip)
		}
	}
}

func writeResultsToFile(
	outfilePath string,
	results chan string,
	done chan bool,
) {
	f, err := os.Create(outfilePath)
	defer f.Close()
	if err != nil {
		log.Panic("Could not open outfile!")
	}
	for res := range results {
		_, err = fmt.Fprintln(f, res)
		if err != nil {
			log.Print(err)
		}
	}
	done <- true
}

func (s *Scanner) Scan() {
	var wg sync.WaitGroup
	outChan := make(chan string)
	done := make(chan bool)
	scanner := bufio.NewScanner(s.subdomainsFile)
	log.Println("Scan running...")
	for scanner.Scan() {
		subdomain := fmt.Sprintf("%s.%s", scanner.Text(), s.domain)
		go lookup(subdomain, &wg, outChan, done)
		wg.Add(1)
	}
	s.subdomainsFile.Close()
	go writeResultsToFile(s.outfilePath, outChan, done)
	go func() {
		wg.Wait()
		close(outChan)
	}()
	d := <- done
	if d == true {
		log.Println("Done!")
	}
}
