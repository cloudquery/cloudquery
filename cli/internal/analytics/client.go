package analytics

import (
	"context"
	"os"
	"strings"
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

func Init() {
	writeKey := env.GetEnvOrDefault("CQ_RUDDERSTACK_WRITE_KEY", "2h38sP5iH58EYKBTRsGByJDDr6r")
	dataPlaneURL := env.GetEnvOrDefault("CQ_RUDDERSTACK_DATA_PLANE_URL", "https://analytics-events.cloudquery.io")
	client = rudderstack.New(writeKey, dataPlaneURL)
}

func getEnvironment() string {
	_, ok := os.LookupEnv("CQ_CLOUD")
	if ok {
		return "cloud"
	}
	return "cli"
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
		environment: getEnvironment(),
	}
}

func redactSource(source specs.Source) specs.Source {
	return specs.Source{
		Metadata: specs.Metadata{
			Path:    source.Metadata.Path,
			Version: source.Metadata.Version,
		},
	}
}

func redactDestination(destination specs.Destination) specs.Destination {
	return specs.Destination{
		Metadata: specs.Metadata{
			Path:    destination.Metadata.Path,
			Version: destination.Metadata.Version,
		},
	}
}

func redactDestinations(destinations []specs.Destination) []specs.Destination {
	redacted := make([]specs.Destination, len(destinations))
	for i, destination := range destinations {
		redacted[i] = redactDestination(destination)
	}
	return redacted
}

func Identify(ctx context.Context, invocationUUID uuid.UUID) {
	if client == nil {
		return
	}

	details := getEventDetails(ctx)
	if details == nil {
		_ = client.Enqueue(rudderstack.Identify{
			AnonymousId: invocationUUID.String(),
			Traits: rudderstack.Traits{
				"environment": getEnvironment(),
			},
		})
		return
	}

	_ = client.Enqueue(rudderstack.Identify{
		AnonymousId: invocationUUID.String(),
		UserId:      details.userId,
		Traits: rudderstack.Traits{
			"team":        details.currentTeam,
			"environment": details.environment,
		},
	})
}

func TrackCommandStart(ctx context.Context, commandName string, invocationUUID uuid.UUID) {
	if client == nil {
		return
	}

	_ = client.Enqueue(rudderstack.Track{
		AnonymousId: invocationUUID.String(),
		Event:       "command_" + strings.ToLower(commandName) + "_start",
		Properties: rudderstack.Properties{
			"invocationUUID": invocationUUID,
		},
	})

}

func TrackCommandEnd(ctx context.Context, commandName string, invocationUUID uuid.UUID, err error) {
	if client == nil {
		return
	}

	status := "success"
	if err != nil {
		status = "error"
	}

	_ = client.Enqueue(rudderstack.Track{
		AnonymousId: invocationUUID.String(),
		Event:       "command_" + strings.ToLower(commandName) + "_end",
		Properties: rudderstack.Properties{
			"invocationUUID": invocationUUID,
			"status":         status,
		},
	})

}

func TrackLoginSuccess(ctx context.Context, invocationUUID uuid.UUID) {
	if client == nil {
		return
	}

	details := getEventDetails(ctx)
	if details == nil {
		return
	}

	_ = client.Enqueue(rudderstack.Track{
		AnonymousId: invocationUUID.String(),
		UserId:      details.userId,
		Event:       "login_success",
		Properties: rudderstack.Properties{
			"invocationUUID": invocationUUID,
			"team":           details.currentTeam,
			"environment":    details.environment,
		},
	})
}

type SyncStartedEvent struct {
	Source       specs.Source
	Destinations []specs.Destination
}

func TrackSyncStarted(ctx context.Context, invocationUUID uuid.UUID, event SyncStartedEvent) {
	if client == nil {
		return
	}

	details := getEventDetails(ctx)
	if details == nil {
		return
	}

	_ = client.Enqueue(rudderstack.Track{
		UserId: details.userId,
		Event:  "sync_run_started",
		Properties: rudderstack.Properties{
			"invocation_uuid": invocationUUID,
			"team":            details.currentTeam,
			"environment":     details.environment,
			"source":          redactSource(event.Source),
			"destinations":    redactDestinations(event.Destinations),
		},
	})
}

type SyncFinishedEvent struct {
	Source        specs.Source
	Destinations  []specs.Destination
	Errors        uint64
	Warnings      uint64
	Duration      time.Duration
	ResourceCount int64
}

func TrackSyncCompleted(ctx context.Context, invocationUUID uuid.UUID, event SyncFinishedEvent) {
	if client == nil {
		return
	}

	details := getEventDetails(ctx)
	if details == nil {
		return
	}

	_ = client.Enqueue(rudderstack.Track{
		UserId: details.userId,
		Event:  "sync_run_completed",
		Properties: rudderstack.Properties{
			"invocation_uuid": invocationUUID,
			"team":            details.currentTeam,
			"environment":     details.environment,
			"source":          redactSource(event.Source),
			"destinations":    redactDestinations(event.Destinations),
			"duration":        event.Duration,
			"status":          "success",
			"resource_count":  event.ResourceCount,
			"errors":          event.Errors,
			"warnings":        event.Warnings,
		},
	})
}

func Close() {
	if client == nil {
		return
	}
	client.Close()
}
