package httpclient

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"
)

const DefaultTimeoutSeconds int = 10

const (
	schemeHTTP  string = "http"
	schemeHTTPS string = "https"

	defaultScheme string = schemeHTTPS
)

func ValidateRequestURL(reqURL string) (string, error) {
	if reqURL == "" || len(strings.TrimSpace(reqURL)) == 0 {
		return "", errors.New("request URL cannot be empty")
	}

	parsedURL, err := url.Parse(reqURL)
	if err != nil {
		return "", errors.New("failed to parse request URL")
	}

	if parsedURL.Scheme == "" {
		parsedURL.Scheme = defaultScheme

		parsedURL, err = url.Parse(parsedURL.String())
		if err != nil {
			return "", errors.New("failed to parse request URL")
		}
	}

	if parsedURL.Scheme != schemeHTTP && parsedURL.Scheme != schemeHTTPS {
		return "", errors.New("request URL contains a scheme that is invalid or not supported by the client")
	}

	if parsedURL.Host == "" {
		return "", errors.New("request URL is missing a host")
	}

	if parsedURL.Fragment != "" {
		return "", errors.New("request URL cannot contain a fragment")
	}

	validReqURL := parsedURL.String()

	return validReqURL, nil
}

func ValidateRequestBody(reqBody string) error {
	reqBodyIsValid := json.Valid([]byte(reqBody))

	if !reqBodyIsValid {
		return errors.New("request body is invalid")
	}

	return nil
}
