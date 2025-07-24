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
