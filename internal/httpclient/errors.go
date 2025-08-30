package httpclient

import (
	"errors"
	"net"
)

var (
	errRequestURLEmpty           = errors.New("is empty")
	errRequestURLParsingFailed   = errors.New("parsing failed")
	errRequestURLSchemeMissing   = errors.New("is missing a scheme")
	errRequestURLSchemeInvalid   = errors.New("contains a scheme that is invalid or not supported by the client")
	errRequestURLHostMissing     = errors.New("is missing a host")
	errRequestURLFragmentPresent = errors.New("contains a fragment")

	errRequestBodyJSONEncodingInvalid = errors.New("not a valid JSON encoding")

	errRequestTimeout             = errors.New("request timed out")
	errRequestDNSResolutionFailed = errors.New("domain name could not be resolved")
)

func normalizeHTTPError(err error) error {
	if isTimeoutError(err) {
		return errRequestTimeout
	}

	if isDNSResolutionError(err) {
		return errRequestDNSResolutionFailed
	}

	return err
}

func isTimeoutError(err error) bool {
	netErr, ok := err.(net.Error)
	return ok && netErr.Timeout()
}

func isDNSResolutionError(err error) bool {
	var dnsErr *net.DNSError
	return errors.As(err, &dnsErr) && dnsErr.IsNotFound
}
