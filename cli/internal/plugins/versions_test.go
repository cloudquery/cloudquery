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

func TestGetLatestCLIRelease(t *testing.T) {
	version, err := GetLatestCLIRelease(context.Background())
	if err != nil {
		t.Fatalf("error calling GetLatestCLIRelease: %v", err)
	}

	if !strings.HasPrefix(version, "v") {
		t.Errorf("got version = %q, want a version starting with 'v'", version)
	}
}

func TestExtractVersionFromTag(t *testing.T) {
	cases := []struct {
		give string
		want string
	}{
		{give: "plugins-source-test-v0.1.21", want: "v0.1.21"},
		{give: "plugins-source-test-v0.1.21-pre.123", want: "v0.1.21-pre.123"},
		{give: "cli-v1.1.0-pre.1", want: "v1.1.0-pre.1"},
		{give: "cli-v123.145.234-pre.123", want: "v123.145.234-pre.123"},
	}
	for _, tc := range cases {
		got := extractVersionFromTag(tc.give)
		if got != tc.want {
			t.Errorf("extractVersionFromTag(%q) = %q, want %q", tc.give, got, tc.want)
		}
	}
}
