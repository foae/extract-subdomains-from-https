package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	domains := os.Args[1:]
	argsLen := len(domains)

	if argsLen < 1 {
		exitWith(fmt.Errorf("error: you need to pass at least 1 domains name"))
	}

	if argsLen > 1 {
		exitWith(fmt.Errorf("error: currently only 1 domains per invocation possible; sorry about this"))
	}

	target := domains[0]

	isDomainValid := regexp.MustCompile(`^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?`)
	if isDomainValid.MatchString(target) == false {
		exitWith(fmt.Errorf("domain is not valid: %v", target))
	}

	if strings.Contains(target, "http://") {
		target = strings.Replace(target, "http://", "https://", 1)
		logToFile("Replaced http:// with https://, now looking for " + target)
	} else if strings.Contains(target, "http://") == false {
		target = fmt.Sprintf("%v%v", "https://", target)
		logToFile("Added protocol https://, now looking for " + target)
	}

	resp, err := http.Get(target)
	if err != nil {
		exitWith(err)
	}
	if resp.TLS == nil {
		exitWith(fmt.Errorf("domain is not served over HTTPS: %v", target))
	}
	logToFile(fmt.Sprintf("Checking domain: %v", target))

	foundDomains := make([]string, 0, 120)
	for _, item := range resp.TLS.PeerCertificates {
		if len(item.DNSNames) > 0 {
			logToFile(fmt.Sprintf("Wrote %v domains to STDOUT. OK", len(item.DNSNames)))
			foundDomains = append(foundDomains, item.DNSNames...)
		}
	}

	json, _ := json.Marshal(foundDomains)
	fmt.Println(string(json))
}

func exitWith(err error) {
	logToFile(err.Error())
	os.Exit(1)
}

func logToFile(s string) {
	s = fmt.Sprintln(s)
	f, err := os.OpenFile("events.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		os.Exit(2)
	}
	defer f.Close()

	if _, err = f.Write([]byte(s)); err != nil {
		os.Exit(3)
	}
}
