package main

import(
	"flag"

	"github.com/annoj/sscan/sscanner"
)

type args struct {
	domain				string
	subdomainsFilePath	string
	resolver			string
	outfilePath			string
	retry 				int
}

func (a* args) parse() {
	flag.StringVar(&a.domain, "d", "", "The name of the domain to scan. Can already contain a subdomain.")
	flag.StringVar(&a.subdomainsFilePath, "l", "", "Subdomains listextfile containing a list of possible subdomains, delimited by newlines.")
	flag.StringVar(&a.resolver, "r", "8.8.8.8", "IP address of the DNS resolver to use.")
	flag.StringVar(&a.outfilePath, "o", "", "Write results to this file. Will be overwritten if already exists.")
	flag.IntVar(&a.retry, "rt", 3, "If lookup fails, retry this many times.")
	flag.Parse()
}

func main() {

	args := new(args)
	args.parse()

	scanner := new(sscanner.Scanner)
	scanner.Init(args.domain, args.subdomainsFilePath, args.resolver, args.outfilePath, args.retry)

	scanner.Scan()
}
