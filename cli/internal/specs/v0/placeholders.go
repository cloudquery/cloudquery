package specs

import (
	"fmt"
	"reflect"
	"regexp"
)

var reReplacement = regexp.MustCompile(`^@([a-zA-Z0-9_\-]+)\.(connection)$`)

type ReplacementValue struct {
	PluginName string
	Connection string
}

func ReplacePlaceholders(spec any, values []ReplacementValue) error {
	value := reflect.ValueOf(spec)
	switch value.Kind() {
	case reflect.Map:
		iter := value.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := reflect.ValueOf(iter.Value().Interface())
			if v.Kind() == reflect.String {
				nv, err := replaceString(v.String(), values)
				if err != nil {
					return err
				}
				if nv != v.String() {
					value.SetMapIndex(k, reflect.ValueOf(nv))
				}
			}
			err := ReplacePlaceholders(iter.Value().Interface(), values)
			if err != nil {
				return err
			}
		}
	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			if value.Index(i).Kind() == reflect.String {
				nv, err := replaceString(value.Index(i).String(), values)
				if err != nil {
					return err
				}
				value.Index(i).Set(reflect.ValueOf(nv))
			} else {
				err := ReplacePlaceholders(value.Index(i).Interface(), values)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func replaceString(s string, values []ReplacementValue) (string, error) {
	// check if the string contains a placeholder, then replace it if it does
	matches := reReplacement.FindStringSubmatch(s)
	if len(matches) == 3 {
		pluginName := matches[1]
		fieldName := matches[2]
		for _, v := range values {
			if v.PluginName == pluginName {
				if fieldName == "connection" {
					return v.Connection, nil
				}
				// we should never get here, it's excluded by the regular expression
				panic("unknown field name: " + fieldName)
			}
		}
		return "", fmt.Errorf("could not find plugin with name %q, referenced in %q", s, pluginName)
	}
	return s, nil
}
