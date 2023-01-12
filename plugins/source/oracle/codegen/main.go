// Prerequisites: A populated `client.OracleClients` struct.
// Outputs resource structs in `resources/services/...`

package main

import (
	"context"
	"fmt"
	"log"
	"path"
	"reflect"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/iancoleman/strcase"
	"golang.org/x/exp/slices"
)

// These services were manually tested, and they work.
var servicesAllowlist = []string{
	"compute",
	"identity",
	"virtualnetwork",
	"database",
	"blockstorage",
	"objectstorage",
	"filestorage",
	"loadbalancer",
	"networkloadbalancer",
	"networkfirewall",
}

var servicesDenylist = []string{
	"identity", // Skipped because most of these resources are global

	"database",      // Skipped because some tables require extra parameters that I don't know where to find...
	"objectstorage", // skipped because it requires usage of 'objectstoragenamespace'..

	"blockstorage", // Skipped because the recipe-autogeneration missed the most important tables (volumes)

	"filestorage", // Skipped because requests are AD-based, not region-based.
}

func main() {
	oracleClientsStruct := client.OracleClients{}

	oracleClientsStructType := reflect.TypeOf(oracleClientsStruct)

	tables := make([]Table, 0)

	for i := 0; i < oracleClientsStructType.NumField(); i++ {
		clientField := oracleClientsStructType.Field(i)
		clientType := clientField.Type

		numTopLevelCompartmentScopeMethods := 0

		pkg := path.Base(clientType.Elem().PkgPath())
		service := extractServiceNameFromClient(clientType.Elem())

		if !slices.Contains(servicesAllowlist, service) {
			continue
		}

		if slices.Contains(servicesDenylist, service) {
			continue
		}

		for j := 0; j < clientType.NumMethod(); j++ {
			method := clientType.Method(j)

			if !strings.HasPrefix(method.Name, "List") {
				continue
			}

			isValid, reason := isListMethodValid(method)
			if !isValid {
				fmt.Println("Skipping method", clientType, method.Name, "because", reason)
				continue
			}

			if !isTopLevelCompartmentScopeMethod(method) {
				continue
			}

			table, err := createTableStruct(pkg, service, method, clientField.Name)
			if err != nil {
				fmt.Println("Skipping method", clientType, method.Name, "because", err)
				continue
			}

			tables = append(tables, table)
			numTopLevelCompartmentScopeMethods++
		}

		if numTopLevelCompartmentScopeMethods == 0 {
			// fmt.Println("WARNING: No top-level copmartment-scope List methods found for", clientType.String())
			continue
		}
	}

	if err := generateTables(tables); err != nil {
		log.Fatal(err)
	}

	if err := generateFetchers(tables); err != nil {
		log.Fatal(err)
	}

	if err := generateTableList(tables); err != nil {
		log.Fatal(err)
	}
}

func extractServiceNameFromClient(clientType reflect.Type) string {
	clientName := clientType.Name()
	pkg := path.Base(clientType.PkgPath())

	if !strings.HasSuffix(clientName, "Client") {
		panic(fmt.Sprintf("Client name '%s' does not end with 'Client'", clientName))
	}

	// The "core" package contains several different "services", but all the
	// rest are pretty-good 1:1 mappings to what we like to call "services".
	if pkg == "core" {
		return strings.ToLower(strings.TrimSuffix(clientName, "Client"))
	} else {
		return pkg
	}
}

