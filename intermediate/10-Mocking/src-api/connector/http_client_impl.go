package connector

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-class/api/config"
	"github.com/golang-class/api/model"
	"log"
	"net/http"
	"strconv"
	"time"
)

type RealCatImageAPIClient struct {
	client  *http.Client
	baseURL string
}

func (c *RealCatImageAPIClient) Search(ctx *gin.Context, limit int) ([]model.CatImage, error) {
	fullUrl := c.baseURL + "/images/search"
	req, err := http.NewRequestWithContext(ctx.Request.Context(), "GET", fullUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json")
	q := req.URL.Query()
	q.Add("limit", strconv.Itoa(limit))
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("HTTP error: %s", resp.Status)
	}

	var result []model.CatImage
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result, nil
}

func NewRealHTTPClient(config *config.Config) CatImageAPIClient {
	client := &http.Client{
		Timeout: time.Second * time.Duration(config.CatAPI.TimeoutSecond),
	}
	return &RealCatImageAPIClient{
		client:  client,
		baseURL: config.CatAPI.Url,
	}
}
