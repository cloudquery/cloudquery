package worker_meta_data

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchWorkerMetaData(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	resp, err := svc.ClientApi.ListWorkerScripts(ctx)
	if err != nil {
		return err
	}
	res <- resp.WorkerList

	return nil
}
func fetchWorkerCronTriggers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	accountId := svc.AccountId
	script := parent.Item.(cloudflare.WorkerMetaData)

	resp, err := svc.ClientApi.ListWorkerCronTriggers(ctx, accountId, script.ID)
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
func fetchWorkersSecrets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	script := parent.Item.(cloudflare.WorkerMetaData)

	resp, err := svc.ClientApi.ListWorkersSecrets(ctx, script.ID)
	if err != nil {
		return err
	}
	res <- resp.Result

	return nil
}
