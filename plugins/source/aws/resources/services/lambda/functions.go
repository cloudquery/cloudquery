package lambda

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Functions() *schema.Table {
	tableName := "aws_lambda_functions"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/lambda/latest/dg/API_GetFunction.html`,
		Resolver:            fetchLambdaFunctions,
		PreResourceResolver: getFunction,
		Transform:           transformers.TransformWithStruct(&lambda.GetFunctionOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "lambda"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.FunctionArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name: "policy_revision_id",
				Type: schema.TypeString,
				// resolved in resolveResourcePolicy
			},
			{
				Name:     "policy_document",
				Type:     schema.TypeJSON,
				Resolver: resolveResourcePolicy,
			},
			{
				Name:     "code_signing_config",
				Type:     schema.TypeJSON,
				Resolver: resolveCodeSigningConfig,
			},
			{
				Name:     "code_repository_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Code.RepositoryType"),
			},
			{
				Name: "update_runtime_on",
				Type: schema.TypeString,
				// resolved in resolveRuntimeManagementConfig
			},
			{
				Name:     "runtime_version_arn",
				Type:     schema.TypeString,
				Resolver: resolveRuntimeManagementConfig,
			},
		},

		Relations: []*schema.Table{
			functionEventInvokeConfigs(),
			functionAliases(),
			functionVersions(),
			functionConcurrencyConfigs(),
			functionEventSourceMappings(),
		},
	}
}

func fetchLambdaFunctions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Lambda
	paginator := lambda.NewListFunctionsPaginator(svc, &lambda.ListFunctionsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Functions
	}
	return nil
}

func getFunction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Lambda
	f := resource.Item.(types.FunctionConfiguration)

	funcResponse, err := svc.GetFunction(ctx, &lambda.GetFunctionInput{
		FunctionName: f.FunctionName,
	})
	if err != nil {
		if c.IsNotFoundError(err) {
			resource.Item = &lambda.GetFunctionOutput{
				Configuration: &f,
			}
			return nil
		}
		// This is intended to handle the case where the user does not have GetFunction permission
		// User should still get an error in the logs, but the data that was able to be fetched should be persisted
		if client.IsAWSError(err, "AccessDenied") || client.IsAWSError(err, "AccessDeniedException") {
			resource.Item = &lambda.GetFunctionOutput{
				Configuration: &f,
			}
			c.Logger().Warn().Err(err).Msg("configuration data retrieved from ListFunctions will still be persisted")
			return nil
		}

		return err
	}

	resource.Item = funcResponse
	return nil
}

func resolveCodeSigningConfig(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(*lambda.GetFunctionOutput)
	if r.Configuration == nil {
		return nil
	}
	c := meta.(*client.Client)
	svc := c.Services().Lambda

	// skip getting CodeSigningConfig since containerized lambda functions does not support this feature
	// value can be nil if the caller doesn't have GetFunctionConfiguration permission and only has List*
	lambdaType := resource.Get("code_repository_type").(*schema.Text)
	if lambdaType != nil && lambdaType.Str == "ECR" {
		return nil
	}

	functionSigning, err := svc.GetFunctionCodeSigningConfig(ctx, &lambda.GetFunctionCodeSigningConfigInput{
		FunctionName: r.Configuration.FunctionName,
	})
	if err != nil {
		return err
	}
	if functionSigning.CodeSigningConfigArn == nil {
		return nil
	}

	signing, err := svc.GetCodeSigningConfig(ctx, &lambda.GetCodeSigningConfigInput{
		CodeSigningConfigArn: functionSigning.CodeSigningConfigArn,
	})
	if err != nil {
		if c.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	if signing.CodeSigningConfig == nil {
		return nil
	}

	return resource.Set("code_signing_config", signing.CodeSigningConfig)
}

func resolveResourcePolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(*lambda.GetFunctionOutput)
	if r.Configuration == nil {
		return nil
	}

	c := meta.(*client.Client)
	svc := c.Services().Lambda

	response, err := svc.GetPolicy(ctx, &lambda.GetPolicyInput{
		FunctionName: r.Configuration.FunctionName,
	})
	if err != nil {
		if client.IsAWSError(err, "ResourceNotFoundException") {
			return nil
		}
		return err
	}

	if response == nil {
		return nil
	}
	if err := resource.Set("policy_revision_id", response.RevisionId); err != nil {
		return err
	}
	var policyDocument map[string]any
	err = json.Unmarshal([]byte(*response.Policy), &policyDocument)
	if err != nil {
		return err
	}
	return resource.Set("policy_document", policyDocument)
}

func resolveRuntimeManagementConfig(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(*lambda.GetFunctionOutput)
	if r.Configuration == nil {
		return nil
	}
	c := meta.(*client.Client)
	svc := c.Services().Lambda

	runtimeManagementConfig, err := svc.GetRuntimeManagementConfig(ctx, &lambda.GetRuntimeManagementConfigInput{
		FunctionName: r.Configuration.FunctionName,
	})

	if err != nil {
		if c.IsNotFoundError(err) {
			return nil
		}
		return err
	}

	if err := resource.Set("runtime_version_arn", runtimeManagementConfig.RuntimeVersionArn); err != nil {
		return err
	}

	return resource.Set("update_runtime_on", runtimeManagementConfig.UpdateRuntimeOn)
}
