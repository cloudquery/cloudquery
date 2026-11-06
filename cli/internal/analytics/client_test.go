package analytics

import (
	"context"
	"errors"
	"testing"

	cqapi "github.com/cloudquery/cloudquery-api-go"
	"github.com/google/uuid"
	rudderstack "github.com/rudderlabs/analytics-go/v4"
)

type fakeClient struct {
	messages []rudderstack.Message
}

func (c *fakeClient) Enqueue(message rudderstack.Message) error {
	c.messages = append(c.messages, message)
	return nil
}

func (*fakeClient) Close() error { return nil }

func withFakeClient(t *testing.T) *fakeClient {
	t.Helper()

	fake := &fakeClient{}
	previousClient, previousDetails := client, cachedSyncEventDetails
	t.Cleanup(func() {
		client, cachedSyncEventDetails = previousClient, previousDetails
	})

	client = fake
	cachedSyncEventDetails = &eventDetails{
		user:        cqapi.User{ID: uuid.New(), Email: "user@acme.com"},
		currentTeam: "acme",
		environment: "cli",
	}
	return fake
}

func trackedProperties(t *testing.T, fake *fakeClient, event string) rudderstack.Properties {
	t.Helper()

	return trackedMessage(t, fake, event).Properties
}

func trackedMessage(t *testing.T, fake *fakeClient, event string) rudderstack.Track {
	t.Helper()

	if len(fake.messages) != 1 {
		t.Fatalf("got %d messages, want 1", len(fake.messages))
	}
	track, ok := fake.messages[0].(rudderstack.Track)
	if !ok {
		t.Fatalf("got %T, want a Track message", fake.messages[0])
	}
	if track.Event != event {
		t.Fatalf("got event %q, want %s", track.Event, event)
	}
	return track
}

func TestTrackSyncCompletedWithIDSummaries(t *testing.T) {
	fake := withFakeClient(t)

	TrackSyncCompleted(context.Background(), uuid.New(), SyncFinishedEvent{
		ResourceCount: 10,
		IDSummaries: map[string]IDSummary{
			AccountDimension:    {Hashes: []string{"aaaa", "bbbb"}, Count: 2},
			RepositoryDimension: {Hashes: []string{"cccc"}, Count: 7, Truncated: true, CountIsFloor: true},
		},
	})

	props := trackedProperties(t, fake, "sync_run_completed")
	if got := props["account_id_count"]; got != 2 {
		t.Errorf("got account_id_count %v, want 2", got)
	}
	if truncated, ok := props["account_ids_truncated"].(bool); !ok || truncated {
		t.Errorf("got account_ids_truncated %v, want false", props["account_ids_truncated"])
	}
	if got := props["repository_id_count"]; got != 7 {
		t.Errorf("got repository_id_count %v, want 7", got)
	}
	if truncated, ok := props["repository_ids_truncated"].(bool); !ok || !truncated {
		t.Errorf("got repository_ids_truncated %v, want true", props["repository_ids_truncated"])
	}
	if floor, ok := props["repository_id_count_is_floor"].(bool); !ok || !floor {
		t.Errorf("got repository_id_count_is_floor %v, want true", props["repository_id_count_is_floor"])
	}
	if floor, ok := props["account_id_count_is_floor"].(bool); !ok || floor {
		t.Errorf("got account_id_count_is_floor %v, want false", props["account_id_count_is_floor"])
	}
	hashes, ok := props["account_id_hashes"].([]string)
	if !ok || len(hashes) != 2 {
		t.Errorf("got account_id_hashes %v, want two hashes", props["account_id_hashes"])
	}
}

func TestTrackSyncCompletedWithoutIDSummaries(t *testing.T) {
	fake := withFakeClient(t)

	TrackSyncCompleted(context.Background(), uuid.New(), SyncFinishedEvent{ResourceCount: 10})

	props := trackedProperties(t, fake, "sync_run_completed")
	for _, key := range []string{"account_id_hashes", "account_id_count", "account_ids_truncated", "account_id_count_is_floor", "repository_id_hashes"} {
		if _, ok := props[key]; ok {
			t.Errorf("got property %q, want it absent", key)
		}
	}
	if got := props["total_rows"]; got != int64(10) {
		t.Errorf("got total_rows %v, want 10", got)
	}
}

func TestIdentity(t *testing.T) {
	fake := withFakeClient(t)
	_ = fake

	userID, team, environment, ok := Identity(context.Background())
	if !ok {
		t.Fatal("got ok false, want true")
	}
	if userID != cachedSyncEventDetails.user.ID.String() {
		t.Errorf("got user %q, want %q", userID, cachedSyncEventDetails.user.ID)
	}
	if team != "acme" || environment != "cli" {
		t.Errorf("got team %q environment %q, want acme/cli", team, environment)
	}
}

