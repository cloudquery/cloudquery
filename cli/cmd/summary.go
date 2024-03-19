package cmd

import (
	"encoding/json"
	"os"
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
	err := checkFile(filename)
	if err != nil {
		return err
	}

	for _, summary := range summaries {
		dataBytes, err := json.Marshal(summary)
		if err != nil {
			return err
		}
		dataBytes = append(dataBytes, []byte("\n")...)
		err = appendToFile(filename, dataBytes)
		if err != nil {
			return err
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

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}
