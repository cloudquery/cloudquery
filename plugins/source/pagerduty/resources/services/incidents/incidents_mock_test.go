package incidents

import (
	"fmt"
	"testing"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
)

func buildMockHttpClient() *client.MockHttpClient {
	mockHttpClient := client.MockHttpClient{}

	response := pagerduty.ListIncidentsResponse{}

	if err := faker.FakeObject(&response); err != nil {
		panic(err)
	}
	if err := client.FakeStringTimestamps(&response.Incidents[0]); err != nil {
		panic(err)
	}

	response.More = false

	mockHttpClient.AddMockResponse("/incidents", response)

	incidentalertResponse := pagerduty.ListAlertsResponse{}
	if err := faker.FakeObject(&incidentalertResponse); err != nil {
		panic(err)
	}
	incidentalertResponse.More = false
	if err := client.FakeStringTimestamps(&incidentalertResponse.Alerts[0]); err != nil {
		panic(err)
	}
	mockHttpClient.AddMockResponse(
		fmt.Sprintf("/incidents/%s/alerts",
			response.Incidents[0].ID),
		incidentalertResponse)

	incidentnoteResponse := make(map[string][]pagerduty.IncidentNote)
	incidennoteSlice := make([]pagerduty.IncidentNote, 1)
	if err := faker.FakeObject(&incidennoteSlice); err != nil {
		panic(err)
	}
	if err := client.FakeStringTimestamps(&incidennoteSlice[0]); err != nil {
		panic(err)
	}
	incidentnoteResponse["notes"] = incidennoteSlice
	mockHttpClient.AddMockResponse(
		fmt.Sprintf("/incidents/%s/notes",
			response.Incidents[0].ID),
		incidentnoteResponse)

	incidentlogentryResponse := pagerduty.ListIncidentLogEntriesResponse{}
	if err := faker.FakeObject(&incidentlogentryResponse); err != nil {
		panic(err)
	}
	incidentlogentryResponse.More = false
	if err := client.FakeStringTimestamps(&incidentlogentryResponse.LogEntries[0]); err != nil {
		panic(err)
	}
	if err := client.FakeStringTimestamps(&incidentlogentryResponse.LogEntries[0].CommonLogEntryField); err != nil {
		panic(err)
	}
	mockHttpClient.AddMockResponse(
		fmt.Sprintf("/incidents/%s/log_entries",
			response.Incidents[0].ID),
		incidentlogentryResponse)

	return &mockHttpClient
}

func TestIncidents(t *testing.T) {
	client.PagerdutyMockTestHelper(t, Incidents(), buildMockHttpClient)
}
