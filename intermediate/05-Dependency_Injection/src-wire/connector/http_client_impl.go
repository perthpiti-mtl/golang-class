package connector

import (
	"io"
	"net/http"
)

type RealHTTPClient struct{}

func (c *RealHTTPClient) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func NewRealHTTPClient() HTTPClient {
	return &RealHTTPClient{}
}
