package database

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/iancoleman/strcase"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var basicTypes = map[reflect.Kind]bool{
	reflect.Bool:   true,
	reflect.Int:    true,
	reflect.Int8:   true,
	reflect.Int16:  true,
	reflect.Int32:  true,
	reflect.Int64:  true,
	reflect.Uint:   true,
	reflect.Uint8:  true,
	reflect.Uint16: true,
	reflect.Uint32: true,
	reflect.Uint64: true,
	reflect.String: true,
}

func isBasicType(v reflect.Type) bool {
	kind := v.Kind()
	if basicTypes[kind] || (kind == reflect.Ptr && basicTypes[v.Elem().Kind()]) {
		return true
	} else if v.String() == "time.Time" || v.String() == "*time.Time" {
		return true
	}

	return false
}

func ValueToInterface(v reflect.Value) interface{} {
	kind := v.Kind()
	if reflect.Int <= kind && kind <= reflect.Int64 {
		return v.Int()
	} else if reflect.Uint <= kind && kind <= reflect.Uint64 {
		return v.Uint()
	} else if reflect.Bool == kind {
		return strconv.FormatBool(v.Bool())
	} else if reflect.String == kind {
		return v.String()
	} else if reflect.Ptr == kind {
		if v.IsNil() {
			return nil
		}
		return ValueToInterface(v.Elem())
	} else if v.Type().String() == "time.Time" {
		return v.Interface().(time.Time)
	} else {
		log.Fatalf("unknown basic type %s", v.Type().String())
	}

	return nil
}

func minus(a, b int) int {
	return a - b
}

var funcMap = template.FuncMap{
	"minus":   minus,
	"deref":   func(i *int64) int64 { return *i },
	"not_nil": func(i *int64) bool { return i != nil },
}

func (d *Database) neo4jInsertOne(v reflect.Value, session neo4j.Session, parentID *int64, relationship string) {
	var uniqueFields []string
	var setFields []string
	var children []reflect.Value
	var childrenNames []string
	var postQueries []string
	var s strings.Builder
	cypherTpl := `
{{- $UniqueFieldsLen := len .UniqueFields }}
{{- if eq 0 $UniqueFieldsLen }}
CREATE (n:{{ .Node }})
{{- else }}
MERGE (n:{{ .Node }} {
	{{- $UniqueFieldsLen = minus $UniqueFieldsLen 1 }}
	{{- range $index, $element := .UniqueFields }}
		{{- if eq $index $UniqueFieldsLen }}
			{{ $element }}: ${{ $element }}
		{{- else }}
			{{ $element }}: ${{ $element }},
		{{- end }}
	{{- end }}
})
{{- end }}
SET
	{{- $SetFieldsLen := len .SetFields }}
	{{- $SetFieldsLen = minus $SetFieldsLen 1 }}
	{{- range $index, $element := .SetFields }}
		{{- if eq $index $SetFieldsLen }}
			n.{{ $element }} = ${{ $element }}
		{{- else }}
			n.{{ $element }} = ${{ $element }},
		{{- end }}
	{{- end }}
{{- if not_nil .ParentID }}
WITH n
MATCH (parent) WHERE id(parent) = {{ deref .ParentID }}
MERGE (parent)-[r:{{ .Relationship }}]->(n)
{{- end }}
RETURN id(n)
`
	NodeName := getNodeName(v)
	params := map[string]interface{}{}
	n := v.NumField()
	for l := 0; l < n; l += 1 {
		fieldType := v.Type().Field(l)
		fieldValue := v.Field(l)
		//kind := fieldType.Type.Kind()
		if strings.ToLower(fieldType.Name) == "id" {
			continue
		} else if fieldType.Name == "_" || isBasicType(fieldType.Type) {
			fieldName := strcase.ToSnake(fieldType.Name)
			tag := fieldType.Tag.Get("neo")
			if strings.Contains(tag, "unique") {
				uniqueFields = append(uniqueFields, fieldName)
			} else if strings.Contains(tag, "ignore") {
				continue
			} else if strings.Contains(tag, "raw") {
				postQueries = append(postQueries, tag[4:])
				continue
			} else {
				setFields = append(setFields, fieldName)
			}
			params[fieldName] = ValueToInterface(fieldValue)
		} else {
			// we are in []*Struct type
			if !fieldValue.IsNil() {
				children = append(children, fieldValue)
				childrenNames = append(childrenNames, strcase.ToScreamingSnake(fieldType.Name))
			}
		}
	}

	t := template.Must(template.New("").Funcs(funcMap).Parse(cypherTpl))
	err := t.Execute(&s, map[string]interface{}{
		"Node":         NodeName,
		"UniqueFields": uniqueFields,
		"SetFields":    setFields,
		"ParentID":     parentID,
		"Relationship": relationship,
	})
	if err != nil {
		log.Fatal(err)
	}

	res, err := session.Run(s.String(), params)
	if err != nil {
		log.Fatal(err)
	}
	record, err := res.Single()
	if err != nil {
		log.Fatal(err)
	}
	nodeID := record.Values[0].(int64)
	for _, query := range postQueries {
		fullQuery := fmt.Sprintf("MATCH (n) WHERE id(n) = %d %s", nodeID, query)
		_, err := session.Run(fullQuery, params)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i, child := range children {
		d.neo4jInsertMany(child, session, &nodeID, childrenNames[i])
	}
}

func (d *Database) neo4jInsertMany(arr reflect.Value, session neo4j.Session, parentID *int64, relationship string) {
	for i := 0; i < arr.Len(); i += 1 {
		v := arr.Index(i).Elem()
		d.neo4jInsertOne(v, session, parentID, relationship)
	}
}
