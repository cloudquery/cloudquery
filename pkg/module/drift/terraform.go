package drift

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/pkg/module/drift/terraform"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/olekukonko/tablewriter"
)

type TFStates []*terraform.Data

// FindType returns all instances of the given type under a given mode
func (t TFStates) FindType(tfType string, tfMode terraform.Mode) TFInstances {
	var ret []terraform.Instance
	for _, d := range t {
		for idx, r := range d.State.Resources {
			if tfMode != "" && r.Mode != tfMode {
				continue
			}
			if r.Type != tfType {
				continue
			}
			ret = append(ret, d.State.Resources[idx].Instances...)
		}
	}
	return ret
}

type TFInstances []terraform.Instance

const tfIDAttribute = "id"

// Attributes returns a map of resource ID vs. attributes
func (r TFInstances) AsResourceList(attributeNames []string, attributeTypes []schema.ValueType) ResourceList {
	ret := make([]*Resource, len(r))
	for i := range r {
		var attributes map[string]interface{}
		if err := json.Unmarshal(r[i].AttributesRaw, &attributes); err != nil {
			panic(err)
		}

		res := &Resource{
			ID: attributes[tfIDAttribute].(string),
		}
		res.Attributes = make([]interface{}, len(attributeNames))
		for i := range attributeNames {
			if val, ok := attributes[attributeNames[i]]; ok {
				res.Attributes[i] = parseTerraformAttribute(val, attributeTypes[i])
			}
		}

		ret[i] = res
	}
	return ret
}

func parseTerraformAttribute(val interface{}, t schema.ValueType) interface{} {
	switch t {
	case schema.TypeTimestamp:
		ts, err := time.Parse(time.RFC3339, val.(string))
		if err != nil {
			ts, err = time.Parse("2006-01-02 15:04:05 -0700 MST", val.(string))
		}
		if err != nil {
			return val // will probably error/detect deep drift
		}
		return fmt.Sprintf("%d", ts.Unix())
	default:
		return val
	}
}

func driftTerraform(ctx context.Context, logger hclog.Logger, conn *pgxpool.Conn, cloudName string, cloudTable *traversedTable, resName string, resources map[string]*ResourceConfig, iacData *IACConfig, states TFStates, runParams RunParams) (*Result, error) {
	res := &Result{
		IAC:       "Terraform",
		Different: nil,
		Equal:     nil,
		Missing:   nil,
		Extra:     nil,
	}

	resData := resources[resName]
	deepMode := runParams.ForceDeep || (resData.Deep != nil && *resData.Deep)

	tfAttributes := make([]string, len(resData.Attributes))
	colTypes := make([]schema.ValueType, len(resData.Attributes))
	for i, a := range resData.Attributes {
		colTypes[i] = cloudTable.Column(resData.Attributes[i]).Type
		if mapped := iacData.attributeMap[a]; mapped != "" {
			tfAttributes[i] = mapped
		} else {
			tfAttributes[i] = a
		}
	}

	tfMode := terraform.Mode(runParams.TfMode)
	if !tfMode.Valid() {
		return nil, fmt.Errorf("invalid tf mode %q", runParams.TfMode)
	}

	tfResources := states.FindType(iacData.Type, tfMode).AsResourceList(tfAttributes, colTypes)

	cloudQueryItems := make([]string, len(resData.Attributes))
	for i := range resData.Attributes {
		switch colTypes[i] {
		case schema.TypeString:
			cloudQueryItems[i] = fmt.Sprintf(`COALESCE("c"."%s",'')`, resData.Attributes[i])
		case schema.TypeTimestamp:
			cloudQueryItems[i] = fmt.Sprintf(`EXTRACT(EPOCH FROM DATE_TRUNC('second', "c"."%s"))::VARCHAR`, resData.Attributes[i])
		default:
			cloudQueryItems[i] = fmt.Sprintf(`"c"."%s"`, resData.Attributes[i])
		}
	}

	var cloudAttrQuery exp.LiteralExpression

	if !deepMode || len(resData.Attributes) == 0 {
		cloudAttrQuery = goqu.L("NULL")
	} else {
		cloudAttrQuery = goqu.L("JSON_BUILD_ARRAY(" + strings.Join(cloudQueryItems, ",") + ")")
	}

	idExp, err := handleIdentifier(resData.Identifiers)
	if err != nil {
		return nil, err
	}

	q := goqu.Dialect("postgres").From(goqu.T(cloudTable.Name).As("c")).Select(idExp, cloudAttrQuery.As("attlist"))
	q = handleSubresource(logger, q, cloudTable, resources, runParams.AccountIDs)
	existing, err := queryIntoResourceList(ctx, logger, conn, q)
	if err != nil {
		return nil, err
	}

	existingMap := existing.Map()
	tfMap := tfResources.Map()

	// Get missing resources
	tfResources.Walk(func(r *Resource) {
		if _, ok := existingMap[r.ID]; !ok {
			res.Missing = append(res.Missing, r)
		}
	})

	// Get extra resources
	{
		q := goqu.Dialect("postgres").From(goqu.T(cloudTable.Name).As("c")).Select(idExp, cloudAttrQuery.As("attlist"))
		q = handleSubresource(logger, q, cloudTable, resources, runParams.AccountIDs)
		q = handleFilters(q, resources[resName]) // This line (the application of filters) is the difference from "existing"
		existingFiltered, err := queryIntoResourceList(ctx, logger, conn, q)
		if err != nil {
			return nil, err
		}

		existingFiltered.Walk(func(r *Resource) {
			if _, ok := tfMap[r.ID]; !ok {
				res.Extra = append(res.Extra, r)
			}
		})
	}

	if !deepMode {
		// Get equal resources
		existing.Walk(func(r *Resource) {
			if _, ok := tfMap[r.ID]; ok {
				res.Equal = append(res.Equal, r)
			}
		})
	} else {
		// Get deepequal and different resources
		existing.Walk(func(r *Resource) {
			tfAttr, ok := tfMap[r.ID]
			if !ok {
				return
			}
			if reflect.DeepEqual(tfAttr, r.Attributes) {
				res.DeepEqual = append(res.DeepEqual, r)
			} else {
				res.Different = append(res.Different, r)
			}
		})
	}
	if deepMode && runParams.Debug && len(res.Different) > 0 {
		if err := RenderDriftTable(resName, resources, cloudName, cloudQueryItems, tfAttributes, res.Different, tfResources); err != nil {
			return nil, err
		}
	}

	return res, nil
}

