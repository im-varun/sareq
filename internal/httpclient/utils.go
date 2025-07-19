package httpclient

import (
	"errors"
	"net/url"
	"strings"
)

func ValidateRequestURL(reqURL string) (string, error) {
	if reqURL == "" {
		return "", errors.New("request URL cannot be empty")
	}

	if isRootRelativeURL(reqURL) {
		return "", errors.New("request URL cannot be root relative")
	}

	parsedURL, err := url.Parse(reqURL)
	if err != nil {
		return "", errors.New("failed to parse request URL")
	}

	if parsedURL.Scheme == "" {
		parsedURL.Scheme = DefaultScheme

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

	parsedReqURL, err := url.ParseRequestURI(parsedURL.String())
	if err != nil {
		return "", errors.New("failed to parse request URL")
	}

	validReqURL := parsedReqURL.String()

	return validReqURL, nil
}

func isRootRelativeURL(url string) bool {
	return strings.HasPrefix(url, "/")
}

func isInvalidScheme(scheme string) bool {
	return scheme != HTTP && scheme != HTTPS
}
