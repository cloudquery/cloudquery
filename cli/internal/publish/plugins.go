package publish

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery/cli/v6/internal/hub"
	"github.com/cloudquery/cloudquery/cli/v6/internal/publish/images"
	"github.com/distribution/reference"
	"github.com/docker/distribution"
	"github.com/docker/distribution/manifest/manifestlist"
	"github.com/docker/distribution/manifest/schema2"
	distributionclient "github.com/docker/distribution/registry/client"
	"github.com/docker/distribution/registry/client/auth/challenge"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/schollz/progressbar/v3"
)

type PackageJSONV1 struct {
	Team             string                    `json:"team"`
	Name             string                    `json:"name"`
	Message          string                    `json:"message"`
	Version          string                    `json:"version"`
	Kind             cloudquery_api.PluginKind `json:"kind"`
	Protocols        []int                     `json:"protocols"`
	SupportedTargets []TargetBuild             `json:"supported_targets"`
	PackageType      string                    `json:"package_type"`
}

type TargetBuild struct {
	OS             string `json:"os"`
	Arch           string `json:"arch"`
	Path           string `json:"path"`
	Checksum       string `json:"checksum"`
	DockerImageTag string `json:"docker_image_tag"`
}

type Opts struct {
	NoProgress bool
}

type LoadResponse struct {
	Stream string `json:"stream"`
}

type dockerProgressReader struct {
	decoder         *json.Decoder
	bar             *progressbar.ProgressBar
	layerPushedByID map[string]int64
	totalBytes      int64
}

type dockerProgressInfo struct {
	Status       string `json:"status"`
	Progress     string `json:"progress"`
	ProgressData struct {
		Current int64 `json:"current"`
		Total   int64 `json:"total"`
	} `json:"progressDetail"`
	LayerID     string `json:"id"`
	ErrorDetail struct {
		Message string `json:"message"`
	} `json:"errorDetail"`
}

type transportWithRegistryAuth struct {
	http.RoundTripper
	baseTransport *http.Transport
	registryAuth  string
}

func newTransportWithRegistryAuth(insecureSkipVerify bool, registryAuth string) *transportWithRegistryAuth {
	baseTransport := http.DefaultTransport.(*http.Transport).Clone()
	baseTransport.TLSClientConfig.InsecureSkipVerify = insecureSkipVerify
	return &transportWithRegistryAuth{
		baseTransport: baseTransport,
		registryAuth:  registryAuth,
	}
}

func (t *transportWithRegistryAuth) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	req.Header.Set("Authorization", "Bearer "+t.registryAuth)
	return t.baseTransport.RoundTrip(req)
}

func pushProgressBar(maxBytes int64, description ...string) *progressbar.ProgressBar {
	desc := ""
	if len(description) > 0 {
		desc = description[0]
	}
	return progressbar.NewOptions64(
		maxBytes,
		progressbar.OptionSetDescription(desc),
		progressbar.OptionSetWriter(os.Stdout),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(10),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionShowCount(),
		progressbar.OptionOnCompletion(func() {
			fmt.Fprint(os.Stdout, "\n")
		}),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionFullWidth(),
		progressbar.OptionSetRenderBlankState(true),
	)
}

