package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const defaultURL = "https://xkcd.com"

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

type Client struct {
	baseURL string
	client  *http.Client
}

type Option func(*Client)

func WithBaseURL(uri string) Option {
	return func(c *Client) {
		c.baseURL = uri
	}
}

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.client = httpClient
	}
}

func New(opts ...Option) (*Client, error) {
	c := &Client{
		baseURL: defaultURL,
		client:  http.DefaultClient,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

func (c *Client) GetComic(num int) (*Comic, error) {
	resp, err := c.client.Get(c.baseURL + "/" + strconv.Itoa(num) + "/info.0.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d for comic %d", resp.StatusCode, num)
	}
	var comic Comic
	err = json.NewDecoder(resp.Body).Decode(&comic)
	if err != nil {
		return nil, err
	}
	return &comic, nil
}

func (c *Client) GetLatest() (*Comic, error) {
	resp, err := c.client.Get(c.baseURL + "/info.0.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d for latest comic", resp.StatusCode)
	}
	var comic Comic
	err = json.NewDecoder(resp.Body).Decode(&comic)
	if err != nil {
		return nil, err
	}
	return &comic, nil
}
