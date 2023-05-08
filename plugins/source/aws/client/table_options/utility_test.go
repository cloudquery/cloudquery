package table_options

import (
	"reflect"
	"time"
)

func findNilOrDefaultFields(v reflect.Value, nilOrDefFields []string) []string {
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := v.Type().Field(i)

		switch fieldValue.Kind() {
		case reflect.Ptr:
			if fieldValue.IsNil() {
				nilOrDefFields = append(nilOrDefFields, fieldType.Name)
			} else {
				if fieldValue.Elem().Kind() == reflect.Struct {
					if fieldValue.Elem().Type() == reflect.TypeOf(time.Time{}) {
						if fieldValue.Elem().Interface().(time.Time).IsZero() {
							nilOrDefFields = append(nilOrDefFields, fieldType.Name)
						}
					} else {
						nilOrDefFields = findNilOrDefaultFields(fieldValue.Elem(), nilOrDefFields)
					}
				} else {
					zeroValue := reflect.Zero(fieldType.Type).Interface()
					if reflect.DeepEqual(fieldValue.Interface(), zeroValue) {
						nilOrDefFields = append(nilOrDefFields, fieldType.Name)
					}
				}
			}
		case reflect.Struct:
			if fieldType.Type == reflect.TypeOf(time.Time{}) {
				if fieldValue.Interface().(time.Time).IsZero() {
					nilOrDefFields = append(nilOrDefFields, fieldType.Name)
				}
			} else {
				nilOrDefFields = findNilOrDefaultFields(fieldValue, nilOrDefFields)
			}

		default:
			zeroValue := reflect.Zero(fieldType.Type).Interface()
			if reflect.DeepEqual(fieldValue.Interface(), zeroValue) {
				nilOrDefFields = append(nilOrDefFields, fieldType.Name)
			}
		}
	}

	return nilOrDefFields
}
