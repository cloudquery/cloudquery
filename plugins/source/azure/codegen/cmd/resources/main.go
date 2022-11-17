package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource/recipes"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/tables"
	"golang.org/x/sync/errgroup"
)

func generateResources() ([]*resource.Resource, error) {
	var resources []*resource.Resource
	for _, resList := range []func() []*resource.Resource{
		recipes.AppService,
		recipes.Authorization,
		recipes.Batch,
		recipes.CDN,
		recipes.Compute,
		recipes.Container,
		recipes.CosmosDB,
		recipes.DataLake,
		recipes.EventHub,
		recipes.FrontDoor,
		recipes.IoTHub,
		recipes.KeyVault,
		recipes.Logic,
		recipes.MariaDB,
		recipes.Monitor,
		recipes.MySQL,
		recipes.Network,
		recipes.PostgreSQL,
		recipes.Redis,
		recipes.Resource,
		recipes.Search,
		recipes.Security,
		recipes.ServiceBus,
		recipes.SQL,
		recipes.Storage,
		recipes.StreamAnalytics,
		recipes.Subscription,
	} {
		resources = append(resources, resList()...)
	}

	grp, _ := errgroup.WithContext(context.Background())
	for _, res := range resources {
		grp.Go(res.Generate)
	}
	return resources, grp.Wait()
}

func main() {
	resources, err := generateResources()
	if err != nil {
		log.Fatal(err)
	}

	err = tables.Generate(resources)
	if err != nil {
		log.Fatal(err)
	}
}
