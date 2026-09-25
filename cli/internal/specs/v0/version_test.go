package specs

import "testing"

func TestPluginPathToOrgName(t *testing.T) {
	tests := []struct {
		path    string
		org     string
		name    string
		wantErr bool
	}{
		{path: "cloudquery/aws", org: "cloudquery", name: "aws"},
		{path: "/.cq", wantErr: true},
		{path: "/.cq/plugins/source/cloudquery/aws/v1.0.0/plugin", wantErr: true},
		{path: "cloudquery/", wantErr: true},
		{path: "/aws", wantErr: true},
		{path: "aws", wantErr: true},
		{path: "", wantErr: true},
		{path: "cloudquery/aws/extra", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			org, name, err := pluginPathToOrgName(tt.path)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error for %q, got org=%q name=%q", tt.path, org, name)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if org != tt.org || name != tt.name {
				t.Fatalf("got org=%q name=%q, want org=%q name=%q", org, name, tt.org, tt.name)
			}
		})
	}
}
