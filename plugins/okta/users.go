package main

import (
	"context"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/okta/okta-sdk-golang/v2/okta"
	//"github.com/okta/okta-sdk-golang/v2/okta/query"
	"go.uber.org/zap"
	"log"
	"reflect"
	"strings"

	//"google.golang.org/api/okta/v1"
	"time"
)

type User struct {
	ID        uint   `gorm:"primarykey"`
	Domain    string `neo:"unique"`
	Activated *time.Time
	Created   *time.Time

	Groups            []*UserGroup `gorm:"constraint:OnDelete:CASCADE;"`
	Login             string
	FirstName         string
	LastName          string
	MiddleName        string
	HonorificPrefix   string
	HonorificSuffix   string
	Email             string
	Title             string
	DisplayName       string
	NickName          string
	ProfileUrl        string
	SecondEmail       string
	MobilePhone       string
	PrimaryPhone      string
	StreetAddress     string
	City              string
	State             string
	ZipCode           string
	CountryCode       string
	PostalAddress     string
	PreferredLanguage string
	Locale            string
	Timezone          string
	UserType          string
	EmployeeNumber    string
	CostCenter        string
	Organization      string
	Division          string
	Department        string
	ManagerId         string
	Manager           string

	CredentialsProviderName string
	CredentialsProviderType string

	ResourceID            string `neo:"unique"`
	LastLogin             *time.Time
	LastUpdated           *time.Time
	PasswordChanged       *time.Time
	Status                string
	StatusChanged         *time.Time
	TransitioningToStatus string
}

func (User) TableName() string {
	return "okta_application_users"
}

type UserGroup struct {
	UserGroupID           uint   `gorm:"primarykey"`
	Domain                string `gorm:"-"`
	UserID                uint   `neo:"ignore"`
	Created               *time.Time
	GroupID               string
	LastMembershipUpdated *time.Time
	LastUpdated           *time.Time
	Name                  string
	Description           string
	Type                  string
}

func (UserGroup) TableName() string {
	return "okta_application_user_groups"
}

func (p *Provider) transformUserGroups(values []*okta.Group) []*UserGroup {
	var tValues []*UserGroup
	for _, v := range values {
		tValues = append(tValues, &UserGroup{
			Domain:                p.config.Domain,
			Created:               v.Created,
			GroupID:               v.Id,
			LastMembershipUpdated: v.LastMembershipUpdated,
			LastUpdated:           v.LastUpdated,
			Type:                  v.Type,
			Name:                  v.Profile.Name,
			Description:           v.Profile.Description,
		})
	}
	return tValues
}

func (p *Provider) transformUser(value *okta.User) *User {
	res := User{
		Domain:                p.config.Domain,
		Activated:             value.Activated,
		Created:               value.Created,
		ResourceID:            value.Id,
		LastLogin:             value.LastLogin,
		LastUpdated:           value.LastUpdated,
		PasswordChanged:       value.PasswordChanged,
		Status:                value.Status,
		StatusChanged:         value.StatusChanged,
		TransitioningToStatus: value.TransitioningToStatus,
	}

	if value.Credentials != nil {
		if value.Credentials.Provider != nil {
			res.CredentialsProviderName = value.Credentials.Provider.Name
			res.CredentialsProviderType = value.Credentials.Provider.Type
		}
	}
	if value.Profile != nil {
		for key, value := range *value.Profile {
			v := reflect.ValueOf(&res).Elem()
			field := v.FieldByName(strings.Title(key))
			field.SetString(fmt.Sprintf("%v", value))
		}
	}

	groups, _, err := p.client.User.ListUserGroups(context.Background(), value.Id)
	if err != nil {
		log.Fatal(err)
	}
	res.Groups = p.transformUserGroups(groups)
	return &res
}

func (p *Provider) transformUsers(values []*okta.User) []*User {
	var tValues []*User
	for _, v := range values {
		tValues = append(tValues, p.transformUser(v))
	}
	return tValues
}

type UserConfig struct {
	Filter string
}

var userTables = []interface{}{
	&User{},
	&UserGroup{},
}

func (p *Provider) users(gConfig interface{}) error {
	var config UserConfig
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	//filter := query.NewQueryParams()
	users, _, err := p.client.User.ListUsers(context.Background(), nil)
	if err != nil {
		return err
	}

	p.db.Where("domain", p.config.Domain).Delete(userTables...)
	p.db.ChunkedCreate(p.transformUsers(users))
	p.log.Info("Fetched resources", zap.Int("count", len(users)))

	return nil
}
