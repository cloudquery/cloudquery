package ssoadmin

import (
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func PermissionSets() *schema.Table {
	tableName := "aws_ssoadmin_permission_sets"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_PermissionSet.html`,
		Resolver:            fetchSsoadminPermissionSets,
		PreResourceResolver: getSsoadminPermissionSet,
		Transform:           client.TransformWithStruct(&types.PermissionSet{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "identitystore"),
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
