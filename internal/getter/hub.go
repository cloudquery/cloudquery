package getter

import (
	"fmt"
	"net/url"
	"strings"
)

type HubDetector struct {
}

func (h HubDetector) Detect(src, _ string) (string, bool, error) {
	if len(src) == 0 {
		return "", false, nil
	}
	return h.detectHTTP(fmt.Sprintf("github.com/cloudquery-policies/%s", src))
}

func (h HubDetector) detectHTTP(src string) (string, bool, error) {
	parts := strings.Split(src, "/")
	if len(parts) < 3 {
		return "", false, fmt.Errorf(
			"CloudQuery Hub URLs should be <policy-name>")
	}

	urlStr := fmt.Sprintf("https://%s", strings.Join(parts[:3], "/"))
	_url, err := url.Parse(urlStr)
	if err != nil {
		return "", true, fmt.Errorf("error parsing GitHub URL: %s", err)
	}

	if !strings.HasSuffix(_url.Path, ".git") {
		_url.Path += ".git"
	}

	if len(parts) > 3 {
		_url.Path += "//" + strings.Join(parts[3:], "/")
	}

	return "git::" + _url.String(), true, nil
}
