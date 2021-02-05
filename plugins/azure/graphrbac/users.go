package graphrbac

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/azure-sdk-for-go/services/preview/storage/mgmt/2018-07-01-preview/storage"
	"github.com/Azure/go-autorest/autorest"
	"github.com/cloudquery/cloudquery/database"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type User struct {
	ID                uint `gorm:"primarykey"`
	AccountID         string
	ImmutableID       *string
	UsageLocation     *string
	GivenName         *string
	Surname           *string
	UserType          string
	AccountEnabled    *bool
	DisplayName       *string
	UserPrincipalName *string
	MailNickname      *string
	Mail              *string
	SignInNames       []*UserSignInName `gorm:"constraint:OnDelete:CASCADE;"`
	ObjectID          *string

	ObjectType string
}

func (User) TableName() string {
	return "azure_graphrbac_users"
}

type UserSignInName struct {
	ID     uint `gorm:"primarykey"`
	UserID uint
	Type   *string
	Value  *string
}

func (UserSignInName) TableName() string {
	return "azure_graphrbac_users_igninnames"
}
func transformUserSignInName(value *graphrbac.SignInName) *UserSignInName {
	return &UserSignInName{
		Type:  value.Type,
		Value: value.Value,
	}
}

func transformUserSignInNames(values []graphrbac.SignInName) []*UserSignInName {
	var tValues []*UserSignInName
	for _, v := range values {
		tValues = append(tValues, transformUserSignInName(&v))
	}
	return tValues
}

func transformUser(value *graphrbac.User) *User {
	return &User{
		ImmutableID:       value.ImmutableID,
		UsageLocation:     value.UsageLocation,
		GivenName:         value.GivenName,
		Surname:           value.Surname,
		UserType:          string(value.UserType),
		AccountEnabled:    value.AccountEnabled,
		DisplayName:       value.DisplayName,
		UserPrincipalName: value.UserPrincipalName,
		MailNickname:      value.MailNickname,
		Mail:              value.Mail,
		SignInNames:       transformUserSignInNames(*value.SignInNames),
		ObjectID:          value.ObjectID,
		ObjectType:        string(value.ObjectType),
	}
}

func transformUsers(values []*graphrbac.User) []*User {
	var tValues []*User
	for _, v := range values {
		tValues = append(tValues, transformUser(v))
	}
	return tValues
}

type UserConfig struct {
	Filter string
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(
		&User{},
		&UserSignInName{},
	)
	if err != nil {
		return err
	}

	return nil
}

func Users(auth autorest.Authorizer, db *database.Database, log *zap.Logger, gConfig interface{}) error {
	var config UserConfig
	ctx := context.Background()
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	s := storage.NewAccountsClient("d61dc1dc-e66f-4429-ae51-4e469be0b0ca")
	s.Authorizer = auth
	o, err := s.List(ctx)
	if err != nil {
		return err
	}
	fmt.Println(o)

	svc := graphrbac.NewUsersClient("2238b30a-6771-4097-af91-481fe9499150")
	svc.Authorizer = auth
	output, err := svc.ListComplete(ctx, "", "")
	if err != nil {
		return err
	}
	fmt.Println("success")
	var tValues []*User
	for output.NotDone() {
		u := output.Value()
		fmt.Println(u)
		tValues = append(tValues, transformUser(&u))
	}

	//db.Where("project_id = ?", c.projectID).Delete(&User{})
	//var tValues []*User
	//for _, items := range output.Items {
	//	tValues = append(tValues, transformUsers(items)...)
	//}
	db.ChunkedCreate(tValues)
	log.Info("populating Users", zap.Int("count", len(tValues)))

	return nil
}