// Checks some prerequisites for a `List` method.
// If the list method is not 'valid', it will also return a string "reason"
func isListMethodValid(method reflect.Method) (bool, string) {
	if method.Func.Type().NumIn() != 3 { // The "receiver" counts as an input
		return false, fmt.Sprintf("it has %d inputs, but it should have exactly 3", method.Func.Type().NumIn())
	}

	if !(method.Func.Type().In(1).Implements(reflect.TypeOf((*context.Context)(nil)).Elem())) {
		return false, "its first input is not a context.Context"
	}

	if method.Func.Type().In(2).Kind() != reflect.Struct {
		return false, "its second input is not a struct"
	}

	if !strings.HasSuffix(method.Func.Type().In(2).Name(), "Request") {
		return false, "its second input is not a request struct"
	}

	_, isPaginated := method.Func.Type().In(2).FieldByName("Page")
	if !isPaginated {
		return false, "it is not paginated"
	}

	if method.Func.Type().NumOut() != 2 {
		return false, fmt.Sprintf("it has %d outputs, but it should have exactly 2", method.Func.Type().NumOut())
	}

	if !method.Func.Type().Out(1).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		return false, "its second output is not an error"
	}

	if method.Func.Type().Out(0).Kind() != reflect.Struct {
		return false, "its first output is not a struct"
	}

	if !strings.HasSuffix(method.Func.Type().Out(0).Name(), "Response") {
		return false, "its first output is not a response struct"
	}

	return true, ""
}

// Top-level `List` methods usually have exactly on `mandatory` field in the request struct - the `compartment_id`.
// Relational `List` methods usually have a `mandatory` field that is the parent's id.
// See examples:
// - https://github.com/oracle/oci-go-sdk/blob/1a34e432f90fade18d83a84b6d8921e8c9ddd5b7/core/list_instances_request_response.go#L19-L22
// - https://github.com/oracle/oci-go-sdk/blob/1a34e432f90fade18d83a84b6d8921e8c9ddd5b7/core/list_instance_devices_request_response.go#L19-L22
func isTopLevelCompartmentScopeMethod(method reflect.Method) bool {
	requestStruct := method.Func.Type().In(2)

	if requestStruct.Kind() == reflect.Pointer {
		requestStruct = requestStruct.Elem()
	}

	numMandatoryFields := 0
	mandatoryFieldIndex := -1
	for i := 0; i < requestStruct.NumField(); i++ {
		field := requestStruct.Field(i)
		if field.Tag.Get("mandatory") == "true" {
			numMandatoryFields++
			mandatoryFieldIndex = i
		}
	}

	if numMandatoryFields != 1 {
		return false
	}

	mandatoryField := requestStruct.Field(mandatoryFieldIndex)

	return mandatoryField.Name == "CompartmentId"
}

func createTableStruct(pkg string, service string, method reflect.Method, clientName string) (Table, error) {
	responseStruct := method.Func.Type().Out(0)
	itemsField, ok := responseStruct.FieldByName("Items")
	if !ok {
		return Table{}, fmt.Errorf("response struct does not have an 'Items' field")
	}

	itemsType := itemsField.Type
	if itemsType.Kind() != reflect.Slice {
		return Table{}, fmt.Errorf("field 'Items' in struct '%s' is not a slice", responseStruct.Name())
	}

	itemsElemType := itemsType.Elem()

	if itemsElemType.Kind() != reflect.Struct {
		return Table{}, fmt.Errorf("item in 'Items' field is not a struct (but a %s)", itemsElemType.Kind())
	}

	_, hasId := itemsElemType.FieldByName("Id")
	if !hasId {
		return Table{}, fmt.Errorf("struct '%s' does not have an 'Id' field", itemsElemType.Name())
	}

	return Table{
		Package:          pkg,
		Service:          service,
		SubService:       extractSubserviceNameFromListMethodName(method.Name),
		StructName:       itemsElemType.Name(),
		ClientName:       clientName,
		ListFunctionName: method.Name,
	}, nil
}

// E.g. ListInstances -> instances
func extractSubserviceNameFromListMethodName(listMethodName string) string {
	if !strings.HasPrefix(listMethodName, "List") {
		panic(fmt.Sprintf("List method name '%s' does not start with 'List'", listMethodName))
	}

	return strcase.ToSnake(strings.TrimPrefix(listMethodName, "List"))
}
