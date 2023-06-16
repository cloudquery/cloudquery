package download

import (
	"context"
	"path"
	"testing"
)

func TestDownloadPluginFromGithubIntegration(t *testing.T) {
	tmp := t.TempDir()
	cases := []struct {
		name       string
		org        string
		plugin     string
		version    string
		pluginType PluginType
		wantErr    bool
	}{
		{name: "monorepo source", org: "cloudquery", plugin: "hackernews", version: "v1.1.4", pluginType: PluginTypeSource},
		{name: "many repo source", org: "cloudquery", plugin: "simple-analytics", version: "v1.0.0", pluginType: PluginTypeSource},
		{name: "monorepo destination", org: "cloudquery", plugin: "postgresql", version: "v2.0.7", pluginType: PluginTypeDestination},
		{name: "community source", org: "hermanschaaf", plugin: "simple-analytics", version: "v1.0.0", pluginType: PluginTypeSource},
		{name: "invalid community source", org: "cloudquery", plugin: "invalid-plugin", version: "v0.0.x", pluginType: PluginTypeSource, wantErr: true},
		{name: "invalid monorepo source", org: "not-cloudquery", plugin: "invalid-plugin", version: "v0.0.x", pluginType: PluginTypeSource, wantErr: true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := DownloadPluginFromGithub(context.Background(), path.Join(tmp, tc.name), tc.org, tc.plugin, tc.version, tc.pluginType)
			if (err != nil) != tc.wantErr {
				t.Errorf("DownloadPluginFromGithub() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
