package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Classifiers() *schema.Table {
	return &schema.Table{
		Name:        "aws_glue_classifiers",
		Description: "Classifiers are triggered during a crawl task",
		Resolver:    fetchGlueClassifiers,
		Multiplex:   client.ServiceAccountRegionMultiplexer("glue"),
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Description:     "The AWS Account ID of the resource",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "region",
				Description:     "The AWS Region of the resource",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "name",
				Description:     "Name of the classifier",
				Type:            schema.TypeString,
				Resolver:        resolveGlueClassifierName,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "csv_classifier",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CsvClassifier"),
			},
			{
				Name:     "grok_classifier",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("GrokClassifier"),
			},
			{
				Name:     "json_classifier",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("JsonClassifier"),
			},
			{
				Name:     "xml_classifier",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("XMLClassifier"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchGlueClassifiers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
