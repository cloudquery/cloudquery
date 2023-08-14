package services

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func PublisherBlockLists() *schema.Table {
	return &schema.Table{
		Name:      "facebookmarketing_publisher_block_lists",
		Resolver:  fetchPublisherBlockLists,
		Transform: client.TransformWithStruct(&rest.PublisherBlockList{}, transformers.WithPrimaryKeys("Id")),
		Columns: []schema.Column{
			{
				Name:       "account_id",
				Resolver:   client.ResolveAccountId,
				Type:       arrow.BinaryTypes.String,
				PrimaryKey: true,
			},
		},
		Description: "https://developers.facebook.com/docs/marketing-api/reference/publisher-block-list/",
	}
}
