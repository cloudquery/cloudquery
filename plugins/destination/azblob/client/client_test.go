package client

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/google/uuid"
)

const storageAccount = "cqdestinationazblob"

func getClient(t *testing.T) *azblob.Client {
	t.Helper()

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Fatal(err)
	}

	storageClient, err := azblob.NewClient("https://"+storageAccount+".blob.core.windows.net", cred, nil)
	if err != nil {
		t.Fatal(err)
	}

	return storageClient
}

func createContainer(t *testing.T, container string) {
	t.Helper()

	storageClient := getClient(t)

	_, err := storageClient.CreateContainer(context.Background(), container, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func deleteContainer(t *testing.T, container string) {
	t.Helper()

	storageClient := getClient(t)

	_, err := storageClient.DeleteContainer(context.Background(), container, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPluginCSV(t *testing.T) {
	container := "test-csv-" + uuid.New().String()
	createContainer(t, container)
	defer deleteContainer(t, container)
	p := destination.NewPlugin("azblob", "development", New, destination.WithManagedWriter())

	destination.PluginTestSuiteRunner(t, p,
		Spec{
			StorageAccount: storageAccount,
			Container:      container,
			Path:           t.TempDir(),
			Format:         FormatTypeCSV,
			NoRotate:       true,
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:    true,
			SkipDeleteStale:  true,
			SkipSecondAppend: true,
		},
	)
}

func TestPluginJSON(t *testing.T) {
	container := "test-json-" + uuid.New().String()
	createContainer(t, container)
	defer deleteContainer(t, container)
	p := destination.NewPlugin("azblob", "development", New, destination.WithManagedWriter())

	destination.PluginTestSuiteRunner(t, p,
		Spec{
			StorageAccount: storageAccount,
			Container:      container,
			Path:           t.TempDir(),
			Format:         FormatTypeJSON,
			NoRotate:       true,
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:    true,
			SkipDeleteStale:  true,
			SkipSecondAppend: true,
		},
	)
}
