package httpclient

import (
	"encoding/json"
	"net/url"
	"strings"
)

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

func ValidateRequestBody(reqBody string) error {
	isValidJSON := json.Valid([]byte(reqBody))

	if !isValidJSON {
		return errRequestBodyJSONEncodingInvalid
	}

	return nil
}
