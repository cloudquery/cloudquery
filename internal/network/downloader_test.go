package network

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testFileDownloadUrl = "https://github.com/cloudquery/cloudquery/releases/download/v0.13.5/cloudquery_Linux_arm64.zip"
const expectedLength int64 = 6052203
const expectedMd5Hash = "85b1c9f43990e78ee3271b129cae34d2"

func TestNewHttpGetDownloader(t *testing.T) {
	// only verifies that download actually works
	d := NewHttpGetDownloader(context.Background(), testFileDownloadUrl)
	b, n, err := d()
	require.NoError(t, err)
	defer b.Close()
	assert.Equal(t, expectedLength, n)

	data, err := io.ReadAll(b)
	require.NoError(t, err)
	s := md5.Sum(data)
	assert.Equal(t, expectedMd5Hash, fmt.Sprintf("%x", s))
}
