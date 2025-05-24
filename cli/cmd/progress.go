package cmd

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/api"
)

type progressBar interface {
	io.Writer
	Add(num int) error
	Finish() error
}

type noopProgressBar struct {
}

func (noopProgressBar) Write(p []byte) (int, error) {
	return len(p), nil
}
func (noopProgressBar) Add(_ int) error {
	return nil
}
func (noopProgressBar) Finish() error {
	return nil
}

const (
	envProgressURL    = "CLOUDQUERY_PROGRESS_URL"
	envProgressApiKey = "CLOUDQUERY_PROGRESS_API_KEY"
)

func getFallbackProgressAPIClient(expectedTokenType auth.TokenType) (*cloudquery_api.ClientWithResponses, error) {
	authClient := auth.NewTokenClient()
	if authClient.GetTokenType() != expectedTokenType {
		return nil, nil
	}

	token, err := authClient.GetToken()
	if err != nil {
		return nil, err
	}
	return api.NewClient(token.Value)
}

func getProgressAPIClient() (*cloudquery_api.ClientWithResponses, error) {
	return getPlatformAPIClient(false)
}

func getPlatformAPIClient(testConn bool) (*cloudquery_api.ClientWithResponses, error) {
	progressURL := os.Getenv(envProgressURL)
	if progressURL == "" {
		return getFallbackProgressAPIClient(
			map[bool]auth.TokenType{
				false: auth.SyncRunAPIKey,
				true:  auth.SyncTestConnectionAPIKey,
			}[testConn])
	}

	key := os.Getenv(envProgressApiKey)
	if key == "" {
		return nil, fmt.Errorf("missing %s environment variable", envProgressApiKey)
	}

	c, err := cloudquery_api.NewClientWithResponses(progressURL,
		cloudquery_api.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", key))
			return nil
		}))
	if err != nil {
		return nil, fmt.Errorf("failed to create api client: %w", err)
	}
	return c, nil
}
