package xkcd

import (
	"encoding/json"
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
	SafeTitle  string `json:"safe_title"`
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

func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.client = client
	}
}

func NewClient(opts ...Option) (*Client, error) {
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
	comic := &Comic{}
	if err := json.NewDecoder(resp.Body).Decode(comic); err != nil {
		return nil, err
	}
	return comic, nil
}

func (c *Client) GetLatestComic(num int) (*Comic, error) {
	resp, err := c.client.Get(c.baseURL + "/info.0.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	comic := &Comic{}
	if err := json.NewDecoder(resp.Body).Decode(comic); err != nil {
		return nil, err
	}
	return comic, nil
}
