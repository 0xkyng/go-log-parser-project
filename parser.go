package main

import (
	"fmt"
	"strconv"
	"strings"
)

type result struct {
	domain string
	visits int
}

type parser struct {
	sum     map[string]result // total visits per domains
	domains []string          //unique domain names
	total   int               // total visit to all domains
	lines   int               // number of parsed files
}

func NewParser() parser {
	return parser{sum: make(map[string]result)}
}

func parse(p parser, line string) (parsed result, err error) {

	// parse the files
	fields := strings.Fields(line)
	if len(fields) != 2 {
		err = fmt.Errorf("wrong input:  %v (%d)", fields, p.lines)
		return 
	}

	parsed.domain = fields[0]

	// Sum the total visits per domain
	parsed.visits, err = strconv.Atoi(fields[1])
	if parsed.visits < 0 || err != nil {
		err = fmt.Errorf("wrong input: %v", fields[1])
		return 
	}

	return 
}
