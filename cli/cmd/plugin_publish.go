package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strings"
	"syscall"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/spf13/cobra"
)

const (
	pluginPublishShort = "Publish to CloudQuery Hub."
	pluginPublishLong  = `Publish to CloudQuery Hub.

This publishes a plugin version to CloudQuery Hub from a local dist directory.
`
	pluginPublishExample = `
# Publish a plugin version from a local dist directory
cloudquery plugin publish my_team/my_plugin`
)

func newCmdPluginPublish() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "publish <team_name>/<plugin_name> [-D dist]",
		Short:   pluginPublishShort,
		Long:    pluginPublishLong,
		Example: pluginPublishExample,
		Hidden:  true,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Set up a channel to listen for OS signals for graceful shutdown.
			ctx, cancel := context.WithCancel(cmd.Context())

			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGTERM)

			go func() {
				<-sigChan
				cancel()
			}()

			return runPluginPublish(ctx, cmd, args)
		},
	}
	cmd.Flags().StringP("dist-dir", "D", "dist", "Path to the dist directory")
	cmd.Flags().BoolP("finalize", "f", false, `Finalize the plugin version after publishing. If false, the plugin version will be marked as draft=true.`)

	return cmd
}

type PackageJSONV1 struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Version string `json:"version"`

	Kind             cloudquery_api.PluginKind `json:"kind"`
	Protocols        []int                     `json:"protocols"`
	SupportedTargets []TargetBuild             `json:"supported_targets"`
	PackageType      string                    `json:"package_type"`
}

type TargetBuild struct {
	OS       string `json:"os"`
	Arch     string `json:"arch"`
	Path     string `json:"path"`
	Checksum string `json:"checksum"`
}

func runPluginPublish(ctx context.Context, cmd *cobra.Command, args []string) error {
	tc := auth.NewTokenClient()
	token, err := tc.GetToken()
	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}

	distDir := cmd.Flag("dist-dir").Value.String()
	pkgJSON, err := readPackageJSON(distDir)
	if err != nil {
		return fmt.Errorf("failed to read package.json: %w", err)
	}

	parts := strings.Split(args[0], "/")
	if len(parts) != 2 {
		return errors.New("invalid plugin name. Must be in format <team_name>/<plugin_name>")
	}
	teamName, pluginName := parts[0], parts[1]

	name := fmt.Sprintf("%s/%s@%s", teamName, pluginName, pkgJSON.Version)
	fmt.Printf("Publishing plugin %s to CloudQuery Hub...\n", name)

	c, err := cloudquery_api.NewClientWithResponses(getEnvOrDefault(envAPIURL, defaultAPIURL),
		cloudquery_api.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			return nil
		}))
	if err != nil {
		return fmt.Errorf("failed to create hub client: %w", err)
	}

	// create new draft version
	err = createNewPluginDraftVersion(ctx, c, teamName, pluginName, pkgJSON)
	if err != nil {
		return fmt.Errorf("failed to create new draft version: %w", err)
	}

	if pkgJSON.Kind == cloudquery_api.Source {
		// upload table schemas
		fmt.Println("Uploading table schemas...")
		tablesJSONPath := filepath.Join(distDir, "tables.json")
		err = uploadTableSchemas(ctx, c, teamName, pluginName, tablesJSONPath, pkgJSON)
		if err != nil {
			return fmt.Errorf("failed to upload table schemas: %w", err)
		}
	}

	// upload docs
	fmt.Println("Uploading docs...")
	docsDir := filepath.Join(distDir, "docs")
	err = uploadDocs(ctx, c, teamName, pluginName, docsDir, pkgJSON)
	if err != nil {
		return fmt.Errorf("failed to upload docs: %w", err)
	}

	// upload binaries
	fmt.Println("Uploading binaries...")
	for _, t := range pkgJSON.SupportedTargets {
		fmt.Printf("- Uploading %s_%s...\n", t.OS, t.Arch)
		err = uploadPluginBinary(ctx, c, teamName, pluginName, t.OS, t.Arch, path.Join(distDir, t.Path), pkgJSON)
		if err != nil {
			return fmt.Errorf("failed to upload binary: %w", err)
		}
	}

	// optional: mark plugin as draft=false
	finalize, err := cmd.Flags().GetBool("finalize")
	if err != nil {
		return err
	}

	if finalize {
		fmt.Println("Finalizing plugin version...")
		draft := false
		resp, err := c.UpdatePluginVersionWithResponse(ctx, teamName, pkgJSON.Kind, pluginName, pkgJSON.Version, cloudquery_api.UpdatePluginVersionJSONRequestBody{
			Draft: &draft,
		})
		if err != nil {
			return fmt.Errorf("failed to finalize plugin version: %w", err)
		}
		if resp.HTTPResponse.StatusCode > 299 {
			return errorFromHTTPResponse(resp.HTTPResponse, resp)
		}
		fmt.Println("Success!")
		fmt.Printf("%s/%s@%s is now available on the CloudQuery Hub.\n", teamName, pluginName, pkgJSON.Version)
		return nil
	}

	fmt.Println("Success!")
	fmt.Println("\nNote: this plugin version is marked as draft=true. You can preview and finalize it on the CloudQuery Hub, or run `cloudquery plugin publish` with the --finalize flag.")

	return nil
}

func createNewPluginDraftVersion(ctx context.Context, c *cloudquery_api.ClientWithResponses, teamName, pluginName string, pkgJSON PackageJSONV1) error {
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
		err := errorFromHTTPResponse(resp.HTTPResponse, resp)
		if resp.HTTPResponse.StatusCode == http.StatusForbidden {
			return fmt.Errorf("%w. Hint: You may need to create the plugin first", err)
		}
		return err
	}
	return nil
}

func uploadTableSchemas(ctx context.Context, c *cloudquery_api.ClientWithResponses, teamName, pluginName, tablesJSONPath string, pkgJSON PackageJSONV1) error {
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
		return errorFromHTTPResponse(resp.HTTPResponse, resp)
	}
	return nil
}

func uploadDocs(ctx context.Context, c *cloudquery_api.ClientWithResponses, teamName, pluginName, docsDir string, pkgJSON PackageJSONV1) error {
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
		contentStr := normalizeContent(string(content))
		pages = append(pages, cloudquery_api.PluginDocsPageCreate{
			Content: contentStr,
			Name:    strings.TrimSuffix(dirEntry.Name(), fileExt),
		})
	}
	body := cloudquery_api.CreatePluginVersionDocsJSONRequestBody{
		Pages: pages,
	}
	resp, err := c.CreatePluginVersionDocsWithResponse(ctx, teamName, pkgJSON.Kind, pluginName, pkgJSON.Version, body)
	if err != nil {
		return fmt.Errorf("failed to upload docs: %w", err)
	}
	if resp.HTTPResponse.StatusCode > 299 {
		return errorFromHTTPResponse(resp.HTTPResponse, resp)
	}
	return nil
}

func uploadPluginBinary(ctx context.Context, c *cloudquery_api.ClientWithResponses, teamName, pluginName, goos, goarch, localPath string, pkgJSON PackageJSONV1) error {
	target := goos + "_" + goarch
	resp, err := c.UploadPluginAssetWithResponse(ctx, teamName, pkgJSON.Kind, pluginName, pkgJSON.Version, target)
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
	err = uploadFile(uploadURL, localPath)
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}
	return nil
}

func readPackageJSON(distDir string) (PackageJSONV1, error) {
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
