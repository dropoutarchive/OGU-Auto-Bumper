package http

import (
	http "github.com/bogdanfinn/fhttp"
	tlsClient "github.com/bogdanfinn/tls-client"
	"io"
	"io/ioutil"
)

func New() (*Client, error) {
	options := []tlsClient.HttpClientOption{
		tlsClient.WithTimeout(30),
		tlsClient.WithClientProfile(tlsClient.Opera_91),
		tlsClient.WithNotFollowRedirects(),
	}

	client, err := tlsClient.NewHttpClient(tlsClient.NewNoopLogger(), options...)
	if err != nil {
		return nil, err
	}

	return &Client{
		TlsClient: client,
	}, err
}

func (c *Client) Request(method, url string, body io.Reader, headers map[string]string) (*Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.TlsClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	readBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		Status: resp.StatusCode,
		Body:   string(readBytes),
	}, nil
}
