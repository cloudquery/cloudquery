// Tests were taken from the original go-getter git_test to validate we support same functionality and expectations
// TODO: Missing tests: SSH and Auth git getting.
package getmodules

import (
	"bytes"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	urlhelper "github.com/hashicorp/go-getter/helper/url"

	"github.com/hashicorp/go-getter"
)

var testHasGit bool

func init() {
	if _, err := exec.LookPath("git"); err == nil {
		testHasGit = true
	}
}

func TestGitGetter_impl(t *testing.T) {
	var _ getter.Getter = new(GoGitGetter)
}

func TestGitGetter(t *testing.T) {
	if !testHasGit {
		t.Skip("git not found, skipping")
	}

	g := new(GoGitGetter)
	dst := tempDir(t)

	repo := testGitRepo(t, "basic")
	repo.commitFile("foo.txt", "hello")

	// With a dir that doesn't exist
	if err := g.Get(dst, repo.url); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Verify the main file exists
	mainPath := filepath.Join(dst, "foo.txt")
	if _, err := os.Stat(mainPath); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestGitGetter_commitID(t *testing.T) {
	if !testHasGit {
		t.Skip("git not found, skipping")
	}

	g := new(GoGitGetter)
	dst := tempDir(t)

	// We're going to create different content on the main branch vs.
	// another branch here, so that below we can recognize if we
	// correctly cloned the commit actually requested (from the
	// "other branch"), not the one at HEAD.
	repo := testGitRepo(t, "commit_id")
	repo.git("checkout", "-b", "main-branch")
	repo.commitFile("wrong.txt", "Nope")
	repo.git("checkout", "-b", "other-branch")
	repo.commitFile("hello.txt", "Yep")
	commitID, err := repo.latestCommit()
	if err != nil {
		t.Fatal(err)
	}
	// Return to the main branch so that HEAD of this repository
	// will be that, rather than "test-branch".
	repo.git("checkout", "main-branch")

	q := repo.url.Query()
	q.Add("ref", commitID)
	repo.url.RawQuery = q.Encode()

	t.Logf("Getting %s", repo.url)
	if err := g.Get(dst, repo.url); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Verify the main file exists
	mainPath := filepath.Join(dst, "hello.txt")
	if _, err := os.Stat(mainPath); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Get again should work
	if err := g.Get(dst, repo.url); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Verify the main file exists
	mainPath = filepath.Join(dst, "hello.txt")
	if _, err := os.Stat(mainPath); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestGitGetter_remoteWithoutMaster(t *testing.T) {
	if !testHasGit {
		t.Log("git not found, skipping")
		t.Skip()
	}

	g := new(GoGitGetter)
	dst := tempDir(t)

	repo := testGitRepo(t, "branch")
	repo.git("checkout", "-b", "test-branch")
	repo.commitFile("branch.txt", "branch")

	q := repo.url.Query()
	repo.url.RawQuery = q.Encode()

	if err := g.Get(dst, repo.url); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Verify the main file exists
	mainPath := filepath.Join(dst, "branch.txt")
	if _, err := os.Stat(mainPath); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Get again should work
	if err := g.Get(dst, repo.url); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Verify the main file exists
	mainPath = filepath.Join(dst, "branch.txt")
	if _, err := os.Stat(mainPath); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestGitGetter_shallowClone(t *testing.T) {
	if !testHasGit {
		t.Log("git not found, skipping")
		t.Skip()
	}

	g := new(GoGitGetter)
	dst := tempDir(t)

	repo := testGitRepo(t, "upstream")
	repo.commitFile("upstream.txt", "0")
	repo.commitFile("upstream.txt", "1")

	// Specify a clone depth of 1
	q := repo.url.Query()
	q.Add("depth", "1")
	repo.url.RawQuery = q.Encode()

	if err := g.Get(dst, repo.url); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Assert rev-list count is '1'
	cmd := exec.Command("git", "rev-list", "HEAD", "--count")
	cmd.Dir = dst
	b, err := cmd.Output()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	out := strings.TrimSpace(string(b))
	if out != "1" {
		t.Fatalf("expected rev-list count to be '1' but got %v", out)
	}
}

func TestGitGetter_shallowCloneWithTag(t *testing.T) {
	if !testHasGit {
		t.Log("git not found, skipping")
		t.Skip()
	}

	g := new(GoGitGetter)
	dst := tempDir(t)

	repo := testGitRepo(t, "upstream")
	repo.commitFile("v1.0.txt", "0")
	repo.git("tag", "v1.0")
	repo.commitFile("v1.1.txt", "1")

	// Specify a clone depth of 1 with a tag
	q := repo.url.Query()
	q.Add("ref", "v1.0")
	q.Add("depth", "1")
	repo.url.RawQuery = q.Encode()

	if err := g.Get(dst, repo.url); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Assert rev-list count is '1'
	cmd := exec.Command("git", "rev-list", "HEAD", "--count")
	cmd.Dir = dst
	b, err := cmd.Output()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	out := strings.TrimSpace(string(b))
	if out != "1" {
		t.Fatalf("expected rev-list count to be '1' but got %v", out)
	}

	// Verify the v1.0 file exists
	mainPath := filepath.Join(dst, "v1.0.txt")
	if _, err := os.Stat(mainPath); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Verify the v1.1 file does not exists
	mainPath = filepath.Join(dst, "v1.1.txt")
	if _, err := os.Stat(mainPath); err == nil {
		t.Fatalf("expected v1.1 file to not exist")
	}
}

func TestGitGetter_shallowCloneWithCommitID(t *testing.T) {
	if !testHasGit {
		t.Log("git not found, skipping")
		t.Skip()
	}

	g := new(GoGitGetter)
	dst := tempDir(t)

	repo := testGitRepo(t, "upstream")
	repo.commitFile("v1.0.txt", "0")
	repo.git("tag", "v1.0")
	repo.commitFile("v1.1.txt", "1")

	commitID, err := repo.latestCommit()
	if err != nil {
		t.Fatal(err)
	}

	// Specify a clone depth of 1 with a naked commit ID
	// This is intentionally invalid: shallow clone always requires a named ref.
	q := repo.url.Query()
	q.Add("ref", commitID[:8])
	q.Add("depth", "1")
	repo.url.RawQuery = q.Encode()

	t.Logf("Getting %s", repo.url)
	err = g.Get(dst, repo.url)
	if err == nil {
		t.Fatalf("success; want error")
	}
	// We use a heuristic to generate an extra hint in the error message if
	// it looks like the user was trying to combine ref=COMMIT with depth.
	if got, want := err.Error(), "(note that setting 'depth' requires 'ref' to be a branch or tag name)"; !strings.Contains(got, want) {
		t.Errorf("missing error message hint\ngot: %s\nwant substring: %s", got, want)
	}
}

func TestGitGetter_branchUpdate(t *testing.T) {
	if !testHasGit {
		t.Skip("git not found, skipping")
	}

	g := new(GoGitGetter)
	dst := tempDir(t)

	// First setup the state with a fresh branch
	repo := testGitRepo(t, "branch-update")
	repo.git("checkout", "-b", "test-branch")
	repo.commitFile("branch.txt", "branch")

	// Get the "test-branch" branch
	q := repo.url.Query()
	q.Add("ref", "test-branch")
	repo.url.RawQuery = q.Encode()
	if err := g.Get(dst, repo.url); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Verify the main file exists
	mainPath := filepath.Join(dst, "branch.txt")
	if _, err := os.Stat(mainPath); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Commit an update to the branch
	repo.commitFile("branch-update.txt", "branch-update")

	// Get again should work
	if err := g.Get(dst, repo.url); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Verify the main file exists
	mainPath = filepath.Join(dst, "branch-update.txt")
	if _, err := os.Stat(mainPath); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestGitGetter_tag(t *testing.T) {
	if !testHasGit {
		t.Skip("git not found, skipping")
	}

	g := new(GoGitGetter)
	dst := tempDir(t)

	repo := testGitRepo(t, "tag")
	repo.commitFile("tag.txt", "tag")
	repo.git("tag", "v1.0")

	q := repo.url.Query()
	q.Add("ref", "v1.0")
	repo.url.RawQuery = q.Encode()

	if err := g.Get(dst, repo.url); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Verify the main file exists
	mainPath := filepath.Join(dst, "tag.txt")
	if _, err := os.Stat(mainPath); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Get again should work
	if err := g.Get(dst, repo.url); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Verify the main file exists
	mainPath = filepath.Join(dst, "tag.txt")
	if _, err := os.Stat(mainPath); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestGitGetter_GetFile(t *testing.T) {
	if !testHasGit {
		t.Skip("git not found, skipping")
	}

	g := new(GoGitGetter)
	dst := tempTestFile(t)
	defer os.RemoveAll(filepath.Dir(dst))

	repo := testGitRepo(t, "file")
	repo.commitFile("file.txt", "hello")

	// Download the file
	repo.url.Path = strings.Join([]string{repo.url.Path, "file.txt"}, "/")
	t.Log(repo.url.String())
	if err := g.GetFile(dst, repo.url); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Verify the main file exists
	if _, err := os.Stat(dst); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestGitGetter_submodule(t *testing.T) {
	if !testHasGit {
		t.Skip("git not found, skipping")
	}

	g := new(GoGitGetter)
	dst := tempDir(t)

	relpath := func(basepath, targpath string) string {
		relpath, err := filepath.Rel(basepath, targpath)
		if err != nil {
			t.Fatal(err)
		}
		return strings.ReplaceAll(relpath, `\`, `/`)
		// on windows git still prefers relatives paths
		// containing `/` for submodules
	}

	// Set up the grandchild
	gc := testGitRepo(t, "grandchild")
	gc.commitFile("grandchild.txt", "grandchild")

	// Set up the child
	c := testGitRepo(t, "child")
	c.commitFile("child.txt", "child")
	c.git("submodule", "add", "-f", relpath(c.dir, gc.dir))
	c.git("commit", "-m", "Add grandchild submodule")

	// Set up the parent
	p := testGitRepo(t, "parent")
	p.commitFile("parent.txt", "parent")
	p.git("submodule", "add", "-f", relpath(p.dir, c.dir))
	p.git("commit", "-m", "Add child submodule")

	// Clone the root repository
	if err := g.Get(dst, p.url); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Check that the files exist
	for _, path := range []string{
		filepath.Join(dst, "parent.txt"),
		filepath.Join(dst, "child", "child.txt"),
		filepath.Join(dst, "child", "grandchild", "grandchild.txt"),
	} {
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("err: %s", err)
		}
	}
}

func TestGitGetter_branch(t *testing.T) {
	if !testHasGit {
		t.Skip("git not found, skipping")
	}

	g := new(GoGitGetter)
	dst := tempDir(t)

	repo := testGitRepo(t, "branch")
	repo.git("checkout", "-b", "test-branch")
	repo.commitFile("branch.txt", "branch")

	q := repo.url.Query()
	q.Add("ref", "test-branch")
	repo.url.RawQuery = q.Encode()

	if err := g.Get(dst, repo.url); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Verify the main file exists
	mainPath := filepath.Join(dst, "branch.txt")
	if _, err := os.Stat(mainPath); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Get again should work
	if err := g.Get(dst, repo.url); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Verify the main file exists
	mainPath = filepath.Join(dst, "branch.txt")
	if _, err := os.Stat(mainPath); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func tempDir(t *testing.T) string {
	dir, err := ioutil.TempDir("", "tf")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if err := os.RemoveAll(dir); err != nil {
		t.Fatalf("err: %s", err)
	}

	return dir
}

func tempTestFile(t *testing.T) string {
	dir := tempDir(t)
	return filepath.Join(dir, "foo")
}

// gitRepo is a helper struct which controls a single temp git repo.
type gitRepo struct {
	t   *testing.T
	url *url.URL
	dir string
}

// testGitRepo creates a new test git repository.
func testGitRepo(t *testing.T, name string) *gitRepo {
	dir, err := ioutil.TempDir("", "go-getter")
	if err != nil {
		t.Fatal(err)
	}
	dir = filepath.Join(dir, name)
	if err := os.Mkdir(dir, 0700); err != nil {
		t.Fatal(err)
	}

	r := &gitRepo{
		t:   t,
		dir: dir,
	}

	url, err := urlhelper.Parse("file://" + r.dir)
	if err != nil {
		t.Fatal(err)
	}
	r.url = url

	t.Logf("initializing git repo in %s", dir)
	r.git("init")
	r.git("config", "user.name", "go-getter")
	r.git("config", "user.email", "go-getter@hashicorp.com")
	t.Logf("initializing git repo in %s done", dir)
	return r
}

// git runs a git command against the repo.
func (r *gitRepo) git(args ...string) {
	cmd := exec.Command("git", args...)
	cmd.Dir = r.dir
	bfr := bytes.NewBuffer(nil)
	cmd.Stderr = bfr
	if err := cmd.Run(); err != nil {
		r.t.Fatal(err, bfr.String())
	}
}

// commitFile writes and commits a text file to the repo.
func (r *gitRepo) commitFile(file, content string) {
	path := filepath.Join(r.dir, file)
	if err := ioutil.WriteFile(path, []byte(content), 0600); err != nil {
		r.t.Fatal(err)
	}
	r.git("add", file)
	r.git("commit", "-m", "Adding "+file)
}

// latestCommit returns the full commit id of the latest commit on the current
// branch.
func (r *gitRepo) latestCommit() (string, error) {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = r.dir
	rawOut, err := cmd.Output()
	if err != nil {
		return "", err
	}
	rawOut = bytes.TrimSpace(rawOut)
	return string(rawOut), nil
}
