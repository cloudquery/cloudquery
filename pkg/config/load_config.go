package config

import (
	"io"
	"net/url"
	"strings"

	"github.com/hairyhenderson/go-fsimpl"
	"github.com/hairyhenderson/go-fsimpl/blobfs"
	"github.com/hairyhenderson/go-fsimpl/filefs"
)

func loadRemoteFile(path string) ([]byte, error) {
	mux := fsimpl.NewMux()
	mux.Add(filefs.FS)
	mux.Add(blobfs.FS)

	sanitizedPath, _ := url.Parse(path)

	// go-fsimpl is looking for a "directory" where it is the full path without the actual file
	// but must have all of the query strings
	directory := path[:strings.LastIndex(path, "/")] + "?" + sanitizedPath.RawQuery
	fileName := sanitizedPath.Path[strings.LastIndex(sanitizedPath.Path, "/")+1:]

	fsys, err := mux.Lookup(directory)
	if err != nil {
		return nil, err
	}

	f, err := fsys.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	contents, err := io.ReadAll(f)
	return contents, err
}
