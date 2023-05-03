package storage

import (
	"context"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azqueue"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func queueAccessPolicy() *schema.Table {
	return &schema.Table{
		Name:        "azure_storage_queue_acl",
		Resolver:    fetchQueueACL,
		Description: "https://learn.microsoft.com/en-us/rest/api/storageservices/get-queue-acl#response-body",
		Transform:   transformers.TransformWithStruct(&azqueue.GetAccessPolicyResponse{}, transformers.WithSkipFields("Date", "RequestID")),
		Columns: schema.ColumnList{
			client.SubscriptionID,
			schema.Column{
				Name:            "queue_id",
				Type:            schema.TypeString,
				Resolver:        schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	}
}

func fetchQueueACL(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	queue := parent.Item.(*armstorage.ListQueue)
	acc := parent.Parent.Item.(*armstorage.Account)

	opts := azqueue.ClientOptions{}
	if cl.Options != nil {
		opts.ClientOptions = cl.Options.ClientOptions
	}

	nm := strings.ReplaceAll(*acc.Name, " ", "_") // This is for the tests, real data will not have spaces
	queueURL := runtime.JoinPaths("https://"+nm+".queue.core.windows.net", url.PathEscape(*queue.Name))
	svc, err := azqueue.NewQueueClient(queueURL, cl.Creds, &opts)
	if err != nil {
		return err
	}

	resp, err := svc.GetAccessPolicy(ctx, nil)
	if err != nil {
		return err
	}

	res <- resp
	return nil
}
