package discover

import (
	"fmt"
	"reflect"
	"strings"
)

// DiscoveredMethod is a method on the Client struct that was discovered
// to have the right output type. If the output type was found in a return struct field,
// that name is specified in OutputFieldName.
type DiscoveredMethod struct {
	Method          reflect.Method
	OutputFieldName string // field name on the output struct
}

// MethodByName returns a method on the client that matches a specific name.
func MethodByName(client any, targetStruct any, name string) (DiscoveredMethod, error) {
	methods := FindMethods(client, targetStruct, []string{name})
	for _, m := range methods {
		if m.Method.Name == name {
			return m, nil
		}
	}
	tt := reflect.TypeOf(targetStruct).Elem()
	return DiscoveredMethod{}, fmt.Errorf("method %v that returns %v not found", name, tt.Name())
}

// FindMethods returns a slice of methods that return the targetStruct (either directly or as a field
// of the returned struct) and start with one of the given prefixes.
func FindMethods(client, targetStruct any, prefixes []string) []DiscoveredMethod {
	v := reflect.ValueOf(client)
	t := v.Type()
	var tt reflect.Type
	if targetStruct != nil {
		tt = reflect.TypeOf(targetStruct).Elem()
	}
	found := make([]DiscoveredMethod, 0)
	// iterate over the methods of the given client
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		hasPrefix := false
		for _, prefix := range prefixes {
			if strings.HasPrefix(method.Name, prefix) {
				hasPrefix = true
			}
		}
		if len(prefixes) > 0 && !hasPrefix {
			// skip methods that don't start with any of the given prefixes
			continue
		}

		if tt == nil {
			// no specific target struct was requested, so return any method
			// that matches the prefix requirements
			found = append(found, DiscoveredMethod{
				Method: method,
			})
			continue
		}

		// iterate over the output types of the method to find matches for targetStruct
		f := v.Method(i).Interface()
		ft := reflect.TypeOf(f)
		for u := 0; u < ft.NumOut(); u++ {
			e := ft.Out(u)
			if ft.Out(u).Kind() == reflect.Ptr || ft.Out(u).Kind() == reflect.Slice {
				e = e.Elem()
			}
			if e.Kind() != reflect.Struct {
				continue
			}
			if tt == nil || reflect.TypeOf(e) == tt {
				// struct is directly returned by the method
				found = append(found, DiscoveredMethod{
					Method: method,
				})
				continue
			}
			for x := 0; x < e.NumField(); x++ {
				ef := e.Field(x)
				eft := ef.Type
				if eft.Kind() == reflect.Ptr || eft.Kind() == reflect.Slice {
					eft = ef.Type.Elem()
				}
				if tt == nil || eft == tt {
					// struct is returned by the method as a field in an Output struct
					found = append(found, DiscoveredMethod{
						Method:          method,
						OutputFieldName: ef.Name,
					})
				}
			}
		}
	}
	return found
}

// FindMethod finds a single method and errors out if zero or more than one is discovered.
func FindMethod(client, targetStruct any, prefixes []string) (DiscoveredMethod, error) {
	f := FindMethods(client, targetStruct, prefixes)
	if len(f) == 0 {
		if targetStruct == nil {
			return DiscoveredMethod{}, fmt.Errorf("found no method with prefix in %v", prefixes)
		}
		name := reflect.TypeOf(targetStruct).Elem().Name()
		return DiscoveredMethod{}, fmt.Errorf("found no method with prefix in %v that also returns %v", prefixes, name)
	}
	if len(f) > 1 {
		if targetStruct == nil {
			return DiscoveredMethod{}, fmt.Errorf("found %d potential methods for prefix: %v", len(f), prefixes)
		}
		name := reflect.TypeOf(targetStruct).Elem().Name()
		names := make([]string, len(f))
		for i := range f {
			names[i] = f[i].Method.Name
		}
		return DiscoveredMethod{}, fmt.Errorf("found %d potential methods that also return %s: %v", len(f), name, strings.Join(names, ","))
	}
	return f[0], nil
}

func FindDescribeMethod(client, targetStruct any) (DiscoveredMethod, error) {
	return FindMethod(client, targetStruct, []string{"Describe", "Get"})
}

func FindListMethod(client, targetStruct any) (DiscoveredMethod, error) {
	return FindMethod(client, targetStruct, []string{"List"})
}

func FindListTagsMethod(client any) (DiscoveredMethod, error) {
	return FindMethod(client, nil, []string{"ListTagsForResource", "ListTagsOfResource"})
}