func (pr *dockerProgressReader) Read(_ []byte) (n int, err error) {
	var progress dockerProgressInfo
	err = pr.decoder.Decode(&progress)
	if err != nil {
		if err == io.EOF {
			return 0, io.EOF
		}
		return 0, fmt.Errorf("failed to parse docker push response: %v", err)
	}
	if progress.ErrorDetail.Message != "" {
		return 0, fmt.Errorf("failed to push image: %s", progress.ErrorDetail.Message)
	}
	if progress.Status == "Pushing" {
		if pr.bar == nil {
			pr.bar = pushProgressBar(1, "Pushing")
			_ = pr.bar.RenderBlank()
		}
		if _, seen := pr.layerPushedByID[progress.LayerID]; !seen {
			pr.layerPushedByID[progress.LayerID] = 0
			pr.totalBytes += progress.ProgressData.Total
			pr.bar.ChangeMax64(pr.totalBytes)
		}
		pr.layerPushedByID[progress.LayerID] = progress.ProgressData.Current
		total := int64(0)
		for _, v := range pr.layerPushedByID {
			total += v
		}
		if total < pr.totalBytes {
			// progressbar stops responding if it reaches 100%, so as a workaround we don't update
			// the bar if we're at 100%, because there may be more layers of the image
			// coming that we don't know about.
			_ = pr.bar.Set64(total)
		}
	}

	return 0, nil
}

func ReadPackageJSON(distDir string) (PackageJSONV1, error) {
	v := SchemaVersion{}
	b, err := os.ReadFile(filepath.Join(distDir, "package.json"))
	if err != nil {
		return PackageJSONV1{}, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return PackageJSONV1{}, err
	}
	if v.SchemaVersion != 1 {
		return PackageJSONV1{}, errors.New("unsupported schema version. This CLI version only supports package.json v1. Try upgrading your CloudQuery CLI version")
	}
	pkgJSON := PackageJSONV1{}
	err = json.Unmarshal(b, &pkgJSON)
	if err != nil {
		return PackageJSONV1{}, err
	}
	return pkgJSON, nil
}

func UploadPluginDocs(ctx context.Context, c *cloudquery_api.ClientWithResponses, teamName, pluginKind, pluginName, version, docsDir string, replace bool) error {
	dirEntries, err := os.ReadDir(docsDir)
	if err != nil {
		return fmt.Errorf("failed to read docs directory: %w", err)
	}

	pages := make([]cloudquery_api.PluginDocsPageCreate, 0, len(dirEntries))
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			continue
		}
		fileExt := filepath.Ext(dirEntry.Name())
		if fileExt != ".md" {
			continue
		}
		content, err := os.ReadFile(filepath.Join(docsDir, dirEntry.Name()))
		if err != nil {
			return fmt.Errorf("failed to read docs file: %w", err)
		}
		contentStr := hub.NormalizeContent(string(content))

		contentStr, err = images.ProcessDocument(ctx, c, teamName, docsDir, contentStr)
		if err != nil {
			return fmt.Errorf("failed to process doc images for %s: %w", dirEntry.Name(), err)
		}

		pages = append(pages, cloudquery_api.PluginDocsPageCreate{
			Content: contentStr,
			Name:    strings.TrimSuffix(dirEntry.Name(), fileExt),
		})
	}

	if replace {
		body := cloudquery_api.ReplacePluginVersionDocsJSONRequestBody{
			Pages: pages,
		}
		resp, err := c.ReplacePluginVersionDocsWithResponse(ctx, teamName, cloudquery_api.PluginKind(pluginKind), pluginName, version, body)
		if err != nil {
			return fmt.Errorf("failed to upload docs: %w", err)
		}
		if resp.HTTPResponse.StatusCode > 299 {
			return hub.ErrorFromHTTPResponse(resp.HTTPResponse, resp)
		}
	} else {
		body := cloudquery_api.CreatePluginVersionDocsJSONRequestBody{
			Pages: pages,
		}
		resp, err := c.CreatePluginVersionDocsWithResponse(ctx, teamName, cloudquery_api.PluginKind(pluginKind), pluginName, version, body)
		if err != nil {
			return fmt.Errorf("failed to upload docs: %w", err)
		}
		if resp.HTTPResponse.StatusCode > 299 {
			return hub.ErrorFromHTTPResponse(resp.HTTPResponse, resp)
		}
	}

	return nil
}

