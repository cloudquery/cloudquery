package config

import (
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/hairyhenderson/go-fsimpl"
	"github.com/hairyhenderson/go-fsimpl/blobfs"
	"github.com/hairyhenderson/go-fsimpl/filefs"
)

// srvURL := setupTestS3Bucket(t)

// 	os.Setenv("AWS_ANON", "true")
// 	defer os.Unsetenv("AWS_ANON")

// 	fsys, err := New(tests.MustURL("s3://mybucket/?region=us-east-1&disableSSL=true&s3ForcePathStyle=true&endpoint=" + srvURL.Host))
// 	assert.NoError(t, err)

// 	de, err := fs.ReadDir(fsys, "dir1")
// 	assert.NoError(t, err)

func loadRemoteFile(path string) ([]byte, error) {
	mux := fsimpl.NewMux()
	mux.Add(filefs.FS)
	mux.Add(blobfs.FS)

	sanitizedPath, _ := url.Parse(path)
	directory := path[:strings.LastIndex(path, "/")] + "?" + sanitizedPath.RawQuery
	fileName := sanitizedPath.Path[strings.LastIndex(sanitizedPath.Path, "/")+1:]
	fmt.Println(directory, fileName)
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
