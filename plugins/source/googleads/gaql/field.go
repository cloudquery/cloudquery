package gaql

import (
	"fmt"
	"reflect"

	"github.com/shenzhencenter/google-ads-pb/services"
)

var googleAdsRow = reflect.TypeOf(new(services.GoogleAdsRow)).Elem()

func FieldName(a any) string {
	want := reflect.TypeOf(a)

	for _, fld := range reflect.VisibleFields(googleAdsRow) {
		if fld.Type == want {
			tag, _ := jsonTag(fld)
			return tag
		}
	}

	panic(fmt.Errorf("failed to find row element for %T", a))
}
