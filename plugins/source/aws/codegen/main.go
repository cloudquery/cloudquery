package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/recipes"
)

func main() {
	resources := recipes.Ec2Resources()
	resources = append(resources, recipes.AccessAnalyzerResources()...)
	resources = append(resources, recipes.ACMResources()...)
	resources = append(resources, recipes.ApiGatewayV2Resources()...)
	resources = append(resources, recipes.ApiGatewayesources()...)
	resources = append(resources, recipes.ApplicationAutoScalingResources()...)
	resources = append(resources, recipes.AppSync()...)
	resources = append(resources, recipes.CloudWatchLogsResources()...)
	resources = append(resources, recipes.CodeBuildResources()...)
	resources = append(resources, recipes.CodePipelineResources()...)
	resources = append(resources, recipes.EcsResources()...)
	resources = append(resources, recipes.EcrResources()...)
	resources = append(resources, recipes.EfsResources()...)
	resources = append(resources, recipes.EksResources()...)
	resources = append(resources, recipes.DaxResources()...)
	resources = append(resources, recipes.ElastiCacheResources()...)

	resources = append(resources, recipes.IOTResources()...)
	resources = append(resources, recipes.KinesisResources()...)
	resources = append(resources, recipes.KMSResources()...)
	resources = append(resources, recipes.LambdaResources()...)
	resources = append(resources, recipes.LightsailResources()...)
	resources = append(resources, recipes.MQResources()...)
	resources = append(resources, recipes.QLDBResources()...)
	resources = append(resources, recipes.RDSResources()...)
	resources = append(resources, recipes.RedshiftResources()...)
	resources = append(resources, recipes.ResourceGroupsResources()...)
	resources = append(resources, recipes.Route53Resources()...)
	resources = append(resources, recipes.S3Resources()...)
	resources = append(resources, recipes.SagemakerResources()...)
	resources = append(resources, recipes.SecretsManagerResources()...)
	resources = append(resources, recipes.SESResources()...)
	resources = append(resources, recipes.ShieldResources()...)
	resources = append(resources, recipes.SNSResources()...)
	resources = append(resources, recipes.SQSResources()...)
	resources = append(resources, recipes.SSMResources()...)
	resources = append(resources, recipes.TransferResources()...)
	resources = append(resources, recipes.WAFResources()...)
	resources = append(resources, recipes.WAFRegionalResources()...)
	resources = append(resources, recipes.WAFv2Resources()...)
	resources = append(resources, recipes.XRayResources()...)
	resources = append(resources, recipes.OrganizationsResources()...)
	resources = append(resources, recipes.WorkspacesResources()...)
	for _, resource := range resources {
		if err := resource.Generate(); err != nil {
			log.Fatal(err)
		}
	}
}
