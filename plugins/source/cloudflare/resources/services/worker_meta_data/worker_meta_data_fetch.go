package worker_meta_data

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchWorkerMetaData(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)

	rc := cloudflare.AccountIdentifier(svc.AccountId)
	params := cloudflare.ListWorkersParams{}
	resp, _, err := svc.ClientApi.ListWorkers(ctx, rc, params)
	if err != nil {
		return err
	}
	res <- resp.WorkerList

	return nil
}
func fetchWorkerCronTriggers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	accountId := svc.AccountId
	script := parent.Item.(cloudflare.WorkerMetaData)

	rc := cloudflare.AccountIdentifier(accountId)
	params := cloudflare.ListWorkerCronTriggersParams{ScriptName: script.ID}
	resp, err := svc.ClientApi.ListWorkerCronTriggers(ctx, rc, params)
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
func fetchWorkersSecrets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	script := parent.Item.(cloudflare.WorkerMetaData)

	rc := cloudflare.AccountIdentifier(svc.AccountId)
	params := cloudflare.ListWorkersSecretsParams{ScriptName: script.ID}
	resp, err := svc.ClientApi.ListWorkersSecrets(ctx, rc, params)
	if err != nil {
		return err
	}
	res <- resp.Result

	return nil
}
