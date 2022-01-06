package getter

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-github/v35/github"
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

func (d *GitHubDetector) detectHTTP(src string) (string, bool, error) {
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

	client := github.NewClient(nil)
	tags, _, err := client.Repositories.ListTags(context.Background(), owner, repo, &github.ListOptions{
		Page:    0,
		PerPage: 1,
	})
	if err != nil {
		return fmt.Errorf("failed to find tags: %w", err)
	}
	if len(tags) == 0 {
		return nil
	}
	q := _url.Query()
	q.Add("ref", tags[0].GetName())
	_url.RawQuery = q.Encode()
	return nil
}
