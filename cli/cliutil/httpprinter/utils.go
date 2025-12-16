package httpprinter

import (
	"bytes"
	"encoding/json"
	"strings"
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

// isTextualMediaType checks if the given media type is textual (non-binary).
func isTextualMediaType(mediaType string) bool {
	if strings.HasPrefix(mediaType, "text/") {
		return true
	}

	switch mediaType {
	case "application/json",
		"application/xml",
		"application/javascript",
		"application/xhtml+xml",
		"application/x-www-form-urlencoded":
		return true
	}

	return false
}
