package httpclient

import (
	"encoding/json"
	"net/url"
	"strings"
)

func ValidateRequestURL(reqURL string) (string, error) {
	if reqURL == "" || len(strings.TrimSpace(reqURL)) == 0 {
		return "", errRequestURLEmpty
	}

	parsedURL, err := url.Parse(reqURL)
	if err != nil {
		return "", errRequestURLParsingFailed
	}

	if parsedURL.Scheme == "" {
		parsedURL.Scheme = defaultScheme

		parsedURL, err = url.Parse(parsedURL.String())
		if err != nil {
			return "", errRequestURLParsingFailed
		}
	}

	if parsedURL.Scheme != schemeHTTP && parsedURL.Scheme != schemeHTTPS {
		return "", errRequestURLInvalidScheme
	}

	if parsedURL.Host == "" {
		return "", errRequestURLMissingHost
	}

	if parsedURL.Fragment != "" {
		return "", errRequestURLContainsFragment
	}

	validReqURL := parsedURL.String()

	return validReqURL, nil
}

func ValidateRequestBody(reqBody string) error {
	reqBodyIsValid := json.Valid([]byte(reqBody))

	if !reqBodyIsValid {
		return errRequestBodyInvalid
	}

	return nil
}
