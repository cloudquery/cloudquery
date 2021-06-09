package file

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/spf13/afero"
)

// OsFs is the struct that defines the os filesystem singleton instance.
type OsFs struct {
	// fs is the afero filesystem instance
	fs afero.Fs
}

var (
	once         sync.Once
	osFsInstance *OsFs
)

func NewOsFs() *OsFs {
	// Singleton instantiation
	once.Do(func() {
		osFsInstance = &OsFs{
			fs: afero.NewOsFs(),
		}
	})
	return osFsInstance
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory. We pass an io.TeeReader
// into Copy() to report progress on the download.
func (o *OsFs) DownloadFile(ctx context.Context, filepath, url string, progressUpdater func(reader io.Reader, total int64) io.Reader) error {
	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	out, err := o.fs.Create(filepath + ".tmp")
	if err != nil {
		return err
	}
	// Get the data
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("got %d http code instead expected %d", resp.StatusCode, http.StatusOK)
	}

	var reader io.Reader = resp.Body
	if progressUpdater != nil {
		reader = progressUpdater(resp.Body, resp.ContentLength)
	}
	// Create our progress reporter and pass it to be used alongside our writer
	if _, err = io.Copy(out, reader); err != nil {
		out.Close()
		return err
	}
	// Close the file without defer so it can happen before Rename()
	out.Close()

	if err = o.fs.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}
	return nil
}

// WalkPathTree is a wrapper around afero.Walk that walks the file tree
// starting with the given path and calls the walker function for every
// object it finds.
func (o *OsFs) WalkPathTree(path string, walker func(path string, info os.FileInfo, err error) error) error {
	return afero.Walk(o.fs, path, walker)
}

// Chmod is a wrapper around afero.Chmod that changes the file/folder permissions.
func (o *OsFs) Chmod(filePath string, mode os.FileMode) error {
	return o.fs.Chmod(filePath, mode)
}

// Remove is a wrapper around afero.Remove that removes a file.
func (o *OsFs) Remove(filePath string) error {
	return o.fs.Remove(filePath)
}

// MkdirAll is a wrapper around afero.Mkdirall that creates the full path
// directory tree.
func (o *OsFs) MkdirAll(path string, perm os.FileMode) error {
	return o.fs.MkdirAll(path, perm)
}

// SetFSInstance sets the FS instance. Should be only used for testing purpose.
func (o *OsFs) SetFSInstance(fs afero.Fs) {
	o.fs = fs
}
