package network

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// Downloader is a function that returns a ReadCloser with downloaded data, an integer specifying
// how much data to expect and an error.
type Downloader func() (io.ReadCloser, int64, error)

// NewHttpGetDownloader returns a new Downloader that performs HTTP GET request against provided URL
func NewHttpGetDownloader(ctx context.Context, url string) Downloader {
	return func() (io.ReadCloser, int64, error) {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return nil, 0, err
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, 0, err
		}
		if resp.StatusCode != http.StatusOK {
			_ = resp.Body.Close()
			return nil, 0, fmt.Errorf("got %d http code instead expected %d", resp.StatusCode, http.StatusOK)
		}
		return resp.Body, resp.ContentLength, nil
	}
}

// NewFakeHttpGetDownloader return a fake downloader that returns given arguments. Useful for tests.
func NewFakeHttpGetDownloader(reader io.ReadCloser, n int64, err error) Downloader {
	return func() (io.ReadCloser, int64, error) {
		return reader, n, err
	}
}
