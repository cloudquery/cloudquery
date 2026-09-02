package platform

import (
	"errors"
	"strings"

	"github.com/cloudquery/cloudquery-api-go/config"
)

const tokenStorePath = "cloudquery/platform_token"

func IsPlatformToken(token string) bool {
	return strings.HasPrefix(token, cqpdPrefix)
}

func SavePlatformToken(token string) error {
	if !IsPlatformToken(token) {
		return errors.New("not a platform token")
	}
	return config.SaveDataString(tokenStorePath, token)
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
