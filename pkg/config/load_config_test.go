package config

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/johannesboyne/gofakes3"
	"github.com/johannesboyne/gofakes3/backend/s3mem"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

const bucketName = "myBucket"
const defaultPermissions = 0644

func putFile(backend gofakes3.Backend, path, mime, content string) error {
	u, err := url.Parse(path)
	if err != nil {
		return err
	}
	_, err = backend.PutObject(
		bucketName,
		strings.TrimPrefix(u.Path, "/"),
		map[string]string{"Content-Type": mime},
		bytes.NewBufferString(content),
		int64(len(content)),
	)

	return err
}

func setupTestS3Bucket(t *testing.T) (*url.URL, *s3mem.Backend) {
	backend := s3mem.New()
	faker := gofakes3.New(backend)

	srv := httptest.NewServer(faker.Server())

	t.Cleanup(srv.Close)

	assert.NoError(t, backend.CreateBucket(bucketName))
	u, err := url.Parse(srv.URL)
	assert.NoError(t, err)
	return u, backend
}

func TestLoadRemoteFile(t *testing.T) {
	srvURL, backend := setupTestS3Bucket(t)
	cases := []struct {
		Path        string
		Name        string
		Configs     string
		Type        string
		SetupFile   bool
		ExpectError bool
	}{
		{
			Name:      "Success-S3Object",
			Type:      "s3",
			Path:      fmt.Sprintf("s3://%s/config.hcl?region=us-east-1&disableSSL=true&s3ForcePathStyle=true&endpoint=%s", bucketName, srvURL.Host),
			Configs:   testConfig,
			SetupFile: true,
		},
		{
			Name:        "Failure-S3Object",
			Type:        "s3",
			Path:        fmt.Sprintf("s3://%s/config2.hcl?region=us-east-1&disableSSL=true&s3ForcePathStyle=true&endpoint=%s", bucketName, srvURL.Host),
			Configs:     testConfig,
			SetupFile:   false,
			ExpectError: true,
		},
		{
			Name:      "Success-RelativePath",
			Type:      "file",
			Path:      "./asdf/asdf/asfd/teddstcdonfig.hcl",
			Configs:   testConfig,
			SetupFile: true,
		},
		{
			Name:        "Failure-RelativePath-NotExists",
			Type:        "file",
			Path:        "./asdf/asdf/asfd/teddstcdonfig.hcl",
			Configs:     testConfig,
			SetupFile:   false,
			ExpectError: true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			var p *Parser
			switch tc.Type {
			case "s3":
				p = NewParser()
				os.Setenv("AWS_ANON", "true")
				defer os.Unsetenv("AWS_ANON")
				if tc.SetupFile {
					assert.NoError(t, putFile(backend, tc.Path, "application/hcl", tc.Configs))
				}

			case "file":
				appFS := afero.NewMemMapFs()
				p = NewParser(func(p *Parser) {
					p.fs = afero.Afero{Fs: appFS}
				})
				if tc.SetupFile {
					p.fs.WriteFile(tc.Path, []byte(tc.Configs), defaultPermissions)
				}
			}
			body, diags := p.LoadConfigFile(tc.Path)
			if !tc.ExpectError {
				assert.Equal(t, 0, len(diags))
				p2 := NewParser()
				cfg, diags := p2.LoadConfigFromSource("test.hcl", []byte(tc.Configs))
				assert.Nil(t, diags)
				assert.Equal(t, cfg, body)
			} else {
				assert.Equal(t, 1, len(diags))
				assert.Equal(t, "Failed to read file", diags[0].Summary)
				assert.Equal(t, fmt.Sprintf("The file \"%s\" could not be read: file does not exist. Hint: Try `cloudquery init <provider>`.", tc.Path), diags[0].Detail)
			}

		})
	}
}
