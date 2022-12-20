package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/codegen/recipes"
)

func generateResources() error {
	resources := []*recipes.Resource{}
	resources = append(resources, recipes.AddonResources()...)
	resources = append(resources, recipes.IncidentResources()...)
	resources = append(resources, recipes.BusinessServicesResources()...)
	resources = append(resources, recipes.EscalationPolicyResources()...)
	resources = append(resources, recipes.ExtensionSchemaResources()...)
	resources = append(resources, recipes.ExtensionResources()...)
	resources = append(resources, recipes.MaintenanceWindowResources()...)
	resources = append(resources, recipes.PriorityResources()...)
	resources = append(resources, recipes.RulesetResources()...)
	resources = append(resources, recipes.ScheduleResources()...)
	resources = append(resources, recipes.ServiceResources()...)
	resources = append(resources, recipes.TagsResources()...)
	resources = append(resources, recipes.TeamResources()...)
	resources = append(resources, recipes.UserResources()...)
	resources = append(resources, recipes.VendorResources()...)

	err := recipes.SetParentChildRelationships(resources)
	if err != nil {
		return err
	}

	for _, resource := range resources {
		err := resource.Generate()
		if err != nil {
			return err
		}
	}

	if err := recipes.GenerateAllTablesList(resources); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := generateResources(); err != nil {
		log.Fatal(err)
	}
}
