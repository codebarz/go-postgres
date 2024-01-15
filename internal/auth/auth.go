package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(header http.Header) (string, error) {
	value := header.Get("Authorization")

	if value == "" {
		return "", errors.New("no authorization header found")
	}

	values := strings.Split(value, " ")

	if len(values) != 2 {
		return "", errors.New("malformed authorization header")
	}

	if values[0] != "ApiKey" {
		return "", errors.New("malformed first part of authorization header")
	}

	return values[1], nil
}
