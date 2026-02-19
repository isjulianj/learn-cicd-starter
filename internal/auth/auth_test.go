package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAuth(t *testing.T) {
	got, _ := GetAPIKey(http.Header{
		"origin":        []string{"google.com"},
		"Authorization": []string{"ApiKey mysecretjey"},
	})
	want := "mysecretjey"

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected %v, got: %v", want, got)
	}
}
