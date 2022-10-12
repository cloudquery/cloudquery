package athena

import (
	"context"
	"errors"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchAthenaWorkGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	input := athena.ListWorkGroupsInput{}
	for {
		response, err := svc.ListWorkGroups(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.WorkGroups
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}

	return nil
}

func getWorkGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena

	wg := resource.Item.(types.WorkGroupSummary)
	dc, err := svc.GetWorkGroup(ctx, &athena.GetWorkGroupInput{
		WorkGroup: wg.Name,
	})
	if err != nil {
		return err
	}
	resource.Item = *dc.WorkGroup
	return nil
}

func resolveAthenaWorkGroupArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	dc := resource.Item.(types.WorkGroup)
	return resource.Set(c.Name, createWorkGroupArn(cl, *dc.Name))
}

func resolveAthenaWorkGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Athena
	wg := resource.Item.(types.WorkGroup)
	arn := createWorkGroupArn(cl, *wg.Name)
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

func fetchAthenaWorkGroupPreparedStatements(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	wg := parent.Item.(types.WorkGroup)
	input := athena.ListPreparedStatementsInput{WorkGroup: wg.Name}
	for {
		response, err := svc.ListPreparedStatements(ctx, &input)
		if err != nil {
			return err
		}
		for _, d := range response.PreparedStatements {
			dc, err := svc.GetPreparedStatement(ctx, &athena.GetPreparedStatementInput{
				WorkGroup:     wg.Name,
				StatementName: d.StatementName,
			})
			if err != nil {
				if c.IsNotFoundError(err) {
					continue
				}
				return err
			}
			res <- *dc.PreparedStatement
			return nil
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func fetchAthenaWorkGroupQueryExecutions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	wg := parent.Item.(types.WorkGroup)
	input := athena.ListQueryExecutionsInput{WorkGroup: wg.Name}
	for {
		response, err := svc.ListQueryExecutions(ctx, &input)
		if err != nil {
			return err
		}
		for _, d := range response.QueryExecutionIds {
			dc, err := svc.GetQueryExecution(ctx, &athena.GetQueryExecutionInput{
				QueryExecutionId: aws.String(d),
			})
			if err != nil {
				if c.IsNotFoundError(err) || isQueryExecutionNotFound(err) {
					continue
				}
				return err
			}
			res <- *dc.QueryExecution
			return nil
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func fetchAthenaWorkGroupNamedQueries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	wg := parent.Item.(types.WorkGroup)
	input := athena.ListNamedQueriesInput{WorkGroup: wg.Name}
	for {
		response, err := svc.ListNamedQueries(ctx, &input)
		if err != nil {
			return err
		}
		for _, d := range response.NamedQueryIds {
			dc, err := svc.GetNamedQuery(ctx, &athena.GetNamedQueryInput{
				NamedQueryId: aws.String(d),
			})
			if err != nil {
				if c.IsNotFoundError(err) {
					continue
				}
				return err
			}
			res <- *dc.NamedQuery
			return nil
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func createWorkGroupArn(cl *client.Client, groupName string) string {
	return cl.ARN(client.Athena, "workgroup", groupName)
}

func isQueryExecutionNotFound(err error) bool {
	var ae smithy.APIError
	if !errors.As(err, &ae) {
		return false
	}
	return ae.ErrorCode() == "InvalidRequestException" && strings.Contains(ae.ErrorMessage(), "was not found")
}