func TestIdentityWithoutTelemetry(t *testing.T) {
	withFakeClient(t)
	client = nil

	if _, _, _, ok := Identity(context.Background()); ok {
		t.Error("got ok true, want false when telemetry is off")
	}
}

func TestIdentityForInternalTeam(t *testing.T) {
	withFakeClient(t)
	cachedSyncEventDetails.isCurrentTeamInternal = true

	if _, _, _, ok := Identity(context.Background()); ok {
		t.Error("got ok true, want false for an internal team")
	}
}

func TestTrackUpgradeCheckCompleted(t *testing.T) {
	fake := withFakeClient(t)
	invocationUUID := uuid.New()

	TrackUpgradeCheckCompleted(context.Background(), invocationUUID, UpgradeCheckEvent{
		SourcePath:       "cloudquery/aws",
		DestinationPaths: []string{"cloudquery/postgresql"},
		FromVersion:      "v32.0.0",
		ToVersion:        "v33.1.0",
		Outcome:          "action_needed",
	})

	track := trackedMessage(t, fake, "upgrade_check_completed")
	if track.UserId != cachedSyncEventDetails.user.ID.String() {
		t.Errorf("got user %q, want %q", track.UserId, cachedSyncEventDetails.user.ID)
	}
	want := map[string]any{
		"invocation_uuid":   invocationUUID,
		"team":              "acme",
		"environment":       "cli",
		"user_id":           cachedSyncEventDetails.user.ID.String(),
		"source_path":       "cloudquery/aws",
		"from_version":      "v32.0.0",
		"to_version":        "v33.1.0",
		"upgrade_available": true,
		"outcome":           "action_needed",
		"status":            "success",
	}
	for key, value := range want {
		if got := track.Properties[key]; got != value {
			t.Errorf("got %s %v, want %v", key, got, value)
		}
	}
	if paths, ok := track.Properties["destination_paths"].([]string); !ok || len(paths) != 1 || paths[0] != "cloudquery/postgresql" {
		t.Errorf("got destination_paths %v, want [cloudquery/postgresql]", track.Properties["destination_paths"])
	}
	if _, ok := track.Properties["user_email"]; ok {
		t.Error("got property user_email, want it absent")
	}
}

func TestTrackUpgradeCheckCompletedWithError(t *testing.T) {
	fake := withFakeClient(t)

	TrackUpgradeCheckCompleted(context.Background(), uuid.New(), UpgradeCheckEvent{
		SourcePath: "cloudquery/aws",
		ToVersion:  "v33.1.0",
		Error:      errors.New("failed to start source"),
	})

	props := trackedProperties(t, fake, "upgrade_check_completed")
	if got := props["status"]; got != "error" {
		t.Errorf("got status %v, want error", got)
	}
	if _, ok := props["outcome"]; ok {
		t.Error("got property outcome, want it absent")
	}
	if available, ok := props["upgrade_available"].(bool); !ok || available {
		t.Errorf("got upgrade_available %v, want false", props["upgrade_available"])
	}
}

func TestUpgradeCheckPropsWithoutLogin(t *testing.T) {
	props := getUpgradeCheckProps(uuid.New(), UpgradeCheckEvent{Outcome: "no_action_needed"}, nil)

	for _, key := range []string{"team", "user_id", "user_email", "environment", "$groups"} {
		if _, ok := props[key]; ok {
			t.Errorf("got property %q, want it absent", key)
		}
	}
	if got := props["outcome"]; got != "no_action_needed" {
		t.Errorf("got outcome %v, want no_action_needed", got)
	}
}

func TestTrackUpgradeCheckCompletedForInternalTeam(t *testing.T) {
	fake := withFakeClient(t)
	cachedSyncEventDetails.isCurrentTeamInternal = true

	TrackUpgradeCheckCompleted(context.Background(), uuid.New(), UpgradeCheckEvent{})

	if len(fake.messages) != 0 {
		t.Errorf("got %d messages, want 0 for an internal team", len(fake.messages))
	}
}

func TestTrackUpgradeCheckCompletedWithoutTelemetry(t *testing.T) {
	withFakeClient(t)
	client = nil

	TrackUpgradeCheckCompleted(context.Background(), uuid.New(), UpgradeCheckEvent{})
}

func TestIsUpgrade(t *testing.T) {
	tests := []struct {
		from, to string
		want     bool
	}{
		{from: "v1.0.0", to: "v1.1.0", want: true},
		{from: "v1.1.0", to: "v1.1.0", want: false},
		{from: "v2.0.0", to: "v1.9.0", want: false},
		{from: "unknown", to: "v1.0.0", want: false},
		{from: "v1.0.0", to: "", want: false},
	}
	for _, tc := range tests {
		if got := isUpgrade(tc.from, tc.to); got != tc.want {
			t.Errorf("isUpgrade(%q, %q) = %v, want %v", tc.from, tc.to, got, tc.want)
		}
	}
}
