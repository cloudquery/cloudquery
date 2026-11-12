package platform

import (
	"errors"
	"fmt"
	"os"
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
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, tokenFileMode)
	if err != nil {
		return fmt.Errorf("failed to open file %q for writing: %w", filePath, err)
	}
	if err := file.Chmod(tokenFileMode); err != nil {
		file.Close()
		return fmt.Errorf("failed to set permissions on %q: %w", filePath, err)
	}
	if _, err := file.WriteString(data); err != nil {
		file.Close()
		return fmt.Errorf("failed to write data to %q: %w", filePath, err)
	}
	return file.Close()
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
