package manifests

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRead(t *testing.T) {
	m, err := Read("testdata/plugin.yaml")
	if err != nil {
		t.Fatal(err)
	}
	want := &Manifest{
		Version: 1,
		Kind:    "plugin",
		Properties: PluginProperties{
			Source:      true,
			Destination: false,
		},
		Targets: []PluginTarget{
			{
				Name: "linux_amd64",
				OS:   "linux",
				Arch: "amd64",
			},
			{
				Name: "windows_amd64",
				OS:   "windows",
				Arch: "amd64",
			},
		},
	}
	if diff := cmp.Diff(want, m); diff != "" {
		t.Fatalf("unexpected manifest (-want +got):\n%s", diff)
	}
}
