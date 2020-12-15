package sql

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/sql/mgmt/2014-04-01/sql"
	"github.com/Azure/go-autorest/autorest"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Server struct {
	ID             uint `gorm:"primarykey"`
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
	ID       uint
	ServerID uint
	Key      string
	Value    *string
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
			Tags:       transformServerTag(value.Tags),
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

func transformServerTag(values map[string]*string) []*ServerTag {
	var tValues []*ServerTag
	for k, v := range values {
		tValue := ServerTag{
			Key:   k,
			Value: v,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

type ServerConfig struct {
	Filter string
}

func MigrateServer(db *gorm.DB) error {
	err := db.AutoMigrate(
		&Server{},
		&ServerTag{},
	)
	if err != nil {
		return err
	}

	return nil
}

func Servers(subscriptionID string, auth autorest.Authorizer, db *gorm.DB, log *zap.Logger, gConfig interface{}) error {
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

	db.Where("subscription_id = ?", subscriptionID).Delete(&Server{})
	tValues := transformServers(subscriptionID, output.Value)
	common.ChunkedCreate(db, tValues)
	log.Info("Fetched resources", zap.Int("count", len(tValues)))

	return nil
}
