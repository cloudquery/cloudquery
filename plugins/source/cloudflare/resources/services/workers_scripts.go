package services

import (
	"context"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cq-provider-cloudflare/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource workers_scripts --config workers.hcl --output .
func WorkersScripts() *schema.Table {
	return &schema.Table{
		Name:         "cloudflare_workers_scripts",
		Description:  "WorkerMetaData contains worker script information such as size, creation & modification dates.",
		Resolver:     fetchWorkersScripts,
		Multiplex:    client.AccountMultiplex,
		DeleteFilter: client.DeleteAccountFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountId,
			},
			{
				Name:        "id",
				Description: "The unique universal identifier for a Cloudflare zone.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:     "eta_g",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ETAG"),
			},
			{
				Name:        "size",
				Description: "Size of the script, in bytes.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "created_on",
				Description: "When the script was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "modified_on",
				Description: "When the script was last modified.",
				Type:        schema.TypeTimestamp,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "cloudflare_workers_script_cron_triggers",
				Description: "WorkerCronTrigger holds an individual cron schedule for a worker.",
				Resolver:    fetchWorkersScriptCronTriggers,
				Columns: []schema.Column{
					{
						Name:        "workers_script_cq_id",
						Description: "Unique CloudQuery ID of cloudflare_workers_scripts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "account_id",
						Description: "The Account ID of the resource.",
						Type:        schema.TypeString,
						Resolver:    client.ResolveAccountId,
					},
					{
						Name:        "cron",
						Description: "Raw cron expression",
						Type:        schema.TypeString,
					},
					{
						Name:        "created_on",
						Description: "When the Cron was created",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "modified_on",
						Description: "When the Cron was last modified",
						Type:        schema.TypeTimestamp,
					},
				},
			},
			{
				Name:        "cloudflare_workers_script_secrets",
				Description: "WorkersSecret contains the name and type of the secret.",
				Resolver:    fetchWorkersScriptSecrets,
				Columns: []schema.Column{
					{
						Name:        "workers_script_cq_id",
						Description: "Unique CloudQuery ID of cloudflare_workers_scripts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "account_id",
						Description: "The Account ID of the resource.",
						Type:        schema.TypeString,
						Resolver:    client.ResolveAccountId,
					},
					{
						Name:        "name",
						Description: "Secret name",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Secret type",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchWorkersScripts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	resp, err := svc.ClientApi.ListWorkerScripts(ctx)
	if err != nil {
		return diag.WrapError(err)
	}
	res <- resp.WorkerList

	return nil
}
func fetchWorkersScriptCronTriggers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	accountId := svc.AccountId
	script := parent.Item.(cloudflare.WorkerMetaData)

	resp, err := svc.ClientApi.ListWorkerCronTriggers(ctx, accountId, script.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	res <- resp

	return nil
}
func fetchWorkersScriptSecrets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	script := parent.Item.(cloudflare.WorkerMetaData)

	resp, err := svc.ClientApi.ListWorkersSecrets(ctx, script.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	res <- resp.Result

	return nil
}
