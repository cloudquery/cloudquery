package plugins

import (
	"context"
	"strings"
	"testing"
)

func TestGetLatestCQPluginRelease(t *testing.T) {
	ctx := context.Background()
	version, err := getLatestCQPluginRelease(ctx, "test", PluginTypeSource)
	if err != nil {
		t.Fatalf("error calling GetLatestPluginRelease: %v", err)
	}
	if !strings.HasPrefix(version, "v") {
		t.Errorf("got version = %q, want a version starting with 'v'", version)
	}
}

func TestGetLatestCommunityPluginRelease(t *testing.T) {
	ctx := context.Background()
	version, err := getLatestCommunityPluginRelease(ctx, "yevgenypats", "test", PluginTypeSource)
	if err != nil {
		t.Fatalf("error calling GetLatestPluginRelease: %v", err)
	}
	if !strings.HasPrefix(version, "v") {
		t.Errorf("got version = %q, want a version starting with 'v'", version)
	}
}

// func TestGetLatestCLIRelease(t *testing.T) {
// 	version, err := GetLatestCLIRelease(context.Background())
// 	if err != nil {
// 		t.Fatalf("error calling GetLatestCLIRelease: %v", err)
// 	}

// 	if !strings.HasPrefix(version, "v") {
// 		t.Errorf("got version = %q, want a version starting with 'v'", version)
// 	}
// }
