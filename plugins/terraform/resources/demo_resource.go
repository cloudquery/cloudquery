package resources

import (
	"context"
	"github.com/cloudquery/cq-provider-sdk/plugin/schema"
	"github.com/cloudquery/cq-provider-template/client"
)

func DemoResource() *schema.Table {
	return &schema.Table{
		Name:         "provider_demo_resources",
		Resolver:     fetchDemoResources,
		// Those are optional
		// DeleteFilter: nil,
		// Multiplex:    nil,
		// IgnoreError:  nil,
		//PostResourceResolver: nil,

		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				//Resolver: provider.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				//Resolver: fetchS3BucketLocation,
			},
			{
				Name: "creation_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "name",
				Type: schema.TypeString,
				Resolver: schema.PathResolver("other_name_in_struct"),
			},
		},
		// A table can have relations
		//Relations: []*schema.Table{
		//	{
		//		Name:     "provider_demo_resource_children",
		//		Resolver: fetchDemoResourceChildren,
		//		Columns: []schema.Column{
		//			{
		//				Name:     "bucket_id",
		//				Type:     schema.TypeUUID,
		//				Resolver: schema.ParentIdResolver,
		//			},
		//			{
		//				Name:     "resource_id",
		//				Type:     schema.TypeString,
		//				Resolver: schema.PathResolver("Grantee.ID"),
		//			},
		//			{
		//				Name:     "type",
		//				Type:     schema.TypeString,
		//				Resolver: schema.PathResolver("Grantee.Type"),
		//			},
		//		},
		//	},
		//},
	}
}


// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchDemoResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	_ = c
	// Fetch using the third party client and put the result in res
	// res <- c.ThirdPartyClient.getDat()
	return nil
}