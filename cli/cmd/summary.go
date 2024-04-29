package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type syncSummary struct {
	CliVersion          string `json:"cli_version"`
	DestinationErrors   uint64 `json:"destination_errors"`
	DestinationName     string `json:"destination_name"`
	DestinationPath     string `json:"destination_path"`
	DestinationVersion  string `json:"destination_version"`
	DestinationWarnings uint64 `json:"destination_warnings"`
	Resources           uint64 `json:"resources"`
	SourceErrors        uint64 `json:"source_errors"`
	SourceName          string `json:"source_name"`
	SourcePath          string `json:"source_path"`
	SourceVersion       string `json:"source_version"`
	SourceWarnings      uint64 `json:"source_warnings"`
	SyncID              string `json:"sync_id"`
}

func persistSummary(filename string, summaries []syncSummary) error {
	// if filename is not specified then we don't need to persist the summary and can return
	if len(filename) == 0 { // just ofr test
		return nil
	}
	err := checkFilePath(filename)
	if err != nil {
		return fmt.Errorf("failed to validate summary file path: %w", err)
	}
	for _, summary := range summaries {
		dataBytes, err := json.Marshal(summary)
		if err != nil {
			return err
		}
		dataBytes = append(dataBytes, []byte("\n")...)
		err = appendToFile(filename, dataBytes)
		if err != nil {
			return fmt.Errorf("failed to append summary to file: %w", err)
		}
	}
	return nil
}

func appendToFile(fileName string, data []byte) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if _, err := f.Write(data); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

func checkFilePath(filename string) error {
	dirPath := filepath.Dir(filename)
	return os.MkdirAll(dirPath, 0755)
}
