package httpclient

import "errors"

var errRequestURLEmpty error = errors.New("request URL cannot be empty")

var errRequestURLParsingFailed error = errors.New("failed to parse request URL")

var errRequestURLInvalidScheme error = errors.New("request URL contains a scheme that is invalid or not supported by the client")

var errRequestURLMissingHost error = errors.New("request URL is missing a host")

var errRequestURLContainsFragment error = errors.New("request URL cannot contain a fragment")

var errRequestBodyInvalid error = errors.New("request body is invalid")
