package main

import (
	"fmt"
	"strconv"
	"strings"
)

// result stores the parsed value for a domain
type result struct {
	domain string
	visits int
}


// parser keeps track of the pasing
type parser struct {
	sum     map[string]result // total visits per domains
	domains []string          //unique domain names
	total   int               // total visit to all domains
	lines   int               // number of parsed files
	lerr    error
}

// newParser constructs, initializes and returns a new parser
func newParser() *parser {
	return &parser{sum: make(map[string]result)}
}


// parse parses a log line and returns the parsed result with an error
func parse(p *parser, line string) (parsed result) {
	if p.lerr != nil {
		return
	}
	p.lines++

	// parse the files
	fields := strings.Fields(line)
	if len(fields) != 2 {
		p.lerr = fmt.Errorf("wrong input:  %v (%d)", fields, p.lines)
		return
	}

	var err error

	parsed.domain = fields[0]
	// Sum the total visits per domain
	parsed.visits, err = strconv.Atoi(fields[1])

	if parsed.visits < 0 || err != nil {
		p.lerr = fmt.Errorf("wrong input: %v", fields[1])

	}

	return
}

// update updates the parser for the given parsing result 
func update(p *parser, parsed result) {
	if p.lerr != nil {
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

func err (p *parser) error {
	return p.lerr
}