func CreateNewPluginDraftVersion(ctx context.Context, c *cloudquery_api.ClientWithResponses, teamName, pluginName string, pkgJSON PackageJSONV1, specJsonSchema *string) error {
	targets := make([]string, len(pkgJSON.SupportedTargets))
	checksums := make([]string, len(pkgJSON.SupportedTargets))
	for i, t := range pkgJSON.SupportedTargets {
		targets[i] = fmt.Sprintf("%s_%s", t.OS, t.Arch)
		checksums[i] = strings.TrimPrefix(t.Checksum, "sha256:")
	}

	body := cloudquery_api.CreatePluginVersionJSONRequestBody{
		Message:          pkgJSON.Message,
		PackageType:      cloudquery_api.PluginPackageType(pkgJSON.PackageType),
		Protocols:        pkgJSON.Protocols,
		SupportedTargets: targets,
		Checksums:        checksums,
		SpecJsonSchema:   specJsonSchema,
	}
	resp, err := c.CreatePluginVersionWithResponse(ctx, teamName, pkgJSON.Kind, pluginName, pkgJSON.Version, body)
	if err != nil {
		return fmt.Errorf("failed to create plugin version: %w", err)
	}
	if resp.HTTPResponse.StatusCode > 299 {
		err := hub.ErrorFromHTTPResponse(resp.HTTPResponse, resp)
		if resp.HTTPResponse.StatusCode == http.StatusForbidden {
			return fmt.Errorf("%w. Hint: You may need to create the plugin first", err)
		}
		return err
	}
	return nil
}

func GetSpecJsonScheme(distDir string) (*string, error) {
	specJsonSchemaPath := filepath.Join(distDir, "spec_json_schema.json")
	content, err := os.ReadFile(specJsonSchemaPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to read spec_json_schema.json: %w", err)
	}
	if _, err := os.ReadFile(specJsonSchemaPath); os.IsNotExist(err) {
		return nil, nil
	}
	contentStr := string(content)
	return &contentStr, nil
}

func UploadTableSchemas(ctx context.Context, c *cloudquery_api.ClientWithResponses, teamName, pluginName, tablesJSONPath string, pkgJSON PackageJSONV1) error {
	b, err := os.ReadFile(tablesJSONPath)
	if err != nil {
		return fmt.Errorf("failed to read tables.json: %w", err)
	}
	tables := make([]cloudquery_api.PluginTableCreate, 0)
	err = json.Unmarshal(b, &tables)
	if err != nil {
		return fmt.Errorf("failed to parse tables.json: %w", err)
	}
	body := cloudquery_api.CreatePluginVersionTablesJSONRequestBody{
		Tables: tables,
	}
	resp, err := c.CreatePluginVersionTablesWithResponse(ctx, teamName, pkgJSON.Kind, pluginName, pkgJSON.Version, body)
	if err != nil {
		return fmt.Errorf("failed to upload table schemas: %w", err)
	}
	if resp.HTTPResponse.StatusCode > 299 {
		return hub.ErrorFromHTTPResponse(resp.HTTPResponse, resp)
	}
	return nil
}

func UploadPluginBinary(ctx context.Context, c *cloudquery_api.ClientWithResponses, goos, goarch, localPath string, pkgJSON PackageJSONV1) error {
	target := goos + "_" + goarch
	resp, err := c.UploadPluginAssetWithResponse(ctx, pkgJSON.Team, pkgJSON.Kind, pkgJSON.Name, pkgJSON.Version, target)
	if err != nil {
		return fmt.Errorf("failed to upload binary: %w", err)
	}
	if resp.HTTPResponse.StatusCode > 299 {
		msg := "failed to upload binary: " + resp.HTTPResponse.Status
		switch {
		case resp.JSON403 != nil:
			msg = fmt.Sprintf("%s: %s", msg, resp.JSON403.Message)
		case resp.JSON401 != nil:
			msg = fmt.Sprintf("%s: %s", msg, resp.JSON401.Message)
		}
		return errors.New(msg)
	}
	if resp.JSON201 == nil {
		return errors.New("upload response is nil, failed to upload binary")
	}
	uploadURL := resp.JSON201.Url
	err = hub.UploadFile(uploadURL, localPath)
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}
	return nil
}

