// persistentdata will look for the given filename in two paths: ~/.cq/ and ./.cq/.
package persistentdata

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

const defaultPermissions = 0644

type Value struct {
	fs      afero.Afero
	Content string
	Created bool
	Path    string
}

func (v Value) Update(content string) error {
	return v.fs.WriteFile(v.Path, []byte(content), defaultPermissions)
}

type Client struct {
	fs  afero.Afero
	fn  string
	gen func() string
}

var errIsDirectory = fmt.Errorf("file is directory")

func New(fs afero.Afero, fn string, gen func() string) *Client {
	return &Client{
		fs:  fs,
		fn:  fn,
		gen: gen,
	}
}

// Get the data (generate if it doesn't exist) and return it
func (c *Client) Get() (v Value, err error) {
	v.fs = c.fs
	for _, prefix := range readOrder() {
		v.Path = filepath.Join(prefix, c.fn)
		v.Content, err = c.read(v.Path)
		if err == nil && v.Content != "" {
			return v, nil
		}
		if err == errIsDirectory {
			return Value{}, err
		}
	}
	if err != nil {
		return Value{}, err
	}

	// value must be empty at this point
	v.Content = c.gen()
	if v.Content == "" {
		return Value{}, nil
	}

	for _, prefix := range writeOrder() {
		v.Path = filepath.Join(prefix, c.fn)
		if err = os.MkdirAll(path.Dir(v.Path), fs.ModePerm); err != nil {
			continue
		}
		if err = c.write(v.Path, v.Content); err != nil {
			continue
		}
		break
	}
	v.Created = err == nil
	return v, err
}

// read the contents of given path in the given fs. Returns errIsDirectory if the file exists but is a directory.
func (c *Client) read(path string) (string, error) {
	exists := true
	fi, err := c.fs.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return "", err
		}
		exists = false
	}
	if exists && fi.IsDir() {
		return "", errIsDirectory
	}
	if !exists {
		return "", nil
	}

	b, err := c.fs.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// write the given payload into the file in the given fs and path
func (c *Client) write(path, payload string) error {
	return c.fs.WriteFile(path, []byte(payload), defaultPermissions)
}

func readOrder() []string {
	order := make([]string, 0, 2)
	if home, err := os.UserHomeDir(); err == nil {
		order = append(order, filepath.Join(home, ".cq"))
	}
	order = append(order, viper.GetString("data-dir"))

	return order
}

func writeOrder() []string {
	return []string{
		viper.GetString("data-dir"),
	}
}
