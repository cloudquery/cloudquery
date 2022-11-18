package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/lightsail/models"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func LightsailResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "alarms",
			Struct:      &types.Alarm{},
			Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Alarm.html",
			SkipFields:  []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:    "arn",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:   "buckets",
			Struct:       &types.Bucket{},
			Description:  "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Bucket.html",
			ExtraColumns: defaultRegionalColumns,
			Relations: []string{
				"BucketAccessKeys()",
			},
		},
		{
			SubService:  "bucket_access_keys",
			Struct:      &types.AccessKey{},
			Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_AccessKey.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "bucket_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "certificates",
			Struct:      &types.Certificate{},
			Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Certificate.html",
			SkipFields:  []string{"Tags"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTags`,
					},
				}...),
		},
		{
			SubService:  "container_services",
			Struct:      &types.ContainerService{},
			Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_ContainerService.html",
			SkipFields:  []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:    "arn",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"ContainerServiceDeployments()",
				"ContainerServiceImages()",
			},
		},
		{
			SubService:  "container_service_deployments",
			Struct:      &types.ContainerServiceDeployment{},
			Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_ContainerServiceDeployment.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "container_service_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "container_service_images",
			Struct:      &types.ContainerImage{},
			Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_ContainerImage.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "container_service_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "database_snapshots",
			Struct:      &types.RelationalDatabaseSnapshot{},
			Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_RelationalDatabaseSnapshot.html",
			SkipFields:  []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:    "arn",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:   "databases",
			Struct:       &types.RelationalDatabase{},
			Description:  "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_RelationalDatabase.html",
			ExtraColumns: defaultRegionalColumns,
			Relations: []string{
				"DatabaseParameters()",
				"DatabaseEvents()",
				"DatabaseLogEvents()",
			},
		},
		{
			SubService:  "database_parameters",
			Struct:      &types.RelationalDatabaseParameter{},
			Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_RelationalDatabaseParameter.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "database_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "database_events",
			Struct:      &types.RelationalDatabaseEvent{},
			Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_RelationalDatabaseEvent.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "database_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "database_log_events",
			Struct:     &models.LogEventWrapper{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "database_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "disks",
			Struct:      &types.Disk{},
			Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Disk.html",
			SkipFields:  []string{"Arn", "Tags"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:    "arn",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTags`,
					},
				}...),
			Relations: []string{
				"DiskSnapshots()",
			},
		},
		{
			SubService:  "disk_snapshots",
			Struct:      &types.DiskSnapshot{},
			Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_DiskSnapshot.html",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "disk_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "distributions",
			Struct:     &models.DistributionWrapper{},
			SkipFields: []string{"Arn", "Tags"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTags`,
					},
				}...),
		},
		{
			SubService:  "instance_snapshots",
			Struct:      &types.InstanceSnapshot{},
			Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_InstanceSnapshot.html",
			SkipFields:  []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:    "arn",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "instances",
			Struct:      &types.Instance{},
			Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Instance.html",
			SkipFields:  []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "access_details",
						Type:     schema.TypeJSON,
						Resolver: `resolveLightsailInstanceAccessDetails`,
					},
					{
						Name:    "arn",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"InstancePortStates()",
			},
		},
		{
			SubService:  "instance_port_states",
			Struct:      &types.InstancePortState{},
			Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_InstancePortState.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "instance_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "load_balancers",
			Struct:      &types.LoadBalancer{},
			Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_LoadBalancer.html",
			SkipFields:  []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:    "arn",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"LoadBalancerTlsCertificates()",
			},
		},
		{
			SubService:  "load_balancer_tls_certificates",
			Struct:      &types.LoadBalancerTlsCertificate{},
			Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_LoadBalancerTlsCertificate.html",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "load_balancer_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "static_ips",
			Struct:      &types.StaticIp{},
			Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_StaticIp.html",
			SkipFields:  []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:    "arn",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "lightsail"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("lightsail")`
		r.UnwrapEmbeddedStructs = true
	}
	return resources
}
