package getmodules

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"

	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"

	"github.com/go-git/go-git/v5/storage/memory"

	"github.com/go-git/go-git/v5"

	"github.com/hashicorp/go-getter"

	urlhelper "github.com/hashicorp/go-getter/helper/url"
	safetemp "github.com/hashicorp/go-safetemp"
)

// getter is our base getter; it regroups
// fields all getters have in common.
type getterCommon struct {
	client *getter.Client
}

func (g *getterCommon) SetClient(c *getter.Client) { g.client = c }

// Context tries to return the Context from the getter's
// client. otherwise, context.Background() is returned.
func (g *getterCommon) Context() context.Context {
	if g == nil || g.client == nil {
		return context.Background()
	}
	return g.client.Ctx
}

// GoGitGetter is a Getter implementation that will download a module from
// a git repository.
type GoGitGetter struct {
	getterCommon
}

var lsRemoteSymRefRegexp = regexp.MustCompile(`ref: refs/heads/([^\s]+).*`)

func (g *GoGitGetter) ClientMode(_ *url.URL) (getter.ClientMode, error) {
	return getter.ClientModeDir, nil
}

func (g *GoGitGetter) Get(dst string, u *url.URL) error {
	// Extract some query parameters we use
	var ref string
	depth := 0 // 0 means "don't use shallow clone"
	q := u.Query()
	if len(q) > 0 {
		ref = q.Get("ref")
		q.Del("ref")

		if n, err := strconv.Atoi(q.Get("depth")); err == nil {
			depth = n
		}
		q.Del("depth")

		// Copy the URL
		var newU = *u
		u = &newU
		u.RawQuery = q.Encode()
	}

	r, err := git.PlainOpen(dst)
	if errors.Is(err, git.ErrRepositoryNotExists) {
		if err := g.clone(g.Context(), dst, u, ref, depth); err != nil {
			return err
		}
		return nil
	}
	return g.update(g.Context(), r, ref, depth)
}

// GetFile for Git doesn't support updating at this time. It will download
// the file every time.
func (g *GoGitGetter) GetFile(dst string, u *url.URL) error {
	td, tdcloser, err := safetemp.Dir("", "getter")
	if err != nil {
		return err
	}
	defer tdcloser.Close()

	// Get the filename, and strip the filename from the URL so we can
	// just get the repository directly.
	filename := filepath.Base(u.Path)
	u.Path = filepath.ToSlash(filepath.Dir(u.Path))
	// Get the full repository
	if err := g.Get(td, u); err != nil {
		return err
	}

	// Copy the single file
	u, err = urlhelper.Parse(fmtFileURL(filepath.Join(td, filename)))
	if err != nil {
		return err
	}

	fg := &getter.FileGetter{Copy: true}
	return fg.GetFile(dst, u)
}

// gitCommitIDRegex is a pattern intended to match strings that seem
// "likely to be" git commit IDs, rather than named refs. This cannot be
// an exact decision because it's valid to name a branch or tag after a series
// of hexadecimal digits too.
//
// We require at least 7 digits here because that's the smallest size git
// itself will typically generate, and so it'll reduce the risk of false
// positives on short branch names that happen to also be "hex words".
var gitCommitIDRegex = regexp.MustCompile("^[0-9a-fA-F]{7,40}$")

func (g *GoGitGetter) checkout(r *git.Repository, ref string) error {
	wt, err := r.Worktree()
	if err != nil {
		return err
	}
	h, err := r.ResolveRevision(plumbing.Revision(ref))
	if err != nil {
		return err
	}
	return wt.Checkout(&git.CheckoutOptions{
		Hash: *h,
	})
}

