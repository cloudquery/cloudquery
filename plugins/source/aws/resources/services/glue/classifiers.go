package glue

import (
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Classifiers() *schema.Table {
	return &schema.Table{
		Name:        "aws_glue_classifiers",
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_Classifier.html`,
		Resolver:    fetchGlueClassifiers,
		Transform:   transformers.TransformWithStruct(&types.Classifier{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: resolveGlueClassifierName,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
