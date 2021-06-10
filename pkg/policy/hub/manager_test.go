package hub

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/cloudquery/cloudquery/internal/file"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/go-git/go-git/v5"
)

func TestManagerImpl_DownloadPolicy(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "TestManagerImpl_DownloadPolicy")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	m := NewManager(&config.Config{
		CloudQuery: config.CloudQuery{
			PolicyDirectory: tmpDir,
		},
	})

	policyHubPath := []string{"michelvocks/my-cq-policy"}
	p, err := m.ParsePolicyHubPath(policyHubPath, "")
	if err != nil {
		t.Fatal(err)
	}

	// Download policy
	if err := m.DownloadPolicy(context.Background(), p); err != nil {
		t.Fatal(err)
	}

	// Make sure downloaded policy folder exists
	policyFolder := filepath.Join(tmpDir, p.Organization, p.Repository)
	osFs := file.NewOsFs()
	if _, err := osFs.Stat(policyFolder); err != nil {
		t.Fatal(err)
	}

	// Change version
	p.Version = "v0.0.1"

	// Download policy again (which should always work)
	if err := m.DownloadPolicy(context.Background(), p); err != nil {
		t.Fatal(err)
	}

	// Make sure version changed e.g. tag was checked out
	r, err := git.PlainOpen(policyFolder)
	if err != nil {
		t.Fatal(err)
	}
	ref, err := r.Head()
	if err != nil {
		t.Fatal(err)
	}
	versionTag, err := r.Tag(p.Version)
	if err != nil {
		t.Fatal(err)
	}
	if ref.Hash() != versionTag.Hash() {
		t.Fatalf("reference is not equal to version. Got %s want %s", ref.Hash().String(), versionTag.Hash().String())
	}
}
