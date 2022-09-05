package helpers

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/recipes"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

func TableAndFetcherNames(r *recipes.Resource) (string, string) {
	cqSubservice := Coalesce(r.CQSubserviceOverride, r.AWSSubService)

	tableNameFromSubService := strcase.ToSnake(cqSubservice)
	fetcherNameFromSubService := strcase.ToCamel(cqSubservice)
	{
		// Generate table and fetcher names using parent info

		prev := tableNameFromSubService
		var (
			preTableNames   []string
			preFetcherNames []string
		)
		rp := r.Parent
		for rp != nil {
			if rp.CQSubserviceOverride != "" {
				preTableNames = append(preTableNames, rp.CQSubserviceOverride)
				preFetcherNames = append(preFetcherNames, strcase.ToCamel(rp.CQSubserviceOverride))
			} else {
				ins := strcase.ToSnake(rp.ItemName)
				if !strings.HasPrefix(prev, ins) {
					preTableNames = append(preTableNames, ins)
					preFetcherNames = append(preFetcherNames, strcase.ToCamel(rp.ItemName))
					prev = ins
				}
			}
			rp = rp.Parent
		}
		if len(preTableNames) > 0 {
			tableNameFromSubService = strings.Join(ReverseStringSlice(preTableNames), "_") + "_" + tableNameFromSubService
			fetcherNameFromSubService = strings.Join(ReverseStringSlice(preFetcherNames), "") + fetcherNameFromSubService
		}
	}

	return tableNameFromSubService, fetcherNameFromSubService
}

type InferResult struct {
	Method     string
	SubService string

	PaginatorTokenField *reflect.StructField

	Fields     map[string]reflect.StructField
	FieldOrder []string
}

// ItemsField returns the field from the struct field candidates that contains the item or items. Only valid in Output type structs.
func (ir *InferResult) ItemsField(singular bool, hint string) reflect.StructField {
	cands := ir.ItemsFieldCandidates(singular)

	if len(cands) != 1 {
		var trc string
		if _, fn, ln, ok := runtime.Caller(1); ok {
			trc = fmt.Sprintf(" (called from %s:%d)", fn, ln)
		}

		cl := make([]string, len(cands))
		for i, c := range cands {
			// if there is a hint, use it
			if hint != "" && (c.Name == hint || c.Name == inflection.Plural(hint)) {
				return cands[i]
			}

			cl[i] = c.Name
		}
		log.Fatal("Could not determine ItemsName for ", ir.Method, ":", len(cands), " candidates: ", strings.Join(cl, ", ")+trc)
	}
	return cands[0]
}

// ItemsFieldCandidates returns candidates for the item or items. Only valid in Output type structs.
func (ir *InferResult) ItemsFieldCandidates(singular bool) []reflect.StructField {
	var cands []reflect.StructField
	for _, n := range ir.FieldOrder {
		f := ir.Fields[n]
		isSlice := BareType(f.Type).Kind() == reflect.Slice
		if singular == !isSlice && f.Name != "ETag" {
			cands = append(cands, f)
		}
	}
	return cands
}

func InferFromStructInput(s interface{}) *InferResult {
	return InferFromType(BareType(reflect.TypeOf(s)), "Input")
}

func InferFromStructOutput(s interface{}) *InferResult {
	return InferFromType(BareType(reflect.TypeOf(s)), "Output")
}

func InferFromType(ist reflect.Type, suffix string) *InferResult {
	res := InferResult{
		Fields: make(map[string]reflect.StructField),
	}

	for _, verbCandidate := range []string{"Get", "Describe", "List"} {
		if !strings.HasPrefix(ist.Name(), verbCandidate) {
			continue
		}

		if suffix != "" && !strings.HasSuffix(ist.Name(), suffix) {
			log.Fatal("Unhandled struct type (bad suffix): ", ist.Name())
		}

		res.Method = strings.TrimSuffix(ist.Name(), suffix)

		res.SubService = strings.TrimPrefix(res.Method, verbCandidate)
		if res.SubService == "" {
			log.Fatal("Failed calculating AWSSubService: empty output for", ist.Name())
		}

		break
	}

	for i := 0; i < ist.NumField(); i++ {
		f := ist.Field(i)
		if f.Name == "noSmithyDocumentSerde" || f.Type.String() == "document.NoSerde" || f.Type.String() == "middleware.Metadata" {
			continue
		}
		if f.Name == "NextToken" || f.Name == "NextMarker" {
			res.PaginatorTokenField = &f
		}

		res.Fields[f.Name] = f
		res.FieldOrder = append(res.FieldOrder, f.Name)
	}

	return &res
}
