package resources

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SQLServerEncryptionProtectors() *schema.Table {
	return &schema.Table{
		Name:        "azure_sql_server_encryption_protectors",
		Description: "EncryptionProtector the server encryption protector",
		Resolver:    fetchSqlServerEncryptionProtectors,
		Columns: []schema.Column{
			{
				Name:        "server_cq_id",
				Description: "Unique ID of azure_sql_servers table (FK)",
				Type:        schema.TypeUUID,
				Resolver:    schema.ParentIdResolver,
			},
			{
				Name:        "kind",
				Description: "Kind of encryption protector This is metadata used for the Azure portal experience",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "Resource location",
				Type:        schema.TypeString,
			},
			{
				Name:        "subregion",
				Description: "Subregion of the encryption protector",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionProtectorProperties.Subregion"),
			},
			{
				Name:        "server_key_name",
				Description: "The name of the server key",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionProtectorProperties.ServerKeyName"),
			},
			{
				Name:        "server_key_type",
				Description: "The encryption protector type like 'ServiceManaged', 'AzureKeyVault' Possible values include: 'ServiceManaged', 'AzureKeyVault'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionProtectorProperties.ServerKeyType"),
			},
			{
				Name:        "uri",
				Description: "The URI of the server key",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionProtectorProperties.URI"),
			},
			{
				Name:        "thumbprint",
				Description: "Thumbprint of the server key",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionProtectorProperties.Thumbprint"),
			},
			{
				Name:        "id",
				Description: "Resource ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource name",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchSqlServerEncryptionProtectors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().SQL.EncryptionProtectors
	server := parent.Item.(sql.Server)
	resourceDetails, err := client.ParseResourceID(*server.ID)
	if err != nil {
		return err
	}
	ep, err := svc.Get(ctx, resourceDetails.ResourceGroup, *server.Name)
	if err != nil {
		return err
	}
	res <- ep
	return nil
}
