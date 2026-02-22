package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestAuth(t *testing.T) {
	tests := map[string]struct {
		key       string
		value     string
		expect    string
		expectErr string
	}{
		"Valid Key":        {key: "Authorization", value: "ApiKey 12345", expect: "12345"},
		"malformed":        {key: "Authorization", value: "-", expectErr: "malformed authorization header"},
		"malformed bearer": {key: "Authorization", value: "Bearer xxxxxxx", expectErr: "malformed authorization header"},
		"no header":        {expectErr: "no authorization header included"},
		"header no value":  {key: "Authorization", expectErr: "no authorization header included"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			header := http.Header{}
			header.Add(tc.key, tc.value)
			out, err := GetAPIKey(header)
			fmt.Println(out, err)

			if err != nil {
				if tc.expectErr != "" && strings.Contains(err.Error(), tc.expectErr) {
					return
				}
				t.Errorf("unexpected error: %v", err)
				return
			}

			if out != tc.expect {
				t.Errorf("Unexpected: TestGetAPIKey:%s", out)
				return
			}
		})

	}
}