func PublishNativeBinaries(ctx context.Context, c *cloudquery_api.ClientWithResponses, distDir string, pkgJSON PackageJSONV1) error {
	fmt.Println("Uploading binaries CloudQuery Hub...")
	for _, t := range pkgJSON.SupportedTargets {
		fmt.Printf("- Uploading %s_%s...\n", t.OS, t.Arch)
		err := UploadPluginBinary(ctx, c, t.OS, t.Arch, path.Join(distDir, t.Path), pkgJSON)
		if err != nil {
			return fmt.Errorf("failed to upload binary: %w", err)
		}
	}
	return nil
}

func getResponseAsString(body io.ReadCloser) (string, error) {
	defer body.Close()
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(body); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func loadDockerImage(ctx context.Context, cli *client.Client, imagePath string) error {
	f, err := os.Open(imagePath)
	if err != nil {
		return fmt.Errorf("failed to open image file: %v", err)
	}
	defer f.Close()
	resp, err := cli.ImageLoad(ctx, f, client.ImageLoadWithQuiet(true))
	if err != nil {
		return fmt.Errorf("failed to load image: %v", err)
	}
	if resp.Body == nil {
		return errors.New("failed to load image: response body is nil")
	}

	respString, err := getResponseAsString(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to get response as string: %w", err)
	}
	loadResponse := LoadResponse{}
	if err := json.Unmarshal([]byte(respString), &loadResponse); err != nil {
		return fmt.Errorf("failed to parse docker load response: %v", err)
	}
	if loadResponse.Stream != "" {
		fmt.Print(loadResponse.Stream)
	}

	return nil
}

func pushImage(ctx context.Context, dockerClient *client.Client, t TargetBuild, opts image.PushOptions, progress bool) error {
	fmt.Printf("Pushing %s\n", t.DockerImageTag)
	opts.Platform = &ocispec.Platform{
		OS:           t.OS,
		Architecture: t.Arch,
	}
	out, err := dockerClient.ImagePush(ctx, t.DockerImageTag, opts)
	if err != nil {
		return fmt.Errorf("failed to push Docker image: %v", err)
	}
	defer out.Close()

	if !progress {
		_, err = io.Copy(io.Discard, out)
		if err != nil {
			return err
		}
		return nil
	}

	// Create a progress reader to display the download progress
	pr := &dockerProgressReader{
		decoder:         json.NewDecoder(out),
		layerPushedByID: map[string]int64{},
	}
	if _, err := io.Copy(io.Discard, pr); err != nil {
		return err
	}
	if pr.bar != nil {
		_ = pr.bar.Finish()
		pr.bar.Close()
	}

	return nil
}

func getDockerToken(ctx context.Context, ref reference.Named, version, team, username, password string, insecureSkipVerify bool) (string, error) {
	// https://distribution.github.io/distribution/spec/auth/token/#how-to-authenticate
	domain := reference.Domain(ref)
	name := reference.Path(ref)

	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: insecureSkipVerify}
	httpClient := &http.Client{Transport: customTransport}

	challengeUrl := fmt.Sprintf("https://%[1]s/v2/%[2]s/manifests/%[3]s", domain, name, version)
	challengeReq, err := http.NewRequestWithContext(ctx, http.MethodPut, challengeUrl, nil)
	if err != nil {
		return "", fmt.Errorf("client: could not create request: %s", err)
	}
	challengeRes, err := httpClient.Do(challengeReq)
	if err != nil {
		return "", fmt.Errorf("client: could not send request: %s", err)
	}
	defer challengeRes.Body.Close()
	if challengeRes.StatusCode != http.StatusUnauthorized {
		return "", fmt.Errorf("client: unexpected status code: %d", challengeRes.StatusCode)
	}
	challengeHeader := challengeRes.Header.Get("WWW-Authenticate")
	if challengeHeader == "" {
		return "", errors.New("client: missing WWW-Authenticate header")
	}
	challenges := challenge.ResponseChallenges(challengeRes)
	if len(challenges) != 1 {
		return "", fmt.Errorf("client: expected 1 challenge header, got %d. Header value %q", len(challenges), challengeHeader)
	}
	realm := challenges[0].Parameters["realm"]
	service := challenges[0].Parameters["service"]
	scope := challenges[0].Parameters["scope"]
	if realm == "" || service == "" || scope == "" {
		return "", fmt.Errorf("client: could not parse challenge header %q", challengeHeader)
	}

	url := fmt.Sprintf("%[1]s?account=cli&service=%[2]s&scope=%[3]s", realm, service, scope)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	basicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	req.Header.Add("X-Meta-Plugin-Version", version)
	req.Header.Add("X-Meta-User-Team-Name", team)
	req.Header.Add("Authorization", "Basic "+basicAuth)
	if err != nil {
		return "", fmt.Errorf("client: could not create request: %s", err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("client: could not send request: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("client: unexpected status code: %d", res.StatusCode)
	}

	tokenResponse := map[string]string{}
	if err := json.NewDecoder(res.Body).Decode(&tokenResponse); err != nil {
		return "", fmt.Errorf("client: could not decode response: %s", err)
	}
	return tokenResponse["token"], nil
}

func getManifestParams(target string) (imageName reference.Named, repository string, err error) {
	ref, err := reference.ParseNamed(target)
	if err != nil {
		return nil, "", fmt.Errorf("failed to parse Docker image tag: %v", err)
	}
	repository = reference.Domain(ref)
	refPath := reference.Path(ref)
	imageName, err = reference.WithName(refPath)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create Docker repository: %v", err)
	}
	return imageName, repository, nil
}

func getManifestService(ctx context.Context, imageName reference.Named, repository string, registryAuth string, insecureSkipVerify bool) (distribution.ManifestService, error) {
	repo, err := distributionclient.NewRepository(imageName, "https://"+repository, newTransportWithRegistryAuth(insecureSkipVerify, registryAuth))
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker repository: %v", err)
	}
	manifestService, err := repo.Manifests(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker manifest: %v", err)
	}
	return manifestService, nil
}

