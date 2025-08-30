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
		return "", errRequestURLSchemeInvalid
	}

	if parsedURL.Host == "" {
		return "", errRequestURLHostMissing
	}

	if parsedURL.Fragment != "" {
		return "", errRequestURLFragmentPresent
	}

	validatedReqURL := parsedURL.String()

	return validatedReqURL, nil
}

func ValidateRequestBody(reqBody string) error {
	reqBodyIsValidJSON := json.Valid([]byte(reqBody))

	if !reqBodyIsValidJSON {
		return errRequestBodyJSONEncodingInvalid
	}

	return nil
}
