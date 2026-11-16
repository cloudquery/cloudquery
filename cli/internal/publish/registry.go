package publish

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/distribution/reference"
	"github.com/opencontainers/go-digest"
)

const (
	mediaTypeManifest     = "application/vnd.docker.distribution.manifest.v2+json"
	mediaTypeManifestList = "application/vnd.docker.distribution.manifest.list.v2+json"
	maxManifestSize       = 4 * 1024 * 1024
)

type manifestPlatform struct {
	Architecture string `json:"architecture"`
	OS           string `json:"os"`
}

type manifestDescriptor struct {
	MediaType string           `json:"mediaType"`
	Size      int64            `json:"size"`
	Digest    digest.Digest    `json:"digest"`
	Platform  manifestPlatform `json:"platform"`
}

type manifestList struct {
	SchemaVersion int                  `json:"schemaVersion"`
	MediaType     string               `json:"mediaType"`
	Manifests     []manifestDescriptor `json:"manifests"`
}

type registryClient struct {
	httpClient *http.Client
	baseURL    string
}

func newRegistryClient(imageName reference.Named, repository string, registryAuth string, insecureSkipVerify bool) *registryClient {
	return &registryClient{
		httpClient: &http.Client{Transport: newTransportWithRegistryAuth(insecureSkipVerify, registryAuth)},
		baseURL:    "https://" + repository + "/v2/" + imageName.Name() + "/manifests/",
	}
}

func (c *registryClient) getManifestDescriptor(ctx context.Context, tag string) (manifestDescriptor, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+tag, nil)
	if err != nil {
		return manifestDescriptor{}, err
	}
	req.Header.Set("Accept", mediaTypeManifest)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return manifestDescriptor{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return manifestDescriptor{}, fmt.Errorf("unexpected status code %d for manifest %q", res.StatusCode, tag)
	}
	mediaType, _, _ := strings.Cut(res.Header.Get("Content-Type"), ";")
	mediaType = strings.TrimSpace(mediaType)
	if mediaType != mediaTypeManifest {
		return manifestDescriptor{}, fmt.Errorf("unexpected media type %q for manifest %q", mediaType, tag)
	}
	payload, err := io.ReadAll(io.LimitReader(res.Body, maxManifestSize+1))
	if err != nil {
		return manifestDescriptor{}, err
	}
	if len(payload) > maxManifestSize {
		return manifestDescriptor{}, fmt.Errorf("manifest %q is larger than %d bytes", tag, maxManifestSize)
	}
	payloadDigest := digest.FromBytes(payload)
	if expected := res.Header.Get("Docker-Content-Digest"); expected != "" && expected != payloadDigest.String() {
		return manifestDescriptor{}, fmt.Errorf("manifest %q digest %s does not match registry digest %s", tag, payloadDigest, expected)
	}
	return manifestDescriptor{MediaType: mediaType, Size: int64(len(payload)), Digest: payloadDigest}, nil
}

func (c *registryClient) putManifestList(ctx context.Context, tag string, list manifestList) (digest.Digest, error) {
	payload, err := json.Marshal(list)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, c.baseURL+tag, bytes.NewReader(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", mediaTypeManifestList)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("unexpected status code %d for manifest list %q", res.StatusCode, tag)
	}
	return digest.FromBytes(payload), nil
}

func getManifestList(ctx context.Context, registry *registryClient, pkgJSON PackageJSONV1) (manifestList, error) {
	namedTags, err := getTagsFromTargets(pkgJSON.SupportedTargets)
	if err != nil {
		return manifestList{}, err
	}

	descriptors := make([]manifestDescriptor, 0, len(namedTags))
	for i, namedTag := range namedTags {
		buildTarget := pkgJSON.SupportedTargets[i]
		descriptor, err := registry.getManifestDescriptor(ctx, namedTag.Tag())
		if err != nil {
			return manifestList{}, fmt.Errorf("failed to create Docker manifest: %w", err)
		}
		descriptor.Platform = manifestPlatform{OS: buildTarget.OS, Architecture: buildTarget.Arch}
		descriptors = append(descriptors, descriptor)
	}
	return manifestList{SchemaVersion: 2, MediaType: mediaTypeManifestList, Manifests: descriptors}, nil
}

func pushManifest(ctx context.Context, pkgJSON PackageJSONV1, dockerToken string, insecureSkipVerify bool) error {
	imageName, repository, err := getManifestParams(pkgJSON.SupportedTargets[0].DockerImageTag)
	if err != nil {
		return err
	}

	registry := newRegistryClient(imageName, repository, dockerToken, insecureSkipVerify)
	list, err := getManifestList(ctx, registry, pkgJSON)
	if err != nil {
		return err
	}

	listDigest, err := registry.putManifestList(ctx, pkgJSON.Version, list)
	if err != nil {
		return fmt.Errorf("failed to create Docker manifest: %w", err)
	}
	fmt.Println("Created manifest:", listDigest.String())
	return nil
}

func parseBearerChallenge(header string) (map[string]string, error) {
	scheme, params, _ := strings.Cut(strings.TrimSpace(header), " ")
	if !strings.EqualFold(scheme, "Bearer") {
		return nil, fmt.Errorf("unsupported challenge %q", header)
	}

	parsed := map[string]string{}
	rest := strings.TrimSpace(params)
	for rest != "" {
		key, afterKey, ok := strings.Cut(rest, "=")
		key = strings.TrimSpace(key)
		if !ok || strings.ContainsAny(key, " \t") {
			return nil, fmt.Errorf("expected 1 Bearer challenge in header %q", header)
		}
		var value string
		if quoted, isQuoted := strings.CutPrefix(afterKey, `"`); isQuoted {
			value, rest, ok = strings.Cut(quoted, `"`)
			if !ok {
				return nil, errors.New("unterminated quoted value in challenge header")
			}
		} else {
			value, rest, _ = strings.Cut(afterKey, ",")
			rest = "," + rest
		}
		parsed[strings.ToLower(key)] = strings.TrimSpace(value)
		rest = strings.TrimSpace(rest)
		if rest != "" && !strings.HasPrefix(rest, ",") {
			return nil, fmt.Errorf("expected 1 Bearer challenge in header %q", header)
		}
		rest = strings.TrimSpace(strings.TrimPrefix(rest, ","))
	}
	return parsed, nil
}
