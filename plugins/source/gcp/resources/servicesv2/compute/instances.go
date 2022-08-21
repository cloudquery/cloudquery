// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"

	"google.golang.org/api/compute/v1"
)

func ComputeInstances() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_instances",
		Resolver:  fetchComputeInstances,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name: "advanced_machine_features",
				Type: schema.TypeJSON,
			},
			{
				Name: "can_ip_forward",
				Type: schema.TypeBool,
			},
			{
				Name: "confidential_instance_config",
				Type: schema.TypeJSON,
			},
			{
				Name: "cpu_platform",
				Type: schema.TypeString,
			},
			{
				Name: "creation_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "deletion_protection",
				Type: schema.TypeBool,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "disks",
				Type: schema.TypeJSON,
			},
			{
				Name: "display_device",
				Type: schema.TypeJSON,
			},
			{
				Name: "fingerprint",
				Type: schema.TypeString,
			},
			{
				Name: "guest_accelerators",
				Type: schema.TypeJSON,
			},
			{
				Name: "hostname",
				Type: schema.TypeString,
			},
			{
				Name: "id",
				Type: schema.TypeInt,
			},
			{
				Name: "key_revocation_action_type",
				Type: schema.TypeString,
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name: "label_fingerprint",
				Type: schema.TypeString,
			},
			{
				Name: "labels",
				Type: schema.TypeJSON,
			},
			{
				Name: "last_start_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "last_stop_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "last_suspended_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "machine_type",
				Type: schema.TypeString,
			},
			{
				Name: "metadata",
				Type: schema.TypeJSON,
			},
			{
				Name: "min_cpu_platform",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "network_interfaces",
				Type: schema.TypeJSON,
			},
			{
				Name: "network_performance_config",
				Type: schema.TypeJSON,
			},
			{
				Name: "params",
				Type: schema.TypeJSON,
			},
			{
				Name: "private_ipv_6_google_access",
				Type: schema.TypeString,
			},
			{
				Name: "reservation_affinity",
				Type: schema.TypeJSON,
			},
			{
				Name: "resource_policies",
				Type: schema.TypeStringArray,
			},
			{
				Name: "satisfies_pzs",
				Type: schema.TypeBool,
			},
			{
				Name: "scheduling",
				Type: schema.TypeJSON,
			},
			{
				Name: "service_accounts",
				Type: schema.TypeJSON,
			},
			{
				Name: "shielded_instance_config",
				Type: schema.TypeJSON,
			},
			{
				Name: "shielded_instance_integrity_policy",
				Type: schema.TypeJSON,
			},
			{
				Name: "source_machine_image",
				Type: schema.TypeString,
			},
			{
				Name: "source_machine_image_encryption_key",
				Type: schema.TypeJSON,
			},
			{
				Name: "start_restricted",
				Type: schema.TypeBool,
			},
			{
				Name: "status",
				Type: schema.TypeString,
			},
			{
				Name: "status_message",
				Type: schema.TypeString,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
			{
				Name: "zone",
				Type: schema.TypeString,
			},
			{
				Name: "server_response",
				Type: schema.TypeJSON,
			},
		},
	}
}

func fetchComputeInstances(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Compute.Instances.AggregatedList(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		var allItems []*compute.Instance
		for _, items := range output.Items {
			allItems = append(allItems, items.Instances...)
		}
		res <- allItems

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
