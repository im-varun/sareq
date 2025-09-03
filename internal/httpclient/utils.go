package httpclient

import (
	"encoding/json"
	"net/url"
	"strings"
)

// ValidateRequestURL validates that the specified request URL is in a proper format.
func ValidateRequestURL(reqURL string) error {
	if reqURL == "" || len(strings.TrimSpace(reqURL)) == 0 {
		return errRequestURLEmpty
	}

	parsedURL, err := url.Parse(reqURL)
	if err != nil {
		return errRequestURLParsingFailed
	}

	if parsedURL.Scheme == "" {
		return errRequestURLSchemeMissing
	}

	if parsedURL.Scheme != schemeHTTP && parsedURL.Scheme != schemeHTTPS {
		return errRequestURLSchemeInvalid
	}

	if parsedURL.Host == "" {
		return errRequestURLHostMissing
	}

	if parsedURL.Fragment != "" {
		return errRequestURLFragmentPresent
	}

	return nil
}

// ValidateRequestBody validates that the request body is a valid JSON encoding.
func ValidateRequestBody(reqBody string) error {
	isValidJSON := json.Valid([]byte(reqBody))

	if !isValidJSON {
		return errRequestBodyJSONEncodingInvalid
	}

	return nil
}
