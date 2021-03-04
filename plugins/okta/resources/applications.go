package resources

import (
	"github.com/mitchellh/mapstructure"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"time"
)

type Application struct {
	Id          string                 `json:"id,omitempty" gorm:"primaryKey"`
	Created     *time.Time             `json:"created,omitempty"`
	Label       string                 `json:"label,omitempty"`
	LastUpdated *time.Time             `json:"lastUpdated,omitempty"`
	Name        string                 `json:"name,omitempty"`
	SignOnMode  string                 `json:"signOnMode,omitempty"`
	Status      string                 `json:"status,omitempty"`
	Visibility  *ApplicationVisibility `json:"visibility,omitempty" gorm:"embedded;embeddedPrefix:visibility_"`
}

func (Application) TableName() string {
	return "okta_applications"
}

func TransformApplication(app okta.App) *Application {
	a, ok := app.(*okta.Application)
	if !ok {
		return nil
	}
	return &Application{
		Created:     a.Created,
		Id:          a.Id,
		Label:       a.Label,
		LastUpdated: a.LastUpdated,
		Name:        a.Name,
		SignOnMode:  a.SignOnMode,
		Status:      a.Status,
		Visibility:  TransformApplicationVisibility(a.Visibility),
	}
}

func TransformApplications(aa []okta.App) []*Application {
	ta := make([]*Application, len(aa))
	for i, a := range aa {
		ta[i] = TransformApplication(a)
	}
	return ta
}

type ApplicationVisibility struct {
	AutoSubmitToolbar *bool `json:"autoSubmitToolbar,omitempty"`
	HideIOS           *bool `json:"iOS,omitempty"`
	HideWeb           *bool `json:"web,omitempty"`
}

func TransformApplicationVisibility(av *okta.ApplicationVisibility) *ApplicationVisibility {
	if av == nil {
		return nil
	}
	if av.Hide == nil {
		av.Hide = &okta.ApplicationVisibilityHide{
			IOS: nil,
			Web: nil,
		}
	}
	return &ApplicationVisibility{
		AutoSubmitToolbar: av.AutoSubmitToolbar,
		HideIOS:           av.Hide.IOS,
		HideWeb:           av.Hide.Web,
	}
}

type ApplicationUser struct {
	AppId           string     `json:"appId,omitempty" gorm:"primaryKey"`
	UserId          string     `json:"userId,omitempty" gorm:"primaryKey"`
	Created         *time.Time `json:"created,omitempty"`
	ExternalId      string     `json:"externalId,omitempty"`
	Id              string     `json:"id,omitempty"`
	LastSync        *time.Time `json:"lastSync,omitempty"`
	LastUpdated     *time.Time `json:"lastUpdated,omitempty"`
	PasswordChanged *time.Time `json:"passwordChanged,omitempty"`
	Scope           string     `json:"scope,omitempty"`
	Status          string     `json:"status,omitempty"`
	StatusChanged   *time.Time `json:"statusChanged,omitempty"`
	SyncState       string     `json:"syncState,omitempty"`
}

func (ApplicationUser) TableName() string {
	return "okta_application_users"
}

func TransformAppUser(appId string, au *okta.AppUser) *ApplicationUser {
	if appId == "" || au == nil {
		return nil
	}
	return &ApplicationUser{
		AppId:           appId,
		UserId:          au.Id,
		Created:         au.Created,
		ExternalId:      au.ExternalId,
		LastSync:        au.LastSync,
		LastUpdated:     au.LastUpdated,
		PasswordChanged: au.PasswordChanged,
		Scope:           au.Scope,
		Status:          au.Status,
		StatusChanged:   au.StatusChanged,
		SyncState:       au.SyncState,
	}
}

func TransformAppUsers(appId string, aau []*okta.AppUser) []*ApplicationUser {
	tau := make([]*ApplicationUser, len(aau))
	for i, au := range aau {
		tau[i] = TransformAppUser(appId, au)
	}
	return tau
}

type ApplicationGroup struct {
	AppId       string                   `json:"appId,omitempty" gorm:"primaryKey"`
	GroupId     string                   `json:"groupId,omitempty" gorm:"primaryKey"`
	LastUpdated *time.Time               `json:"lastUpdated,omitempty"`
	Priority    int64                    `json:"priority,omitempty"`
	Profile     *ApplicationGroupProfile `json:"profile,omitempty" gorm:"embedded"`
}

func (ApplicationGroup) TableName() string {
	return "okta_application_groups"
}

func TransformAppGroup(appId string, ag *okta.ApplicationGroupAssignment) *ApplicationGroup {
	if appId == "" || ag == nil {
		return nil
	}
	return &ApplicationGroup{
		AppId:       appId,
		GroupId:     ag.Id,
		LastUpdated: ag.LastUpdated,
		Priority:    ag.Priority,
		Profile:     TransformApplicationGroupProfile(ag.Profile),
	}
}

func TransformAppGroups(appId string, aag []*okta.ApplicationGroupAssignment) []*ApplicationGroup {
	tag := make([]*ApplicationGroup, len(aag))
	for i, ag := range aag {
		tag[i] = TransformAppGroup(appId, ag)
	}
	return tag
}

type ApplicationGroupProfile struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func TransformApplicationGroupProfile(profile interface{}) *ApplicationGroupProfile {
	if profile == nil {
		return nil
	}
	var up ApplicationGroupProfile
	err := mapstructure.Decode(profile, &up)
	if err != nil {
		return nil
	}
	return &up
}
