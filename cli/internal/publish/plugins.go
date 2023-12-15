package publish

import (
	"bytes"
	"context"
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
	"github.com/cloudquery/cloudquery/cli/internal/hub"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
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

func CreateNewPluginDraftVersion(ctx context.Context, c *cloudquery_api.ClientWithResponses, teamName, pluginName string, pkgJSON PackageJSONV1) error {
	targets := make([]string, len(pkgJSON.SupportedTargets))
	checksums := make([]string, len(pkgJSON.SupportedTargets))
	for i, t := range pkgJSON.SupportedTargets {
		targets[i] = fmt.Sprintf("%s_%s", t.OS, t.Arch)
		checksums[i] = strings.TrimPrefix(t.Checksum, "sha256:")
	}

	body := cloudquery_api.CreatePluginVersionJSONRequestBody{
		Message:          pkgJSON.Message,
		PackageType:      cloudquery_api.CreatePluginVersionJSONBodyPackageType(pkgJSON.PackageType),
		Protocols:        pkgJSON.Protocols,
		SupportedTargets: targets,
		Checksums:        checksums,
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
	resp, err := c.UploadPluginAssetWithResponse(ctx, pkgJSON.Team, pkgJSON.Kind, pkgJSON.Team, pkgJSON.Version, target)
	if err != nil {
		return fmt.Errorf("failed to upload binary: %w", err)
	}
	if resp.HTTPResponse.StatusCode > 299 {
		msg := fmt.Sprintf("failed to upload binary: %s", resp.HTTPResponse.Status)
		switch {
		case resp.JSON403 != nil:
			msg = fmt.Sprintf("%s: %s", msg, resp.JSON403.Message)
		case resp.JSON401 != nil:
			msg = fmt.Sprintf("%s: %s", msg, resp.JSON401.Message)
		}
		return fmt.Errorf(msg)
	}
	if resp.JSON201 == nil {
		return fmt.Errorf("upload response is nil, failed to upload binary")
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

func getResponseAsString(body io.ReadCloser) string {
	defer body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	return buf.String()
}

func loadDockerImage(ctx context.Context, cli *client.Client, imagePath string) error {
	f, err := os.Open(imagePath)
	if err != nil {
		return fmt.Errorf("failed to open image file: %v", err)
	}
	defer f.Close()
	resp, err := cli.ImageLoad(ctx, f, true)
	if err != nil {
		return fmt.Errorf("failed to load image: %v", err)
	}
	if resp.Body == nil {
		return fmt.Errorf("failed to load image: response body is nil")
	}

	respString := getResponseAsString(resp.Body)
	loadResponse := LoadResponse{}
	if err := json.Unmarshal([]byte(respString), &loadResponse); err != nil {
		return fmt.Errorf("failed to parse docker load response: %v", err)
	}
	if loadResponse.Stream != "" {
		fmt.Print(loadResponse.Stream)
	}

	return nil
}

func pushImage(ctx context.Context, dockerClient *client.Client, t TargetBuild, opts types.ImagePushOptions) error {
	fmt.Printf("Pushing %s\n", t.DockerImageTag)
	opts.Platform = fmt.Sprintf("%s/%s", t.OS, t.Arch)
	out, err := dockerClient.ImagePush(ctx, t.DockerImageTag, opts)
	if err != nil {
		return fmt.Errorf("failed to push Docker image: %v", err)
	}
	defer out.Close()

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

func PublishToDockerRegistry(ctx context.Context, token, distDir string, pkgJSON PackageJSONV1) error {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %v", err)
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
	authConfig := registry.AuthConfig{
		Username: "cli",
		Password: token,
	}
	encodedAuth, err := registry.EncodeAuthConfig(authConfig)
	if err != nil {
		return fmt.Errorf("failed to encode Docker auth config: %v", err)
	}
	opts := types.ImagePushOptions{
		RegistryAuth: encodedAuth,
	}
	for _, t := range pkgJSON.SupportedTargets {
		if err := pushImage(ctx, dockerClient, t, opts); err != nil {
			return err
		}
	}

	return nil
}
