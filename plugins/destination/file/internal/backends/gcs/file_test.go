package gcs

import (
	"bytes"
	"context"
	"testing"

	"cloud.google.com/go/storage"
)

const testBufferSize = 1024

func TestFileNoLimit(t *testing.T) {
	ctx := context.Background()
	storageClient, err := storage.NewClient(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	name := "test_file_no_limit.txt"
	writer, err := OpenAppendOnly(ctx, storageClient, "cq-yev-test", name, 0)
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

	reader, err := OpenReadOnly(ctx, storageClient, "cq-yev-test", name)
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