package analytics

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	cqauth "github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/env"
	"github.com/rs/zerolog/log"
)

const (
	defaultAPIURL  = "https://api.cloudquery.io"
	envAPIURL      = "CLOUDQUERY_API_URL"
	saltTimeout    = 5 * time.Second
	saltPathFormat = "%s/teams/%s/analytics-salt"
)

var (
	saltOnce   sync.Once
	cachedSalt string
)

func TeamAnalyticsSalt(ctx context.Context, team string) string {
	if team == "" {
		return ""
	}

	saltOnce.Do(func() {
		cachedSalt = fetchTeamAnalyticsSalt(ctx, team)
	})
	return cachedSalt
}

func fetchTeamAnalyticsSalt(ctx context.Context, team string) string {
	token, err := cqauth.NewTokenClient().GetToken()
	if err != nil {
		log.Debug().Err(err).Msg("analytics: no token for the analytics salt request")
		return ""
	}

	requestURL := fmt.Sprintf(saltPathFormat, env.GetEnvOrDefault(envAPIURL, defaultAPIURL), url.PathEscape(team))

	ctx, cancel := context.WithTimeout(ctx, saltTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		log.Debug().Err(err).Msg("analytics: failed to build the analytics salt request")
		return ""
	}
	req.Header.Set("Authorization", "Bearer "+token.Value)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Debug().Err(err).Msg("analytics: analytics salt request failed")
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Debug().Int("status", resp.StatusCode).Msg("analytics: analytics salt request returned non-200")
		return ""
	}

	var body struct {
		Salt string `json:"salt"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		log.Debug().Err(err).Msg("analytics: failed to decode the analytics salt")
		return ""
	}
	return body.Salt
}
