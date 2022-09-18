//go:generate mockgen -destination=./mocks/security.go -package=mocks . SecurityAutoProvisioningSettingsClient,SecurityContactsClient,SecurityPricingsClient,SecuritySettingsClient,SecurityJitNetworkAccessPoliciesClient,SecurityAssessmentsClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
	"github.com/Azure/go-autorest/autorest"
)

type SecurityAutoProvisioningSettingsClient interface {
	List(ctx context.Context) (result security.AutoProvisioningSettingListPage, err error)
}

type SecurityContactsClient interface {
	List(ctx context.Context) (result security.ContactListPage, err error)
}

type SecurityPricingsClient interface {
	List(ctx context.Context) (result security.PricingList, err error)
}

type SecuritySettingsClient interface {
	List(ctx context.Context) (result security.SettingsListPage, err error)
}

type SecurityJitNetworkAccessPoliciesClient interface {
	List(ctx context.Context) (result security.JitNetworkAccessPoliciesListPage, err error)
}

type SecurityAssessmentsClient interface {
	List(ctx context.Context, scope string) (result security.AssessmentListPage, err error)
}

type SecurityClient struct {
	Assessments              SecurityAssessmentsClient
	AutoProvisioningSettings SecurityAutoProvisioningSettingsClient
	Contacts                 SecurityContactsClient
	Pricings                 SecurityPricingsClient
	Settings                 SecuritySettingsClient
	JitNetworkAccessPolicies SecurityJitNetworkAccessPoliciesClient
}

func NewSecurityClient(subscriptionId string, auth autorest.Authorizer) SecurityClient {
	// New*Client requires that the ASC location is passed as an argument, but API methods
	// that we actually use do not use that location for performing the request (judging by REST API docs).
	// That is why we are passing an empty string here.
	aps := security.NewAutoProvisioningSettingsClient(subscriptionId, "")
	aps.Authorizer = auth
	pricings := security.NewPricingsClient(subscriptionId, "")
	pricings.Authorizer = auth
	contacts := security.NewContactsClient(subscriptionId, "")
	contacts.Authorizer = auth
	settings := security.NewSettingsClient(subscriptionId, "")
	settings.Authorizer = auth
	jitNetworkAccessPolicies := security.NewJitNetworkAccessPoliciesClient(subscriptionId, "")
	jitNetworkAccessPolicies.Authorizer = auth
	assessments := security.NewAssessmentsClient(subscriptionId, "")
	assessments.Authorizer = auth
	return SecurityClient{
		Assessments:              assessments,
		AutoProvisioningSettings: aps,
		Contacts:                 contacts,
		Pricings:                 pricings,
		Settings:                 settings,
		JitNetworkAccessPolicies: jitNetworkAccessPolicies,
	}
}
