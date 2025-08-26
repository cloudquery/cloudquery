package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/client/spec"
	"github.com/tmc/langchaingo/textsplitter"
)

const EmbeddingsEndpoint = "https://api.openai.com/v1/embeddings"

// EmbeddingsRequester calls an OpenAI-compatible embeddings endpoint.
// It derives its settings from PgVector configuration in the destination Spec.
type EmbeddingsRequester struct {
	apiKey       string
	modelName    string
	dimensions   int
	textSplitter *spec.PgVectorTextSplitter
	httpClient   *http.Client
	endpointBase string
}

// EmbeddingsRequesterOption allows customizing the requester.
type EmbeddingsRequesterOption func(*EmbeddingsRequester)

// WithHTTPClient overrides the default HTTP client.
func WithHTTPClient(client *http.Client) EmbeddingsRequesterOption {
	return func(r *EmbeddingsRequester) { r.httpClient = client }
}

// WithEmbeddingsEndpointBase overrides the default embeddings endpoint base (for tests).
// When set, requests are sent to base+"/v1/embeddings" instead of the default EmbeddingsEndpoint.
func WithEmbeddingsEndpointBase(base string) EmbeddingsRequesterOption {
	return func(r *EmbeddingsRequester) { r.endpointBase = base }
}

// NewEmbeddingsRequesterFromSpec constructs an EmbeddingsRequester from the full Spec.
func NewEmbeddingsRequesterFromSpec(s *spec.Spec, opts ...EmbeddingsRequesterOption) (*EmbeddingsRequester, error) {
	if s == nil || s.PgVectorConfig == nil {
		return nil, errors.New("pgvector config is required")
	}
	return NewEmbeddingsRequester(s.PgVectorConfig, opts...)
}

// NewEmbeddingsRequester constructs an EmbeddingsRequester from PgVectorConfig.
func NewEmbeddingsRequester(cfg *spec.PgVectorConfig, opts ...EmbeddingsRequesterOption) (*EmbeddingsRequester, error) {
	if cfg == nil {
		return nil, errors.New("pgvector config is required")
	}
	emb := cfg.Embedding
	if emb.APIKey == "" || emb.ModelName == "" {
		return nil, errors.New("invalid embedding configuration")
	}
	// Enforce allowed models and dimensions
	dims, err := specEmbeddingDimensionsForModel(emb.ModelName)
	if err != nil {
		return nil, err
	}

	r := &EmbeddingsRequester{
		apiKey:     emb.APIKey,
		modelName:  emb.ModelName,
		dimensions: dims,
		httpClient: &http.Client{Timeout: 60 * time.Second},
	}
	// Default text splitter if none provided
	if cfg.TextSplitter != nil {
		r.textSplitter = cfg.TextSplitter
	} else {
		r.textSplitter = &spec.PgVectorTextSplitter{
			RecursiveText: spec.PgVectorRecursiveText{
				ChunkSize:    1000,
				ChunkOverlap: 500,
			},
		}
	}
	for _, o := range opts {
		o(r)
	}
	if r.httpClient == nil {
		r.httpClient = &http.Client{Timeout: 60 * time.Second}
	}
	return r, nil
}

// CreateEmbeddings requests embeddings for the provided inputs.
// It returns the per-input slices of {chunk, vector} in the same order as inputs.
func (r *EmbeddingsRequester) CreateEmbeddings(ctx context.Context, inputs []string) ([][]EmbeddingResponse, error) {
	if len(inputs) == 0 {
		return [][]EmbeddingResponse{}, nil
	}

	splitter := textsplitter.NewRecursiveCharacter(
		textsplitter.WithChunkSize(r.textSplitter.RecursiveText.ChunkSize),
		textsplitter.WithChunkOverlap(r.textSplitter.RecursiveText.ChunkOverlap),
	)

	// Split each input and flatten into a single batch while tracking boundaries
	flatInputs := make([]string, 0)
	boundaries := make([][2]int, len(inputs)) // start, count per original input
	for i, input := range inputs {
		chunks, err := splitter.SplitText(input)
		if err != nil {
			return nil, err
		}
		start := len(flatInputs)
		flatInputs = append(flatInputs, chunks...)
		boundaries[i] = [2]int{start, len(chunks)}
	}

	// Single request for all chunks
	flatEmbeddings, err := r.doEmbeddingsRequest(ctx, flatInputs)
	if err != nil {
		return nil, err
	}
	if len(flatEmbeddings) != len(flatInputs) {
		return nil, fmt.Errorf("embeddings count mismatch: got %d want %d", len(flatEmbeddings), len(flatInputs))
	}

	// Map back to per-input results
	all := make([][]EmbeddingResponse, len(inputs))
	for i := range inputs {
		start, count := boundaries[i][0], boundaries[i][1]
		all[i] = append(all[i], flatEmbeddings[start:start+count]...)
	}
	return all, nil
}

type openAIEmbeddingRequest struct {
	Model string   `json:"model"`
	Input []string `json:"input"`
}

type openAIEmbeddingResponse struct {
	Data []struct {
		Embedding []float32 `json:"embedding"`
	} `json:"data"`
}

// doEmbeddingsRequest sends a single HTTP request to the embeddings endpoint
// for the provided inputs and returns the embeddings in the same order, paired with their chunks.
func (r *EmbeddingsRequester) doEmbeddingsRequest(ctx context.Context, inputs []string) ([]EmbeddingResponse, error) {
	payload := openAIEmbeddingRequest{Model: r.modelName, Input: inputs}
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal embeddings request: %w", err)
	}
	endpoint := EmbeddingsEndpoint
	if r.endpointBase != "" {
		base := strings.TrimRight(r.endpointBase, "/")
		endpoint = base + "/v1/embeddings"
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+r.apiKey)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("embeddings request failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("embeddings request failed with status %s: %s", resp.Status, string(body))
	}
	var out openAIEmbeddingResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("failed to decode embeddings response: %w", err)
	}
	if len(out.Data) != len(inputs) {
		return nil, fmt.Errorf("embeddings response length mismatch: got %d, want %d", len(out.Data), len(inputs))
	}
	result := make([]EmbeddingResponse, len(out.Data))
	for i := range out.Data {
		result[i] = EmbeddingResponse{Chunk: inputs[i], Vector: out.Data[i].Embedding}
	}
	return result, nil
}

// helper mirrors spec.embeddingDimensionsForModel without exporting it from spec
func specEmbeddingDimensionsForModel(model string) (int, error) {
	switch model {
	case "text-embedding-3-small":
		return 1536, nil
	case "text-embedding-3-large":
		return 3072, nil
	default:
		return 0, fmt.Errorf("unsupported model_name: %s", model)
	}
}

type EmbeddingResponse struct {
	Chunk  string
	Vector []float32
}
