package recipes

import (
	"github.com/crowdstrike/gofalcon/falcon/models"
)

func Alerts() []*Resource {
	resources := []*Resource{
		{
			Service:    "alerts",
			SubService: "query",
			Struct:     &models.MsaQueryResponse{},
		},
	}

	return resources
}