func getTagsFromTargets(targets []TargetBuild) ([]reference.NamedTagged, error) {
	namedTags := make([]reference.NamedTagged, len(targets))
	for i, t := range targets {
		namedRef, err := reference.ParseNormalizedNamed(t.DockerImageTag)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Docker image tag: %v", err)
		}
		nameWithTag, ok := namedRef.(reference.NamedTagged)
		if !ok {
			return nil, fmt.Errorf("failed to parse Docker image tag: %v", err)
		}
		namedTags[i] = nameWithTag
	}
	return namedTags, nil
}

func getManifestList(ctx context.Context, manifestService distribution.ManifestService, pkgJSON PackageJSONV1) (*manifestlist.DeserializedManifestList, error) {
	namedTags, err := getTagsFromTargets(pkgJSON.SupportedTargets)
	if err != nil {
		return nil, err
	}

	descriptors := make([]manifestlist.ManifestDescriptor, 0)
	for i, namedTag := range namedTags {
		buildTarget := pkgJSON.SupportedTargets[i]
		manifest, err := manifestService.Get(ctx, "", distribution.WithTag(namedTag.Tag()), distribution.WithManifestMediaTypes([]string{schema2.MediaTypeManifest}))
		if err != nil {
			return nil, fmt.Errorf("failed to create Docker manifest: %v", err)
		}
		mediaType, canonical, err := manifest.Payload()
		if err != nil {
			return nil, fmt.Errorf("failed to create Docker manifest: %v", err)
		}
		manifestDesc := manifestlist.ManifestDescriptor{
			Descriptor: distribution.Descriptor{
				Digest:    digest.FromBytes(canonical),
				Size:      int64(len(canonical)),
				MediaType: mediaType},
			Platform: manifestlist.PlatformSpec{OS: buildTarget.OS, Architecture: buildTarget.Arch},
		}
		descriptors = append(descriptors, manifestDesc)
	}
	list, err := manifestlist.FromDescriptors(descriptors)
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker manifest: %v", err)
	}
	return list, nil
}

