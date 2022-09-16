package lambda

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type AliasWrapper struct {
	*types.AliasConfiguration
	UrlConfig *lambda.GetFunctionUrlConfigOutput
}

func fetchLambdaFunctions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input lambda.ListFunctionsInput
	c := meta.(*client.Client)
	svc := c.Services().Lambda
	for {
		response, err := svc.ListFunctions(ctx, &input)
		if err != nil {
			return err
		}

		for _, f := range response.Functions {
			getFunctionInput := lambda.GetFunctionInput{
				FunctionName: f.FunctionName,
			}
			funcResponse, err := svc.GetFunction(ctx, &getFunctionInput, func(options *lambda.Options) {
				options.Region = c.Region
			})

			if err != nil {
				if c.IsNotFoundError(err) || c.IsAccessDeniedError(err) {
					c.Logger().Warn().Err(err).Msg("Failed to get function")
					res <- &lambda.GetFunctionOutput{
						Configuration: &f,
					}
					continue
				}
				return err
			}
			res <- funcResponse
		}

		if aws.ToString(response.NextMarker) == "" {
			break
		}
		input.Marker = response.NextMarker
	}
	return nil
}
func resolvePolicyCodeSigningConfig(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
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

	if response != nil {
		if err := resource.Set("policy_revision_id", response.RevisionId); err != nil {
			return err
		}
		var policyDocument map[string]interface{}
		err = json.Unmarshal([]byte(*response.Policy), &policyDocument)
		if err != nil {
			return err
		}
		if err := resource.Set("policy_document", policyDocument); err != nil {
			return err
		}
	}

	// skip getting CodeSigningConfig since containerized lambda functions does not support this feature
	// value can be nil if the caller doesn't have GetFunctionConfiguration permission and only has List*
	lambdaType, ok := resource.Get("code_repository_type").(*string)
	if !ok || *lambdaType == "ECR" {
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

func fetchLambdaFunctionEventInvokeConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*lambda.GetFunctionOutput)
	if p.Configuration == nil {
		return nil
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Lambda
	config := lambda.ListFunctionEventInvokeConfigsInput{
		FunctionName: p.Configuration.FunctionName,
	}

	for {
		output, err := svc.ListFunctionEventInvokeConfigs(ctx, &config)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- output.FunctionEventInvokeConfigs
		if output.NextMarker == nil {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}
func fetchLambdaFunctionAliases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*lambda.GetFunctionOutput)
	if p.Configuration == nil {
		return nil
	}

	c := meta.(*client.Client)
	svc := c.Services().Lambda
	config := lambda.ListAliasesInput{
		FunctionName: p.Configuration.FunctionName,
	}

	for {
		output, err := svc.ListAliases(ctx, &config)
		if err != nil {
			return err
		}
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		aliases := make([]AliasWrapper, 0, len(output.Aliases))
		for _, a := range output.Aliases {
			alias := a
			urlConfig, err := svc.GetFunctionUrlConfig(ctx, &lambda.GetFunctionUrlConfigInput{
				FunctionName: p.Configuration.FunctionName,
				Qualifier:    alias.Name,
			})
			if err != nil && !c.IsNotFoundError(err) {
				return err
			}
			aliases = append(aliases, AliasWrapper{&alias, urlConfig})
		}
		res <- aliases
		if output.NextMarker == nil {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}

func fetchLambdaFunctionVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*lambda.GetFunctionOutput)
	if p.Configuration == nil {
		return nil
	}

	svc := meta.(*client.Client).Services().Lambda
	config := lambda.ListVersionsByFunctionInput{
		FunctionName: p.Configuration.FunctionName,
	}

	for {
		output, err := svc.ListVersionsByFunction(ctx, &config)
		if err != nil {
			if meta.(*client.Client).IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- output.Versions
		if output.NextMarker == nil {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}

func fetchLambdaFunctionConcurrencyConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*lambda.GetFunctionOutput)
	if p.Configuration == nil {
		return nil
	}

	cl := meta.(*client.Client)
	svc := cl.Services().Lambda
	config := lambda.ListProvisionedConcurrencyConfigsInput{
		FunctionName: p.Configuration.FunctionName,
	}

	for {
		output, err := svc.ListProvisionedConcurrencyConfigs(ctx, &config)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- output.ProvisionedConcurrencyConfigs
		if output.NextMarker == nil {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}
func fetchLambdaFunctionEventSourceMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*lambda.GetFunctionOutput)
	if p.Configuration == nil {
		return nil
	}

	cl := meta.(*client.Client)
	svc := cl.Services().Lambda
	config := lambda.ListEventSourceMappingsInput{
		FunctionName: p.Configuration.FunctionName,
	}

	for {
		output, err := svc.ListEventSourceMappings(ctx, &config)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- output.EventSourceMappings
		if output.NextMarker == nil {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}
