package crawler

import (
	"testing"
)

func TestScanAvailable(t *testing.T) {
	target := "facebookcorewwwi.onion"
	r := scan(target)
	if r != "available" {
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
