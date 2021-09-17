package resources

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func AdUsers() *schema.Table {
	return &schema.Table{
		Name:         "azure_ad_users",
		Description:  "User active Directory user information",
		Resolver:     fetchAdUsers,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "object_id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "immutable_id",
				Description: "This must be specified if you are using a federated domain for the user's userPrincipalName (UPN) property when creating a new user account It is used to associate an on-premises Active Directory user account with their Azure AD user object",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ImmutableID"),
			},
			{
				Name:        "usage_location",
				Description: "A two letter country code (ISO standard 3166) Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries Examples include: \"US\", \"JP\", and \"GB\"",
				Type:        schema.TypeString,
			},
			{
				Name:        "given_name",
				Description: "The given name for the user",
				Type:        schema.TypeString,
			},
			{
				Name:        "surname",
				Description: "The user's surname (family name or last name)",
				Type:        schema.TypeString,
			},
			{
				Name:        "user_type",
				Description: "A string value that can be used to classify user types in your directory, such as 'Member' and 'Guest' Possible values include: 'Member', 'Guest'",
				Type:        schema.TypeString,
			},
			{
				Name:        "account_enabled",
				Description: "Whether the account is enabled",
				Type:        schema.TypeBool,
			},
			{
				Name:        "display_name",
				Description: "The display name of the user",
				Type:        schema.TypeString,
			},
			{
				Name:        "user_principal_name",
				Description: "The principal name of the user",
				Type:        schema.TypeString,
			},
			{
				Name:        "mail_nickname",
				Description: "The mail alias for the user",
				Type:        schema.TypeString,
			},
			{
				Name:        "mail",
				Description: "The primary email address of the user",
				Type:        schema.TypeString,
			},
			{
				Name:        "additional_properties",
				Description: "Unmatched properties from the message are deserialized this collection",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "object_id",
				Description: "The object ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectID"),
			},
			{
				Name:     "deletion_timestamp_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DeletionTimestamp.Time"),
			},
			{
				Name:        "object_type",
				Description: "Possible values include: 'ObjectTypeDirectoryObject', 'ObjectTypeApplication', 'ObjectTypeGroup', 'ObjectTypeServicePrincipal', 'ObjectTypeUser'",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_ad_user_sign_in_names",
				Description: "SignInName contains information about a sign-in name of a local account user in an Azure Active Directory B2C tenant",
				Resolver:    fetchAdUserSignInNames,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"user_cq_id", "signin_type", "signin_value"}},
				Columns: []schema.Column{
					{
						Name:        "user_cq_id",
						Description: "Unique ID of azure_ad_users table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "additional_properties",
						Description: "Unmatched properties from the message are deserialized this collection",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "signin_type",
						Description: "A string value that can be used to classify user sign-in types in your directory, such as 'emailAddress' or 'userName'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Type"),
					},
					{
						Name:        "signin_value",
						Description: "The sign-in used by the local account Must be unique across the company/tenant For example, 'johnc@examplecom'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Value"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchAdUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().AD.Users
	response, err := svc.List(ctx, "", "")
	if err != nil {
		return err
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}

func fetchAdUserSignInNames(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	user, ok := parent.Item.(graphrbac.User)
	if !ok {
		return fmt.Errorf("not a graphrbac.User instance: %#v", parent.Item)
	}
	if user.SignInNames != nil {
		for _, item := range *user.SignInNames {
			res <- item
		}
	}
	return nil
}
