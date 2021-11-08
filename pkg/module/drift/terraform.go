package drift

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/cloudquery/cloudquery/pkg/module/drift/terraform"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/olekukonko/tablewriter"
)

type TFStates []*terraform.Data

// FindType returns all instances of the given type under a given mode
func (t TFStates) FindType(tfType, tfMode string) TFInstances {
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

// Attributes returns a map of resource ID vs. attributes
func (r TFInstances) AsResourceList(attributeNames []string) ResourceList {
	ret := make([]*Resource, len(r))
	for i := range r {
		var attributes map[string]interface{}
		if err := json.Unmarshal(r[i].AttributesRaw, &attributes); err != nil {
			panic(err)
		}

		res := &Resource{
			ID: attributes["id"].(string),
		}
		res.Attributes = make([]interface{}, len(attributeNames))
		for i := range attributeNames {
			if val, ok := attributes[attributeNames[i]]; ok {
				res.Attributes[i] = val
			}
		}

		ret[i] = res
	}
	return ret
}

func (d *Drift) driftTerraform(ctx context.Context, conn *pgxpool.Conn, cloudName string, cloudTable *traversedTable, resName string, resources map[string]*ResourceConfig, iacData *IACConfig, states TFStates) (*Result, error) {
	res := &Result{
		IAC:       "Terraform",
		Different: nil,
		Equal:     nil,
		Missing:   nil,
		Extra:     nil,
	}

	resData := resources[resName]
	deepMode := d.params.ForceDeep || (resData.Deep != nil && *resData.Deep)

	tfAttributes := make([]string, len(resData.Attributes))
	for i, a := range resData.Attributes {
		if mapped := iacData.attributeMap[a]; mapped != "" {
			tfAttributes[i] = mapped
		} else {
			tfAttributes[i] = a
		}
	}

	tfResources := states.FindType(iacData.Type, d.params.TfMode).AsResourceList(tfAttributes)

	cloudQueryItems := make([]string, len(resData.Attributes))
	for i := range resData.Attributes {
		if cloudTable.Column(resData.Attributes[i]).Type == schema.TypeString {
			cloudQueryItems[i] = fmt.Sprintf(`COALESCE("c"."%s",'')`, resData.Attributes[i])
		} else {
			cloudQueryItems[i] = fmt.Sprintf(`"c"."%s"`, resData.Attributes[i])
		}
	}

	var cloudAttrQuery exp.LiteralExpression

	if !deepMode || len(resData.Attributes) == 0 {
		cloudAttrQuery = goqu.L("NULL")
	} else {
		cloudAttrQuery = goqu.L("JSONB_BUILD_ARRAY(" + strings.Join(cloudQueryItems, ",") + ")")
	}

	idExp, err := d.handleIdentifier(resData.Identifiers)
	if err != nil {
		return nil, err
	}

	q := goqu.Dialect("postgres").From(goqu.T(cloudTable.Name).As("c")).Select(idExp, cloudAttrQuery.As("attlist"))
	q = d.handleSubresource(q, cloudTable, resources)
	existing, err := d.queryIntoResourceList(ctx, conn, q)
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
		q = d.handleSubresource(q, cloudTable, resources)
		q = d.handleFilters(q, resources[resName]) // This line (the application of filters) is the difference from "existing"
		existingFiltered, err := d.queryIntoResourceList(ctx, conn, q)
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
	if deepMode && d.params.Debug && len(res.Different) > 0 {
		if err := d.terraformDebugDifferentResources(resName, resources, cloudName, cloudQueryItems, tfAttributes, res.Different, tfResources); err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (d *Drift) terraformDebugDifferentResources(resName string, resources map[string]*ResourceConfig, cloudName string, cloudQueryItems, tfAttributes []string, differentIDs, tfRes ResourceList) error {
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
			continue // Resource exists only in cloud. This is already handled by the "Extra" resource/check
		}

		tfAttrs := tfMap[k]
		table := makeTable(fmt.Sprintf("DIFF RESOURCE: %s", k))
		var (
			matchingAttr []string
			matchingVal  []string
		)
		for i := range tfAttrs {
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
