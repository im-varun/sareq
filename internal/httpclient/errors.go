package httpclient

import (
	"errors"
	"net"
)

var (
	errRequestTimeout             = errors.New("request timed out")
	errRequestDNSResolutionFailed = errors.New("domain name could not be resolved")
)

// normalizeNetworkError standardizes known network-related errors into a more user-friendly
// error.
func normalizeNetworkError(err error) error {
	if isTimeoutError(err) {
		return errRequestTimeout
	}

	if isDNSResolutionError(err) {
		return errRequestDNSResolutionFailed
	}

	return err
}

// isTimeoutError checks if the given error is a network timeout error.
func isTimeoutError(err error) bool {
	netErr, ok := err.(net.Error)
	return ok && netErr.Timeout()
}

// isDNSResolutionError checks if the given error is a DNS resolution error.
func isDNSResolutionError(err error) bool {
	var dnsErr *net.DNSError
	return errors.As(err, &dnsErr) && dnsErr.IsNotFound
}