func (g *GoGitGetter) clone(ctx context.Context, dst string, u *url.URL, ref string, depth int) error {
	originalRef := ref // we handle an unspecified ref differently than explicitly selecting the default branch below
	if ref == "" {
		ref = findRemoteDefaultBranch(u)
	}
	opts := &git.CloneOptions{
		URL:               u.String(),
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	}
	if depth > 0 {
		opts.Depth = depth
		opts.ReferenceName = findRemoteRef(u, ref)
	}
	r, err := git.PlainCloneContext(ctx, dst, false, opts)
	if err != nil {
		if depth > 0 && originalRef != "" {
			// If we're creating a shallow clone then the given ref must be
			// a named ref (branch or tag) rather than a commit directly.
			// We can't accurately recognize the resulting error here without
			// hard-coding assumptions about git's human-readable output, but
			// we can at least try a heuristic.
			if gitCommitIDRegex.MatchString(originalRef) {
				return fmt.Errorf("%w (note that setting 'depth' requires 'ref' to be a branch or tag name)", err)
			}
		}
		return err
	}

	if depth < 1 && originalRef != "" {
		// If we didn't add --depth and --branch above then we will now be
		// on the remote repository's default branch, rather than the selected
		// ref, so we'll need to fix that before we return.
		return g.checkout(r, originalRef)
	}
	return nil
}

func (g *GoGitGetter) update(ctx context.Context, r *git.Repository, ref string, depth int) error {
	wt, err := r.Worktree()
	if err != nil {
		return err
	}
	head, err := r.Head()
	if err != nil {
		return err
	}
	if head.Name() != plumbing.ReferenceName("refs/heads/"+ref) {
		if ref == "" {
			ref = findDefaultBranch(r).Short()
		}
		if err := g.checkout(r, ref); err != nil {
			return err
		}
	}

	err = wt.PullContext(ctx, &git.PullOptions{
		ReferenceName:     head.Name(),
		Depth:             depth,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Force:             false,
	})
	if errors.Is(err, git.NoErrAlreadyUpToDate) {
		return nil
	}
	return err
}

// findDefaultBranch checks the repo's origin remote for its default branch
// (generally "master"). "master" is returned if an origin default branch
// can't be determined.
func findDefaultBranch(r *git.Repository) plumbing.ReferenceName {
	refs, err := r.References()
	if err != nil {
		return plumbing.Master
	}
	for {
		b, err := refs.Next()
		if b == nil || err != nil {
			return plumbing.Master
		}
		if b.Name() == plumbing.HEAD {
			return b.Target()
		}
	}
}

// findRemoteDefaultBranch checks the remote repo's HEAD symref to return the remote repo's
// default branch. "master" is returned if no HEAD symref exists.
func findRemoteDefaultBranch(u *url.URL) string {
	rfs, err := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name:  "",
		URLs:  []string{u.String()},
		Fetch: nil,
	}).List(&git.ListOptions{
		Auth:            nil,
		InsecureSkipTLS: false,
		CABundle:        nil,
	})
	if err != nil {
		return "master"
	}
	for _, rf := range rfs {
		matches := lsRemoteSymRefRegexp.FindStringSubmatch(rf.String())
		if matches != nil {
			return matches[len(matches)-1]
		}
	}
	return "master"
}

func findRemoteRef(u *url.URL, refName string) plumbing.ReferenceName {
	rfs, err := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name:  "",
		URLs:  []string{u.String()},
		Fetch: nil,
	}).List(&git.ListOptions{
		Auth:            nil,
		InsecureSkipTLS: false,
		CABundle:        nil,
	})
	if err != nil {
		return plumbing.ReferenceName(refName)
	}
	for _, rf := range rfs {
		if rf.Name().String() == refName || rf.Name().Short() == refName {
			return rf.Name()
		}

	}
	return plumbing.ReferenceName(refName)
}

func fmtFileURL(path string) string {
	if runtime.GOOS == "windows" {
		// Make sure we're using "/" on Windows. URLs are "/"-based.
		path = filepath.ToSlash(path)
		return fmt.Sprintf("file://%s", path)
	}

	// Make sure that we don't start with "/" since we add that below.
	if path[0] == '/' {
		path = path[1:]
	}
	return fmt.Sprintf("file:///%s", path)
}
