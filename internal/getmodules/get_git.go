package getmodules

import (
	"net/url"
	"os/exec"

	"github.com/hashicorp/go-getter"
)

type GitGetter struct {
	getter.GitGetter
	GoGitGetter
}

// Get downloads the given URL into the given directory. This always
// assumes that we're updating and gets the latest version that it can.
//
// The directory may already exist (if we're updating). If it is in a
// format that isn't understood, an error should be returned. Get shouldn't
// simply nuke the directory.
func (g *GitGetter) Get(dst string, u *url.URL) error {
	if _, err := exec.LookPath("git"); err != nil {
		return g.GitGetter.Get(dst, u)
	}
	return g.GoGitGetter.Get(dst, u)
}

// GetFile downloads the give URL into the given path. The URL must
// reference a single file. If possible, the Getter should check if
// the remote end contains the same file and no-op this operation.
func (g *GitGetter) GetFile(dst string, u *url.URL) error {
	if _, err := exec.LookPath("git"); err != nil {
		return g.GitGetter.GetFile(dst, u)
	}
	return g.GoGitGetter.GetFile(dst, u)
}

func (g *GitGetter) ClientMode(url *url.URL) (getter.ClientMode, error) {
	return getter.ClientModeDir, nil
}

func (g *GitGetter) SetClient(client *getter.Client) {
	g.GitGetter.SetClient(client)
	g.GoGitGetter.SetClient(client)
}
