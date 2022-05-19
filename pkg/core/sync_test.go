package core

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/pkg/core/state"
	"github.com/jackc/pgx/v4"

	"github.com/cloudquery/cloudquery/internal/firebase"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/stretchr/testify/assert"
)

// TODO: test edge case where "latest" provider is not download and we want to downgrade from it -> expect correct diag

func Test_Sync(t *testing.T) {
	dsn := setupDB(t)
	pManager, err := plugin.NewManager(registry.NewRegistryHub(firebase.CloudQueryRegistryURL))
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	// Download plugin before sync call
	_, diags := Download(context.Background(), pManager, &DownloadOptions{[]registry.Provider{{Name: "test", Version: "v0.0.10", Source: "cloudquery"}, {Name: "test", Version: "latest", Source: "cloudquery"}}, false})
	assert.False(t, diags.HasErrors())

	sta, err := state.NewClient(context.Background(), dsn)
	if err != nil {
		assert.NoError(t, err)
	}
	defer sta.Close()

	result, diags := Sync(context.Background(), sta, pManager, registry.Provider{
		Name:    "test",
		Version: "v0.0.10",
		Source:  "cloudquery",
	})
	assert.False(t, diags.HasErrors())
	assert.Equal(t, &SyncResult{State: Installed, OldVersion: "", NewVersion: "v0.0.10"}, result)
	// Verify tables were created
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		assert.FailNow(t, "failed to create connection")
		return
	}
	defer conn.Close(context.Background())
	_, err = conn.Exec(context.Background(), "select some_bool from slow_resource")
	assert.Nil(t, err)

	// Download plugin before sync call
	_, diags = Download(context.Background(), pManager, &DownloadOptions{[]registry.Provider{{Name: "test", Version: "v0.0.11", Source: "cloudquery"}}, false})
	assert.False(t, diags.HasErrors())

	// upgrade
	result, diags = Sync(context.Background(), sta, pManager, registry.Provider{
		Name:    "test",
		Version: "v0.0.11",
		Source:  "cloudquery",
	})
	assert.False(t, diags.HasErrors())
	assert.Equal(t, &SyncResult{State: Upgraded, OldVersion: "v0.0.10", NewVersion: "v0.0.11"}, result)
	_, err = conn.Exec(context.Background(), "select some_bool from slow_resource")
	assert.Nil(t, err)

	result, diags = Sync(context.Background(), sta, pManager, registry.Provider{
		Name:    "test",
		Version: "v0.0.10",
		Source:  "cloudquery",
	})
	assert.False(t, diags.HasErrors())
	assert.Equal(t, &SyncResult{State: Downgraded, OldVersion: "v0.0.11", NewVersion: "v0.0.10"}, result)
	_, err = conn.Exec(context.Background(), "select some_bool from slow_resource")
	assert.Nil(t, err)
}

// TODO: in general we have an issue here if we drop with a previous version we must pick the latest similar to migrations
// The correct solution here is to save installed providers in the database

func Test_Drop(t *testing.T) {
	dsn := setupDB(t)
	pManager, err := plugin.NewManager(registry.NewRegistryHub(firebase.CloudQueryRegistryURL), plugin.WithAllowReattach())
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	sta, err := state.NewClient(context.Background(), dsn)
	if err != nil {
		assert.NoError(t, err)
	}
	defer sta.Close()

	// Download plugin before sync call
	_, diags := Download(context.Background(), pManager, &DownloadOptions{[]registry.Provider{{Name: "test", Version: "v0.0.10", Source: "cloudquery"}}, false})
	assert.False(t, diags.HasErrors())

	diags = Drop(context.Background(), sta, pManager, registry.Provider{
		Name:    "test",
		Version: "v0.0.10",
		Source:  "cloudquery",
	})
	assert.False(t, diags.HasErrors())

	result, diags := Sync(context.Background(), sta, pManager, registry.Provider{
		Name:    "test",
		Version: "v0.0.10",
		Source:  "cloudquery",
	})
	assert.False(t, diags.HasErrors())
	assert.Equal(t, &SyncResult{State: Installed, OldVersion: "", NewVersion: "v0.0.10"}, result)
	// Verify tables were created
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		assert.FailNow(t, "failed to create connection")
		return
	}
	defer conn.Close(context.Background())
	_, err = conn.Exec(context.Background(), "select some_bool from slow_resource")
	assert.Nil(t, err)

	diags = Drop(context.Background(), sta, pManager, registry.Provider{
		Name:    "test",
		Version: "v0.0.10",
		Source:  "cloudquery",
	})
	assert.False(t, diags.HasErrors())
}
