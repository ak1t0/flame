package reader

import (
	"reflect"
	"testing"
)

func TestReadJson(t *testing.T) {
	expected := []string{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.onion", "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb.onion"}
	target := "test.json"
	actual, _ := ReadJson(target)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestReadJson")
	}
}
