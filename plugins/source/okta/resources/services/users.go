package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:     "okta_users",
		Resolver: fetchUsers,
		Columns: []schema.Column{
			{
				Name: "activated",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "created",
				Type: schema.TypeTimestamp,
			},
			{
				Name:     "credentials_password_hash_algorithm",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Credentials.Password.Hash.Algorithm"),
			},
			{
				Name:     "credentials_password_hash_salt",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Credentials.Password.Hash.Salt"),
			},
			{
				Name:     "credentials_password_hash_salt_order",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Credentials.Password.Hash.SaltOrder"),
			},
			{
				Name:     "credentials_password_hash_value",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Credentials.Password.Hash.Value"),
			},
			{
				Name:     "credentials_password_hash_work_factor",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Credentials.Password.Hash.WorkFactor"),
			},
			{
				Name:     "credentials_password_hook_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Credentials.Password.Hook.Type"),
			},
			{
				Name:     "credentials_password_value",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Credentials.Password.Value"),
			},
			{
				Name:     "credentials_provider_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Credentials.Provider.Name"),
			},
			{
				Name:     "credentials_provider_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Credentials.Provider.Type"),
			},
			{
				Name:     "credentials_recovery_question_answer",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Credentials.RecoveryQuestion.Answer"),
			},
			{
				Name:     "credentials_recovery_question",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Credentials.RecoveryQuestion.Question"),
			},
			{
				Name:            "id",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name: "last_login",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "last_updated",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "password_changed",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "profile",
				Type: schema.TypeJSON,
			},
			{
				Name: "status",
				Type: schema.TypeString,
			},
			{
				Name: "status_changed",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "transitioning_to_status",
				Type: schema.TypeString,
			},
			{
				Name:     "type_created",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Type.Created"),
			},
			{
				Name:     "type_created_by",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type.CreatedBy"),
			},
			{
				Name:     "type_default",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Type.Default"),
			},
			{
				Name:     "type_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type.Description"),
			},
			{
				Name:     "type_display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type.DisplayName"),
			},
			{
				Name:     "type_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type.Id"),
			},
			{
				Name:     "type_last_updated",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Type.LastUpdated"),
			},
			{
				Name:     "type_last_updated_by",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type.LastUpdatedBy"),
			},
			{
				Name:     "type_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type.Name"),
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	api := meta.(*client.Client)
	users, resp, err := api.Okta.User.ListUsers(ctx, query.NewQueryParams(query.WithLimit(200), query.WithAfter("")))
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return nil
	}
	res <- users
	for resp != nil && resp.HasNextPage() {
		var nextUserSet []*okta.User
		resp, err = resp.Next(ctx, &nextUserSet)
		if err != nil {
			return err
		}
		res <- nextUserSet
	}
	return nil
}
