package getter

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/cloudquery/cloudquery/internal/firebase"
)

type HubDetector struct {
}

func (h HubDetector) Detect(src, pwd string) (string, bool, error) {
	if len(src) == 0 {
		return "", false, nil
	}
	fileDetector := fileDetector{}
	if _, ok, _ := fileDetector.Detect(src, pwd); ok {
		return "", false, nil
	}
	return h.detectHTTP(fmt.Sprintf("github.com/cloudquery-policies/%s", src))
}

func (HubDetector) detectHTTP(src string) (string, bool, error) {
	parts := strings.Split(src, "/")
	if len(parts) < 3 {
		return "", false, fmt.Errorf(
			"CloudQuery Hub URLs should be <policy-name>")
	}

	urlStr := fmt.Sprintf("https://%s", strings.Join(parts[:3], "/"))
	_url, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return "", true, fmt.Errorf("error parsing GitHub URL: %s", err)
	}
	resp, err := http.Get(_url.String())
	if err != nil {
		return "", false, fmt.Errorf("failed to check if policy in hub: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return "", false, nil
	}

	if !_url.Query().Has("ref") {
		if err := addLatestTag(_url, parts[1], parts[2]); err != nil {
			return "", false, err
		}
	}

	if !strings.HasSuffix(_url.Path, ".git") {
		_url.Path += ".git"
	}

	if len(parts) > 3 {
		_url.Path += "//" + strings.Join(parts[3:], "/")
	}

	return "git::" + _url.String(), true, nil
}

func addLatestTag(_url *url.URL, owner, repo string) error {
	client := firebase.New(firebase.CloudQueryRegistryURL)
	org, ok := repoToFirebasePath[owner]
	if !ok {
		org = owner
	}
	latest, err := client.GetLatestPolicyRelease(context.Background(), org, repo)

	if err != nil {
		return fmt.Errorf("failed to find latest version: %w", err)
	}

	if latest == "" {
		return nil
	}

	q := _url.Query()
	q.Add("ref", latest)
	_url.RawQuery = q.Encode()
	return nil
}
