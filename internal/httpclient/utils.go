package httpclient

import (
	"errors"
	"net/url"
)

func ValidateRequestURL(reqURL string) (string, error) {
	if reqURL == "" {
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

	if isInvalidScheme(parsedURL.Scheme) {
		return "", errors.New("scheme in request URL is invalid or not supported by the client")
	}

	if parsedURL.Host == "" {
		return "", errors.New("request URL is missing a host")
	}

	validReqURL := parsedURL.String()

	return validReqURL, nil
}

func isInvalidScheme(scheme string) bool {
	return scheme != schemeHTTP && scheme != schemeHTTPS
}
