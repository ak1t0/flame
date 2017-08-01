package reader

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

type Parsed struct {
	Domain string `json:"domain"`
}

func ReadJson(name string) []string {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	
	// reshape given log data
	s := string(data[:])
	s = "[" + strings.Replace(s, "\n", ",\n", -1) + "]"
	s = strings.Replace(s, ",\n]", "\n]", 1)
	b := []byte(s)

	var logs []Parsed
	if err := json.Unmarshal(b, &logs); err != nil {
		log.Fatal(err)
	}

	var domains []string
	for _, log := range logs {
		domains = append(domains, log.Domain)
	}
	
	return domains
}
