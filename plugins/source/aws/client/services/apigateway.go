// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
)

//go:generate mockgen -package=mocks -destination=../mocks/apigateway.go -source=apigateway.go ApigatewayClient
type ApigatewayClient interface {
	GetAccount(context.Context, *apigateway.GetAccountInput, ...func(*apigateway.Options)) (*apigateway.GetAccountOutput, error)
	GetApiKey(context.Context, *apigateway.GetApiKeyInput, ...func(*apigateway.Options)) (*apigateway.GetApiKeyOutput, error)
	GetApiKeys(context.Context, *apigateway.GetApiKeysInput, ...func(*apigateway.Options)) (*apigateway.GetApiKeysOutput, error)
	GetAuthorizer(context.Context, *apigateway.GetAuthorizerInput, ...func(*apigateway.Options)) (*apigateway.GetAuthorizerOutput, error)
	GetAuthorizers(context.Context, *apigateway.GetAuthorizersInput, ...func(*apigateway.Options)) (*apigateway.GetAuthorizersOutput, error)
	GetBasePathMapping(context.Context, *apigateway.GetBasePathMappingInput, ...func(*apigateway.Options)) (*apigateway.GetBasePathMappingOutput, error)
	GetBasePathMappings(context.Context, *apigateway.GetBasePathMappingsInput, ...func(*apigateway.Options)) (*apigateway.GetBasePathMappingsOutput, error)
	GetClientCertificate(context.Context, *apigateway.GetClientCertificateInput, ...func(*apigateway.Options)) (*apigateway.GetClientCertificateOutput, error)
	GetClientCertificates(context.Context, *apigateway.GetClientCertificatesInput, ...func(*apigateway.Options)) (*apigateway.GetClientCertificatesOutput, error)
	GetDeployment(context.Context, *apigateway.GetDeploymentInput, ...func(*apigateway.Options)) (*apigateway.GetDeploymentOutput, error)
	GetDeployments(context.Context, *apigateway.GetDeploymentsInput, ...func(*apigateway.Options)) (*apigateway.GetDeploymentsOutput, error)
	GetDocumentationPart(context.Context, *apigateway.GetDocumentationPartInput, ...func(*apigateway.Options)) (*apigateway.GetDocumentationPartOutput, error)
	GetDocumentationParts(context.Context, *apigateway.GetDocumentationPartsInput, ...func(*apigateway.Options)) (*apigateway.GetDocumentationPartsOutput, error)
	GetDocumentationVersion(context.Context, *apigateway.GetDocumentationVersionInput, ...func(*apigateway.Options)) (*apigateway.GetDocumentationVersionOutput, error)
	GetDocumentationVersions(context.Context, *apigateway.GetDocumentationVersionsInput, ...func(*apigateway.Options)) (*apigateway.GetDocumentationVersionsOutput, error)
	GetDomainName(context.Context, *apigateway.GetDomainNameInput, ...func(*apigateway.Options)) (*apigateway.GetDomainNameOutput, error)
	GetDomainNames(context.Context, *apigateway.GetDomainNamesInput, ...func(*apigateway.Options)) (*apigateway.GetDomainNamesOutput, error)
	GetExport(context.Context, *apigateway.GetExportInput, ...func(*apigateway.Options)) (*apigateway.GetExportOutput, error)
	GetGatewayResponse(context.Context, *apigateway.GetGatewayResponseInput, ...func(*apigateway.Options)) (*apigateway.GetGatewayResponseOutput, error)
	GetGatewayResponses(context.Context, *apigateway.GetGatewayResponsesInput, ...func(*apigateway.Options)) (*apigateway.GetGatewayResponsesOutput, error)
	GetIntegration(context.Context, *apigateway.GetIntegrationInput, ...func(*apigateway.Options)) (*apigateway.GetIntegrationOutput, error)
	GetIntegrationResponse(context.Context, *apigateway.GetIntegrationResponseInput, ...func(*apigateway.Options)) (*apigateway.GetIntegrationResponseOutput, error)
	GetMethod(context.Context, *apigateway.GetMethodInput, ...func(*apigateway.Options)) (*apigateway.GetMethodOutput, error)
	GetMethodResponse(context.Context, *apigateway.GetMethodResponseInput, ...func(*apigateway.Options)) (*apigateway.GetMethodResponseOutput, error)
	GetModel(context.Context, *apigateway.GetModelInput, ...func(*apigateway.Options)) (*apigateway.GetModelOutput, error)
	GetModelTemplate(context.Context, *apigateway.GetModelTemplateInput, ...func(*apigateway.Options)) (*apigateway.GetModelTemplateOutput, error)
	GetModels(context.Context, *apigateway.GetModelsInput, ...func(*apigateway.Options)) (*apigateway.GetModelsOutput, error)
	GetRequestValidator(context.Context, *apigateway.GetRequestValidatorInput, ...func(*apigateway.Options)) (*apigateway.GetRequestValidatorOutput, error)
	GetRequestValidators(context.Context, *apigateway.GetRequestValidatorsInput, ...func(*apigateway.Options)) (*apigateway.GetRequestValidatorsOutput, error)
	GetResource(context.Context, *apigateway.GetResourceInput, ...func(*apigateway.Options)) (*apigateway.GetResourceOutput, error)
	GetResources(context.Context, *apigateway.GetResourcesInput, ...func(*apigateway.Options)) (*apigateway.GetResourcesOutput, error)
	GetRestApi(context.Context, *apigateway.GetRestApiInput, ...func(*apigateway.Options)) (*apigateway.GetRestApiOutput, error)
	GetRestApis(context.Context, *apigateway.GetRestApisInput, ...func(*apigateway.Options)) (*apigateway.GetRestApisOutput, error)
	GetSdk(context.Context, *apigateway.GetSdkInput, ...func(*apigateway.Options)) (*apigateway.GetSdkOutput, error)
	GetSdkType(context.Context, *apigateway.GetSdkTypeInput, ...func(*apigateway.Options)) (*apigateway.GetSdkTypeOutput, error)
	GetSdkTypes(context.Context, *apigateway.GetSdkTypesInput, ...func(*apigateway.Options)) (*apigateway.GetSdkTypesOutput, error)
	GetStage(context.Context, *apigateway.GetStageInput, ...func(*apigateway.Options)) (*apigateway.GetStageOutput, error)
	GetStages(context.Context, *apigateway.GetStagesInput, ...func(*apigateway.Options)) (*apigateway.GetStagesOutput, error)
	GetTags(context.Context, *apigateway.GetTagsInput, ...func(*apigateway.Options)) (*apigateway.GetTagsOutput, error)
	GetUsage(context.Context, *apigateway.GetUsageInput, ...func(*apigateway.Options)) (*apigateway.GetUsageOutput, error)
	GetUsagePlan(context.Context, *apigateway.GetUsagePlanInput, ...func(*apigateway.Options)) (*apigateway.GetUsagePlanOutput, error)
	GetUsagePlanKey(context.Context, *apigateway.GetUsagePlanKeyInput, ...func(*apigateway.Options)) (*apigateway.GetUsagePlanKeyOutput, error)
	GetUsagePlanKeys(context.Context, *apigateway.GetUsagePlanKeysInput, ...func(*apigateway.Options)) (*apigateway.GetUsagePlanKeysOutput, error)
	GetUsagePlans(context.Context, *apigateway.GetUsagePlansInput, ...func(*apigateway.Options)) (*apigateway.GetUsagePlansOutput, error)
	GetVpcLink(context.Context, *apigateway.GetVpcLinkInput, ...func(*apigateway.Options)) (*apigateway.GetVpcLinkOutput, error)
	GetVpcLinks(context.Context, *apigateway.GetVpcLinksInput, ...func(*apigateway.Options)) (*apigateway.GetVpcLinksOutput, error)
}
