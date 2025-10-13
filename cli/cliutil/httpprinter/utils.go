package httpprinter

import (
	"bytes"
	"encoding/json"
)

// prettifyResponseBody prettifies a JSON response body.
func prettifyResponseBody(respBody string) (string, error) {
	var prettyBodyBuf bytes.Buffer

	err := json.Indent(&prettyBodyBuf, []byte(respBody), "", "  ")
	if err != nil {
		return "", err
	}

	prettyBody := prettyBodyBuf.String()

	return prettyBody, nil
}
