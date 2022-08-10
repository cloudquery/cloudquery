package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource classifiers --config classifiers.hcl --output .
func Classifiers() *schema.Table {
	return &schema.Table{
		Name:         "aws_glue_classifiers",
		Description:  "Classifiers are triggered during a crawl task",
		Resolver:     fetchGlueClassifiers,
		Multiplex:    client.ServiceAccountRegionMultiplexer("glue"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region", "name"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "name",
				Description: "Name of the classifier",
				Type:        schema.TypeString,
				Resolver:    resolveGlueClassifierName,
			},
			{
				Name:        "csv_classifier_name",
				Description: "The name of the classifier",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CsvClassifier.Name"),
			},
			{
				Name:        "csv_classifier_allow_single_column",
				Description: "Enables the processing of files that contain only one column",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("CsvClassifier.AllowSingleColumn"),
			},
			{
				Name:        "csv_classifier_contains_header",
				Description: "Indicates whether the CSV file contains a header",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CsvClassifier.ContainsHeader"),
			},
			{
				Name:        "csv_classifier_creation_time",
				Description: "The time that this classifier was registered",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("CsvClassifier.CreationTime"),
			},
			{
				Name:        "csv_classifier_delimiter",
				Description: "A custom symbol to denote what separates each column entry in the row",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CsvClassifier.Delimiter"),
			},
			{
				Name:        "csv_classifier_disable_value_trimming",
				Description: "Specifies not to trim values before identifying the type of column values",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("CsvClassifier.DisableValueTrimming"),
			},
			{
				Name:        "csv_classifier_header",
				Description: "A list of strings representing column names",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("CsvClassifier.Header"),
			},
			{
				Name:        "csv_classifier_last_updated",
				Description: "The time that this classifier was last updated",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("CsvClassifier.LastUpdated"),
			},
			{
				Name:        "csv_classifier_quote_symbol",
				Description: "A custom symbol to denote what combines content into a single column value",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CsvClassifier.QuoteSymbol"),
			},
			{
				Name:        "csv_classifier_version",
				Description: "The version of this classifier",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("CsvClassifier.Version"),
			},
			{
				Name:        "grok_classifier_classification",
				Description: "An identifier of the data format that the classifier matches, such as Twitter, JSON, Omniture logs, and so on",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("GrokClassifier.Classification"),
			},
			{
				Name:        "grok_classifier_grok_pattern",
				Description: "The grok pattern applied to a data store by this classifier",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("GrokClassifier.GrokPattern"),
			},
			{
				Name:        "grok_classifier_name",
				Description: "The name of the classifier",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("GrokClassifier.Name"),
			},
			{
				Name:        "grok_classifier_creation_time",
				Description: "The time that this classifier was registered",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("GrokClassifier.CreationTime"),
			},
			{
				Name:        "grok_classifier_custom_patterns",
				Description: "Optional custom grok patterns defined by this classifier",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("GrokClassifier.CustomPatterns"),
			},
			{
				Name:        "grok_classifier_last_updated",
				Description: "The time that this classifier was last updated",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("GrokClassifier.LastUpdated"),
			},
			{
				Name:        "grok_classifier_version",
				Description: "The version of this classifier",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("GrokClassifier.Version"),
			},
			{
				Name:        "json_classifier_json_path",
				Description: "A JsonPath string defining the JSON data for the classifier to classify",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("JsonClassifier.JsonPath"),
			},
			{
				Name:        "json_classifier_name",
				Description: "The name of the classifier",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("JsonClassifier.Name"),
			},
			{
				Name:        "json_classifier_creation_time",
				Description: "The time that this classifier was registered",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("JsonClassifier.CreationTime"),
			},
			{
				Name:        "json_classifier_last_updated",
				Description: "The time that this classifier was last updated",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("JsonClassifier.LastUpdated"),
			},
			{
				Name:        "json_classifier_version",
				Description: "The version of this classifier",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("JsonClassifier.Version"),
			},
			{
				Name:        "xml_classifier_classification",
				Description: "An identifier of the data format that the classifier matches",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("XMLClassifier.Classification"),
			},
			{
				Name:        "xml_classifier_name",
				Description: "The name of the classifier",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("XMLClassifier.Name"),
			},
			{
				Name:        "xml_classifier_creation_time",
				Description: "The time that this classifier was registered",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("XMLClassifier.CreationTime"),
			},
			{
				Name:        "xml_classifier_last_updated",
				Description: "The time that this classifier was last updated",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("XMLClassifier.LastUpdated"),
			},
			{
				Name:        "xml_classifier_row_tag",
				Description: "The XML tag designating the element that contains each record in an XML document being parsed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("XMLClassifier.RowTag"),
			},
			{
				Name:        "xml_classifier_version",
				Description: "The version of this classifier",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("XMLClassifier.Version"),
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
			return diag.WrapError(err)
		}
		res <- output.Classifiers

		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}

// nolint:gocritic
func resolveGlueClassifierName(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Classifier)
	if r.CsvClassifier != nil {
		return diag.WrapError(resource.Set(c.Name, r.CsvClassifier.Name))
	} else if r.JsonClassifier != nil {
		return diag.WrapError(resource.Set(c.Name, r.JsonClassifier.Name))
	} else if r.GrokClassifier != nil {
		return diag.WrapError(resource.Set(c.Name, r.GrokClassifier.Name))
	} else if r.XMLClassifier != nil {
		return diag.WrapError(resource.Set(c.Name, r.XMLClassifier.Name))
	}
	return nil
}
