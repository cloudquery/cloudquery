package services

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Images() *schema.Table {
	return &schema.Table{
		Name:        "cloudflare_images",
		Description: "Image represents a Cloudflare Image.",
		Resolver:    fetchImages,
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountID,
			},
			{
				Name:            "id",
				Description:     "Image unique identifier",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "filename",
				Description: "Image file name",
				Type:        schema.TypeString,
			},
			{
				Name:        "metadata",
				Description: "User modifiable key-value store. Can be used for keeping references to another system of record for managing images. Metadata must not exceed 1024 bytes.",
				Type:        schema.TypeJSON,
			},
			{
				Name:     "require_signed_urls",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("RequireSignedURLs"),
			},
			{
				Name:        "variants",
				Description: "Object specifying available variants for an image.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "uploaded",
				Description: "When the media item was uploaded.",
				Type:        schema.TypeTimestamp,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	accountId := svc.AccountId

	records, err := svc.ClientApi.ListImages(ctx, accountId, cloudflare.PaginationOptions{})
	if err != nil {
		return err
	}
	res <- records
	return nil
}
