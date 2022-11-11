// Auto generated code - DO NOT EDIT.

package web

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
)

func functions() *schema.Table {
	return &schema.Table{
		Name:        "azure_web_functions",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web#FunctionEnvelope`,
		Resolver:    fetchWebFunctions,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "web_app_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "function_app_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FunctionAppID"),
			},
			{
				Name:     "script_root_path_href",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ScriptRootPathHref"),
			},
			{
				Name:     "script_href",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ScriptHref"),
			},
			{
				Name:     "config_href",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConfigHref"),
			},
			{
				Name:     "test_data_href",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TestDataHref"),
			},
			{
				Name:     "secrets_file_href",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecretsFileHref"),
			},
			{
				Name:     "href",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Href"),
			},
			{
				Name:     "files",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Files"),
			},
			{
				Name:     "test_data",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TestData"),
			},
			{
				Name:     "invoke_url_template",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InvokeURLTemplate"),
			},
			{
				Name:     "language",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Language"),
			},
			{
				Name:     "is_disabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsDisabled"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}

func fetchWebFunctions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Web.Functions

	site := parent.Item.(web.Site)
	response, err := svc.ListFunctions(ctx, *site.ResourceGroup, *site.Name)

	if err != nil {
		return err
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}

	return nil
}
