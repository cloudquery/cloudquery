package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/adrg/xdg"
)

const (
	firebaseAPIKey = "AIzaSyCxsrwjABEF-dWLzUqmwiL-ct02cnG9GCs"
	tokenURL       = "https://securetoken.googleapis.com/v1/token?key=" + firebaseAPIKey

	EnvVarCloudQueryAPIKey = "CLOUDQUERY_API_KEY"
)

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	UserID       string `json:"user_id"`
	ProjectID    string `json:"project_id"`
}

func GetToken() (string, error) {
	token := os.Getenv(EnvVarCloudQueryAPIKey)
	if token == "" {
		refreshToken, err := readRefreshToken()
		if err != nil {
			return "", fmt.Errorf("%w. Hint: You may need to run `cloudquery login` or set %s", err, EnvVarCloudQueryAPIKey)
		}
		if refreshToken == "" {
			return "", fmt.Errorf("could not find authentication token. Hint: You may need to run `cloudquery login` or set %s", EnvVarCloudQueryAPIKey)
		}
		token, err = getIDToken(refreshToken)
		if err != nil {
			return "", fmt.Errorf("failed to sign in with custom token: %w", err)
		}
	}
	return token, nil
}

func removeRefreshToken() error {
	tokenFilePath, err := xdg.DataFile("cloudquery/token")
	if err != nil {
		return fmt.Errorf("can't determine a proper location for token file: %w", err)
	}
	if err := os.RemoveAll(tokenFilePath); err != nil {
		return fmt.Errorf("failed to remove token file %q: %w", tokenFilePath, err)
	}
	return nil
}

func SaveRefreshToken(refreshToken string) error {
	tokenFilePath, err := xdg.DataFile("cloudquery/token")
	if err != nil {
		return fmt.Errorf("can't determine a proper location for token file: %w", err)
	}
	tokenFile, err := os.OpenFile(tokenFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return fmt.Errorf("can't open token file %q for writing: %w", tokenFilePath, err)
	}
	defer func() {
		e := tokenFile.Close()
		if err == nil && e != nil {
			err = fmt.Errorf("can't close token file %q after writing: %w", tokenFilePath, e)
		}
	}()
	if _, err = tokenFile.WriteString(refreshToken); err != nil {
		return fmt.Errorf("failed to write token to %q: %w", tokenFilePath, err)
	}
	return nil
}

func getIDToken(refreshToken string) (string, error) {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)

	resp, err := http.PostForm(tokenURL, data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return "", fmt.Errorf("failed to read response body: %w", readErr)
		}
		return "", fmt.Errorf("failed to refresh token: %s: %s", resp.Status, body)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	tokenResp, err := parseToken(body)
	if err != nil {
		return "", err
	}
	err = SaveRefreshToken(tokenResp.RefreshToken)
	if err != nil {
		return "", fmt.Errorf("failed to save refresh token: %w", err)
	}

	return tokenResp.IDToken, nil
}
func parseToken(response []byte) (tokenResponse, error) {
	var tr tokenResponse
	err := json.Unmarshal(response, &tr)
	if err != nil {
		return tokenResponse{}, err
	}
	return tr, nil
}

func readRefreshToken() (string, error) {
	tokenFilePath, err := xdg.DataFile("cloudquery/token")
	if err != nil {
		return "", fmt.Errorf("failed to get token file path: %w", err)
	}
	b, err := os.ReadFile(tokenFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to read token file: %w", err)
	}
	return strings.TrimSpace(string(b)), nil
}
