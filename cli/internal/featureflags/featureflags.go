package featureflags

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/cloudquery/cloudquery/cli/v6/internal/env"
	"github.com/rs/zerolog/log"
)

const (
	AccountIDAnalytics = "cq-cli-account-id-analytics"

	defaultClientSideID = "657068acb95b56102349695e"
	defaultBaseURL      = "https://clientsdk.launchdarkly.com"
	requestTimeout      = 2 * time.Second
)

type Context struct {
	UserID      string
	Team        string
	Environment string
	CLIVersion  string
}

type evaluation struct {
	Value json.RawMessage `json:"value"`
}

var (
	fetchOnce sync.Once
	cached    map[string]evaluation
)

func BoolFlag(ctx context.Context, key string, defaultValue bool, evalCtx Context) bool {
	if evalCtx.UserID == "" {
		return defaultValue
	}

	fetchOnce.Do(func() {
		cached = fetch(ctx, evalCtx)
	})

	eval, ok := cached[key]
	if !ok {
		return defaultValue
	}
	var value bool
	if err := json.Unmarshal(eval.Value, &value); err != nil {
		return defaultValue
	}
	return value
}

func fetch(ctx context.Context, evalCtx Context) map[string]evaluation {
	clientSideID := env.GetEnvOrDefault("CQ_LD_CLIENT_ID", defaultClientSideID)
	baseURL := env.GetEnvOrDefault("CQ_LD_BASE_URL", defaultBaseURL)

	encodedContext, err := encodeContext(evalCtx)
	if err != nil {
		log.Debug().Err(err).Msg("feature flags: failed to encode evaluation context")
		return nil
	}
	url := baseURL + "/sdk/evalx/" + clientSideID + "/contexts/" + encodedContext

	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Debug().Err(err).Msg("feature flags: failed to build request")
		return nil
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Debug().Err(err).Msg("feature flags: evaluation request failed")
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Debug().Int("status", resp.StatusCode).Msg("feature flags: evaluation returned non-200")
		return nil
	}

	var flags map[string]evaluation
	if err := json.NewDecoder(resp.Body).Decode(&flags); err != nil {
		log.Debug().Err(err).Msg("feature flags: failed to decode evaluation response")
		return nil
	}
	return flags
}

func encodeContext(evalCtx Context) (string, error) {
	body, err := json.Marshal(map[string]string{
		"kind":        "user",
		"key":         evalCtx.UserID,
		"team":        evalCtx.Team,
		"environment": evalCtx.Environment,
		"cliVersion":  evalCtx.CLIVersion,
	})
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(body), nil
}
