package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)
func main() {
	
	p := NewParser()
	// crete a new scanner that uses standard input
	in := bufio.NewScanner(os.Stdin)

	// scan the log file line by line by calling the scan method
	for in.Scan() {
		p.lines++

		parsed, err := parse(p, in.Text())
		if err != nil {
			fmt.Println(err)
			return
		}

		domain, visits := parsed.domain, parsed.visits

		// Collect the unique domains
		if _, ok := p.sum[domain]; !ok {
			p.domains = append(p.domains, domain)
		}

		// Keep track of total and per domain visits
		p.total += visits
		p.sum[domain] = result{
			domain: domain,
			visits: visits + p.sum[domain].visits,
		}
	}

	// Print the visits
	sort.Strings(p.domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Printf(strings.Repeat("-", 45))

	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}
	fmt.Printf("%-30s %10d\n", "TOTAL", p.total)

	if err := in.Err(); err != nil{
		fmt.Println("> Err:", err)
	}
}