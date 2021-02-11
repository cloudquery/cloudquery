package mysql

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
	"github.com/Azure/go-autorest/autorest"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-azure/utils"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"regexp"
	"time"
)

type Server struct {
	_              interface{} `neo:"raw:MERGE (a:AzureSubscription {subscription_id: $subscription_id}) MERGE (a) - [:Resource] -> (n)"`
	ID             uint        `gorm:"primarykey"`
	SubscriptionID string      `neo:"unique"`

	IdentityPrincipalID *string
	IdentityType        string
	IdentityTenantID    *string

	SkuName     *string
	SkuTier     string
	SkuCapacity *int32
	SkuSize     *string
	SkuFamily   *string

	AdministratorLogin       *string
	Version                  string
	SslEnforcement           string
	MinimalTLSVersion        string
	ByokEnforcement          *string
	InfrastructureEncryption string
	UserVisibleState         string
	FullyQualifiedDomainName *string
	EarliestRestoreDate      *time.Time

	StorageProfileBackupRetentionDays *int32
	StorageProfileGeoRedundantBackup  string
	StorageProfileStorageMB           *int32
	StorageProfileStorageAutogrow     string
	ReplicationRole                   *string
	MasterServerID                    *string
	ReplicaCapacity                   *int32
	PublicNetworkAccess               string
	PrivateEndpointConnections        []*ServerPrivateEndpointConnection `gorm:"constraint:OnDelete:CASCADE;"`
	Tags                              []*ServerTag                       `gorm:"constraint:OnDelete:CASCADE;"`
	Configurations                    []*ServerConfiguration             `gorm:"constraint:OnDelete:CASCADE;"`
	Location                          *string
	ResourceID                        *string `neo:"unique"`
	Name                              *string
	Type                              *string
}

func (Server) TableName() string {
	return "azure_mysql_servers"
}

type ServerConfiguration struct {
	ID             uint   `gorm:"primarykey"`
	ServerID       uint   `neo:"ignore"`
	SubscriptionID string `gorm:"ignore"`

	Value         *string
	DefaultValue  *string
	DataType      *string
	AllowedValues *string
	Source        *string
	ResourceID    *string
	Name          *string
	Type          *string
}

func (ServerConfiguration) TableName() string {
	return "azure_mysql_server_configurations"
}

type ServerPrivateEndpointConnection struct {
	ID             uint   `gorm:"primarykey"`
	ServerID       uint   `neo:"ignore"`
	SubscriptionID string `gorm:"ignore"`

	ResourceID                                       *string
	PrivateEndpointID                                *string
	PrivateLinkServiceConnectionStateStatus          string
	PrivateLinkServiceConnectionStateDescription     *string
	PrivateLinkServiceConnectionStateActionsRequired string
	ProvisioningState                                string
}

func (ServerPrivateEndpointConnection) TableName() string {
	return "azure_mysql_server_private_endpoint_connections"
}

type ServerTag struct {
	ID             uint   `gorm:"primarykey"`
	ServerID       uint   `neo:"ignore"`
	SubscriptionID string `gorm:"ignore"`

	Key   string
	Value *string
}

func (ServerTag) TableName() string {
	return "azure_mysql_tags"
}

