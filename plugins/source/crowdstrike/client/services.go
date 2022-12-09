package client

import (
	"github.com/crowdstrike/gofalcon/falcon/client/alerts"
	"github.com/crowdstrike/gofalcon/falcon/client/incidents"
)

//go:generate mockgen -package=mocks -destination=./mocks/mock_incidents_client.go . Incidents
type Incidents interface {
	incidents.ClientService
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_alerts_client.go . Alerts
type Alerts interface {
	alerts.ClientService
}
