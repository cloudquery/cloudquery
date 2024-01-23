package api

import (
	"context"
	"fmt"
	"net/http"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery/cli/internal/env"
)

const (
	defaultAPIURL = "https://api.cloudquery.io"
	envAPIURL     = "CLOUDQUERY_API_URL"
)

func NewClient(token string) (*cloudquery_api.ClientWithResponses, error) {
	c, err := cloudquery_api.NewClientWithResponses(env.GetEnvOrDefault(envAPIURL, defaultAPIURL),
		cloudquery_api.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			return nil
		}))
	if err != nil {
		return nil, fmt.Errorf("failed to create api client: %w", err)
	}
	return c, nil
}
