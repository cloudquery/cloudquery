package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ec2/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RegionalConfigs() *schema.Table {
	tableName := "aws_ec2_regional_configs"
	return &schema.Table{
		Name:      tableName,
		Resolver:  fetchEc2RegionalConfigs,
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform: transformers.TransformWithStruct(&models.RegionalConfig{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}
