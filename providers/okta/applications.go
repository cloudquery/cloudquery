package main

import (
	"context"
	"github.com/mitchellh/mapstructure"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"go.uber.org/zap"
	"log"
	"time"
)

type Application struct {
	ID     uint   `gorm:"primarykey"`
	Domain string `neo:"unique"`

	AccessibilityErrorRedirectUrl string
	AccessibilityLoginRedirectUrl string
	AccessibilitySelfService      *bool

	Created *time.Time

	CredentialsSigningLastRotated  *time.Time
	CredentialsSigningNextRotation *time.Time
	CredentialsSigningRotationMode string

	CredentialsUserNameTemplateSuffix   string
	CredentialsUserNameTemplateTemplate string
	CredentialsUserNameTemplateType     string

	Features      []*ApplicationFeatures `gorm:"constraint:OnDelete:CASCADE;"`
	ApplicationID string                 `neo:"unique"`
	Label         string
	LastUpdated   *time.Time

	LicensingSeatCount int64
	Name               string

	SettingsImplicitAssignment      *bool
	SettingsInlineHookId            string
	SettingsNotificationsVpnHelpUrl string
	SettingsNotificationsVpnMessage string

	SignOnMode string
	Status     string

	VisibilityAutoSubmitToolbar *bool
	VisibilityHideIOS           *bool
	VisibilityHideWeb           *bool
}

func (Application) TableName() string {
	return "okta_applications"
}

type ApplicationFeatures struct {
	ID            uint   `gorm:"primarykey"`
	ApplicationID uint   `neo:"ignore"`
	Domain        string `gorm:"-"`
	Value         string
}

func (ApplicationFeatures) TableName() string {
	return "okta_application_features"
}

func (p *Provider) transformApplicationFeatures(values []string) []*ApplicationFeatures {
	var tValues []*ApplicationFeatures
	for _, v := range values {
		tValues = append(tValues, &ApplicationFeatures{
			Domain: p.config.Domain,

			Value: v,
		})
	}
	return tValues
}

func (p *Provider) transformApplication(value *okta.Application) *Application {
	res := Application{
		Domain: p.config.Domain,

		AccessibilityErrorRedirectUrl: value.Accessibility.ErrorRedirectUrl,
		AccessibilityLoginRedirectUrl: value.Accessibility.LoginRedirectUrl,
		AccessibilitySelfService:      value.Accessibility.SelfService,

		Created: value.Created,

		Features:      p.transformApplicationFeatures(value.Features),
		ApplicationID: value.Id,
		Label:         value.Label,
		LastUpdated:   value.LastUpdated,
		//LicensingSeatCount: value.Licensing.SeatCount,
		Name: value.Name,

		SignOnMode: value.SignOnMode,
		Status:     value.Status,
	}

	if value.Licensing != nil {
		res.LicensingSeatCount = value.Licensing.SeatCount
	}

	if value.Credentials != nil {
		if value.Credentials.Signing != nil {
			res.CredentialsSigningLastRotated = value.Credentials.Signing.LastRotated
			res.CredentialsSigningNextRotation = value.Credentials.Signing.NextRotation
			res.CredentialsSigningRotationMode = value.Credentials.Signing.RotationMode
		}
		if value.Credentials.UserNameTemplate != nil {
			res.CredentialsUserNameTemplateSuffix = value.Credentials.UserNameTemplate.Suffix
			res.CredentialsUserNameTemplateTemplate = value.Credentials.UserNameTemplate.Template
			res.CredentialsUserNameTemplateType = value.Credentials.UserNameTemplate.Type
		}
	}

	if value.Settings != nil {
		res.SettingsImplicitAssignment = value.Settings.ImplicitAssignment
		res.SettingsInlineHookId = value.Settings.InlineHookId
		if value.Settings.Notifications != nil && value.Settings.Notifications.Vpn != nil {
			res.SettingsNotificationsVpnHelpUrl = value.Settings.Notifications.Vpn.HelpUrl
			res.SettingsNotificationsVpnMessage = value.Settings.Notifications.Vpn.Message
		}
	}

	if value.Visibility != nil {
		res.VisibilityAutoSubmitToolbar = value.Visibility.AutoSubmitToolbar
		if value.Visibility.Hide != nil {
			res.VisibilityHideIOS = value.Visibility.Hide.IOS
			res.VisibilityHideWeb = value.Visibility.Hide.Web
		}
	}
	log.Println(value.Profile)

	return &res
}

func (p *Provider) transformApplications(values []okta.App) []*Application {
	var tValues []*Application
	for _, v := range values {
		tValues = append(tValues, p.transformApplication(v.(*okta.Application)))
	}
	return tValues
}

type ApplicationConfig struct {
	Filter string
}

var applicationTables = []interface{}{
	&Application{},
	&ApplicationFeatures{},
}

func (p *Provider) applications(gConfig interface{}) error {
	var config ApplicationConfig
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	//filter := query.NewQueryParams()
	applications, _, err := p.client.Application.ListApplications(context.Background(), nil)
	if err != nil {
		return err
	}

	p.db.Where("domain", p.config.Domain).Delete(applicationTables...)
	p.db.ChunkedCreate(p.transformApplications(applications))
	p.log.Info("Fetched resources", zap.Int("count", len(applications)))
	return nil
}
