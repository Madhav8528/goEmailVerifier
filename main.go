package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain hasDMX hasSPF SPFrecords hasDMARC DMARCrecords")

	for scanner.Scan() {
		emailVerifier(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error occured during taking input", err)
	}
}

func emailVerifier(domain string) {

	var hasDMX, hasSPF, hasDMARC bool
	var SPFrecords, DMARCrecords string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println("Error while looking dmx!!", err)
	}

	if len(mxRecords) > 0 {
		hasDMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Println("Error while looking TXT!!", err)
	}

	for _, dom := range txtRecords {
		if strings.HasPrefix(dom, "v=spf1") {
			SPFrecords = dom
			hasSPF = true
			break
		}
	}
	dmarcString := "_dmarc." + domain

	dmarcRecord, err := net.LookupTXT(dmarcString)
	if err != nil {
		fmt.Println("Error while looking for dmarc!!", err)
	}

	for _, record := range dmarcRecord {
		if strings.HasPrefix(record, "v=DMARC1") {
			DMARCrecords = record
			hasDMARC = true
			break
		}
	}

	fmt.Printf("For domain: %v\n", domain)
	fmt.Printf("MX is: %v\n", hasDMX)
	fmt.Printf("SPF is: %v\n", hasSPF)
	if hasSPF {
		fmt.Printf("SPF records: %v\n", SPFrecords)
	} else {
		fmt.Println("No SPF records found for this domain.")
	}
	fmt.Printf("DMARC is: %v\n", hasDMARC)
	if hasDMARC {
		fmt.Printf("DMARC records: %v\n", DMARCrecords)
	} else {
		fmt.Println("NO DMARC record found for this domain.")
	}

}
