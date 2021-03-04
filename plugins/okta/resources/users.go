package resources

import (
	"github.com/mitchellh/mapstructure"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"time"
)

type User struct {
	Activated             *time.Time   `json:"activated,omitempty"`
	Created               *time.Time   `json:"created,omitempty"`
	Id                    string       `json:"id,omitempty"`
	LastLogin             *time.Time   `json:"lastLogin,omitempty"`
	LastUpdated           *time.Time   `json:"lastUpdated,omitempty"`
	PasswordChanged       *time.Time   `json:"passwordChanged,omitempty"`
	Profile               *UserProfile `json:"profile,omitempty" gorm:"embedded;embeddedPrefix:profile_"`
	Status                string       `json:"status,omitempty"`
	StatusChanged         *time.Time   `json:"statusChanged,omitempty"`
	TransitioningToStatus string       `json:"transitioningToStatus,omitempty"`
	UserTypeId            string       `json:"userTypeId,omitempty"`
}

func (u User) TableName() string {
	return "okta_users"
}

func TransformUser(u *okta.User) *User {
	return &User{
		Activated:             u.Activated,
		Created:               u.Created,
		Id:                    u.Id,
		LastLogin:             u.LastLogin,
		LastUpdated:           u.LastUpdated,
		PasswordChanged:       u.PasswordChanged,
		Profile:               TransformUserProfile(u.Profile),
		Status:                u.Status,
		StatusChanged:         u.StatusChanged,
		TransitioningToStatus: u.TransitioningToStatus,
		UserTypeId:            u.Type.Id,
	}
}

func TransformUsers(uu []*okta.User) []*User {
	tuu := make([]*User, len(uu))
	for i, u := range uu {
		tuu[i] = TransformUser(u)
	}
	return tuu
}

type UserProfile struct {
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
}

func TransformUserProfile(profile *okta.UserProfile) *UserProfile {
	if profile == nil {
		return nil
	}
	var up UserProfile
	err := mapstructure.Decode(profile, &up)
	if err != nil {
		return nil
	}
	return &up
}

type UserType struct {
	Id            string     `json:"id,omitempty" gorm:"primaryKey"`
	Created       *time.Time `json:"created,omitempty"`
	CreatedBy     string     `json:"createdBy,omitempty"`
	Default       *bool      `json:"default,omitempty"`
	Description   string     `json:"description,omitempty"`
	DisplayName   string     `json:"displayName,omitempty"`
	LastUpdated   *time.Time `json:"lastUpdated,omitempty"`
	LastUpdatedBy string     `json:"lastUpdatedBy,omitempty"`
	Name          string     `json:"name,omitempty"`
}

func (UserType) TableName() string {
	return "okta_user_types"
}

func TransformUserType(userType *okta.UserType) *UserType {
	if userType == nil {
		return nil
	}
	return &UserType{
		Id:            userType.Id,
		Created:       userType.Created,
		CreatedBy:     userType.CreatedBy,
		Default:       userType.Default,
		Description:   userType.Description,
		DisplayName:   userType.DisplayName,
		LastUpdated:   userType.LastUpdated,
		LastUpdatedBy: userType.LastUpdatedBy,
		Name:          userType.Name,
	}
}

func TransformUserTypes(uu []*okta.UserType) []*UserType {
	tut := make([]*UserType, len(uu))
	for i, u := range uu {
		tut[i] = TransformUserType(u)
	}
	return tut
}
