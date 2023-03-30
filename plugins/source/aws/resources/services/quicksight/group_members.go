package quicksight

import (
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func groupMembers() *schema.Table {
	tableName := "aws_quicksight_group_members"
	return &schema.Table{
		Name:        tableName,
		Description: "https://docs.aws.amazon.com/quicksight/latest/APIReference/API_GroupMember.html",
		Resolver:    fetchQuicksightGroupMembers,
		Transform:   client.TransformWithStruct(&types.GroupMember{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "quicksight"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:            "group_arn",
				Type:            schema.TypeString,
				Resolver:        schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	}
}
