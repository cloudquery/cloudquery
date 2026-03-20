package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const mockClusterInfoResponse = `{
	"cluster_name": "test-cluster",
	"cluster_uuid": "test-uuid-1234",
	"name": "test-node",
	"tagline": "You Know, for Search",
	"version": {
		"number": "9.0.0",
		"build_flavor": "default",
		"build_hash": "abc123def456",
		"build_snapshot": false,
		"build_type": "docker",
		"lucene_version": "10.1.0",
		"minimum_wire_compatibility_version": "8.18.0",
		"minimum_index_compatibility_version": "8.0.0",
		"build_date": "2026-01-01T00:00:00.000Z"
	}
}`

func TestNew_LogsAllClusterInfoFields(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Elastic-Product", "Elasticsearch")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, mockClusterInfoResponse)
	}))
	defer server.Close()

	var buf bytes.Buffer
	logger := zerolog.New(&buf).Level(zerolog.DebugLevel)

	specBytes, err := json.Marshal(&Spec{Addresses: []string{server.URL}})
	require.NoError(t, err)

	c, err := New(context.Background(), logger, specBytes, plugin.NewClientOptions{})
	require.NoError(t, err)
	defer c.Close(context.Background())

	logOutput := buf.String()

	// All top-level and version fields must be present in the log output
	assert.Contains(t, logOutput, "test-cluster", "cluster_name should be logged")
	assert.Contains(t, logOutput, "test-uuid-1234", "cluster_uuid should be logged")
	assert.Contains(t, logOutput, "test-node", "node name should be logged")
	assert.Contains(t, logOutput, "You Know, for Search", "tagline should be logged")
	assert.Contains(t, logOutput, "9.0.0", "version number should be logged")
	assert.Contains(t, logOutput, "abc123def456", "build hash should be logged")
	assert.Contains(t, logOutput, "default", "build flavor should be logged")
	assert.Contains(t, logOutput, "docker", "build type should be logged")
	assert.Contains(t, logOutput, "10.1.0", "lucene version should be logged")
	assert.Contains(t, logOutput, "8.18.0", "minimum wire compatibility version should be logged")
	assert.Contains(t, logOutput, "8.0.0", "minimum index compatibility version should be logged")
	assert.Contains(t, logOutput, "Elasticsearch cluster info", "log message should be present")
}
