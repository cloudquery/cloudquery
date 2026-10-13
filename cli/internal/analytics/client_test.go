package analytics

import (
	"context"
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

func trackedProperties(t *testing.T, fake *fakeClient) rudderstack.Properties {
	t.Helper()

	if len(fake.messages) != 1 {
		t.Fatalf("got %d messages, want 1", len(fake.messages))
	}
	track, ok := fake.messages[0].(rudderstack.Track)
	if !ok {
		t.Fatalf("got %T, want a Track message", fake.messages[0])
	}
	if track.Event != "sync_run_completed" {
		t.Fatalf("got event %q, want sync_run_completed", track.Event)
	}
	return track.Properties
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

	props := trackedProperties(t, fake)
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

	props := trackedProperties(t, fake)
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
