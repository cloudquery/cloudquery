package publish

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	cqapiauth "github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery-api-go/config"
	"github.com/cloudquery/cloudquery/cli/internal/hub"
	"github.com/cloudquery/cloudquery/cli/internal/team"
)

type ManifestJSONV1 struct {
	Type        string `json:"type"` // always "addon"
	TeamName    string `json:"team_name"`
	AddonName   string `json:"addon_name"`
	AddonType   string `json:"addon_type"`
	AddonFormat string `json:"addon_format"` // unused

	Message   string `json:"message"`
	PathToZip string `json:"path"`
	PathToDoc string `json:"doc"`

	PluginDeps []string `json:"plugin_deps"`
	AddonDeps  []string `json:"addon_deps"`
}

func ReadManifestJSON(manifestPath string) (ManifestJSONV1, error) {
	v := SchemaVersion{}
	b, err := os.ReadFile(manifestPath)
	if err != nil {
		return ManifestJSONV1{}, err
	}

	if err := json.Unmarshal(b, &v); err != nil {
		return ManifestJSONV1{}, err
	}
	if v.SchemaVersion != 1 {
		return ManifestJSONV1{}, errors.New("unsupported schema version. This CLI version only supports manifest.json v1. Try upgrading your CloudQuery CLI version")
	}

	manifest := ManifestJSONV1{}
	if err := json.Unmarshal(b, &manifest); err != nil {
		return ManifestJSONV1{}, err
	}
	return manifest, nil
}

func CreateNewAddonDraftVersion(ctx context.Context, c *cloudquery_api.ClientWithResponses, manifest ManifestJSONV1, version, manifestDir, zipPath string) error {
	if manifest.PluginDeps == nil {
		manifest.PluginDeps = []string{}
	}
	if manifest.AddonDeps == nil {
		manifest.AddonDeps = []string{}
	}
	body := cloudquery_api.CreateAddonVersionJSONRequestBody{
		AddonDeps:  &manifest.AddonDeps,
		PluginDeps: &manifest.PluginDeps,
	}

	if manifest.PathToDoc != "" {
		absDocFile := filepath.Join(manifestDir, manifest.PathToDoc)
		b, err := os.ReadFile(absDocFile)
		if err != nil {
			return fmt.Errorf("failed to read doc file: %w", err)
		}
		body.Doc, err = processDocumentImages(ctx, c, manifest.TeamName, filepath.Dir(absDocFile), string(b))
		if err != nil {
			return fmt.Errorf("failed to process doc images: %w", err)
		}
	}

	if manifest.Message != "" {
		if strings.HasPrefix(manifest.Message, "@") {
			messageFile := filepath.Join(manifestDir, strings.TrimPrefix(manifest.Message, "@"))
			messageBytes, err := os.ReadFile(messageFile)
			if err != nil {
				return fmt.Errorf("failed to read message file: %w", err)
			}
			body.Message = string(messageBytes)
		} else {
			body.Message = manifest.Message
		}
	}

	f, err := os.Open(zipPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()
	s := sha256.New()
	if _, err := io.Copy(s, f); err != nil {
		return fmt.Errorf("failed to calculate checksum: %w", err)
	}
	body.Checksum = fmt.Sprintf("%x", s.Sum(nil))

	resp, err := c.CreateAddonVersionWithResponse(ctx, manifest.TeamName, cloudquery_api.AddonType(manifest.AddonType), manifest.AddonName, version, body)
	if err != nil {
		return fmt.Errorf("failed to create addon version: %w", err)
	}
	if resp.HTTPResponse.StatusCode > 299 {
		err := hub.ErrorFromHTTPResponse(resp.HTTPResponse, resp)
		if resp.HTTPResponse.StatusCode == http.StatusForbidden {
			return fmt.Errorf("%w. Hint: You may need to create the addon first", err)
		}
		return err
	}
	return nil
}

func UploadAddon(ctx context.Context, c *cloudquery_api.ClientWithResponses, manifest ManifestJSONV1, version, localPath string) error {
	resp, err := c.UploadAddonAssetWithResponse(ctx, manifest.TeamName, cloudquery_api.AddonType(manifest.AddonType), manifest.AddonName, version)
	if err != nil {
		return fmt.Errorf("failed to upload addon: %w", err)
	}
	if resp.HTTPResponse.StatusCode > 299 {
		msg := fmt.Sprintf("failed to upload addon: %s", resp.HTTPResponse.Status)
		switch {
		case resp.JSON403 != nil:
			msg = fmt.Sprintf("%s: %s", msg, resp.JSON403.Message)
		case resp.JSON401 != nil:
			msg = fmt.Sprintf("%s: %s", msg, resp.JSON401.Message)
		}
		return fmt.Errorf(msg)
	}
	if resp.JSON201 == nil {
		return fmt.Errorf("upload response is nil, failed to upload addon")
	}
	uploadURL := resp.JSON201.Url

	if err := hub.UploadFile(uploadURL, localPath); err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}
	return nil
}

