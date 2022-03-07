package security

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SecurityAssessments() *schema.Table {
	return &schema.Table{
		Name:         "azure_security_assessments",
		Description:  "Assessment security assessment on a resource",
		Resolver:     fetchSecurityAssessments,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "display_name",
				Description: "User friendly display name of the assessment.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssessmentProperties.DisplayName"),
			},
			{
				Name:        "code",
				Description: "Programmatic code for the status of the assessment.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssessmentProperties.Status.Code"),
			},
			{
				Name:        "cause",
				Description: "Programmatic code for the cause of the assessment status.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssessmentProperties.Status.Cause"),
			},
			{
				Name:        "description",
				Description: "Human readable description of the assessment status.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssessmentProperties.Status.Description"),
			},
			{
				Name:        "additional_data",
				Description: "Additional data regarding the assessment.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("AssessmentProperties.AdditionalData"),
			},
			{
				Name:        "azure_portal_uri",
				Description: "Link to assessment in Azure Portal.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssessmentProperties.Links.AzurePortalURI"),
			},
			{
				Name:        "metadata_display_name",
				Description: "User friendly display name of the assessment.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssessmentProperties.Metadata.DisplayName"),
			},
			{
				Name:        "metadata_policy_definition_id",
				Description: "Azure resource ID of the policy definition that turns this assessment calculation on.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssessmentProperties.Metadata.PolicyDefinitionID"),
			},
			{
				Name:        "metadata_description",
				Description: "Human readable description of the assessment.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssessmentProperties.Metadata.Description"),
			},
			{
				Name:        "metadata_remediation_description",
				Description: "Human readable description of what you should do to mitigate this security issue.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssessmentProperties.Metadata.RemediationDescription"),
			},
			{
				Name:     "metadata_categories",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AssessmentProperties.Metadata.Categories"),
			},
			{
				Name:        "metadata_severity",
				Description: "The severity level of the assessment.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssessmentProperties.Metadata.Severity"),
			},
			{
				Name:        "metadata_user_impact",
				Description: "The user impact of the assessment.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssessmentProperties.Metadata.UserImpact"),
			},
			{
				Name:        "metadata_implementation_effort",
				Description: "The implementation effort required to remediate this assessment.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssessmentProperties.Metadata.ImplementationEffort"),
			},
			{
				Name:     "metadata_threats",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AssessmentProperties.Metadata.Threats"),
			},
			{
				Name:        "metadata_preview",
				Description: "True if this assessment is in preview release status.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AssessmentProperties.Metadata.Preview"),
			},
			{
				Name:        "metadata_assessment_type",
				Description: "BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssessmentProperties.Metadata.AssessmentType"),
			},
			{
				Name:        "metadata_partner_data_partner_name",
				Description: "Name of the company of the partner.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssessmentProperties.Metadata.PartnerData.PartnerName"),
			},
			{
				Name:        "metadata_partner_data_product_name",
				Description: "Name of the product of the partner that created the assessment.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssessmentProperties.Metadata.PartnerData.ProductName"),
			},
			{
				Name:        "partner_name",
				Description: "Name of the company of the partner",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssessmentProperties.PartnersData.PartnerName"),
			},
			{
				Name:        "id",
				Description: "Resource Id.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type.",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_details",
				Description: "Assessed resource details.",
				Type:        schema.TypeJSON,
				Resolver:    resolveSecurityAssessmentResourceDetails,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSecurityAssessments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Security.Assessments
	result, err := svc.List(ctx, fmt.Sprintf("/subscriptions/%s", cl.SubscriptionId))
	if err != nil {
		return err
	}
	for result.NotDone() {
		res <- result.Values()
		if err := result.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}

func resolveSecurityAssessmentResourceDetails(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	a := resource.Item.(security.Assessment)
	if a.AssessmentProperties == nil {
		return nil
	}
	b, err := json.Marshal(a.AssessmentProperties.ResourceDetails)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
