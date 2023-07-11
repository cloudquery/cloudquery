package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

const TestSubscription = "12345678-1234-1234-1234-123456789000"
const LegacyAccountName = "9971ccb0-02bb-45e2-bd6a-9e340372dcba"
const ModernAccountName = "7c05a543-80ff-571e-9f98-1063b3b53cf2:99ad03ad-2d1b-4889-a452-090ad407d25f_2019-05-31"

var testResourceGroup = "test-resource-group"

type MockCreds struct {
}

func (*MockCreds) GetToken(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return azcore.AccessToken{
		Token:     "SomeToken",
		ExpiresOn: time.Now().Add(time.Hour * 24),
	}, nil
}

type MockHttpClient struct {
	rootURL string
	scheme  string
	host    string
	client  *http.Client
}

func NewMockHttpClient(cl *http.Client, rootURL string) *MockHttpClient {
	u, err := url.Parse(rootURL)
	if err != nil {
		panic(err)
	}
	return &MockHttpClient{
		client:  cl,
		rootURL: rootURL,
		scheme:  u.Scheme,
		host:    u.Host,
	}
}

func (c *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	req.URL.Host = c.host
	req.URL.Scheme = c.scheme
	return c.client.Do(req)
}

func MockTestHelper(t *testing.T, table *schema.Table, createServices func(*mux.Router) error) {
	// version := "vDev"
	t.Helper()
	debug = true
	table.IgnoreInTests = false
	creds := &MockCreds{}
	router := mux.NewRouter()
	h := httptest.NewServer(router)
	defer h.Close()
	mockClient := NewMockHttpClient(h.Client(), h.URL)

	var legacyAccount armbilling.Account
	if err := faker.FakeObject(&legacyAccount); err != nil {
		t.Fatal(err)
	}
	legacyAccount.ID = to.Ptr("/providers/Microsoft.Billing/billingAccounts/" + LegacyAccountName)
	legacyAccount.Name = to.Ptr(LegacyAccountName)
	legacyAccount.Properties.BillingProfiles = nil

	var modernAccount armbilling.Account
	if err := faker.FakeObject(&modernAccount); err != nil {
		t.Fatal(err)
	}
	modernAccount.ID = to.Ptr("/providers/Microsoft.Billing/billingAccounts/" + ModernAccountName)
	modernAccount.Name = to.Ptr(ModernAccountName)
	modernAccount.Properties.BillingProfiles.Value[0].ID = to.Ptr("/providers/Microsoft.Billing/billingAccounts/account-id/billingProfiles/profile-id")
	modernAccount.Properties.BillingProfiles.Value[0].Name = to.Ptr("profile-id")

	var billingPeriod armbilling.Period
	if err := faker.FakeObject(&billingPeriod); err != nil {
		t.Fatal(err)
	}
	billingPeriod.ID = to.Ptr("/subscriptions/" + TestSubscription + "/providers/Microsoft.Billing/billingPeriods/202205-1")

	l := zerolog.New(zerolog.NewTestWriter(t)).Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro}).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	err := createServices(router)
	if err != nil {
		t.Fatal(err)
	}
	registeredNamespaces := make(map[string]map[string]bool)
	registeredNamespaces[TestSubscription] = make(map[string]bool)
	for _, namespace := range namespaces {
		registeredNamespaces[TestSubscription][namespace] = true
	}

	resourceGroup := &armresources.ResourceGroup{}
	err = faker.FakeObject(resourceGroup)
	if err != nil {
		t.Fatal(err)
	}
	resourceGroup.Name = &testResourceGroup

	c := &Client{
		logger: l,
		Options: &arm.ClientOptions{
			ClientOptions: policy.ClientOptions{
				Transport: mockClient,
			},
		},
		registeredNamespaces: registeredNamespaces,
		Creds:                creds,
		subscriptions:        []string{TestSubscription},
		ResourceGroups: map[string][]*armresources.ResourceGroup{
			TestSubscription: {resourceGroup},
		},
		BillingAccounts: []*armbilling.Account{&legacyAccount, &modernAccount},
		BillingPeriods: map[string][]*armbilling.Period{
			TestSubscription: {&billingPeriod},
		},
		storageAccountKeys: &sync.Map{},
		pluginSpec: &Spec{
			NormalizeIDs: true,
		},
	}
	sched := scheduler.NewScheduler(scheduler.WithLogger(l))
	messages, err := sched.SyncAll(context.Background(), c, schema.Tables{table})
	if err != nil {
		t.Fatalf("failed to sync: %v", err)
	}

	records := messages.GetInserts().GetRecordsForTable(table)
	emptyColumns := schema.FindEmptyColumns(table, records)
	if len(emptyColumns) > 0 {
		t.Fatalf("empty columns: %v", emptyColumns)
	}
}
