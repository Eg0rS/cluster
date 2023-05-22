package target

import (
	"auth-gateway/config"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func NewHTTPRequestForwarder(settings *config.Settings) *HTTPRequestForwarder {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	client := http.Client{
		Timeout:   time.Duration(settings.RequestForwarderTimeout) * time.Second,
		Transport: t,
	}

	return &HTTPRequestForwarder{
		client: client,
	}
}

type HTTPRequestForwarder struct {
	client http.Client
}

func (c *HTTPRequestForwarder) Forward(r *http.Request, payload string) (*http.Response, error) {
	path := getServiceURL(r)

	pathStr, err := url.QueryUnescape(path.String())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	forwardingRequest, _ := http.NewRequestWithContext(r.Context(), r.Method, pathStr, r.Body)
	for headerName, headerValues := range r.Header {
		for _, headerValue := range headerValues {
			forwardingRequest.Header.Add(headerName, headerValue)
		}
	}

	forwardingRequest.Header.Add("UserContext", payload)
	res, err := c.client.Do(forwardingRequest)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getServiceURL(r *http.Request) url.URL {
	path := url.URL{
		Scheme:   "http",
		Host:     strings.TrimPrefix(r.URL.Path, "/"),
		RawQuery: r.URL.RawQuery,
	}
	return path
}
