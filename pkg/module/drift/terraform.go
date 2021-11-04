package drift

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/olekukonko/tablewriter"
)

func (d *Drift) driftTerraform(ctx context.Context, conn *pgxpool.Conn, cloudName string, cloudTable *traversedTable, resName string, resources map[string]*ResourceConfig, iacData *IACConfig) (*Result, error) {
	res := &Result{
		IAC:         "Terraform",
		Different:   nil,
		Equal:       nil,
		Missing:     nil,
		Extra:       nil,
		ListManaged: d.params.ListManaged,
	}

	tfProvider := d.params.TfProvider
	if tfProvider == "" {
		tfProvider = cloudName
	}

	resData := resources[resName]

	tfAttributes := make([]string, len(resData.Attributes))
	tfQueryItems := make([]string, len(resData.Attributes))
	for i, a := range resData.Attributes {
		if mapped := iacData.attributeMap[a]; mapped != "" {
			tfAttributes[i] = mapped
		} else {
			tfAttributes[i] = a
		}
		switch cloudTable.Column(a).Type {
		case schema.TypeString:
			tfQueryItems[i] = fmt.Sprintf(`COALESCE(i.attributes->>'%s','')`, tfAttributes[i])
		case schema.TypeJSON:
			tfQueryItems[i] = fmt.Sprintf(`(i.attributes->>'%s')::json`, tfAttributes[i])
		default:
			tfQueryItems[i] = fmt.Sprintf(`i.attributes->>'%s'`, tfAttributes[i])
		}
	}

	cloudQueryItems := make([]string, len(resData.Attributes))
	for i := range resData.Attributes {
		if cloudTable.Column(resData.Attributes[i]).Type == schema.TypeString {
			cloudQueryItems[i] = fmt.Sprintf(`COALESCE("c"."%s",'')`, resData.Attributes[i])
		} else {
			cloudQueryItems[i] = fmt.Sprintf(`"c"."%s"`, resData.Attributes[i])
		}
	}

	tfAttrQuery := goqu.L("JSONB_BUILD_ARRAY(" + strings.Join(tfQueryItems, ",") + ")")
	cloudAttrQuery := goqu.L("JSONB_BUILD_ARRAY(" + strings.Join(cloudQueryItems, ",") + ")")

	if len(resData.Attributes) == 0 {
		tfAttrQuery = goqu.L("''")
		cloudAttrQuery = goqu.L("''")
	}

	tfSelect := goqu.Dialect("postgres").From(goqu.T("tf_resource_instances").As("i")).
		Select("i.instance_id", tfAttrQuery.As("attlist")).
		Join(goqu.T("tf_resources").As("r"), goqu.On(goqu.Ex{"r.cq_id": goqu.I("i.resource_id")})).
		Join(goqu.T("tf_data").As("d"), goqu.On(goqu.Ex{"d.cq_id": goqu.I("r.running_id")})).
		Where(goqu.Ex{"r.provider": goqu.V(tfProvider)}).
		Where(goqu.Ex{"r.mode": goqu.V(d.params.TfMode)}).
		Where(goqu.Ex{"r.type": goqu.V(iacData.Type)})

	if len(d.params.TfBackendNames) > 0 {
		tfSelect = tfSelect.Where(goqu.Ex{"d.backend_name": d.params.TfBackendNames})
	}

	deepMode := d.params.ForceDeep || (resData.Deep != nil && *resData.Deep)

	rawIdExp, idExp, err := d.handleIdentifier(resData.Identifiers)
	if err != nil {
		return nil, err
	}
	matchExp := goqu.Ex{"tf.instance_id": rawIdExp}

	if !deepMode {
		// Get equal resources
		res.Equal, err = d.terraformEqualResources(ctx, conn, cloudTable, resources, tfSelect, matchExp)
		if err != nil {
			return nil, err
		}
	}

	// Get missing resources
	res.Missing, err = d.terraformMissingResources(ctx, conn, cloudTable, resources, tfSelect, matchExp)
	if err != nil {
		return nil, err
	}

	// Get extra resources
	res.Extra, err = d.terraformExtraResources(ctx, conn, cloudTable, resName, resources, tfSelect, matchExp, idExp)
	if err != nil {
		return nil, err
	}

	if deepMode {
		// Get different resources
		res.Different, err = d.terraformDifferentResources(ctx, conn, cloudTable, resources, tfSelect, matchExp, idExp, cloudAttrQuery, append(res.Missing, res.Extra...))
		if err != nil {
			return nil, err
		}

		if d.params.Debug && len(res.Different) > 0 {
			if err := d.terraformDebugDifferentResources(ctx, conn, cloudTable, resName, resources, tfSelect, idExp, cloudAttrQuery, cloudName, cloudQueryItems, tfQueryItems, res.Different); err != nil {
				return nil, err
			}
		}
	}

	if deepMode {
		// Get deepequal resources
		res.DeepEqual, err = d.terraformDeepEqualResources(ctx, conn, cloudTable, resources, tfSelect, matchExp, cloudAttrQuery)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (d *Drift) terraformEqualResources(ctx context.Context, conn *pgxpool.Conn, cloudTable *traversedTable, resources map[string]*ResourceConfig, tfSelect *goqu.SelectDataset, matchExp goqu.Ex) (ResourceList, error) {
	q := goqu.Dialect("postgres").From(goqu.T(cloudTable.Name).As("c"))
	q = d.handleSubresource(q, cloudTable, resources)
	q = q.With("tf", tfSelect).Join(goqu.T("tf"), goqu.On(matchExp)).
		Select("tf.instance_id")
	return d.queryIntoResourceList(ctx, conn, q, "equals", nil)
}

func (d *Drift) terraformMissingResources(ctx context.Context, conn *pgxpool.Conn, cloudTable *traversedTable, resources map[string]*ResourceConfig, tfSelect *goqu.SelectDataset, matchExp goqu.Ex) (ResourceList, error) {
	q := goqu.Dialect("postgres").From(goqu.T(cloudTable.Name).As("c"))
	q = d.handleSubresource(q, cloudTable, resources)
	q = q.With("tf", tfSelect).LeftJoin(goqu.T("tf"), goqu.On(matchExp)).
		Select("tf.instance_id").Where(goqu.Ex{"c.cq_id": nil})
	return d.queryIntoResourceList(ctx, conn, q, "missing", nil)
}

func (d *Drift) terraformExtraResources(ctx context.Context, conn *pgxpool.Conn, cloudTable *traversedTable, resName string, resources map[string]*ResourceConfig, tfSelect *goqu.SelectDataset, matchExp goqu.Ex, idExp exp.Expression) (ResourceList, error) {
	q := goqu.Dialect("postgres").From(goqu.T(cloudTable.Name).As("c"))
	q = d.handleSubresource(q, cloudTable, resources)
	q = q.With("tf", tfSelect).LeftJoin(goqu.T("tf"), goqu.On(matchExp)).
		Select(idExp).Where(goqu.Ex{"tf.instance_id": nil})
	q = d.handleFilters(q, resources[resName])
	return d.queryIntoResourceList(ctx, conn, q, "extras", nil)
}

func (d *Drift) terraformDifferentResources(ctx context.Context, conn *pgxpool.Conn, cloudTable *traversedTable, resources map[string]*ResourceConfig, tfSelect *goqu.SelectDataset, matchExp goqu.Ex, idExp goqu.Expression, cloudAttrQuery exp.LiteralExpression, ignoreRes ResourceList) (ResourceList, error) {
	q := goqu.Dialect("postgres").From(goqu.T(cloudTable.Name).As("c"))
	q = d.handleSubresource(q, cloudTable, resources)
	q = q.With("tf", tfSelect).LeftJoin(goqu.T("tf"),
		goqu.On(
			matchExp,
			goqu.L("? @> ?", goqu.I("tf.attlist"), cloudAttrQuery),
			goqu.L("? <@ ?", goqu.I("tf.attlist"), cloudAttrQuery),
		),
	).
		Select(idExp).Where(goqu.Ex{"tf.instance_id": nil})

	return d.queryIntoResourceList(ctx, conn, q, "differs", ignoreRes)
}

func (d *Drift) terraformDeepEqualResources(ctx context.Context, conn *pgxpool.Conn, cloudTable *traversedTable, resources map[string]*ResourceConfig, tfSelect *goqu.SelectDataset, matchExp goqu.Ex, cloudAttrQuery exp.LiteralExpression) (ResourceList, error) {
	q := goqu.Dialect("postgres").From(goqu.T(cloudTable.Name).As("c"))
	q = d.handleSubresource(q, cloudTable, resources)
	q = q.With("tf", tfSelect).Join(goqu.T("tf"),
		goqu.On(
			matchExp,
			goqu.L("? @> ?", goqu.I("tf.attlist"), cloudAttrQuery),
			goqu.L("? <@ ?", goqu.I("tf.attlist"), cloudAttrQuery),
		),
	).
		Select("tf.instance_id")
	return d.queryIntoResourceList(ctx, conn, q, "deepequal", nil)
}

func (d *Drift) terraformDebugDifferentResources(ctx context.Context, conn *pgxpool.Conn, cloudTable *traversedTable, resName string, resources map[string]*ResourceConfig, tfSelect *goqu.SelectDataset, idExp goqu.Expression, cloudAttrQuery exp.LiteralExpression, cloudName string, cloudQueryItems, tfQueryItems []string, differentIDs ResourceList) error {
	resData := resources[resName]

	// get tf side
	sel := goqu.Dialect("postgres").From("tf").With("tf", tfSelect).Select(goqu.I("tf.instance_id").As("id"), goqu.I("tf.attlist").As("attlist")).Where(
		goqu.Ex{
			"tf.instance_id": differentIDs.IDs(),
		})
	tfAttList, err := d.queryIntoAttributeList(ctx, conn, sel, "attlist-tf")
	if err != nil {
		return err
	}

	// get cloud side
	sel = goqu.Dialect("postgres").From(goqu.T(cloudTable.Name).As("c")).Select(idExp, cloudAttrQuery.As("attlist")).Where(
		exp.NewBooleanExpression(exp.InOp, goqu.I("c."+resData.Identifiers[0]), differentIDs.IDs()),
	)
	sel = d.handleSubresource(sel, cloudTable, resources)
	cloudAttList, err := d.queryIntoAttributeList(ctx, conn, sel, "attlist-cloud")
	if err != nil {
		return err
	}

	makeTable := func(title string) *tablewriter.Table {
		fmt.Println(title)
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{strings.ToUpper(cloudName) + " EXPR", strings.ToUpper(cloudName) + " VAL", "TERRAFORM VAL", "TERRAFORM EXPR"})
		table.SetBorder(true)
		return table
	}

	for k, tfAttrs := range tfAttList {
		cloudAttrs, ok := cloudAttList[k]
		if !ok {
			continue // Resource exists only in TF. This is already handled by the "Missing" resource/check
		}
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
					tfQueryItems[i],
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
