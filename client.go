package viapool

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"
)

type Client struct {
	APIKey    string        `json:"apiKey"`
	APISecret string        `json:"apiSecret"`
	UseUTC    bool          `json:"useUTC"`
	Timeout   time.Duration `json:"timeount"`
}

func (c *Client) hash(data string) string {
	h := hmac.New(sha256.New, []byte(c.APISecret))
	h.Write([]byte(data))

	return hex.EncodeToString(h.Sum(nil))
}

func (c *Client) request(req *http.Request) (*http.Response, error) {
	client := http.DefaultClient
	client.Timeout = c.Timeout
	if c.UseUTC {
		query := req.URL.Query()
		query.Add("utc", "true")
		req.URL.RawQuery = query.Encode()
	}
	req.Header.Add("X-API-KEY", c.APIKey)
	if req.Method == http.MethodPost || req.Method == http.MethodPut {
		data := req.URL.Query().Encode()
		hash := c.hash(data)
		req.Header.Add("X-SIGNATURE", hash)
	}

	return client.Do(req)
}