func pushManifest(ctx context.Context, pkgJSON PackageJSONV1, dockerToken string, insecureSkipVerify bool) error {
	imageName, repository, err := getManifestParams(pkgJSON.SupportedTargets[0].DockerImageTag)
	if err != nil {
		return err
	}

	manifestService, err := getManifestService(ctx, imageName, repository, dockerToken, insecureSkipVerify)
	if err != nil {
		return err
	}

	manifestList, err := getManifestList(ctx, manifestService, pkgJSON)
	if err != nil {
		return err
	}

	digest, err := manifestService.Put(ctx, manifestList, distribution.WithTag(pkgJSON.Version))
	if err != nil {
		return fmt.Errorf("failed to create Docker manifest: %v", err)
	}
	fmt.Println("Created manifest:", digest.String())
	return nil
}

func PublishToDockerRegistry(ctx context.Context, token, distDir string, pkgJSON PackageJSONV1, popts Opts) error {
	// We use a mix of the Docker Go SDK that implements the Docker Engine API https://docs.docker.com/engine/api/latest/
	// and talking to the registry directly since the Docker Engine API doesn't support the manifest API yet

	// Loading and pushing the images is done via the Docker Go SDK
	additionalHeaders := map[string]string{"X-Meta-Plugin-Version": pkgJSON.Version, "X-Meta-User-Team-Name": pkgJSON.Team}
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithHTTPHeaders(additionalHeaders), client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %v", err)
	}
	if len(pkgJSON.SupportedTargets) == 0 {
		return errors.New("no supported targets found")
	}
	fmt.Println("Pushing images to CloudQuery Docker Registry...")
	for _, t := range pkgJSON.SupportedTargets {
		fmt.Printf("Loading %s...\n", t.Path)
		imagePath := path.Join(distDir, t.Path)
		err := loadDockerImage(ctx, dockerClient, imagePath)
		if err != nil {
			return err
		}
	}
	username, password := "cli", token
	authConfig := registry.AuthConfig{
		Username: username,
		Password: password,
	}
	encodedAuth, err := registry.EncodeAuthConfig(authConfig)
	if err != nil {
		return fmt.Errorf("failed to encode Docker auth config: %v", err)
	}
	opts := image.PushOptions{
		RegistryAuth: encodedAuth,
	}
	for _, t := range pkgJSON.SupportedTargets {
		if err := pushImage(ctx, dockerClient, t, opts, !popts.NoProgress); err != nil {
			return err
		}
	}

	// Pushing the manifest is done by talking to the registry directly, so we need to get a bearer token
	ref, err := reference.ParseNamed(pkgJSON.SupportedTargets[0].DockerImageTag)
	if err != nil {
		return fmt.Errorf("failed to parse Docker image tag: %v", err)
	}

	var insecureSkipVerify = false
	if strings.HasPrefix(reference.Domain(ref), "localhost") {
		insecureSkipVerify = true
	}

	dockerToken, err := getDockerToken(ctx, ref, pkgJSON.Version, pkgJSON.Team, username, password, insecureSkipVerify)
	if err != nil {
		return fmt.Errorf("failed to get bearer token: %v", err)
	}

	if err := pushManifest(ctx, pkgJSON, dockerToken, insecureSkipVerify); err != nil {
		return fmt.Errorf("failed to tag image: %v", err)
	}

	return nil
}
