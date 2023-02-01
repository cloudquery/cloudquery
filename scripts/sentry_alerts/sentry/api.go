package sentry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/atlassian/go-sentry-api"
)

type API struct {
	client *sentry.Client
}

type SlackAlertRuleAction struct {
	sentry.AlertRuleAction
	Workspace string `json:"workspace"`
	Channel   string `json:"channel"`
}

type SlackAlertRule struct {
	sentry.AlertRule
	Actions []SlackAlertRuleAction `json:"actions"`
	Owner   string                 `json:"owner"`
}

func encodeOrError(in interface{}) (io.Reader, error) {
	bytedata, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(bytedata), nil
}

func newRequest(token string, method, endpoint string, in interface{}) (*http.Request, error) {
	bodyreader, err := encodeOrError(in)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, endpoint, bodyreader)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Accept", "application/json")
	req.Close = true

	return req, nil
}

func (api API) addAlertRule(org sentry.Organization, project sentry.Project, slackWorkspace string, slackChannelName string) error {
	// The Go API supports only generic alerts, so we need to wrap the generic one in a Slack alert and send the request ourselves.
	// See https://gist.github.com/nikolaik/85e19b89223686b9ab560822fb63bc01 for the Python version of this code.
	fmt.Printf("Creating alert for project %s\n", *project.Slug)
	action := sentry.AlertRuleAction{
		ID: "sentry.integrations.slack.notify_action.SlackNotifyServiceAction",
	}
	slackAction := SlackAlertRuleAction{
		AlertRuleAction: action,
		Workspace:       slackWorkspace,
		Channel:         slackChannelName,
	}
	rule := sentry.AlertRule{ID: *project.Slug,
		Name:      *project.Slug,
		Frequency: 30,
		Conditions: []sentry.AlertRuleCondition{
			{
				ID:   "sentry.rules.conditions.first_seen_event.FirstSeenEventCondition",
				Name: "A new issue is created",
			},
			{
				ID:   "sentry.rules.conditions.regression_event.RegressionEventCondition",
				Name: "The issue changes state from resolved to unresolved",
			},
			{
				ID:   "sentry.rules.conditions.regression_event.RegressionEventCondition",
				Name: "The issue changes state from ignored to unresolved",
			},
		},
		ActionMatch: sentry.AlertRuleMatchAny,
	}
	slackRule := SlackAlertRule{
		AlertRule: rule,
		Actions:   []SlackAlertRuleAction{slackAction},
		Owner:     fmt.Sprintf("team:%s", *(*org.Teams)[0].ID),
	}

	endpoint := fmt.Sprintf("%sprojects/%s/%s/rules/", api.client.Endpoint, *org.Slug, *project.Slug)
	req, err := newRequest(api.client.AuthToken, "POST", endpoint, slackRule)
	if err != nil {
		return err
	}
	res, err := api.client.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= 300 {
		return fmt.Errorf("failed to create alert %s. Status code %d", slackRule.Name, res.StatusCode)
	}
	fmt.Printf("Alert created successfully: %s\n", slackRule.Name)
	return nil
}

func New(token string) (*API, error) {
	client, err := sentry.NewClient(token, nil, nil)
	if err != nil {
		return nil, err
	}
	return &API{client: client}, nil
}

func (api API) GetOrganization(orgName string) (sentry.Organization, error) {
	org, err := api.client.GetOrganization(orgName)
	if err != nil {
		return sentry.Organization{}, err
	}
	return org, nil
}

func (api API) GetProjects(org sentry.Organization) ([]sentry.Project, error) {
	team := (*org.Teams)[0]
	projects, err := api.client.GetTeamProjects(org, team)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (api API) getProjectAlerts(org sentry.Organization, project sentry.Project) ([]sentry.AlertRule, error) {
	alerts, link, err := api.client.GetAlertRules(org, project)
	if err != nil {
		return nil, err
	}
	allAlerts := make([]sentry.AlertRule, 0, len(alerts))
	allAlerts = append(allAlerts, alerts...)
	for link != nil && link.Next.Results {
		link, err = api.client.GetPage(link.Next, &alerts)
		if err != nil {
			return nil, err
		}
		allAlerts = append(allAlerts, alerts...)
	}
	return allAlerts, nil
}

func isAlertExists(alerts []sentry.AlertRule, alertName string) bool {
	for _, alert := range alerts {
		if alert.Name == alertName {
			return true
		}
	}
	return false
}

func (api API) SetAlerts(org sentry.Organization, projects []sentry.Project, slackWorkspace string, slackChannelName string) error {
	for _, project := range projects {
		alerts, err := api.getProjectAlerts(org, project)
		if err != nil {
			return err
		}
		if isAlertExists(alerts, *project.Slug) {
			fmt.Printf("Alert for project %s already exists. Skipping\n", *project.Slug)
			continue
		}

		err = api.addAlertRule(org, project, slackWorkspace, slackChannelName)
		if err != nil {
			return err
		}
	}

	return nil
}
