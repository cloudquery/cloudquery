package cmd

import (
	"encoding/json"
	"os"
)

type syncSummary struct {
	CliVersion          string
	DestinationErrors   uint64
	DestinationName     string
	DestinationPath     string
	DestinationVersion  string
	DestinationWarnings uint64
	Resources           uint64
	SourceErrors        uint64
	SourceName          string
	SourcePath          string
	SourceVersion       string
	SourceWarnings      uint64
	SyncID              string
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
