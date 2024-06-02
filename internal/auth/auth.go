package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Parses out the API Key from the HTTP header
// Example: `AUTHORIZATION` : {API key}
func GetAPIKey(header http.Header) (string, error) {
	if val := header.Get("AUTHORIZATION"); val == "" {
		return "", errors.New("no Authorization header present")
	} else if vals := strings.Split(val, " "); len(vals) != 2 {
		return "", errors.New("malformed Authorization header")
	} else if vals[0] != "ApiKey" {
		return "", errors.New("malformed Authorization header")
	} else {
		return vals[1], nil
	}
}
