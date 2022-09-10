package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Connections() *schema.Table {
	return &schema.Table{
		Name:        "aws_glue_connections",
		Description: "Defines a connection to a data source",
		Resolver:    fetchGlueConnections,
		Multiplex:   client.ServiceAccountRegionMultiplexer("glue"),
		Columns: []schema.Column{
			{
				Name:            "arn",
				Description:     "ARN of the resource",
				Type:            schema.TypeString,
				Resolver:        resolveGlueConnectionArn,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "connection_properties",
				Description: "Key-value pairs that define parameters for the connection",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "connection_type",
				Description: "The type of the connection",
				Type:        schema.TypeString,
			},
			{
				Name:        "creation_time",
				Description: "The time that this connection definition was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "The description of the connection",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_updated_by",
				Description: "The user, group, or role that last updated this connection definition",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_updated_time",
				Description: "The last time that this connection definition was updated",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "match_criteria",
				Description: "A list of criteria that can be used in selecting this connection",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "name",
				Description: "The name of the connection definition",
				Type:        schema.TypeString,
			},
			{
				Name:     "physical_connection_requirements",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PhysicalConnectionRequirements"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchGlueConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Glue
	input := glue.GetConnectionsInput{}
	for {
		output, err := svc.GetConnections(ctx, &input)
		if err != nil {
			return err
		}
		res <- output.ConnectionList

		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}
func resolveGlueConnectionArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	r := resource.Item.(types.Connection)
	arn := aws.String(connectionARN(cl, &r))
	return resource.Set(c.Name, arn)
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func connectionARN(cl *client.Client, c *types.Connection) string {
	return cl.ARN(client.GlueService, "connection", *c.Name)
}
