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
