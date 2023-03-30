package ssoadmin

import (
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func AccountAssignments() *schema.Table {
	tableName := "aws_ssoadmin_account_assignments"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_AccountAssignment.html`,
		Resolver:    fetchSsoadminAccountAssignments,
		Transform:   transformers.TransformWithStruct(&types.AccountAssignment{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "identitystore"),
	}
}
