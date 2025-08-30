package httpclient

import (
	"errors"
	"testing"
)

func TestValidateRequestURL(t *testing.T) {
	var tests = []struct {
		name          string
		reqURL        string
		expectedError error
	}{
		{
			name:          "empty request url with no spaces",
			reqURL:        "",
			expectedError: errRequestURLEmpty,
		},
		{
			name:          "empty request url with spaces",
			reqURL:        "    ",
			expectedError: errRequestURLEmpty,
		},
		{
			name:          "request url containing only scheme delimiter",
			reqURL:        "://",
			expectedError: errRequestURLParsingFailed,
		},
		{
			name:          "request url with scheme delimiter and without scheme",
			reqURL:        "://example.com",
			expectedError: errRequestURLParsingFailed,
		},
		{
			name:          "request url without scheme",
			reqURL:        "example.com",
			expectedError: errRequestURLSchemeMissing,
		},
		{
			name:          "request url with valid http scheme",
			reqURL:        "http://example.com",
			expectedError: nil,
		},
		{
			name:          "request url with valid https scheme",
			reqURL:        "https://example.com",
			expectedError: nil,
		},
		{
			name:          "request url with invalid scheme",
			reqURL:        "hello://example.com",
			expectedError: errRequestURLSchemeInvalid,
		},
		{
			name:          "request url with invalid http like scheme",
			reqURL:        "htt://example.com",
			expectedError: errRequestURLSchemeInvalid,
		},
		{
			name:          "request url with invalid https like scheme",
			reqURL:        "htps://example.com",
			expectedError: errRequestURLSchemeInvalid,
		},
		{
			name:          "request url with valid http scheme and missing host",
			reqURL:        "http:///path/to/example/resource",
			expectedError: errRequestURLHostMissing,
		},
		{
			name:          "request url with valid https scheme and missing host",
			reqURL:        "https:///path/to/example/resource",
			expectedError: errRequestURLHostMissing,
		},
		{
			name:          "request url without valid scheme and missing host",
			reqURL:        "/path/to/example/resource",
			expectedError: errRequestURLSchemeMissing,
		},
		{
			name:          "request url with valid http scheme and fragment",
			reqURL:        "http://example.com/path/to/resource#hash",
			expectedError: errRequestURLFragmentPresent,
		},
		{
			name:          "request url with valid https scheme and fragment",
			reqURL:        "https://example.com/path/to/resource#hash",
			expectedError: errRequestURLFragmentPresent,
		},
		{
			name:          "request url with localhost",
			reqURL:        "http://localhost/hello",
			expectedError: nil,
		},
		{
			name:          "request url with localhost and port",
			reqURL:        "http://localhost:8000/hello",
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateRequestURL(test.reqURL)
			if err == nil {
				err = errors.New("")
			}

			testErr := test.expectedError
			if testErr == nil {
				testErr = errors.New("")
			}

			if err.Error() != testErr.Error() {
				t.Errorf("expected error: \"%s\", got error: \"%s\"", test.expectedError, err.Error())
			}
		})
	}
}

func TestValidateRequestBody(t *testing.T) {
	var tests = []struct {
		name          string
		reqBody       string
		expectedError error
	}{
		{
			name:          "invalid request body",
			reqBody:       `{"id": 1, "name": "john doe"}}`,
			expectedError: errRequestBodyJSONEncodingInvalid,
		},
		{
			name:          "valid request body",
			reqBody:       `{"id": 1, "name": "john doe"}`,
			expectedError: nil,
		},
		{
			name:          "invalid nested request body",
			reqBody:       `{"id": 1, "person": {"firstname": "john", "lastname": "doe"}`,
			expectedError: errRequestBodyJSONEncodingInvalid,
		},
		{
			name:          "valid nested request body",
			reqBody:       `{"id": 1, "person": {"firstname": "john", "lastname": "doe"}}`,
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateRequestBody(test.reqBody)
			if err == nil {
				err = errors.New("")
			}

			testErr := test.expectedError
			if testErr == nil {
				testErr = errors.New("")
			}

			if err.Error() != testErr.Error() {
				t.Errorf("expected error: \"%s\", got error: \"%s\"", test.expectedError, err.Error())
			}
		})
	}
}
