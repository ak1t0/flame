package crawler

import (
	"testing"
)

func TestScanAvailable(t *testing.T) {
	target := "facebookcorewwwi.onion"
	r := scan(target)
	if r == "connection refused" {
		t.Errorf("TestScanAvailable")
	}
	if r == "connection forbidden" {
		t.Errorf("TestScanAvailable")
	}
	if r == "general failure" {
		t.Errorf("TestScanAvailable")
	}
	if r == "TTL expierd" {
		t.Errorf("TestScanAvailable")
	}
}

func TestScanTimeout(t *testing.T) {
	target := "aaaaaaaaaaaaaaaa.onion"
	r := scan(target)
	if r != "timeout" {
		t.Errorf("TestScanTimuout")
	}
}
