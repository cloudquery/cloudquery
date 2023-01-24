package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Get(ctx context.Context, url string, res any) error {
	request, err := http.NewRequestWithContext(ctx, "GET", c.PluginSpec.BaseURL+"/"+url, nil)
	if err != nil {
		return err
	}

	// Set the content type header
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+c.PluginSpec.ApiKey)

	response, err := c.Client.Do(request)
	if err != nil {
		return err
	}

	// Read the response body
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get url %s: %s: %s", url, response.Status, body)
	}

	if err := json.Unmarshal(body, &res); err != nil {
		return fmt.Errorf("failed to unmarshal response body from %s: %w", url, err)
	}

	return nil
}
