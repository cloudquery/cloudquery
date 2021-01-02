package utils

import (
	"github.com/Azure/go-autorest/autorest/date"
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

func AzureDateToTime(t *date.Time) *time.Time {
	if t != nil {
		location, err := time.LoadLocation("UTC")
		if err != nil {
			log.Fatal(err)
		}
		v := t.In(location)
		return &v
	}
	return nil
}

func AzureUUIDToString(v *uuid.UUID) *string {
	if v != nil {
		s := v.String()
		return &s
	}
	return nil
}
