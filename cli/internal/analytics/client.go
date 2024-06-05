package analytics

import (
	"context"
	"time"

	cqauth "github.com/cloudquery/cloudquery-api-go/auth"
	internalAuth "github.com/cloudquery/cloudquery/cli/internal/auth"
	"github.com/cloudquery/cloudquery/cli/internal/env"
	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
	"github.com/google/uuid"
	rudderstack "github.com/rudderlabs/analytics-go/v4"
)

var (
	client rudderstack.Client
)

type eventDetails struct {
	userId      string
	currentTeam string
	environment string
}

func initRudderStack() {
	writeKey := env.GetEnvOrDefault("CQ_RUDDERSTACK_WRITE_KEY", "2h38sP5iH58EYKBTRsGByJDDr6r")
	dataPlaneURL := env.GetEnvOrDefault("CQ_RUDDERSTACK_DATA_PLANE_URL", "https://analytics-events.cloudquery.io")
	client = rudderstack.New(writeKey, dataPlaneURL)
}

func Init(ctx context.Context) error {
	initRudderStack()
	return nil
}

func getEnvironment(token cqauth.Token) string {
	switch token.Type {
	case cqauth.SyncRunAPIKey, cqauth.SyncTestConnectionAPIKey:
		return "cloud"
	case cqauth.APIKey, cqauth.BearerToken:
		return "cli"
	default:
		return "unknown"
	}
}

func getEventDetails(ctx context.Context) *eventDetails {
	tc := cqauth.NewTokenClient()
	token, err := tc.GetToken()
	if err != nil {
		return nil
	}
	userId, _ := internalAuth.GetUserId(ctx, token)
	if userId == "" {
		return nil
	}
	currentTeam, _ := internalAuth.GetTeamForToken(ctx, token)

	return &eventDetails{
		userId:      userId,
		currentTeam: currentTeam,
		environment: getEnvironment(token),
	}
}

func redactSource(source *specs.Source) *specs.Source {
	return &specs.Source{
		Metadata: specs.Metadata{
			Path:    source.Metadata.Path,
			Version: source.Metadata.Version,
		},
	}
}

func redactSources(sources []*specs.Source) []*specs.Source {
	redacted := make([]*specs.Source, len(sources))
	for i, source := range sources {
		redacted[i] = redactSource(source)
	}
	return redacted
}

func redactDestination(destination *specs.Destination) *specs.Destination {
	return &specs.Destination{
		Metadata: specs.Metadata{
			Path:    destination.Metadata.Path,
			Version: destination.Metadata.Version,
		},
	}

}

func redactDestinations(destinations []*specs.Destination) []*specs.Destination {
	redacted := make([]*specs.Destination, len(destinations))
	for i, destination := range destinations {
		redacted[i] = redactDestination(destination)
	}
	return redacted

}

func Identify(ctx context.Context) {
	if client == nil {
		return
	}

	details := getEventDetails(ctx)
	if details == nil {
		return
	}

	client.Enqueue(rudderstack.Identify{
		UserId: details.userId,
		Traits: rudderstack.Traits{
			"team":        details.currentTeam,
			"environment": details.environment,
		},
	})
}

func TrackLogin(ctx context.Context, invocationUUID uuid.UUID) {
	if client == nil {
		return
	}

	details := getEventDetails(ctx)
	if details == nil {
		return
	}

	client.Enqueue(rudderstack.Track{
		UserId: details.userId,
		Event:  "Login",
		Properties: rudderstack.Properties{
			"invocationUUID": invocationUUID,
			"team":           details.currentTeam,
			"environment":    details.environment,
		},
	})
}

type SyncStartedEvent struct {
	Sources      []*specs.Source
	Destinations []*specs.Destination
}

func TrackSyncStarted(ctx context.Context, invocationUUID uuid.UUID, event SyncStartedEvent) {
	if client == nil {
		return
	}

	details := getEventDetails(ctx)
	if details == nil {
		return
	}

	client.Enqueue(rudderstack.Track{
		UserId: details.userId,
		Event:  "Sync Started",
		Properties: rudderstack.Properties{
			"invocationUUID": invocationUUID,
			"environment":    details.environment,
			"sources":        redactSources(event.Sources),
			"destinations":   redactDestinations(event.Destinations),
		},
	})
}

type SyncFinishedEvent struct {
	Source        specs.Source
	Destinations  []specs.Destination
	Errors        uint64
	Warnings      uint64
	Duration      time.Duration
	Result        string
	ResourceCount int64
}

func TrackSyncFinished(ctx context.Context, invocationUUID uuid.UUID, event SyncFinishedEvent) {
	if client == nil {
		return
	}

	details := getEventDetails(ctx)
	if details == nil {
		return
	}

	toRedact := make([]*specs.Destination, len(event.Destinations))
	for i, destination := range event.Destinations {
		toRedact[i] = &destination
	}

	client.Enqueue(rudderstack.Track{
		UserId: details.userId,
		Event:  "Sync Finished",
		Properties: rudderstack.Properties{
			"invocationUUID": invocationUUID,
			"environment":    details.environment,
			"source":         redactSource(&event.Source),
			"destinations":   redactDestinations(toRedact),
			"duration":       event.Duration,
			"result":         event.Result,
			"resourceCount":  event.ResourceCount,
			"errors":         event.Errors,
			"warnings":       event.Warnings,
		},
	})
}

func Close() {
	if client == nil {
		return
	}
	client.Close()
}
