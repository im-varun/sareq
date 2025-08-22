package httpclient

import "testing"

func TestValidateRequestURL(t *testing.T) {
	var tests = []struct {
		name           string
		reqURL         string
		expectedOutput string
		expectedError  string
	}{
		{
			name:           "empty request url with no spaces",
			reqURL:         "",
			expectedOutput: "",
			expectedError:  "request URL cannot be empty",
		},
		{
			name:           "empty request url with spaces",
			reqURL:         "    ",
			expectedOutput: "",
			expectedError:  "request URL cannot be empty",
		},
		{
			name:           "request url containing only scheme delimiter",
			reqURL:         "://",
			expectedOutput: "",
			expectedError:  "failed to parse request URL",
		},
		{
			name:           "request url with scheme delimiter and without scheme",
			reqURL:         "://example.com",
			expectedOutput: "",
			expectedError:  "failed to parse request URL",
		},
		{
			name:           "request url without scheme",
			reqURL:         "example.com",
			expectedOutput: defaultScheme + "://example.com",
			expectedError:  "",
		},
		{
			name:           "request url with valid http scheme",
			reqURL:         "http://example.com",
			expectedOutput: "http://example.com",
			expectedError:  "",
		},
		{
			name:           "request url with valid https scheme",
			reqURL:         "https://example.com",
			expectedOutput: "https://example.com",
			expectedError:  "",
		},
		{
			name:           "request url with invalid scheme",
			reqURL:         "hello://example.com",
			expectedOutput: "",
			expectedError:  "request URL contains a scheme that is invalid or not supported by the client",
		},
		{
			name:           "request url with invalid http like scheme",
			reqURL:         "htt://example.com",
			expectedOutput: "",
			expectedError:  "request URL contains a scheme that is invalid or not supported by the client",
		},
		{
			name:           "request url with invalid https like scheme",
			reqURL:         "htps://example.com",
			expectedOutput: "",
			expectedError:  "request URL contains a scheme that is invalid or not supported by the client",
		},
		{
			name:           "request url with valid http scheme and missing host",
			reqURL:         "http:///path/to/example/resource",
			expectedOutput: "",
			expectedError:  "request URL is missing a host",
		},
		{
			name:           "request url with valid https scheme and missing host",
			reqURL:         "https:///path/to/example/resource",
			expectedOutput: "",
			expectedError:  "request URL is missing a host",
		},
		{
			name:           "request url without valid scheme and missing host",
			reqURL:         "/path/to/example/resource",
			expectedOutput: "",
			expectedError:  "request URL is missing a host",
		},
		{
			name:           "request url with fragment",
			reqURL:         "example.com/path/to/resource#hash",
			expectedOutput: "",
			expectedError:  "request URL cannot contain a fragment",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := ValidateRequestURL(test.reqURL)
			if err != nil {
				if err.Error() != test.expectedError {
					t.Errorf("expected error: \"%s\", got error: \"%s\"", test.expectedError, err.Error())
				}
			}

			if output != test.expectedOutput {
				t.Errorf("expected output: \"%s\", got output: \"%s\"", test.expectedOutput, output)
			}
		})
	}
}

func TestValidateRequestBody(t *testing.T) {
	var tests = []struct {
		name          string
		reqBody       string
		expectedError string
	}{
		{
			name:          "invalid request body",
			reqBody:       `{"id": 1, "name": "john doe"}}`,
			expectedError: "request body is invalid",
		},
		{
			name:          "valid request body",
			reqBody:       `{"id": 1, "name": "john doe"}`,
			expectedError: "",
		},
		{
			name:          "invalid nested request body",
			reqBody:       `{"id": 1, "person": {"firstname": "john", "lastname": "doe"}`,
			expectedError: "request body is invalid",
		},
		{
			name:          "valid nested request body",
			reqBody:       `{"id": 1, "person": {"firstname": "john", "lastname": "doe"}}`,
			expectedError: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateRequestBody(test.reqBody)
			if err != nil {
				if err.Error() != test.expectedError {
					t.Errorf("expected error: \"%s\", got error: \"%s\"", test.expectedError, err.Error())
				}
			}
		})
	}
}
