package utils

import (
	"io"
	"net/http"
)

type HTTPClient struct {
	client *http.Client
}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{},
	}
}

func (hc *HTTPClient) Get(url string) (*http.Response, error) {
	// create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// send the request and return the response
	return hc.client.Do(req)
}

func (hc *HTTPClient) Post(url string, body io.Reader) (*http.Response, error) {
	// create a new request
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	// set the content type header
	req.Header.Set("Content-Type", "application/json")

	// send the request and return the response
	return hc.client.Do(req)
}

func (hc *HTTPClient) Put(url string, body io.Reader) (*http.Response, error) {
	// create a new request
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}

	// set the content type header
	req.Header.Set("Content-Type", "application/json")

	// send the request and return the response
	return hc.client.Do(req)
}

func (hc *HTTPClient) Patch(url string, body io.Reader) (*http.Response, error) {
	// create a new request
	req, err := http.NewRequest("PATCH", url, body)
	if err != nil {
		return nil, err
	}

	// set the content type header
	req.Header.Set("Content-Type", "application/json")

	// send the request and return the response
	return hc.client.Do(req)
}

func (hc *HTTPClient) Delete(url string) (*http.Response, error) {
	// create a new request
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	// send the request and return the response
	return hc.client.Do(req)
}
