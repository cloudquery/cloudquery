package policy

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/go-hclog"

	"github.com/cloudquery/cloudquery/internal/file"
	"github.com/stretchr/testify/assert"
)

func TestManagerImpl_DownloadPolicy(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "TestManagerImpl_DownloadPolicy")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Setup database
	pool, tearDownFunc := setupDatabase(t, "test_policy_table")
	defer tearDownFunc(t)

	m := NewManager(tmpDir, pool, hclog.New(&hclog.LoggerOptions{}))

	// TODO: Add test for official cloudquery org
	// Checkout repo with "main" branch
	policyHubPath := []string{"michelvocks/my-cq-policy"}
	p, err := m.ParsePolicyHubPath(policyHubPath, "")
	assert.NoError(t, err)

	// Download policy
	err = m.DownloadPolicy(context.Background(), p)
	assert.NoError(t, err)

	// Make sure downloaded policy folder exists
	policyFolder := filepath.Join(tmpDir, defaultLocalSubPath, p.Organization, p.Repository)
	osFs := file.NewOsFs()
	_, err = osFs.Stat(policyFolder)
	assert.NoError(t, err)

	// Download policy again (which should always work)
	err = m.DownloadPolicy(context.Background(), p)
	assert.NoError(t, err)

	// Checkout repo with "master" branch
	policyHubPath = []string{"michelvocks/cq-test-policy"}
	p, err = m.ParsePolicyHubPath(policyHubPath, "")
	assert.NoError(t, err)

	err = m.DownloadPolicy(context.Background(), p)
	assert.NoError(t, err)

	// Make sure downloaded policy folder exists
	policyFolder = filepath.Join(tmpDir, defaultLocalSubPath, p.Organization, p.Repository)
	_, err = osFs.Stat(policyFolder)
	assert.NoError(t, err)
}

func TestManagerImpl_RunPolicy(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "TestManagerImpl_RunPolicy")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Setup database
	pool, tearDownFunc := setupDatabase(t, "test_policy_table")
	defer tearDownFunc(t)

	m := NewManager(tmpDir, pool, hclog.New(&hclog.LoggerOptions{}))

	// TODO: Add test for official cloudquery org
	policyHubPath := []string{"michelvocks/my-cq-policy"}
	p, err := m.ParsePolicyHubPath(policyHubPath, "")
	assert.NoError(t, err)

	// Download policy
	if err := m.DownloadPolicy(context.Background(), p); err != nil {
		t.Fatal(err)
	}

	// Run policy with specific version
	p.Version = "v0.0.2"
	results, err := m.RunPolicy(context.Background(), &ExecuteRequest{
		Policy:         p,
		UpdateCallback: nil,
		StopOnFailure:  true,
	})
	assert.NoError(t, err)
	assert.True(t, results.Passed)

	// Make sure all expected keys are contained
	expectedKeys := []string{
		"test-policy/top-level-query",
		"test-policy/sub-policy-1/sub-level-query",
		"test-policy/sub-policy-2/sub-level-query",
	}
	for k := range results.Results {
		assert.Contains(t, expectedKeys, k)
	}
}
