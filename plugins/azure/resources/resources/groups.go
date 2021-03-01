package resources

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-06-01/resources"
	"github.com/Azure/go-autorest/autorest"
	"github.com/cloudquery/cloudquery/database"
	"github.com/hashicorp/go-hclog"
	"github.com/mitchellh/mapstructure"
)

type Group struct {
	ID             uint `gorm:"primarykey"`
	SubscriptionID string
	ResourceID     *string
	Name           *string
	Type           *string

	PropertiesProvisioningState *string
	Location                    *string
	ManagedBy                   *string
	Tags                        []*GroupTag `gorm:"constraint:OnDelete:CASCADE;"`
}

func (Group) TableName() string {
	return "azure_resources_groups"
}

type GroupTag struct {
	ID uint `gorm:"primarykey"`
	GroupID uint
	Key     string
	Value   *string
}

func (GroupTag) TableName() string {
	return "azure_resources_group_tags"
}

func transformGroups(subscriptionID string, values *[]resources.Group) []*Group {
	var tValues []*Group
	for _, value := range *values {
		tValue := Group{
			SubscriptionID: subscriptionID,
			ResourceID:     value.ID,
			Name:           value.Name,
			Type:           value.Type,

			PropertiesProvisioningState: value.Properties.ProvisioningState,
			Location:                    value.Location,
			ManagedBy:                   value.ManagedBy,
			Tags:                        transformGroupTags(value.Tags),
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func transformGroupTags(values map[string]*string) []*GroupTag {
	var tValues []*GroupTag
	for k, v := range values {
		tValue := GroupTag{
			Key:   k,
			Value: v,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

type GroupConfig struct {
	Filter string
}

var GroupTables = []interface{}{
	&Group{},
	&GroupTag{},
}

func Groups(subscriptionID string, auth autorest.Authorizer, db *database.Database, log hclog.Logger, gConfig interface{}) error {
	var config GroupConfig
	ctx := context.Background()
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	svc := resources.NewGroupsClient(subscriptionID)
	svc.Authorizer = auth
	output, err := svc.List(ctx, "", nil)
	if err != nil {
		return err
	}

	db.Where("subscription_id", subscriptionID).Delete(GroupTables...)
	for output.NotDone() {
		values := output.Values()
		err := output.NextWithContext(ctx)
		if err != nil {
			return err
		}
		tValues := transformGroups(subscriptionID, &values)
		db.ChunkedCreate(tValues)
		log.Info("Fetched resources", "count", len(tValues))
	}

	return nil
}
