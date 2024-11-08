package auth

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {

	headerCorrect := make(http.Header)
	headerCorrect.Add("Authorization", "ApiKey 123abc")

	headerMissing := make(http.Header)
	headerMissing.Add("as", "ds")

	headerMalformed := make(http.Header)
	headerMalformed.Add("Authorization", "ApiLock 321321321")

	tests := map[string]struct {
		key    string
		err    error
		header http.Header
	}{
		"passing":        {key: "123abc", err: nil, header: headerCorrect},
		"missing header": {key: "123abc", err: errors.New("no authorization header included"), header: headerMissing},
		"malformed":      {key: "123abc", err: errors.New("malformed authorization header"), header: headerMalformed},
	}

	fmt.Println(tests)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			got, err := GetAPIKey(tc.header)

			if (tc.err == nil) != (err == nil) {
				t.Fatalf("expected error: %v, got: %v", tc.err, err)
			}

			if tc.err != nil {
				if tc.err.Error() != err.Error() {
					t.Fatalf("expected: %#v, got: %#v", tc.err, err)
				}
			}

			if tc.err == nil {
				if tc.key != got {
					t.Fatalf("expected key: %v, got: %v", tc.key, got)
				}
			}
		})
	}
}
