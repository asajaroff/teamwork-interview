package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/mail"
	"os"
	"strings"
)

type Customer struct {
	fistName  string
	lastName  string
	email     string
	gender    string
	ipAddress string
}

type DomainList struct {
	domain string
	count  int
}

func main() {
	filename := os.Getenv("FILENAME")
	log.Println("Starting main program")
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	log.Printf("Importing from file: %s", filename)

	var customerSlice []Customer
	var domainCount []DomainList
	var currentCustomerLine Customer

	csvReader := csv.NewReader(f)
	emailCount := 0
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			log.Println("Reached EOF")
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if validateCustomerLine(rec) != nil {
			currentCustomerLine.fistName, currentCustomerLine.lastName, currentCustomerLine.email, currentCustomerLine.gender, currentCustomerLine.ipAddress = rec[0], rec[1], rec[2], rec[3], rec[4]
			customerSlice = append(customerSlice, currentCustomerLine)
			email := stripDomain(currentCustomerLine.email)
			if email == "" {
				emailCount++
			}
			domainCount = addEmailCount(email, domainCount)
		} else {
			continue
		}
	}
	log.Printf("Filtered through %d records", emailCount)
	sortedList := sortList(domainCount)
	fmt.Println(sortedList)
}

// Returns error if Customer is not valid -invalid email or IP address-
func validateCustomerLine(line []string) error {
	_, err := mail.ParseAddress(line[2])
	if err != nil {
		log.Printf("Error importing the record %s", line)
		return nil
	}
	if net.ParseIP(line[4]) == nil {
		log.Printf("ERR: \"%s\" is not a valid IP address, at position. Skipping", line[4])
		return errors.New("Invalid IP Address")
	}
	return errors.New("Invalid email")
}

// Returns a string with the domain if valid, else returns error
func stripDomain(addr string) string {
	domain := strings.Split(addr, "@")
	return domain[1]
}

// Check for domain in array and increase counter or add it
// TODO Improve performance by passing pointer instead of array
func addEmailCount(addr string, list []DomainList) []DomainList {
	var found bool
	for i := range list {
		if addr == list[i].domain {
			list[i].count++
			found = true
			// log.Printf("Found domain %s, increasing counter", addr)
			break
		}
	}
	if !found {
		c := DomainList{domain: addr, count: 1}
		list = append(list, c)
		// log.Printf("Domain %s was not in the array, creating new entry.", addr)
	}
	return list
}

// Returns a sorted list of domains
// Performance on customer
//real	0m1.344s
//user	0m0.536s
//sys	0m0.511s
func sortList(list []DomainList) []DomainList {
	var swap DomainList
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list); j++ {
			if list[i].count > list[j].count {
				continue
			} else {
				swap = list[i]
				list[i] = list[j]
				list[j] = swap
				// log.Printf("Swapped %s since it has %d which is larger that", list[i].domain, list[i].count)
			}
		}
	}
	return list
}
