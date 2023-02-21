package main

import (
	"cloudquery/sentryalerts/sentry"
	"os"
)

func main() {
	sentryToken := os.Getenv("SENTRY_TOKEN")
	sentryOrg := os.Getenv("SENTRY_ORG")
	// This is the ID in Sentry
	// You can find it by going to https://sentry.io/settings/<org>/integrations/slack/, then Configurations->Configure, and see the id in the URL.
	slackWorkspace := os.Getenv("SLACK_WORKSPACE")
	slackChannelName := os.Getenv("SLACK_CHANNEL_NAME")

	if sentryToken == "" {
		panic("SENTRY_TOKEN is not set")
	}

	if sentryOrg == "" {
		panic("SENTRY_ORG is not set")
	}

	if slackWorkspace == "" {
		panic("SLACK_WORKSPACE is not set")
	}

	if slackChannelName == "" {
		panic("SLACK_CHANNEL is not set")
	}

	sentry, err := sentry.New(sentryToken)
	if err != nil {
		panic(err)
	}

	org, err := sentry.GetOrganization(sentryOrg)
	if err != nil {
		panic(err)
	}

	projects, err := sentry.GetProjects(org)
	if err != nil {
		panic(err)
	}

	err = sentry.SetAlerts(org, projects, slackWorkspace, slackChannelName)
	if err != nil {
		panic(err)
	}
}