func transformServers(subscriptionID string, auth autorest.Authorizer, values *[]mysql.Server) ([]*Server, error) {
	var tValues []*Server
	ctx := context.Background()
	resourceGroupRe := regexp.MustCompile("resourceGroups/([a-zA-Z0-9-_]+)/")
	for _, value := range *values {
		tValue := Server{
			SubscriptionID: subscriptionID,
			Tags:           transformServerTags(subscriptionID, value.Tags),
			Location:       value.Location,
			ResourceID:     value.ID,
			Name:           value.Name,
			Type:           value.Type,
		}
		match := resourceGroupRe.FindStringSubmatch(*value.ID)
		if len(match) < 2 {
			return nil, fmt.Errorf("coultn't extract resource group from %s", *value.ID)
		}
		svc := mysql.NewConfigurationsClient(subscriptionID)
		svc.Authorizer = auth
		configResult, err := svc.ListByServer(ctx, match[1], *value.Name)
		if err != nil {
			return nil, err
		}
		if configResult.Value != nil {
			tValue.Configurations = transformConfigurations(subscriptionID, configResult.Value)
		}

		if value.Identity != nil {

			tValue.IdentityPrincipalID = utils.AzureUUIDToString(value.Identity.PrincipalID)
			tValue.IdentityType = string(value.Identity.Type)
			tValue.IdentityTenantID = utils.AzureUUIDToString(value.Identity.TenantID)
		}
		if value.Sku != nil {

			tValue.SkuName = value.Sku.Name
			tValue.SkuTier = string(value.Sku.Tier)
			tValue.SkuCapacity = value.Sku.Capacity
			tValue.SkuSize = value.Sku.Size
			tValue.SkuFamily = value.Sku.Family
		}
		if value.ServerProperties != nil {

			tValue.AdministratorLogin = value.ServerProperties.AdministratorLogin
			tValue.Version = string(value.ServerProperties.Version)
			tValue.SslEnforcement = string(value.ServerProperties.SslEnforcement)
			tValue.MinimalTLSVersion = string(value.ServerProperties.MinimalTLSVersion)
			tValue.ByokEnforcement = value.ServerProperties.ByokEnforcement
			tValue.InfrastructureEncryption = string(value.ServerProperties.InfrastructureEncryption)
			tValue.UserVisibleState = string(value.ServerProperties.UserVisibleState)
			tValue.FullyQualifiedDomainName = value.ServerProperties.FullyQualifiedDomainName
			tValue.EarliestRestoreDate = utils.AzureDateToTime(value.ServerProperties.EarliestRestoreDate)
			tValue.ReplicationRole = value.ServerProperties.ReplicationRole
			tValue.MasterServerID = value.ServerProperties.MasterServerID
			tValue.ReplicaCapacity = value.ServerProperties.ReplicaCapacity
			tValue.PublicNetworkAccess = string(value.ServerProperties.PublicNetworkAccess)
			if value.ServerProperties.PrivateEndpointConnections != nil {
				tValue.PrivateEndpointConnections = transformServerPrivateEndpointConnections(subscriptionID, value.ServerProperties.PrivateEndpointConnections)
			}
		}
		tValues = append(tValues, &tValue)
	}
	return tValues, nil
}

func transformConfigurations(subscriptionID string, values *[]mysql.Configuration) []*ServerConfiguration {
	var tValues []*ServerConfiguration
	for _, value := range *values {
		tValue := ServerConfiguration{
			SubscriptionID: subscriptionID,
			ResourceID:     value.ID,
			Name:           value.Name,
			Type:           value.Type,
		}

		if value.ConfigurationProperties != nil {
			tValue.Value = value.ConfigurationProperties.Value
			tValue.DefaultValue = value.ConfigurationProperties.DefaultValue
			tValue.DataType = value.ConfigurationProperties.DataType
			tValue.AllowedValues = value.ConfigurationProperties.AllowedValues
			tValue.Source = value.ConfigurationProperties.Source
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func transformServerPrivateEndpointConnections(subscriptionID string, values *[]mysql.ServerPrivateEndpointConnection) []*ServerPrivateEndpointConnection {
	var tValues []*ServerPrivateEndpointConnection
	for _, value := range *values {
		tValue := ServerPrivateEndpointConnection{
			SubscriptionID: subscriptionID,
			ResourceID:     value.ID,
		}
		if value.Properties != nil {
			tValue.ProvisioningState = string(value.Properties.ProvisioningState)
			if value.Properties.PrivateEndpoint != nil {
				tValue.PrivateEndpointID = value.Properties.PrivateEndpoint.ID
			}
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func transformServerTags(subscriptionID string, values map[string]*string) []*ServerTag {
	var tValues []*ServerTag
	for k, v := range values {
		tValue := ServerTag{
			SubscriptionID: subscriptionID,
			Key:            k,
			Value:          v,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

type ServerConfig struct {
	Filter string
}

var ServerTables = []interface{}{
	&Server{},
	&ServerConfiguration{},
	&ServerPrivateEndpointConnection{},
	&ServerTag{},
}

func Servers(subscriptionID string, auth autorest.Authorizer, db *database.Database, log *zap.Logger, gConfig interface{}) error {
	var config ServerConfig
	ctx := context.Background()
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	svc := mysql.NewServersClient(subscriptionID)
	svc.Authorizer = auth
	output, err := svc.List(ctx)
	if err != nil {
		return err
	}

	db.Where("subscription_id", subscriptionID).Delete(ServerTables...)
	if output.Value != nil {
		tValues, err := transformServers(subscriptionID, auth, output.Value)
		if err != nil {
			return err
		}
		db.ChunkedCreate(tValues)
		log.Info("Fetched resources", zap.Int("count", len(tValues)))
	}

	return nil
}
