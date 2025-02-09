package client

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	http *resty.Client
}

func New(backendURL string, requestTimeout time.Duration) Client {
	return Client{
		http: resty.New().
			SetTimeout(requestTimeout).
			SetBaseURL(backendURL),
	}
}
