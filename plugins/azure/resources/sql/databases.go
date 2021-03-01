package sql

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/sql/mgmt/2014-04-01/sql"
	"github.com/Azure/go-autorest/autorest"
	"github.com/cloudquery/cloudquery/database"
	"github.com/hashicorp/go-hclog"
	"github.com/mitchellh/mapstructure"
	"regexp"
	"time"
)

type Database struct {
	ID             uint   `gorm:"primarykey"`
	SubscriptionID string `neo:"unique"`
	Kind           *string

	Collation                               *string
	CreationDate                            time.Time
	ContainmentState                        *int64
	CurrentServiceObjectiveID               string
	DatabaseID                              string `neo:"unique"`
	EarliestRestoreDate                     time.Time
	CreateMode                              string
	SourceDatabaseID                        *string
	SourceDatabaseDeletionDate              time.Time
	RestorePointInTime                      time.Time
	RecoveryServicesRecoveryPointResourceID *string
	Edition                                 string
	MaxSizeBytes                            *string
	RequestedServiceObjectiveID             string
	RequestedServiceObjectiveName           string
	ServiceLevelObjective                   string
	Status                                  *string
	ElasticPoolName                         *string
	DefaultSecondaryLocation                *string
	TransparentDataEncryption               []*DatabaseTransparentDataEncryption `gorm:"constraint:OnDelete:CASCADE;"`
	FailoverGroupID                         *string
	ReadScale                               string
	SampleName                              string
	ZoneRedundant                           *bool

	Location   *string
	Tags       []*DatabaseTag `gorm:"constraint:OnDelete:CASCADE;"`
	ResourceID *string
	Name       *string
	Type       *string
}

func (Database) TableName() string {
	return "azure_sql_databases"
}

type DatabaseTransparentDataEncryption struct {
	ID             uint   `gorm:"primarykey"`
	DatabaseID     uint   `neo:"ignore"`
	SubscriptionID string `gorm:"-"`

	Location   *string
	Status     string
	ResourceID *string
	Name       *string
	Type       *string
}

func (DatabaseTransparentDataEncryption) TableName() string {
	return "azure_sql_database_transparent_data_encryptions"
}

type DatabaseTag struct {
	ID             uint   `gorm:"primarykey"`
	DatabaseID     uint   `neo:"ignore"`
	SubscriptionID string `gorm:"-"`

	Key   string
	Value *string
}

func (DatabaseTag) TableName() string {
	return "azure_sql_database_tags"
}

