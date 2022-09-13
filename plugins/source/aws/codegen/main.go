package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/recipes"
)

func main() {
	resources := recipes.Ec2Resources()
	resources = append(resources, recipes.AccessAnalyzerResources()...)
	resources = append(resources, recipes.ACMResources()...)
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
