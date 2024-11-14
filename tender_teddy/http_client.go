package economy_of_tests

import (
	"encoding/json"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	client *resty.Client
}

func (c *Client) WithURL(url string) *Client {
	c.client.BaseURL = url
	return c
}

func (c *Client) Publish(ad ClassifiedAd) error {
	body, _ := json.Marshal(ad)
	_, err := c.client.R().SetBody(body).Post("/ads")
	return err
}

func (c *Client) ListAds() (ads []ClassifiedAd, err error) {
	response, _ := c.client.R().Get("/list")
	err = json.Unmarshal(response.Body(), &ads)
	return ads, err
}

func (c *Client) Search(word string) (ads []ClassifiedAd, err error) {
	response, _ := c.client.R().SetQueryParam("word", word).Get("/ads")
	err = json.Unmarshal(response.Body(), &ads)
	return ads, err
}

func NewClient(client *http.Client) *Client {
	return &Client{client: resty.NewWithClient(client)}
}
