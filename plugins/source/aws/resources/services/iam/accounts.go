package iam

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Accounts() *schema.Table {
	tableName := "aws_iam_accounts"
	return &schema.Table{
		Name:      tableName,
		Resolver:  fetchIamAccounts,
		Transform: client.TransformWithStruct(&models.Account{}),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}
