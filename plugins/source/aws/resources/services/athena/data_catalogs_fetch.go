package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchAthenaDataCatalogs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return client.ListAndDetailResolver(ctx, meta, res, listDataCatalogs, dataCatalogDetail)
}
func resolveAthenaDataCatalogArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	dc := resource.Item.(types.DataCatalog)
	return resource.Set(c.Name, createDataCatalogArn(cl, *dc.Name))
}
func resolveAthenaDataCatalogTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Athena
	dc := resource.Item.(types.DataCatalog)
	arn := createDataCatalogArn(cl, *dc.Name)
	params := athena.ListTagsForResourceInput{ResourceARN: &arn}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTagsForResource(ctx, &params)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		client.TagsIntoMap(result.Tags, tags)
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return resource.Set(c.Name, tags)
}
func fetchAthenaDataCatalogDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	input := athena.ListDatabasesInput{
		CatalogName: parent.Item.(types.DataCatalog).Name,
	}
	for {
		response, err := svc.ListDatabases(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.DatabaseList

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func fetchAthenaDataCatalogDatabaseTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Athena
	input := athena.ListTableMetadataInput{
		CatalogName:  parent.Parent.Item.(types.DataCatalog).Name,
		DatabaseName: parent.Item.(types.Database).Name,
	}
	for {
		response, err := svc.ListTableMetadata(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.TableMetadataList

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func listDataCatalogs(ctx context.Context, meta schema.ClientMeta, detailChan chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	input := athena.ListDataCatalogsInput{}
	for {
		response, err := svc.ListDataCatalogs(ctx, &input, func(options *athena.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		for _, item := range response.DataCatalogsSummary {
			detailChan <- item
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func dataCatalogDetail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, listInfo interface{}) {
	c := meta.(*client.Client)
	catalogSummary := listInfo.(types.DataCatalogSummary)
	svc := c.Services().Athena
	dc, err := svc.GetDataCatalog(ctx, &athena.GetDataCatalogInput{
		Name: catalogSummary.CatalogName,
	})
	if err != nil {
		// retrieving of default data catalog (AwsDataCatalog) returns "not found error" but it exists and its
		// relations can be fetched by its name
		if *catalogSummary.CatalogName == "AwsDataCatalog" {
			resultsChan <- types.DataCatalog{Name: catalogSummary.CatalogName, Type: catalogSummary.Type}
			return
		}
		if c.IsNotFoundError(err) {
			return
		}
		errorChan <- err
		return
	}
	resultsChan <- *dc.DataCatalog
}
func createDataCatalogArn(cl *client.Client, catalogName string) string {
	return cl.ARN(client.Athena, "datacatalog", catalogName)
}
