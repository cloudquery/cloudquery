package ssoadmin

import (
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func PermissionSets() *schema.Table {
	return &schema.Table{
		Name:                "aws_ssoadmin_permission_sets",
		Description:         `https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_PermissionSet.html`,
		Resolver:            fetchSsoadminPermissionSets,
		PreResourceResolver: getSsoadminPermissionSet,
		Transform:           transformers.TransformWithStruct(&types.PermissionSet{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("identitystore"),
		Columns: []schema.Column{
			{
				Name:     "inline_policy",
				Type:     schema.TypeJSON,
				Resolver: getSsoadminPermissionSetInlinePolicy,
			},
		},

		Relations: []*schema.Table{
			AccountAssignments(),
		},
	}
}
