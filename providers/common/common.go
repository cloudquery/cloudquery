package common

import (
	"gorm.io/gorm"
	"reflect"
	"strings"
)

const chunkSize = 100

func ChunkedCreate(db *gorm.DB, value interface{}) {
	arr := reflect.ValueOf(value)
	for i := 0; i < arr.Len(); i += chunkSize {
		end := i + chunkSize
		if i+chunkSize > arr.Len() {
			end = arr.Len()
		}
		db.Create(arr.Slice(i, end).Interface())
	}
}

func StringListToString(arr []*string) *string {
	s := strings.Builder{}
	isEmpty := true
	for _, val := range arr {
		if val != nil {
			s.WriteString(*val)
			s.WriteString(",")
			isEmpty = false
		}
	}
	if isEmpty {
		return nil
	} else {
		res := s.String()
		return &res
	}
}

type ClientInterface interface {
	CollectResource(resource string, config interface{}) error
}