func GetTeamForAnyToken(ctx context.Context, c *cloudquery_api.ClientWithResponses, tokenType cqapiauth.TokenType) (string, error) {
	switch tokenType {
	case cqapiauth.BearerToken:
		currentTeam, err := config.GetValue("team")
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("failed to get team from config: %w", err)
		}
		if currentTeam == "" {
			return "", fmt.Errorf("team is required. Hint: use `cloudquery switch` to set a team")
		}
		return currentTeam, nil
	case cqapiauth.APIKey:
		teams, err := team.NewClientFromAPI(c).ListAllTeams(ctx)
		if err != nil {
			return "", err
		}
		switch l := len(teams); l {
		case 0:
			return "", errors.New("api key has no assigned team")
		case 1:
			return teams[0], nil
		default:
			return "", fmt.Errorf("api key has more than one team: %s", strings.Join(teams, ", "))
		}
	default:
		return "", fmt.Errorf("unknown token type %v", tokenType)
	}
}

func GetAddonMetadata(ctx context.Context, c *cloudquery_api.ClientWithResponses, currentTeam, addonTeam, addonType, addonName, addonVersion string) (location, checksum string, retErr error) {
	aj := "application/json"
	resp, err := c.DownloadAddonAssetByTeamWithResponse(ctx, currentTeam, addonTeam, cloudquery_api.AddonType(addonType), addonName, addonVersion, &cloudquery_api.DownloadAddonAssetByTeamParams{Accept: &aj})
	if err != nil {
		return "", "", fmt.Errorf("failed to get team addon metadata: %w", err)
	}
	if resp.StatusCode() > 299 || resp.JSON200 == nil {
		return "", "", fmt.Errorf("failed to read team addon metadata: %w", hub.ErrorFromHTTPResponse(resp.HTTPResponse, resp))
	}
	return resp.JSON200.Location, resp.JSON200.Checksum, nil
}

func DownloadAddonFromResponse(res *http.Response, expectedChecksum, targetDir string) (retErr error) {
	var (
		fileWriter io.WriteCloser
		size       int64
		err        error
	)

	switch targetDir {
	case "-":
		fileWriter = os.Stdout
	default:
		zipPath := filepath.Join(targetDir, path.Base(res.Request.URL.Path))
		if st, err := os.Stat(zipPath); err == nil {
			if st.IsDir() {
				return fmt.Errorf("file %s already exists: is a directory", zipPath)
			}
			return fmt.Errorf("file %s already exists", zipPath)
		}

		f, err := os.Create(zipPath)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		fileWriter = f

		defer func() {
			if retErr != nil {
				_ = os.Remove(zipPath)
				return
			}
			fmt.Fprintf(os.Stderr, "Wrote %d bytes to %s\n", size, zipPath)
		}()
	}

	shaWriter := sha256.New()
	w := io.MultiWriter(fileWriter, shaWriter)
	if size, err = io.Copy(w, res.Body); err != nil {
		_ = fileWriter.Close()
		return fmt.Errorf("failed to write: %w", err)
	}
	if err := fileWriter.Close(); err != nil {
		return fmt.Errorf("failed to close: %w", err)
	}
	if err := res.Body.Close(); err != nil {
		return fmt.Errorf("failed to close response body: %w", err)
	}

	writtenChecksum := fmt.Sprintf("%x", shaWriter.Sum(nil))
	if writtenChecksum != expectedChecksum {
		return fmt.Errorf("checksum mismatch: expected %s, got %s", expectedChecksum, writtenChecksum)
	}

	return nil
}
