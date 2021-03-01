package sql

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/sql/mgmt/2014-04-01/sql"
	"github.com/Azure/go-autorest/autorest"
	"github.com/cloudquery/cloudquery/database"
	"github.com/hashicorp/go-hclog"
	"github.com/mitchellh/mapstructure"
)

type Server struct {
	_              interface{} `neo:"raw:MERGE (a:AzureSubscription {subscription_id: $subscription_id}) MERGE (a) - [:Resource] -> (n)"`
	ID             uint        `gorm:"primarykey"`
	SubscriptionID string
	Kind           *string

	// ServerProperties
	FullyQualifiedDomainName   *string
	Version                    string
	AdministratorLogin         *string
	ExternalAdministratorSid   string
	ExternalAdministratorLogin *string
	State                      string

	Location   *string
	Tags       []*ServerTag `gorm:"constraint:OnDelete:CASCADE;"`
	ResourceID *string
	Name       *string
	Type       *string
}

func (Server) TableName() string {
	return "azure_sql_servers"
}

type ServerTag struct {
	ID             uint   `gorm:"primarykey"`
	ServerID       uint   `neo:"ignore"`
	SubscriptionID string `gorm:"-"`

	Key   string
	Value *string
}

func (ServerTag) TableName() string {
	return "azure_sql_server_tags"
}

func transformServers(subscriptionID string, values *[]sql.Server) []*Server {
	var tValues []*Server
	for _, value := range *values {
		tValue := Server{
			SubscriptionID: subscriptionID,
			Kind:           value.Kind,

			Location:   value.Location,
			Tags:       transformServerTag(subscriptionID, value.Tags),
			ResourceID: value.ID,
			Name:       value.Name,
			Type:       value.Type,
		}

		if value.ServerProperties != nil {
			tValue.FullyQualifiedDomainName = value.ServerProperties.FullyQualifiedDomainName
			tValue.Version = string(value.ServerProperties.Version)
			tValue.AdministratorLogin = value.ServerProperties.AdministratorLogin
			tValue.ExternalAdministratorLogin = value.ServerProperties.ExternalAdministratorLogin
			tValue.State = string(value.ServerProperties.State)

			if value.ExternalAdministratorSid != nil {
				tValue.ExternalAdministratorSid = value.ServerProperties.ExternalAdministratorSid.String()
			}
		}

		tValues = append(tValues, &tValue)
	}
	return tValues
}

func transformServerTag(subscriptionID string, values map[string]*string) []*ServerTag {
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
	&ServerTag{},
}

func Servers(subscriptionID string, auth autorest.Authorizer, db *database.Database, log hclog.Logger, gConfig interface{}) error {
	var config ServerConfig
	ctx := context.Background()
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	svc := sql.NewServersClient(subscriptionID)
	svc.Authorizer = auth
	output, err := svc.List(ctx)
	if err != nil {
		return err
	}

	db.Where("subscription_id", subscriptionID).Delete(ServerTables...)
	tValues := transformServers(subscriptionID, output.Value)
	db.ChunkedCreate(tValues)
	log.Info("Fetched resources", "count", len(tValues))

	return nil
}
