package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type result struct {
	domain string
	visits int
}

type parser struct {
	sum map[string]result // total visits per domains
	domains []string //unique domain names
	total int       // total visit to all domains
	lines int      // number of parsed files
}

func main() {
	
	p := parser {
		sum : make(map[string]result),
	}

	// crete a new scanner that uses standard input
	in := bufio.NewScanner(os.Stdin)

	// scan the log file line by line by calling the scan method
	for in.Scan() {
		p.lines++

		// parse the files
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("wrong input:  %v (%d)\n", fields, p.lines)
			return
		}

		domain := fields[0]

		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %v\n", fields[1])
			return
		}

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