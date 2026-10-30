package publish

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/opencontainers/go-digest"
	"github.com/stretchr/testify/require"
)

func TestParseBearerChallenge(t *testing.T) {
	challenge, err := parseBearerChallenge(`Bearer realm="https://auth.example.com/token",service="registry",scope="repository:team/plugin:pull,push"`)
	require.NoError(t, err)
	require.Equal(t, map[string]string{
		"realm":   "https://auth.example.com/token",
		"service": "registry",
		"scope":   "repository:team/plugin:pull,push",
	}, challenge)
}

func TestParseBearerChallengeErrors(t *testing.T) {
	for _, header := range []string{
		`Basic realm="registry"`,
		`Bearer realm="unterminated`,
		`Bearer realm="a", Basic realm="b"`,
	} {
		_, err := parseBearerChallenge(header)
		require.Error(t, err, header)
	}
}

func TestPushManifest(t *testing.T) {
	manifests := map[string][]byte{
		"1.0.0-linux-amd64": []byte(`{"schemaVersion":2,"layers":["amd64"]}`),
		"1.0.0-linux-arm64": []byte(`{"schemaVersion":2,"layers":["arm64"]}`),
	}
	var pushed manifestList
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "Bearer token", r.Header.Get("Authorization"))
		tag := strings.TrimPrefix(r.URL.Path, "/v2/team/plugin/manifests/")
		switch r.Method {
		case http.MethodGet:
			require.Equal(t, mediaTypeManifest, r.Header.Get("Accept"))
			payload := manifests[tag]
			w.Header().Set("Content-Type", mediaTypeManifest)
			w.Header().Set("Docker-Content-Digest", digest.FromBytes(payload).String())
			_, _ = w.Write(payload)
		case http.MethodPut:
			require.Equal(t, "1.0.0", tag)
			require.Equal(t, mediaTypeManifestList, r.Header.Get("Content-Type"))
			body, err := io.ReadAll(r.Body)
			require.NoError(t, err)
			require.NoError(t, json.Unmarshal(body, &pushed))
			w.WriteHeader(http.StatusCreated)
		}
	}))
	defer ts.Close()

	host := strings.TrimPrefix(ts.URL, "https://")
	pkgJSON := PackageJSONV1{
		Version: "1.0.0",
		SupportedTargets: []TargetBuild{
			{OS: "linux", Arch: "amd64", DockerImageTag: host + "/team/plugin:1.0.0-linux-amd64"},
			{OS: "linux", Arch: "arm64", DockerImageTag: host + "/team/plugin:1.0.0-linux-arm64"},
		},
	}

	require.NoError(t, pushManifest(context.Background(), pkgJSON, "token", true))
	require.Equal(t, manifestList{
		SchemaVersion: 2,
		MediaType:     mediaTypeManifestList,
		Manifests: []manifestDescriptor{
			{
				MediaType: mediaTypeManifest,
				Size:      int64(len(manifests["1.0.0-linux-amd64"])),
				Digest:    digest.FromBytes(manifests["1.0.0-linux-amd64"]),
				Platform:  manifestPlatform{OS: "linux", Architecture: "amd64"},
			},
			{
				MediaType: mediaTypeManifest,
				Size:      int64(len(manifests["1.0.0-linux-arm64"])),
				Digest:    digest.FromBytes(manifests["1.0.0-linux-arm64"]),
				Platform:  manifestPlatform{OS: "linux", Architecture: "arm64"},
			},
		},
	}, pushed)
}

func TestGetManifestDescriptorRejectsDigestMismatch(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", mediaTypeManifest)
		w.Header().Set("Docker-Content-Digest", digest.FromString("other").String())
		_, _ = w.Write([]byte(`{"schemaVersion":2}`))
	}))
	defer ts.Close()

	imageName, repository, err := getManifestParams(strings.TrimPrefix(ts.URL, "https://") + "/team/plugin:1.0.0")
	require.NoError(t, err)
	registry := newRegistryClient(imageName, repository, "token", true)

	_, err = registry.getManifestDescriptor(context.Background(), "1.0.0")
	require.ErrorContains(t, err, "does not match registry digest")
}
