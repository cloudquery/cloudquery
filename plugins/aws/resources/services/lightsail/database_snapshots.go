package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource database_snapshots --config gen.hcl --output .
func DatabaseSnapshots() *schema.Table {
	return &schema.Table{
		Name:          "aws_lightsail_database_snapshots",
		Description:   "Describes a database snapshot",
		Resolver:      fetchLightsailDatabaseSnapshots,
		Multiplex:     client.ServiceAccountRegionMultiplexer("lightsail"),
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		IgnoreInTests: true, // can't be created using terraform.
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the database snapshot",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_at",
				Description: "The timestamp when the database snapshot was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "engine",
				Description: "The software of the database snapshot (for example, MySQL)",
				Type:        schema.TypeString,
			},
			{
				Name:        "engine_version",
				Description: "The database engine version for the database snapshot (for example, 5723)",
				Type:        schema.TypeString,
			},
			{
				Name:        "from_relational_database_arn",
				Description: "The Amazon Resource Name (ARN) of the database from which the database snapshot was created",
				Type:        schema.TypeString,
			},
			{
				Name:        "from_relational_database_blueprint_id",
				Description: "The blueprint ID of the database from which the database snapshot was created",
				Type:        schema.TypeString,
			},
			{
				Name:        "from_relational_database_bundle_id",
				Description: "The bundle ID of the database from which the database snapshot was created",
				Type:        schema.TypeString,
			},
			{
				Name:        "from_relational_database_name",
				Description: "The name of the source database from which the database snapshot was created",
				Type:        schema.TypeString,
			},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Location.AvailabilityZone"),
			},
			{
				Name:        "name",
				Description: "The name of the database snapshot",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_type",
				Description: "The Lightsail resource type",
				Type:        schema.TypeString,
			},
			{
				Name:        "size_in_gb",
				Description: "The size of the disk in GB (for example, 32) for the database snapshot",
				Type:        schema.TypeInt,
			},
			{
				Name:        "state",
				Description: "The state of the database snapshot",
				Type:        schema.TypeString,
			},
			{
				Name:        "support_code",
				Description: "The support code for the database snapshot",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The tag keys and optional values for the resource",
				Type:        schema.TypeJSON,
				Resolver:    resolveDatabaseSnapshotsTags,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchLightsailDatabaseSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input lightsail.GetRelationalDatabaseSnapshotsInput
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	for {
		response, err := svc.GetRelationalDatabaseSnapshots(ctx, &input, func(options *lightsail.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.RelationalDatabaseSnapshots
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
func resolveDatabaseSnapshotsTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.RelationalDatabaseSnapshot)
	tags := make(map[string]string)
	client.TagsIntoMap(r.Tags, tags)
	return diag.WrapError(resource.Set(c.Name, tags))
}
