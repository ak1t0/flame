package reader

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type Parsed struct {
	Domain string `json:"domain"`
}

func ReadData(name string) ([]string, error) {
	if strings.Contains(name, ".json") {
		return ReadJson(name)
	}
	if strings.Contains(name, ".txt") {
		return ReadText(name)
	} else {
		var s []string
		return s, nil
	}
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

func ReadText(name string) ([]string, error) {
	fp, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var domains []string
	for scanner.Scan() {
		domains = append(domains, scanner.Text())
	}

	return domains, nil
}
