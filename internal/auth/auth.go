package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(h http.Header) (string, error) {
	val := h.Get("Authorization")

	if val == "" {
		return "", errors.New("error: No Auth info found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("error: Malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("error: Malformed auth header")
	}

	return vals[1], nil

}
