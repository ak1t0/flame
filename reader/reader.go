package reader

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type Parsed struct {
	Domain string `json:"domain"`
}

func ReadJson(name string) ([]string, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	// reshape given log data
	s := string(data[:])
	s = "[" + strings.Replace(s, "\n", ",\n", -1) + "]"
	s = strings.Replace(s, ",\n]", "\n]", 1)
	b := []byte(s)

	var logs []Parsed
	if err := json.Unmarshal(b, &logs); err != nil {
		return nil, err
	}

	var domains []string
	for _, log := range logs {
		domains = append(domains, log.Domain)
	}

	return domains, nil
}
