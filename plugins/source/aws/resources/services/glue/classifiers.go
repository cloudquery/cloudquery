package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Classifiers() *schema.Table {
	tableName := "aws_glue_classifiers"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_Classifier.html`,
		Resolver:    fetchGlueClassifiers,
		Transform:   transformers.TransformWithStruct(&types.Classifier{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "glue"),
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

func fetchGlueClassifiers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Glue
	input := glue.GetClassifiersInput{}
	for {
		output, err := svc.GetClassifiers(ctx, &input)
		if err != nil {
			return err
		}
		res <- output.Classifiers

		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}
func resolveGlueClassifierName(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Classifier)
	if r.CsvClassifier != nil {
		return resource.Set(c.Name, r.CsvClassifier.Name)
	}
	if r.JsonClassifier != nil {
		return resource.Set(c.Name, r.JsonClassifier.Name)
	}
	if r.GrokClassifier != nil {
		return resource.Set(c.Name, r.GrokClassifier.Name)
	}
	if r.XMLClassifier != nil {
		return resource.Set(c.Name, r.XMLClassifier.Name)
	}
	return nil
}
