package azure_blob_storage

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

const testBufferSize = 1024
const (
	storageAccount = "https://filebackendtest.blob.core.windows.net/"
	containerName  = "test-container"
)

func TestFileNoLimit(t *testing.T) {
	ctx := context.Background()
	// authenticate with Azure Active Directory
	cred, err := azidentity.NewAzureCLICredential(nil)
	if err != nil {
		t.Fatal(err)
	}
	// create a client for the specified storage account
	storageClient, err := azblob.NewClient(storageAccount, cred, nil)
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.OpenFile("testdata/file.txt", os.O_RDONLY, 0)
	// TODO: handle error
	defer file.Close()

	name := "test_file_no_limit.txt"

	// upload the file to the specified container with the specified blob name
	_, err = storageClient.UploadFile(context.TODO(), containerName, name, file, nil)
	if err != nil {
		t.Fatalf("could not upload file: %v", err)
	}
	// TODO: handle error

	writer, err := OpenAppendOnly(ctx, storageClient, containerName, name, 0)
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
		t.Fatal("error closing writer:", err)
	}

	reader, err := OpenReadOnly(ctx, storageClient, containerName, name)
	if err != nil {
		t.Fatal(err)
	}
	content := make([]byte, testBufferSize)
	n, err := reader.Read(content)
	if err != nil {
		t.Fatal("error reading file:", err)
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
