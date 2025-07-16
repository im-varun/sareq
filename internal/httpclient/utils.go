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

	if strings.HasPrefix(reqURL, "/") {
		return "", errors.New("request URL must be absolute for an HTTP request")
	}

	parsedURL, err := url.Parse(reqURL)
	if err != nil {
		return "", errors.New("failed to parse request URL")
	}

	if parsedURL.Scheme == "" {
		parsedURL.Scheme = DefaultScheme
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return "", errors.New("scheme in request URL is not supported by the client")
	}

	if !strings.HasPrefix(parsedURL.String(), "http://") && !strings.HasPrefix(parsedURL.String(), "https://") {
		return "", errors.New("request URL does not contain a valid scheme delimiter")
	}

	parsedReqURL, err := url.ParseRequestURI(parsedURL.String())
	if err != nil {
		return "", errors.New("failed to parse request URL")
	}

	return parsedReqURL.String(), nil
}
