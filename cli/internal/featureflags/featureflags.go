// Package featureflags evaluates LaunchDarkly flags from the CLI.
//
// The CLI is a public binary, so it cannot carry a LaunchDarkly server SDK key
// and LaunchDarkly ships no Go client-side SDK. Instead this package calls the
// client-side evaluation endpoint with the environment's client-side ID, the
// same public identifier the web apps ship.
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
	// AccountIDAnalytics gates attaching hashed cloud account identifiers to
	// sync_run_completed. Temporary: it exists for a pricing research window.
	AccountIDAnalytics = "cq-cli-account-id-analytics"

	defaultClientSideID = "657068acb95b56102349695e"
	defaultBaseURL      = "https://clientsdk.launchdarkly.com"
	requestTimeout      = 2 * time.Second
)

// Context is the evaluation context flag targeting rules see. It mirrors the
// user context the web apps build, so a rule written once matches both.
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

// BoolFlag returns the boolean variation of key, or defaultValue when the flag
// is missing or LaunchDarkly cannot be reached. Flags are fetched once per
// process; every failure mode resolves to defaultValue.
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
	url := baseURL + "/sdk/evalx/" + clientSideID + "/context/" + encodedContext

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
