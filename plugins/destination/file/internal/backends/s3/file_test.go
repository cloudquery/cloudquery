package s3

import (
	"bytes"
	"context"
	"path"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const testBufferSize = 1024

func TestFileNoLimit(t *testing.T) {
	ctx := context.Background()
	awsCfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		t.Fatal(err)
	}
	awsClient := s3.NewFromConfig(awsCfg)
	uploader := manager.NewUploader(awsClient)
	downloader := manager.NewDownloader(awsClient)

	name := path.Join(t.TempDir(), "test_file_no_limit.txt")
	bucket := "cq-playground-test"
	writer, err := OpenAppendOnly(ctx, uploader, bucket, name, 0)
	if err != nil {
		t.Fatal(err)
	}
	testContent := []byte("test_file_no_limit")
	for i := 0; i < 2; i++ {
		n, err := writer.Write(testContent)
		if err != nil {
			t.Fatal(err)
		}
		if n != len(testContent) {
			t.Fatalf("expected %d bytes written, got %d", len(testContent), n)
		}
	}

	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}

	reader, err := OpenReadOnly(ctx, downloader, bucket, name)
	if err != nil {
		t.Fatal(err)
	}
	content := make([]byte, testBufferSize)
	n, err := reader.Read(content)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(testContent)*2 {
		t.Fatalf("expected %d bytes read, got %d", len(testContent), n)
	}
	expectedContent := append(testContent, testContent...)
	if !bytes.Equal(content[:n], expectedContent) {
		t.Fatalf("expected %s, got %s", string(expectedContent), string(content))
	}
	if err := reader.Close(); err != nil {
		t.Fatal(err)
	}
}
