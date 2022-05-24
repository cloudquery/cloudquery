package getter

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/cloudquery/cloudquery/internal/firebase"
)

var (
	repoToFirebasePath = map[string]string{
		"cloudquery-policies": "cloudquery",
	}
)

// GitHubDetector implements Detector to detect GitHub URLs and turn
// them into URLs that the Git Getter can understand.
type GitHubDetector struct{}

func (d *GitHubDetector) Detect(src, _ string) (string, bool, error) {
	if len(src) == 0 {
		return "", false, nil
	}

	if strings.HasPrefix(src, "github.com/") {
		return d.detectHTTP(src)
	}

	return "", false, nil
}

func (*GitHubDetector) detectHTTP(src string) (string, bool, error) {
	parts := strings.Split(src, "/")
	if len(parts) < 3 {
		return "", false, fmt.Errorf(
			"GitHub URLs should be github.com/username/repo")
	}

	urlStr := fmt.Sprintf("https://%s", strings.Join(parts[:3], "/"))
	_url, err := url.Parse(urlStr)
	if err != nil {
		return "", true, fmt.Errorf("error parsing GitHub URL: %s", err)
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
