package platform

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
	"github.com/cloudquery/cloudquery-api-go/config"
)

const (
	tokenStorePath = "cloudquery/platform_token"
	tokenFileMode  = 0600
)

func IsPlatformToken(token string) bool {
	return strings.HasPrefix(token, cqpdPrefix)
}

func SavePlatformToken(token string) error {
	if !IsPlatformToken(token) {
		return errors.New("not a platform token")
	}
	filePath, err := xdg.DataFile(tokenStorePath)
	if err != nil {
		return fmt.Errorf("failed to get file path: %w", err)
	}
	return writeOwnerOnlyFile(filePath, token)
}

func writeOwnerOnlyFile(filePath string, data string) error {
	tempFile, err := os.CreateTemp(filepath.Dir(filePath), filepath.Base(filePath)+".tmp-*")
	if err != nil {
		return fmt.Errorf("failed to create temporary file for %q: %w", filePath, err)
	}
	tempPath := tempFile.Name()
	if err := writeAndCloseOwnerOnly(tempFile, data); err != nil {
		os.Remove(tempPath)
		return fmt.Errorf("failed to write data to %q: %w", filePath, err)
	}
	if err := os.Rename(tempPath, filePath); err != nil {
		os.Remove(tempPath)
		return fmt.Errorf("failed to replace %q: %w", filePath, err)
	}
	return nil
}

func writeAndCloseOwnerOnly(file *os.File, data string) error {
	err := file.Chmod(tokenFileMode)
	if err == nil {
		_, err = file.WriteString(data)
	}
	if err == nil {
		err = file.Sync()
	}
	return errors.Join(err, file.Close())
}

func ReadPlatformToken() string {
	token, err := config.ReadDataString(tokenStorePath)
	if err != nil || !IsPlatformToken(token) {
		return ""
	}
	return token
}

func RemovePlatformToken() error {
	return config.DeleteDataString(tokenStorePath)
}
