// persistentdata will look for the given filename in two paths: ~/.cq/ and ./.cq/.
package persistentdata

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

type Client struct {
	fs  afero.Afero
	fn  string
	gen func() string
}

var ErrIsDirectory = fmt.Errorf("file is directory")

func New(fs afero.Afero, fn string, gen func() string) *Client {
	return &Client{
		fs:  fs,
		fn:  fn,
		gen: gen,
	}
}

// Get the data (generate if it doesn't exist) and return it
func (c *Client) Get() (value string, newlyCreated bool, err error) {
	if home, err := os.UserHomeDir(); err == nil {
		// special case for having a persistent ID, read-only access
		id, err := c.read(filepath.Join(home, ".cq"))
		if err == ErrIsDirectory {
			// finish early if directory encountered
			return "", false, err
		}
		if id != "" {
			return id, false, nil
		}
	}

	path := filepath.Join(".", ".cq")
	if id, err := c.read(path); err != nil {
		return "", false, err
	} else if id != "" {
		return id, false, nil
	}

	id := c.gen()
	if id == "" {
		return "", false, nil
	}

	if err := c.write(path, id); err != nil {
		return "", false, err
	}

	return id, true, nil
}

// read the contents of given fn in the given fs. Returns ErrIsDirectory if the file exists but is a directory.
func (c *Client) read(path string) (string, error) {
	fn := filepath.Join(path, c.fn)

	exists := true
	fi, err := c.fs.Stat(fn)
	if err != nil {
		if !os.IsNotExist(err) {
			return "", err
		}
		exists = false
	}
	if exists && fi.IsDir() {
		return "", ErrIsDirectory
	}
	if !exists {
		return "", nil
	}

	b, err := c.fs.ReadFile(fn)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// write the given payload into the file in the given fs and path
func (c *Client) write(path, payload string) error {
	fn := filepath.Join(path, c.fn)
	return c.fs.WriteFile(fn, []byte(payload), 0644)
}
