package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
)

type PackageJSONVersion struct {
	SchemaVersion int `json:"schema_version"`
}

type PackageJSONV1 struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Version string `json:"version"`

	PackageJSONPluginProperties
	PackageJSONAddonProperties
}

type PackageJSONPluginProperties struct {
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

type PackageJSONAddonProperties struct {
	Type        string   `json:"type"` // "addon"
	AddonType   string   `json:"addon_type"`
	AddonFormat string   `json:"addon_format"`
	PluginDeps  []string `json:"plugin_deps"`
	AddonDeps   []string `json:"addon_deps"`
	Doc         string   `json:"doc"`
}

func (p PackageJSONV1) IsPlugin() bool {
	return p.PackageJSONAddonProperties.Type == "" && p.PackageJSONPluginProperties.Kind != ""
}

func (p PackageJSONV1) IsAddon() bool {
	return p.PackageJSONAddonProperties.Type == "addon" && p.PackageJSONPluginProperties.Kind == ""
}

func errorFromHTTPResponse(httpResp *http.Response, resp any) error {
	fields := make(map[string]any)
	el := reflect.ValueOf(resp).Elem()
	for i := 0; i < el.NumField(); i++ {
		f := el.Field(i)
		fields[el.Type().Field(i).Name] = f.Interface()
	}
	for k, v := range fields {
		if !strings.HasPrefix(k, "JSON") || v == nil || reflect.ValueOf(v).Elem().Kind() != reflect.Struct {
			continue
		}
		msg := reflect.ValueOf(v).Elem().FieldByName("Message")
		if msg.IsValid() {
			return fmt.Errorf("%s: %s", strings.TrimPrefix(k, "JSON"), msg.String())
		}
	}
	return fmt.Errorf("error code: %v", httpResp.StatusCode)
}

func uploadFile(uploadURL, localPath string) error {
	file, err := os.Open(localPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	req, err := http.NewRequest(http.MethodPut, uploadURL, file)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return fmt.Errorf("failed to read response body: %w", readErr)
		}
		return fmt.Errorf("status %s: %s", resp.Status, body)
	}
	return nil
}

func normalizeContent(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.ReplaceAll(s, "\r", "\n")
	return s
}

func readPackageJSON(distDir string) (PackageJSONV1, error) {
	v := PackageJSONVersion{}
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
