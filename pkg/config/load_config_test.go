package config

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/johannesboyne/gofakes3"
	"github.com/johannesboyne/gofakes3/backend/s3mem"
	"github.com/stretchr/testify/assert"
)

const bucketName = "myBucket"

func putFile(backend gofakes3.Backend, file, mime, content string) error {
	_, err := backend.PutObject(
		bucketName,
		file,
		map[string]string{"Content-Type": mime},
		bytes.NewBufferString(content),
		int64(len(content)),
	)

	return err
}

func setupTestS3Bucket(t *testing.T) *url.URL {
	backend := s3mem.New()
	faker := gofakes3.New(backend)

	srv := httptest.NewServer(faker.Server())

	t.Cleanup(srv.Close)

	assert.NoError(t, backend.CreateBucket(bucketName))
	assert.NoError(t, putFile(backend, "config.hcl", "application/hcl", testConfig))

	u, err := url.Parse(srv.URL)
	assert.NoError(t, err)
	return u
}

func TestLoadRemoteFile(t *testing.T) {
	srvURL := setupTestS3Bucket(t)
	os.Setenv("AWS_ANON", "true")
	defer os.Unsetenv("AWS_ANON")
	fmt.Println(srvURL)
	p := NewParser()
	body, diags := p.LoadConfigFile(fmt.Sprintf("s3://%s/config.hcl?region=us-east-1&disableSSL=true&s3ForcePathStyle=true&endpoint=%s", bucketName, srvURL.Host))
	assert.Equal(t, 0, len(diags))

	cfg, diags := p.LoadConfigFromSource("test.hcl", []byte(testConfig))
	assert.Nil(t, diags)
	assert.Equal(t, cfg, body)
}
