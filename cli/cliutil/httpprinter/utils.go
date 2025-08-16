package httpprinter

import (
	"bytes"
	"encoding/json"
)

func prettifyResponseBody(respBody string) (string, error) {
	var prettyBodyBuf bytes.Buffer

	err := json.Indent(&prettyBodyBuf, []byte(respBody), "", "  ")
	if err != nil {
		return "", err
	}

	prettyBody := prettyBodyBuf.String()

	return prettyBody, nil
}
