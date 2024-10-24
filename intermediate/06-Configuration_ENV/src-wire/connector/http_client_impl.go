package connector

import (
	"example.com/05-di/m/config"
	"fmt"
	"io"
	"net/http"
)

type RealHTTPClient struct {
}

func (c *RealHTTPClient) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func NewRealHTTPClient(cfg *config.Config) HTTPClient {
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	print(addr)
	return &RealHTTPClient{}
}
