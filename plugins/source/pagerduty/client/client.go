package client

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
)

// By default, the maximum `limit` parameter allowed by pagerduty API is 100.
const MaxPaginationLimit = 100

type Client struct {
	PagerdutyClient *pagerduty.Client
	logger          zerolog.Logger
}

func Configure(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
	authToken, err := getAuthToken()
	if err != nil {
		return nil, err
	}

	pagerdutyClient := pagerduty.NewClient(authToken)

	cqClient := Client{
		PagerdutyClient: pagerdutyClient,
		logger:          logger,
	}

	return &cqClient, nil
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

type PdYmlStruct struct {
	Authtoken string
}

// There is no easy way to get account-id or similar from the pagerduty API.
func (*Client) ID() string {
	return "pagerduty"
}

// getAuthToken returns the pagerduty auth token.
// It supports the following methods (in the following order of precedence):
// - Reading from `PAGERDUTY_AUTH_TOKEN` environment variable.
// - Reading from `~/.pd.yml` or `~/.pd.yaml` file. (Similar to the pagerduty CLI).
func getAuthToken() (string, error) {
	failedToGetAuthTokenErrorMessage := "failed to get pagerduty auth token. Please provide an auth-token (see https://www.cloudquery.io/docs/plugins/sources/pagerduty/overview#authentication)"

	envAuthToken := os.Getenv("PAGERDUTY_AUTH_TOKEN")
	if envAuthToken != "" {
		return envAuthToken, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("%s: %w", failedToGetAuthTokenErrorMessage, err)
	}

	ymlPath := path.Join(home, ".pd.yml")
	yamlPath := path.Join(home, ".pd.yaml")

	var contents []byte

	switch {
	case doesFileExist(ymlPath):
		contents, err = os.ReadFile(ymlPath)
		if err != nil {
			return "", err
		}
	case doesFileExist(yamlPath):
		contents, err = os.ReadFile(yamlPath)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf(failedToGetAuthTokenErrorMessage)
	}

	var pdYmlStruct PdYmlStruct
	err = yaml.Unmarshal(contents, &pdYmlStruct)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal `.pd.yml` file: %w", err)
	}

	if len(pdYmlStruct.Authtoken) == 0 {
		return "", fmt.Errorf(failedToGetAuthTokenErrorMessage)
	}

	return pdYmlStruct.Authtoken, nil
}

func doesFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
