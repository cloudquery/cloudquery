package bigquery

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
	"google.golang.org/api/bigquery/v2"
)

func BigqueryDatasetAccesses() *schema.Table {
	return &schema.Table{
		Name:     "gcp_bigquery_dataset_accesses",
		Resolver: fetchBigqueryDatasetAccesses,
		Columns: []schema.Column{
			{
				Name:     "dataset_cq_id",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIdResolver,
			},
			{
				Name:     "dataset_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("id"),
			},
			{
				Name:          "target_types",
				Description:   "Which resources in the dataset this entry applies to.",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
				Resolver:      resolveBigqueryDatasetAccessTargetTypes,
			},
			{
				Name:        "domain",
				Description: "A domain to grant access to Any users signed in with the domain specified will be granted the specified access Example: \"examplecom\" Maps to IAM policy member \"domain:DOMAIN\"",
				Type:        schema.TypeString,
			},
			{
				Name:        "group_by_email",
				Description: "An email address of a Google Group to grant access to Maps to IAM policy member \"group:GROUP\"",
				Type:        schema.TypeString,
			},
			{
				Name:        "iam_member",
				Description: "Some other type of member that appears in the IAM Policy but isn't a user, group, domain, or special group",
				Type:        schema.TypeString,
			},
			{
				Name:        "role",
				Description: "An IAM role ID that should be granted to the user, group, or domain specified in this access entry The following legacy mappings will be applied: OWNER  roles/bigquerydataOwner WRITER roles/bigquerydataEditor READER  roles/bigquerydataViewer This field will accept any of the above formats, but will return only the legacy format For example, if you set this field to \"roles/bigquerydataOwner\", it will be returned back as \"OWNER\"",
				Type:        schema.TypeString,
			},
			{
				Name:        "routine_dataset_id",
				Description: "The ID of the dataset containing this routine",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Routine.DatasetId"),
			},
			{
				Name:        "routine_project_id",
				Description: "The ID of the project containing this routine",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Routine.ProjectId"),
			},
			{
				Name:        "routine_id",
				Description: "The ID of the routine The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_) The maximum length is 256 characters",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Routine.RoutineId"),
			},
			{
				Name:        "special_group",
				Description: "A special group to grant access to Possible values include: projectOwners: Owners of the enclosing project projectReaders: Readers of the enclosing project projectWriters: Writers of the enclosing project allAuthenticatedUsers: All authenticated BigQuery users Maps to similarly-named IAM members",
				Type:        schema.TypeString,
			},
			{
				Name:        "user_by_email",
				Description: "An email address of a user to grant access to For example: fred@examplecom Maps to IAM policy member \"user:EMAIL\" or \"serviceAccount:EMAIL\"",
				Type:        schema.TypeString,
			},
			{
				Name:        "view_dataset_id",
				Description: "The ID of the dataset containing this table",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("View.DatasetId"),
			},
			{
				Name:        "view_project_id",
				Description: "The ID of the project containing this table",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("View.ProjectId"),
			},
			{
				Name:        "view_table_id",
				Description: "The ID of the table The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_) The maximum length is 1,024 characters",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("View.TableId"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchBigqueryDatasetAccesses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*bigquery.Dataset)
	res <- p.Access
	return nil
}
func resolveBigqueryDatasetAccessTargetTypes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(*bigquery.DatasetAccess)
	if p.Dataset == nil {
		return nil
	}
	return errors.WithStack(resource.Set(c.Name, p.Dataset.TargetTypes))
}
