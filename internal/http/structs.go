package http

import (
	tlsClient "github.com/bogdanfinn/tls-client"
)

type Client struct {
	TlsClient tlsClient.HttpClient
}

type Response struct {
	Body   string
	Status int
}
