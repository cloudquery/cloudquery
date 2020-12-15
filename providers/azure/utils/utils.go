package utils

import (
	"github.com/Azure/go-autorest/autorest/date"
	uuid "github.com/satori/go.uuid"
	"time"
)

func AzureDateToTime(t *date.Time) *time.Time {
	if t != nil {
		v := t.ToTime()
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
