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
			name:           "Empty Request URL",
			reqURL:         "",
			expectedOutput: "",
			expectedError:  "request URL cannot be empty",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			answer, err := ValidateRequestURL(test.reqURL)
			if err != nil {
				if err.Error() != test.expectedError {
					t.Errorf("expected error: \"%s\", got error: \"%s\"", test.expectedError, err.Error())
				}
			}

			if answer != test.expectedOutput {
				t.Errorf("expected output: \"%s\", got output: \"%s\"", test.expectedOutput, answer)
			}
		})
	}
}
