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

import(
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
	subdomainsFile	*os.File
}

func (s* Scanner) Init(subdomainsFile string) {
	s.subdomainsFile = readSubdomainsFile(subdomainsFile)
}

