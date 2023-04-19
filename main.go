package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Clean(data string) *[]string {
	var result []string
	var ports []string
	ipRegex := regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`)
	portsRegex := regexp.MustCompile(`\b\d{1,5}/open`)

	ip := ipRegex.FindString(data)
	portsRaw := portsRegex.FindAllString(data, -1)

	for _, portRaw := range portsRaw {
		port := strings.Split(portRaw, "/")[0]
		ports = append(ports, port)
	}

	for _, port := range ports {
		result = append(result, fmt.Sprintf("%s:%s\n", ip, port))
	}

	return &result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Text()
		list := Clean(data)

		for _, ip := range *list {
			fmt.Print(ip)
		}
	}
}
