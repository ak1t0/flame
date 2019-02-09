package crawler

import (
	"fmt"
	"github.com/ak1t0/flame/format"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func Crawl(logs format.OnionLogs) format.OnionLogs {
	var scanned format.OnionLogs

	// data channel
	c := make(chan format.OnionLog, 6000)
	// semaphore
	sem := 10
	semc := make(chan bool, sem)

	// crawl concurrently
	for _, v := range logs {
		semc <- true
		go func(v format.OnionLog) {
			defer func() { <-semc }()
			v.Result = scan(v.Address)
			c <- v
		}(v)
	}

	// crawled data
	for i := 0; i < len(logs); i++ {
		data := <-c
		scanned = append(scanned, data)
	}

	return scanned
}

func scan(target string) string {
	hostPort := "127.0.0.1:9050"
	p, err := proxy.SOCKS5("tcp", hostPort, nil, proxy.Direct)
	if err != nil {
		fmt.Println(err)
	}

	client := http.DefaultClient
	client.Transport = &http.Transport{
		Dial: p.Dial,
	}

	// connection and timeout
	c := make(chan string, 1)
	limit := 30

	go func() {
		resp, err := client.Get("http://" + target)

		if err != nil {
			log.Println(err)

			e := err.Error()
			if strings.Contains(e, "connection refused") {
				c <- "connection refused"
			} else if strings.Contains(e, "connection forbidden") {
				c <- "connection forbidden"
			} else if strings.Contains(e, "host unreachable") {
				c <- "general failure"
			} else if strings.Contains(e, "TTL expired") {
				c <- "TTL expired"
			} else {
				c <- e //"Unknown Error"
			}

			return
		}

		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c <- "available"
		}
		c <- string(contents)
		return
	}()

	select {
	case res := <-c:
		return res
	// timeout
	case <-time.After(time.Second * time.Duration(limit)):
		return "timeout"
	}
}
