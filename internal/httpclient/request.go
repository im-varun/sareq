package httpclient

type RequestConfig struct {
	method string
	url    string
	body   string
}

func NewRequestConfig(method string, url string, body string) *RequestConfig {
	return &RequestConfig{
		method: method,
		url:    url,
		body:   body,
	}
}
