package teams

import (
	"fmt"
	"testing"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
)

func buildMockHttpClient() *client.MockHttpClient {
	mockHttpClient := client.MockHttpClient{}

	response := pagerduty.ListTeamResponse{}

	if err := faker.FakeObject(&response); err != nil {
		panic(err)
	}
	if err := client.FakeStringTimestamps(&response.Teams[0]); err != nil {
		panic(err)
	}

	response.More = false

	mockHttpClient.AddMockResponse("/teams", response)

	membersResponse := pagerduty.ListTeamMembersResponse{}
	if err := faker.FakeObject(&membersResponse); err != nil {
		panic(err)
	}
	membersResponse.More = false
	mockHttpClient.AddMockResponse(
		fmt.Sprintf("/teams/%s/members", response.Teams[0].ID),
		membersResponse)

	return &mockHttpClient
}

func TestTeams(t *testing.T) {
	client.PagerdutyMockTestHelper(t, Teams(), buildMockHttpClient)
}
