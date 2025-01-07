package hub

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
)

type HubPluginRef struct {
	TeamName string
	Kind     string
	Name     string
	Version  string
}

func (h HubPluginRef) String() string {
	return fmt.Sprintf("%s/%s/%s@%s", h.TeamName, h.Kind, h.Name, h.Version)
}

func ParseHubPluginRef(ref string) (*HubPluginRef, error) {
	versionParts := strings.Split(ref, "@")
	if len(versionParts) != 2 {
		return nil, errors.New("invalid plugin version: Must be in format <team_name>/<kind>/<plugin_name>@<version>")
	}
	if !strings.HasPrefix(versionParts[1], "v") {
		return nil, errors.New("invalid plugin version: version must start with 'v'")
	}

	parts := strings.Split(versionParts[0], "/")
	if len(parts) != 3 {
		return nil, errors.New("invalid plugin name: Must be in format <team_name>/<kind>/<plugin_name>@<version>")
	}

	if parts[1] != "source" && parts[1] != "destination" {
		return nil, errors.New("invalid plugin kind: must be either 'source' or 'destination'")
	}

	return &HubPluginRef{
		TeamName: parts[0],
		Kind:     parts[1],
		Name:     parts[2],
		Version:  versionParts[1],
	}, nil
}

func ErrorFromHTTPResponse(httpResp *http.Response, resp any) error {
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

func UploadFile(uploadURL, localPath string) error {
	return UploadFileWithContentType(context.Background(), uploadURL, localPath, "application/octet-stream")
}

func UploadFileWithContentType(ctx context.Context, uploadURL, localPath, contentType string) error {
	file, err := os.Open(localPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	return UploadReaderWithContentType(ctx, uploadURL, file, contentType)
}

func UploadReaderWithContentType(ctx context.Context, uploadURL string, reader io.Reader, contentType string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, uploadURL, reader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", contentType)

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

func NormalizeContent(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.ReplaceAll(s, "\r", "\n")
	return s
}
