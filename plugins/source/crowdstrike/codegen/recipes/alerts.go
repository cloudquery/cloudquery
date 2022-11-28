package recipes

import (
	"github.com/crowdstrike/gofalcon/falcon/models"
)

func Alerts() []*Resource {
	resources := []*Resource{
		{
			Service:    "Alerts",
			SubService: "Query",
			Struct:     &models.MsaQueryResponse{},
			// PKColumns:  []string{"id"},
		},
	}

	return resources
}
