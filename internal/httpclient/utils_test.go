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
			name:           "empty request url",
			reqURL:         "",
			expectedOutput: "",
			expectedError:  "request URL cannot be empty",
		},
		{
			name:           "empty request url with whitespaces",
			reqURL:         "    ",
			expectedOutput: "",
			expectedError:  "failed to parse request URL",
		},
		{
			name:           "empty request url with tab spaces",
			reqURL:         "				",
			expectedOutput: "",
			expectedError:  "failed to parse request URL",
		},
		{
			name:           "empty request url with whitespaces and tab spaces",
			reqURL:         "  			  		",
			expectedOutput: "",
			expectedError:  "failed to parse request URL",
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
