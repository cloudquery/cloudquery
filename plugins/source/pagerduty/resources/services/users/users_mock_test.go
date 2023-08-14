package users

import (
	"fmt"
	"testing"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
)

func buildMockHttpClient() *client.MockHttpClient {
	mockHttpClient := client.MockHttpClient{}

	response := pagerduty.ListUsersResponse{}

	if err := faker.FakeObject(&response); err != nil {
		panic(err)
	}
	if err := client.FakeStringTimestamps(&response.Users[0]); err != nil {
		panic(err)
	}

	response.More = false

	mockHttpClient.AddMockResponse("/users", response)

	contactmethodResponse := pagerduty.ListContactMethodsResponse{}
	if err := faker.FakeObject(&contactmethodResponse); err != nil {
		panic(err)
	}
	contactmethodResponse.More = false
	if err := client.FakeStringTimestamps(&contactmethodResponse.ContactMethods[0]); err != nil {
		panic(err)
	}
	mockHttpClient.AddMockResponse(
		fmt.Sprintf("/users/%s/contact_methods",
			response.Users[0].ID),
		contactmethodResponse)

	notificationruleResponse := pagerduty.ListUserNotificationRulesResponse{}
	if err := faker.FakeObject(&notificationruleResponse); err != nil {
		panic(err)
	}
	notificationruleResponse.More = false
	if err := client.FakeStringTimestamps(&notificationruleResponse.NotificationRules[0]); err != nil {
		panic(err)
	}
	mockHttpClient.AddMockResponse(
		fmt.Sprintf("/users/%s/notification_rules",
			response.Users[0].ID),
		notificationruleResponse)

	return &mockHttpClient
}

func TestUsers(t *testing.T) {
	client.PagerdutyMockTestHelper(t, Users(), buildMockHttpClient)
}
