package lambda

import (
	"context"
	"encoding/json"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func Functions() *schema.Table {
	tableName := "aws_lambda_functions"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/lambda/latest/dg/API_GetFunction.html`,
		Resolver:            fetchLambdaFunctions,
		PreResourceResolver: getFunction,
		Transform:           transformers.TransformWithStruct(&lambda.GetFunctionOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "lambda"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Configuration.FunctionArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name: "policy_revision_id",
				Type: arrow.BinaryTypes.String,
				// resolved in resolveResourcePolicy
			},
			{
				Name:     "policy_document",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveResourcePolicy,
			},
			{
				Name:     "code_signing_config",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveCodeSigningConfig,
			},
			{
				Name:     "code_repository_type",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Code.RepositoryType"),
			},
			{
				Name: "update_runtime_on",
				Type: arrow.BinaryTypes.String,
				// resolved in resolveRuntimeManagementConfig
			},
			{
				Name:     "runtime_version_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: resolveRuntimeManagementConfig,
			},
			{
				Name: "code",
				Type: sdkTypes.ExtensionTypes.JSON,
			},
			{
				Name:     "concurrency",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveConcurrency,
			},
			{
				Name: "configuration",
				Type: sdkTypes.ExtensionTypes.JSON,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveTags,
			},
		},

		Relations: []*schema.Table{
			functionAliases(),
			functionConcurrencyConfigs(),
			functionEventInvokeConfigs(),
			functionEventSourceMappings(),
			functionURLConfigs(),
			functionVersions(),
		},
	}
}

func fetchLambdaFunctions(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceLambda).Lambda
	paginator := lambda.NewListFunctionsPaginator(svc, &lambda.ListFunctionsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *lambda.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Functions
	}
	return nil
}

func getFunction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceLambda).Lambda
	f := resource.Item.(types.FunctionConfiguration)
	funcResponse, err := svc.GetFunction(ctx, &lambda.GetFunctionInput{
		FunctionName: f.FunctionName,
	}, func(options *lambda.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
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
			cl.Logger().Warn().Err(err).Msg("configuration data retrieved from ListFunctions will still be persisted")
			return nil
		}
		return err
	}
	if funcResponse.Code != nil {
		cl.Logger().Warn().Msg("location of lambda function redacted for security purposes")
		funcResponse.Code.Location = aws.String("REDACTED_FOR_SECURITY_PURPOSES")
	}
	resource.Item = funcResponse
	return nil
}

func resolveCodeSigningConfig(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(*lambda.GetFunctionOutput)
	if r.Configuration == nil {
		return nil
	}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceLambda).Lambda

	// skip getting CodeSigningConfig since containerized lambda functions does not support this feature
	if r.Configuration.PackageType == types.PackageTypeImage {
		return nil
	}

	functionSigning, err := svc.GetFunctionCodeSigningConfig(ctx, &lambda.GetFunctionCodeSigningConfigInput{
		FunctionName: r.Configuration.FunctionName,
	}, func(options *lambda.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	if functionSigning.CodeSigningConfigArn == nil {
		return nil
	}

	signing, err := svc.GetCodeSigningConfig(ctx, &lambda.GetCodeSigningConfigInput{
		CodeSigningConfigArn: functionSigning.CodeSigningConfigArn,
	}, func(options *lambda.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
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

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceLambda).Lambda

	response, err := svc.GetPolicy(ctx, &lambda.GetPolicyInput{
		FunctionName: r.Configuration.FunctionName,
	}, func(options *lambda.Options) {
		options.Region = cl.Region
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
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceLambda).Lambda

	// skip getting GetRuntimeManagementConfig since containerized lambda functions does not support this feature
	if r.Configuration.PackageType == types.PackageTypeImage {
		return nil
	}

	runtimeManagementConfig, err := svc.GetRuntimeManagementConfig(ctx, &lambda.GetRuntimeManagementConfigInput{
		FunctionName: r.Configuration.FunctionName,
	}, func(options *lambda.Options) {
		options.Region = cl.Region
	})

	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}

	if err := resource.Set("runtime_version_arn", runtimeManagementConfig.RuntimeVersionArn); err != nil {
		return err
	}

	return resource.Set("update_runtime_on", runtimeManagementConfig.UpdateRuntimeOn)
}

func resolveConcurrency(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	r := resource.Item.(*lambda.GetFunctionOutput)
	// No way of getting functionName
	if r.Configuration == nil {
		return nil
	}

	// setting concurrency value from GetFunction call
	if r.Code != nil {
		return resource.Set(col.Name, r.Concurrency)
	}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceLambda).Lambda

	functionConcurrency, err := svc.GetFunctionConcurrency(ctx, &lambda.GetFunctionConcurrencyInput{
		FunctionName: r.Configuration.FunctionName,
	}, func(options *lambda.Options) {
		options.Region = cl.Region
	})

	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}

	// convert from lambda.GetFunctionConcurrencyOutput to types.Concurrency

	data, err := json.Marshal(functionConcurrency)
	if err != nil {
		return err
	}
	var funcConcurrency types.Concurrency
	err = json.Unmarshal(data, &funcConcurrency)
	if err != nil {
		return err
	}

	return resource.Set(col.Name, functionConcurrency)
}

func resolveTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	r := resource.Item.(*lambda.GetFunctionOutput)
	// No way of getting functionName
	if r.Configuration == nil {
		return nil
	}

	// setting tags value from GetFunction call
	if r.Code != nil {
		return resource.Set(col.Name, r.Tags)
	}

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceLambda).Lambda

	funcTags, err := svc.ListTags(ctx, &lambda.ListTagsInput{
		Resource: r.Configuration.FunctionArn,
	}, func(options *lambda.Options) {
		options.Region = cl.Region
	})

	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(col.Name, funcTags.Tags)
}
