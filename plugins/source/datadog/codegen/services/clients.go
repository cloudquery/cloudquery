package services

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

var clients = []interface{}{
	&datadogV2.UsersApi{},
}
