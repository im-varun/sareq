package httpprinter

import (
	"bytes"
	"encoding/json"
)

func prettifyResponseBody(respBody string) (string, error) {
	var prettyBody bytes.Buffer

	err := json.Indent(&prettyBody, []byte(respBody), "", "  ")
	if err != nil {
		return "", err
	}

	return prettyBody.String(), nil
}
