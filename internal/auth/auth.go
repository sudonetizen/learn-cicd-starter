package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrMalformedAuthzHeader = errors.New("malformed authorization header")
var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

// GetAPIKey -
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", ErrMalformedAuthzHeader
	}

	return splitAuth[1], nil
}
