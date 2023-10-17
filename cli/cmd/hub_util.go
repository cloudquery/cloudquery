package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
)

type SchemaVersion struct {
	SchemaVersion int `json:"schema_version"`
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
