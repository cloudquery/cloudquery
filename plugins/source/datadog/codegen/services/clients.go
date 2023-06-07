package services

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

var clients = []any{
	&datadogV1.DashboardListsApi{},
	&datadogV1.DashboardsApi{},
	&datadogV1.DowntimesApi{},
	&datadogV1.HostsApi{},
	&datadogV2.IncidentsApi{},
	&datadogV1.MonitorsApi{},
	&datadogV1.NotebooksApi{},
	&datadogV2.RolesApi{},
	&datadogV1.ServiceLevelObjectivesApi{},
	&datadogV1.ServiceLevelObjectiveCorrectionsApi{},
	&datadogV1.SyntheticsApi{},
	&datadogV2.UsersApi{},
	&datadogV2.RUMApi{},
}
