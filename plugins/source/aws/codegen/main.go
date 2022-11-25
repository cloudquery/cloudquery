package main

import (
	"fmt"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/recipes"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/services"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/tables"
)

func generateResources() ([]*recipes.Resource, error) {
	var resources []*recipes.Resource
	resources = append(resources, recipes.AccessAnalyzerResources()...)
	resources = append(resources, recipes.AccountResources()...)
	resources = append(resources, recipes.ACMResources()...)
	resources = append(resources, recipes.APIGatewayResources()...)
	resources = append(resources, recipes.APIGatewayV2Resources()...)
	resources = append(resources, recipes.ApplicationAutoScalingResources()...)
	resources = append(resources, recipes.ApprunnerResources()...)
	resources = append(resources, recipes.AppstreamResources()...)
	resources = append(resources, recipes.AppSync()...)
	resources = append(resources, recipes.AthenaResources()...)
	resources = append(resources, recipes.AutoscalingResources()...)
	resources = append(resources, recipes.BackupResources()...)
	resources = append(resources, recipes.CloudformationResources()...)
	resources = append(resources, recipes.CloudfrontResources()...)
	resources = append(resources, recipes.CloudHSMV2()...)
	resources = append(resources, recipes.CloudtrailResources()...)
	resources = append(resources, recipes.CloudWatchLogsResources()...)
	resources = append(resources, recipes.CloudwatchResources()...)
	resources = append(resources, recipes.CodeBuildResources()...)
	resources = append(resources, recipes.CodePipelineResources()...)
	resources = append(resources, recipes.CognitoResources()...)
	resources = append(resources, recipes.ConfigResources()...)
	resources = append(resources, recipes.DaxResources()...)
	resources = append(resources, recipes.DirectConnectResources()...)
	resources = append(resources, recipes.DMSResources()...)
	resources = append(resources, recipes.DocumentDBResources()...)
	resources = append(resources, recipes.DynamoDBResources()...)
	resources = append(resources, recipes.EC2Resources()...)
	resources = append(resources, recipes.ECRPublicResources()...)
	resources = append(resources, recipes.ECRResources()...)
	resources = append(resources, recipes.ECSResources()...)
	resources = append(resources, recipes.EFSResources()...)
	resources = append(resources, recipes.EKSResources()...)
	resources = append(resources, recipes.ElastiCacheResources()...)
	resources = append(resources, recipes.ElasticbeanstalkResources()...)
	resources = append(resources, recipes.ElasticsearchResources()...)
	resources = append(resources, recipes.ELBv1Resources()...)
	resources = append(resources, recipes.ELBv2Resources()...)
	resources = append(resources, recipes.EMRResources()...)
	resources = append(resources, recipes.EventbridgeResources()...)
	resources = append(resources, recipes.FirehoseResources()...)
	resources = append(resources, recipes.FraudDetectorResources()...)
	resources = append(resources, recipes.FSXResources()...)
	resources = append(resources, recipes.GlacierResources()...)
	resources = append(resources, recipes.GlueResources()...)
	resources = append(resources, recipes.GuarddutyResources()...)
	resources = append(resources, recipes.IAMResources()...)
	resources = append(resources, recipes.IdentitystoreResources()...)
	resources = append(resources, recipes.Inspector2Resources()...)
	resources = append(resources, recipes.InspectorResources()...)
	resources = append(resources, recipes.IOTResources()...)
	resources = append(resources, recipes.KafkaResources()...)
	resources = append(resources, recipes.KinesisResources()...)
	resources = append(resources, recipes.KMSResources()...)
	resources = append(resources, recipes.LambdaResources()...)
	resources = append(resources, recipes.LightsailResources()...)
	resources = append(resources, recipes.MQResources()...)
	resources = append(resources, recipes.MWAAResources()...)
	resources = append(resources, recipes.NeptuneResources()...)
	resources = append(resources, recipes.OrganizationsResources()...)
	resources = append(resources, recipes.QLDBResources()...)
	resources = append(resources, recipes.QuickSightResources()...)
	resources = append(resources, recipes.RAMResources()...)
	resources = append(resources, recipes.RDSResources()...)
	resources = append(resources, recipes.RedshiftResources()...)
	resources = append(resources, recipes.ResourceGroupsResources()...)
	resources = append(resources, recipes.Route53Resources()...)
	resources = append(resources, recipes.S3Resources()...)
	resources = append(resources, recipes.SagemakerResources()...)
	resources = append(resources, recipes.SchedulerResources()...)
	resources = append(resources, recipes.SecretsManagerResources()...)
	resources = append(resources, recipes.ServiceCatalogResources()...)
	resources = append(resources, recipes.ServiceQuotasResources()...)
	resources = append(resources, recipes.SESResources()...)
	resources = append(resources, recipes.ShieldResources()...)
	resources = append(resources, recipes.SNSResources()...)
	resources = append(resources, recipes.SQSResources()...)
	resources = append(resources, recipes.SSMResources()...)
	resources = append(resources, recipes.SSOAdminResources()...)
	resources = append(resources, recipes.StepFunctionResources()...)
	resources = append(resources, recipes.TimestreamResources()...)
	resources = append(resources, recipes.TransferResources()...)
	resources = append(resources, recipes.WAFRegionalResources()...)
	resources = append(resources, recipes.WAFResources()...)
	resources = append(resources, recipes.WAFv2Resources()...)
	resources = append(resources, recipes.WorkspacesResources()...)
	resources = append(resources, recipes.XRayResources()...)

	err := recipes.SetParentChildRelationships(resources)
	if err != nil {
		return nil, fmt.Errorf("failed to set parent-child relationships: %w", err)
	}
	for _, resource := range resources {
		if err := resource.Generate(); err != nil {
			return nil, err
		}
	}

	return resources, nil
}

func main() {
	err := services.Generate()
	if err != nil {
		log.Fatal(err)
	}

	resources, err := generateResources()
	if err != nil {
		log.Fatal(err)
	}

	err = tables.Generate(resources)
	if err != nil {
		log.Fatal(err)
	}
}
