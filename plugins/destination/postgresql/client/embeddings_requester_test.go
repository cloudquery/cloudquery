package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/client/spec"
	"github.com/stretchr/testify/require"
)

func testSpec() *spec.Spec {
	return &spec.Spec{
		ConnectionString: "postgres://user:pass@localhost:5432/db?sslmode=disable",
		PgVectorConfig: &spec.PgVectorConfig{
			Tables: []spec.PgVectorTableConfig{{
				TableName:       "t",
				EmbedColumns:    []string{"c"},
				MetadataColumns: []string{"id"},
			}},
			Embedding: spec.PgVectorEmbedding{
				Dimensions: 1536,
				APIKey:     "k",
				ModelName:  "text-embedding-3-small",
			},
		},
	}
}

func TestNewEmbeddingsRequesterFromSpec_Validation(t *testing.T) {
	s := &spec.Spec{ConnectionString: "postgres://u:pw@h/db?sslmode=disable"}
	_, err := NewEmbeddingsRequesterFromSpec(s)
	require.Error(t, err)

	s = testSpec()
	r, err := NewEmbeddingsRequesterFromSpec(s)
	require.NoError(t, err)
	require.NotNil(t, r)
}

func TestCreateEmbeddings_SendsRequestAndKeepsOrder(t *testing.T) {
	// Arrange mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/v1/embeddings", r.URL.Path)
		require.Equal(t, "Bearer k", r.Header.Get("Authorization"))
		var in struct {
			Model string   `json:"model"`
			Input []string `json:"input"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&in))
		require.Equal(t, "text-embedding-3-small", in.Model)
		require.Equal(t, []string{"a", "b"}, in.Input)

		// respond keeping order and using dimension 3
		type resp struct {
			Data []struct {
				Embedding []float32 `json:"embedding"`
			} `json:"data"`
		}
		out := resp{Data: []struct {
			Embedding []float32 `json:"embedding"`
		}{
			{Embedding: []float32{1, 2, 3}},
			{Embedding: []float32{4, 5, 6}},
		}}
		_ = json.NewEncoder(w).Encode(out)
	}))
	defer server.Close()

	// Build requester with endpoint override
	s := testSpec()
	r, err := NewEmbeddingsRequester(s.PgVectorConfig, WithEmbeddingsEndpointBase(server.URL))
	require.NoError(t, err)

	// Act
	got, err := r.CreateEmbeddings(context.Background(), []string{"a", "b"})
	require.NoError(t, err)

	// Assert
	require.Len(t, got, 2)
	require.Len(t, got[0], 1)
	require.Len(t, got[1], 1)
	require.Equal(t, "a", got[0][0].Chunk)
	require.Equal(t, "b", got[1][0].Chunk)
	require.Equal(t, []float32{1, 2, 3}, got[0][0].Vector)
	require.Equal(t, []float32{4, 5, 6}, got[1][0].Vector)
}

func TestCreateEmbeddings_EmptyInput(t *testing.T) {
	s := testSpec()
	r, err := NewEmbeddingsRequesterFromSpec(s)
	require.NoError(t, err)
	out, err := r.CreateEmbeddings(context.Background(), nil)
	require.NoError(t, err)
	require.Len(t, out, 0)
}
