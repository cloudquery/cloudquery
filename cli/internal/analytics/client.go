package analytics

import (
	"context"
	"os"
	"time"

	cqapi "github.com/cloudquery/cloudquery-api-go"
	cqauth "github.com/cloudquery/cloudquery-api-go/auth"
	internalAuth "github.com/cloudquery/cloudquery/cli/internal/auth"
	"github.com/cloudquery/cloudquery/cli/internal/env"
	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
	"github.com/google/uuid"
	rudderstack "github.com/rudderlabs/analytics-go/v4"
)

var (
	client                 rudderstack.Client
	cachedSyncEventDetails *eventDetails
)

type eventDetails struct {
	user                  cqapi.User
	currentTeam           string
	isCurrentTeamInternal bool
	environment           string
}

type noOpLogger struct{}

func (noOpLogger) Logf(format string, args ...any)   {}
func (noOpLogger) Errorf(format string, args ...any) {}

func InitClient() {
	writeKey := env.GetEnvOrDefault("CQ_RUDDERSTACK_WRITE_KEY", "2h38sP5iH58EYKBTRsGByJDDr6r")
	dataPlaneURL := env.GetEnvOrDefault("CQ_RUDDERSTACK_DATA_PLANE_URL", "https://analytics-events.cloudquery.io")
	config := rudderstack.Config{
		DataPlaneUrl: dataPlaneURL,
		Logger:       noOpLogger{},
	}
	var err error
	client, err = rudderstack.NewWithConfig(writeKey, config)
	if err != nil {
		client = nil
	}
}

func getEnvironment() string {
	_, ok := os.LookupEnv("CQ_CLOUD")
	if ok {
		return "cloud"
	}
	return "cli"
}

// getSyncEventDetails returns the cached event details if available, otherwise it fetches the details from the API
func getSyncEventDetails(ctx context.Context) *eventDetails {
	if cachedSyncEventDetails == nil {
		refreshSyncEventDetails(ctx)
	}
	return cachedSyncEventDetails
}

func refreshSyncEventDetails(ctx context.Context) *eventDetails {
	tc := cqauth.NewTokenClient()
	token, err := tc.GetToken()
	if err != nil {
		return nil
	}
	user, _ := internalAuth.GetUser(ctx, token)
	if user == nil {
		return nil
	}
	currentTeam, _ := internalAuth.GetTeamForToken(ctx, token)

	currentTeamInternal, _ := internalAuth.IsTeamInternal(ctx, currentTeam)

	eventDetails := &eventDetails{
		user:                  *user,
		currentTeam:           currentTeam,
		isCurrentTeamInternal: currentTeamInternal,
		environment:           getEnvironment(),
	}

	// Cache event details for future use
	cachedSyncEventDetails = eventDetails

	return eventDetails
}

func TrackLoginSuccess(ctx context.Context, invocationUUID uuid.UUID) {
	if client == nil {
		return
	}

	details := refreshSyncEventDetails(ctx)
	if details == nil {
		return
	}

	if details.isCurrentTeamInternal {
		return
	}

	_ = client.Enqueue(rudderstack.Track{
		UserId: details.user.ID.String(),
		Event:  "login_success",
		Properties: rudderstack.Properties{
			"invocation_uuid": invocationUUID,
			"team":            details.currentTeam,
			"environment":     details.environment,
			"$groups": rudderstack.Properties{
				"team": details.currentTeam,
			},
		},
	})
}

type SyncStartedEvent struct {
	Source       specs.Source
	Destinations []specs.Destination
}

func getSyncCommonProps(invocationUUID uuid.UUID, event SyncStartedEvent, details *eventDetails) rudderstack.Properties {
	destinationPaths := make([]string, len(event.Destinations))
	for i, d := range event.Destinations {
		destinationPaths[i] = d.Path
	}

	props := rudderstack.NewProperties().
		// we are using the same invocation_uuid for sync_run_id
		// invocation_uuid to be consistent with the rest of the events
		// sync_run_id to match with cloud events
		Set("invocation_uuid", invocationUUID).
		Set("sync_run_id", invocationUUID).
		Set("team", details.currentTeam).
		Set("$groups", rudderstack.NewProperties().
			Set("team", details.currentTeam)).
		Set("environment", details.environment).
		Set("sync_name", event.Source.Name).
		Set("source_path", event.Source.Path).
		Set("destination_paths", destinationPaths).
		Set("user_id", details.user.ID).
		Set("user_email", details.user.Email)

	return props
}

func TrackSyncStarted(ctx context.Context, invocationUUID uuid.UUID, event SyncStartedEvent) {
	if client == nil {
		return
	}

	details := getSyncEventDetails(ctx)
	if details == nil {
		return
	}

	if details.isCurrentTeamInternal {
		return
	}

	_ = client.Enqueue(rudderstack.Track{
		UserId:     details.user.ID.String(),
		Event:      "sync_run_started",
		Properties: getSyncCommonProps(invocationUUID, event, details),
	})
}

type SyncFinishedEvent struct {
	SyncStartedEvent
	Errors            uint64
	Warnings          uint64
	Duration          time.Duration
	ResourceCount     int64
	AbortedDueToError error
}

func TrackSyncCompleted(ctx context.Context, invocationUUID uuid.UUID, event SyncFinishedEvent) {
	if client == nil {
		return
	}

	details := getSyncEventDetails(ctx)
	if details == nil {
		return
	}

	if details.isCurrentTeamInternal {
		return
	}

	status := "success"
	if event.AbortedDueToError != nil {
		status = "error"
	}

	props := getSyncCommonProps(invocationUUID, event.SyncStartedEvent, details).
		Set("duration", event.Duration).
		Set("status", status).
		Set("total_rows", event.ResourceCount).
		Set("errors", event.Errors).
		Set("warnings", event.Warnings).
		Set("aborted_due_to_error", event.AbortedDueToError)

	_ = client.Enqueue(rudderstack.Track{
		UserId:     details.user.ID.String(),
		Event:      "sync_run_completed",
		Properties: props,
	})
}

func Close() {
	if client == nil {
		return
	}
	client.Close()
}