func RenderDriftTable(resName string, resources map[string]*ResourceConfig, cloudName string, cloudQueryItems, tfAttributes []string, differentIDs, tfRes ResourceList) error {
	resData := resources[resName]

	makeTable := func(title string) *tablewriter.Table {
		fmt.Println(title)
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{strings.ToUpper(cloudName) + " EXPR", strings.ToUpper(cloudName) + " VAL", "TERRAFORM VAL", "TERRAFORM EXPR"})
		table.SetBorder(true)
		return table
	}

	tfMap, cloudMap := tfRes.Map(), differentIDs.Map()
	for _, k := range tfRes.IDs() {
		cloudAttrs, ok := cloudMap[k]
		if !ok {
			continue // Resource exists only in TF. This is already handled by the "Missing" resource/check
		}

		tfAttrs := tfMap[k]
		table := makeTable(fmt.Sprintf("DIFF RESOURCE: %s:%s", resName, k))
		var (
			matchingAttr []string
			matchingVal  []string
		)
		for i := range tfAttrs {
			if tfAttributes[i] == tfIDAttribute {
				continue // don't print ID attributes (cloud side might not match due to use of composite IDs)
			}

			if !reflect.DeepEqual(cloudAttrs[i], tfAttrs[i]) {
				table.Append([]string{
					cloudQueryItems[i],
					fmt.Sprintf("%v", cloudAttrs[i]),
					fmt.Sprintf("%v", tfAttrs[i]),
					tfAttributes[i],
				})
			} else {
				matchingAttr = append(matchingAttr, `"`+resData.Attributes[i]+`"`)
				matchingVal = append(matchingVal, fmt.Sprintf("%v", cloudAttrs[i]))
			}
		}
		table.Render()
		fmt.Println("Matching attributes " + strings.Join(matchingAttr, ", "))
		table = tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ATTRIBUTE", "MATCHING VALUE"})
		table.SetBorder(true)
		for i := range matchingAttr {
			table.Append([]string{strings.Trim(matchingAttr[i], `"`), matchingVal[i]})
		}
		table.Render()
	}

	return nil
}