func transformDatabases(subscriptionID string, values *[]sql.Database) []*Database {
	var tValues []*Database
	for _, value := range *values {
		tValue := Database{
			SubscriptionID: subscriptionID,
			Kind:           value.Kind,

			Location:   value.Location,
			Tags:       transformDatabaseTags(subscriptionID, value.Tags),
			ResourceID: value.ID,
			Name:       value.Name,
			Type:       value.Type,
		}

		if value.DatabaseProperties != nil {
			tValue.Collation = value.DatabaseProperties.Collation
			tValue.ContainmentState = value.DatabaseProperties.ContainmentState
			tValue.CurrentServiceObjectiveID = value.DatabaseProperties.CurrentServiceObjectiveID.String()
			tValue.DatabaseID = value.DatabaseProperties.DatabaseID.String()
			tValue.CreateMode = string(value.DatabaseProperties.CreateMode)
			tValue.SourceDatabaseID = value.DatabaseProperties.SourceDatabaseID
			tValue.RecoveryServicesRecoveryPointResourceID = value.DatabaseProperties.RecoveryServicesRecoveryPointResourceID
			tValue.Edition = string(value.DatabaseProperties.Edition)
			tValue.MaxSizeBytes = value.DatabaseProperties.MaxSizeBytes
			tValue.RequestedServiceObjectiveID = value.DatabaseProperties.RequestedServiceObjectiveID.String()
			tValue.RequestedServiceObjectiveName = string(value.DatabaseProperties.RequestedServiceObjectiveName)
			tValue.ServiceLevelObjective = string(value.DatabaseProperties.ServiceLevelObjective)
			tValue.Status = value.DatabaseProperties.Status
			tValue.ElasticPoolName = value.DatabaseProperties.ElasticPoolName
			tValue.DefaultSecondaryLocation = value.DatabaseProperties.DefaultSecondaryLocation
			tValue.FailoverGroupID = value.DatabaseProperties.FailoverGroupID
			tValue.ReadScale = string(value.DatabaseProperties.ReadScale)
			tValue.SampleName = string(value.DatabaseProperties.SampleName)
			tValue.ZoneRedundant = value.DatabaseProperties.ZoneRedundant

			if value.DatabaseProperties.CreationDate != nil {
				tValue.CreationDate = value.DatabaseProperties.CreationDate.ToTime()
			}

			if value.DatabaseProperties.EarliestRestoreDate != nil {
				tValue.EarliestRestoreDate = value.DatabaseProperties.EarliestRestoreDate.ToTime()
			}

			if value.DatabaseProperties.SourceDatabaseDeletionDate != nil {
				tValue.SourceDatabaseDeletionDate = value.DatabaseProperties.SourceDatabaseDeletionDate.ToTime()
			}

			if value.DatabaseProperties.RestorePointInTime != nil {
				tValue.RestorePointInTime = value.DatabaseProperties.RestorePointInTime.ToTime()
			}

			if value.DatabaseProperties.TransparentDataEncryption != nil {
				tValue.TransparentDataEncryption = transformDatabaseTransparentDataEncryptions(subscriptionID, value.DatabaseProperties.TransparentDataEncryption)
			}

		}

		tValues = append(tValues, &tValue)
	}
	return tValues
}

func transformDatabaseTransparentDataEncryptions(subscriptionID string, values *[]sql.TransparentDataEncryption) []*DatabaseTransparentDataEncryption {
	var tValues []*DatabaseTransparentDataEncryption
	for _, value := range *values {
		tValue := DatabaseTransparentDataEncryption{
			Location:       value.Location,
			SubscriptionID: subscriptionID,

			ResourceID: value.ID,
			Name:       value.Name,
			Type:       value.Type,
		}
		if value.TransparentDataEncryptionProperties != nil {
			tValue.Status = string(value.TransparentDataEncryptionProperties.Status)
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func transformDatabaseTags(subscriptionID string, values map[string]*string) []*DatabaseTag {
	var tValues []*DatabaseTag
	for k, v := range values {
		tValue := DatabaseTag{
			SubscriptionID: subscriptionID,
			Key:            k,
			Value:          v,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

type DatabaseConfig struct {
	Filter string
}

var DatabaseTables = []interface{}{
	&Database{},
	&DatabaseTransparentDataEncryption{},
	&DatabaseTag{},
}

func Databases(subscriptionID string, auth autorest.Authorizer, db *database.Database, log hclog.Logger, gConfig interface{}) error {
	var config DatabaseConfig
	ctx := context.Background()
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	// First get all the servers
	serverClient := sql.NewServersClient(subscriptionID)
	serverClient.Authorizer = auth
	serverResult, err := serverClient.List(ctx)
	if err != nil {
		return err
	}
	resourceGroupRe := regexp.MustCompile("resourceGroups/([a-zA-Z0-9-_]+)/")
	db.Where("subscription_id", subscriptionID).Delete(DatabaseTables...)
	if len(*serverResult.Value) == 0 {
		log.Info("Fetched resources", "count", 0)
	}
	for _, server := range *serverResult.Value {
		svc := sql.NewDatabasesClient(subscriptionID)
		svc.Authorizer = auth
		match := resourceGroupRe.FindStringSubmatch(*server.ID)
		if len(match) < 2 {
			return fmt.Errorf("coultn't extract resource group from %s", *server.ID)
		}
		output, err := svc.ListByServer(ctx, match[1], *server.Name, "", "")
		if err != nil {
			return err
		}
		tValues := transformDatabases(subscriptionID, output.Value)
		db.ChunkedCreate(tValues)
		log.Info("Fetched resources", "count", len(tValues))
	}

	return nil
}
