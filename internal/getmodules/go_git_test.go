package getmodules

import (
	"context"
	"net/url"
	"testing"
)

func TestGoGitFindRemoteBranch(t *testing.T) {

	g := GoGit{}
	u, _ := url.Parse("https://github.com/hashicorp/go-getter")
	g.FindRemoteDefaultBranch(context.TODO(), u)
}
