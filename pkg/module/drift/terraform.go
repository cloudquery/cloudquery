package drift

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/cloudquery/cloudquery/pkg/module/drift/terraform"
	"github.com/cloudquery/cq-provider-sdk/provider/execution"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/go-hclog"
	"github.com/olekukonko/tablewriter"
	"github.com/tidwall/gjson"
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

type Attribute struct {
	ID        string           // Identifier from config
	SQL       string           // SQL expression
	Type      schema.ValueType // Type in DB as reported by provider
	TFName    string           // TF attribute name
	Unordered bool             // True if unordered slice

	CloudMod func(interface{}) interface{} // Modifier function after fetching the SQL
}

type AttrList []Attribute

func (a AttrList) SQLs() []string {
	ret := make([]string, len(a))
	for i := range a {
		ret[i] = a[i].SQL
	}
	return ret
}

func (a AttrList) TypeOf(id string) schema.ValueType {
	for i := range a {
		if a[i].ID == id {
			return a[i].Type
		}
	}
	return schema.TypeInvalid
}

const tfIDAttribute = "id"

// AsResourceList returns a map of resource ID vs. attributes
func (r TFInstances) AsResourceList(identifiers []string, alist AttrList, path string) ResourceList {
	if len(identifiers) == 0 {
		identifiers = []string{tfIDAttribute}
	}

	ret := make([]*Resource, 0, len(r))
	for i := range r {
		ret = append(ret, parseTerraformInstance(r[i], identifiers, alist, path)...)
	}
	return ret
}

func parseTerraformInstance(ins terraform.Instance, identifiers []string, alist AttrList, path string) ResourceList {
	registerGJsonHelpers()

	var elems []gjson.Result

	root := gjson.ParseBytes(ins.AttributesRaw)

	if path != "" {
		arr := gjson.GetBytes(ins.AttributesRaw, path)
		if !arr.IsArray() {
			panic("invalid path " + path + ": not an array")
		}
		elems = arr.Array()
	} else {
		elems = append(elems, root)
	}

	ret := make([]*Resource, 0, len(elems))

	for _, el := range elems {
		if !el.IsObject() {
			panic("invalid array element: not an object: " + el.Type.String())
		}

		getAttributes := func(id string) (interface{}, bool) {
			const rootPathPrefix = "root."

			var v gjson.Result
			if strings.HasPrefix(id, rootPathPrefix) {
				v = root.Get(strings.TrimPrefix(id, rootPathPrefix))
			} else {
				v = el.Get(id)
			}

			return v.Value(), v.Exists()
		}

		idVals := make([][]string, len(identifiers)) // identifiers -> [list of possible values]
		numResources := 1
		for i, idName := range identifiers {
			v, _ := getAttributes(idName)
			var (
				vv []interface{}
				ok bool
			)
			if vv, ok = v.([]interface{}); !ok {
				vv = []interface{}{v}
			}
			for j := range vv {
				v := parseTerraformAttribute(vv[j], alist.TypeOf(idName))
				idVals[i] = append(idVals[i], efaceToString(v))
			}
			numResources *= len(vv)
		}

		attributes := make([]interface{}, len(alist))
		for i := range alist {
			if val, ok := getAttributes(alist[i].TFName); ok {
				attributes[i] = parseTerraformAttribute(val, alist[i].Type)
			}
		}

		combinations := make([][]string, 0, numResources)
		for i := range idVals { // iterate each identifier, pick one from each element (which is the array of possible values) and repeat
			combinations = MatrixProduct(combinations, idVals[i])
		}

		for _, idVals := range combinations {
			res := &Resource{
				ID:         strings.Join(idVals, idSeparator),
				Attributes: attributes,
			}
			ret = append(ret, res)
		}

	}

	return ret
}

func parseTerraformAttribute(val interface{}, t schema.ValueType) interface{} {
	if val == nil {
		return nil
	}

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

func driftTerraform(ctx context.Context, logger hclog.Logger, conn execution.QueryExecer, cloudName string, cloudTable *traversedTable, resName string, resources map[string]*ResourceConfig, iacData *IACConfig, states TFStates, runParams RunParams, accountIDs []string) (*Result, error) {
	registerGJsonHelpers()

	res := &Result{
		Different: nil,
		Equal:     nil,
		Missing:   nil,
		Extra:     nil,
	}

	resData := resources[resName]
	deepMode := runParams.ForceDeep || (resData.Deep != nil && *resData.Deep)

	alist := make(AttrList, len(resData.Attributes))

	setMap := make(map[string]struct{}, len(resData.Sets))
	for i := range resData.Sets {
		setMap[resData.Sets[i]] = struct{}{}
	}

	var tagExp exp.LiteralExpression

	for i, a := range resData.Attributes {
		tfName, cloudMod := getTFNameFromMap(a, iacData.attributeMap)

		alist[i] = Attribute{
			ID:       resData.Attributes[i],
			Type:     cloudTable.Column(a).Type,
			TFName:   tfName,
			CloudMod: cloudMod,
		}

		switch alist[i].Type {
		case schema.TypeString:
			alist[i].SQL = fmt.Sprintf(`COALESCE("c"."%s",'')`, resData.Attributes[i])
		case schema.TypeTimestamp:
			alist[i].SQL = fmt.Sprintf(`EXTRACT(EPOCH FROM DATE_TRUNC('second', "c"."%s"))::VARCHAR`, resData.Attributes[i])
		default:
			alist[i].SQL = fmt.Sprintf(`"c"."%s"`, resData.Attributes[i])
		}

		_, alist[i].Unordered = setMap[alist[i].ID]

		if alist[i].ID == "tags" && alist[i].Type == schema.TypeJSON {
			tagExp = goqu.L(alist[i].SQL)
		}
	}

	if tagExp == nil {
		tagExp = goqu.L("NULL")

		if resData.acl.HasTagFilters() {
			logger.Warn("tag based filtering not possible on this resource type", "resource", resName)
		}
	}

	tfResources := states.FindType(iacData.Type, terraform.ModeManaged).AsResourceList(iacData.Identifiers, alist, iacData.Path)

	var cloudAttrQuery exp.LiteralExpression

	if !deepMode || len(alist) == 0 {
		cloudAttrQuery = goqu.L("NULL")
	} else {
		cloudAttrQuery = goqu.L("JSON_BUILD_ARRAY(" + strings.Join(alist.SQLs(), ",") + ")")
	}

	idExp, err := handleIdentifiers(resData.Identifiers)
	if err != nil {
		return nil, err
	}

	q := goqu.Dialect("postgres").From(goqu.T(cloudTable.Name).As("c")).Select(idExp, cloudAttrQuery.As("attlist"), tagExp.As("tags"))
	q = handleSubresource(logger, q, cloudTable, resources, accountIDs)
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
	}, resData.acl.ShouldSkip)

	// Get extra resources
	{
		q := goqu.Dialect("postgres").From(goqu.T(cloudTable.Name).As("c")).Select(idExp, cloudAttrQuery.As("attlist"), tagExp.As("tags"))
		q = handleSubresource(logger, q, cloudTable, resources, accountIDs)
		q = handleFilters(q, resources[resName]) // This line (the application of filters) is the difference from "existing"
		existingFiltered, err := queryIntoResourceList(ctx, logger, conn, q)
		if err != nil {
			return nil, err
		}

		existingFiltered.Walk(func(r *Resource) {
			if _, ok := tfMap[r.ID]; !ok {
				res.Extra = append(res.Extra, r)
			}
		}, resData.acl.ShouldSkip)
	}

	if !deepMode {
		// Get equal resources
		existing.Walk(func(r *Resource) {
			if _, ok := tfMap[r.ID]; ok {
				res.Equal = append(res.Equal, r)
			}
		}, resData.acl.ShouldSkip)
	} else {
		// Get deepequal and different resources
		existing.Walk(func(r *Resource) {
			tfAttr, ok := tfMap[r.ID]
			if !ok {
				return
			}
			if EqualAttributes(r.Attributes, tfAttr, alist) {
				res.DeepEqual = append(res.DeepEqual, r)
			} else {
				res.Different = append(res.Different, r)
			}
		}, resData.acl.ShouldSkip)
	}
	if deepMode && runParams.Debug && len(res.Different) > 0 {
		if err := RenderDriftTable(resName, resources, cloudName, alist, res.Different, tfResources); err != nil {
			return nil, err
		}
	}

	return res, nil
}

func RenderDriftTable(resName string, resources map[string]*ResourceConfig, cloudName string, alist AttrList, differentIDs, tfRes ResourceList) error {
	resData := resources[resName]

	makeTable := func(title string) *tablewriter.Table {
		fmt.Println(title)
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{strings.ToUpper(cloudName) + " EXPR", strings.ToUpper(cloudName) + " VAL", "TERRAFORM VAL", "TERRAFORM EXPR"})
		table.SetBorder(true)
		return table
	}

	// If there are any cloud modifiers in the columns, process them before comparing/printing
	for i := range differentIDs {
		for j := range differentIDs[i].Attributes {
			if alist[j].CloudMod != nil {
				differentIDs[i].Attributes[j] = alist[j].CloudMod(differentIDs[i].Attributes[j])
			}
		}
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
			if !EqualAttributes([]interface{}{cloudAttrs[i]}, []interface{}{tfAttrs[i]}, []Attribute{alist[i]}) {
				cStr := efaceToString(cloudAttrs[i])
				tStr := efaceToString(tfAttrs[i])
				if cStr == tStr {
					cStr += fmt.Sprintf(" %T", cloudAttrs[i])
					tStr += fmt.Sprintf(" %T", tfAttrs[i])
				}
				table.Append([]string{
					alist[i].SQL,
					cStr,
					tStr,
					alist[i].TFName,
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

// EqualAttributes compares a and b slice values according to given attribute list. For a, alist.CloudMod is called (if non-nil) before proceeding with the comparison.
func EqualAttributes(a []interface{}, b []interface{}, alist AttrList) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		aval := a[i]
		if alist[i].CloudMod != nil {
			aval = alist[i].CloudMod(aval)
		}

		if alist[i].Unordered {
			aSlc, ok1 := aval.([]interface{})
			bSlc, ok2 := b[i].([]interface{})
			if !ok1 || !ok2 {
				// not slices
				return false
			}
			if !EqualSets(aSlc, bSlc) {
				return false
			}
			continue
		}

		if !equals(aval, b[i]) {
			return false
		}
	}

	return true
}

func EqualSets(a []interface{}, b []interface{}) bool {
	less := func(a, b interface{}) bool { return efaceToString(a) < efaceToString(b) }
	return cmp.Equal(a, b, cmpopts.SortSlices(less))
}

func efaceToString(a interface{}) string {
	return fmt.Sprintf("%v", a)
}

func equals(a, b interface{}) bool {
	// case of assume_role_policy fields: string and map[string]interface{} needs to be compared
	if _, ok := b.(map[string]interface{}); ok {
		if aMap, ok := isStringAMap(a); ok {
			a = aMap
		}
	} else if _, ok := a.(map[string]interface{}); ok {
		if bMap, ok := isStringAMap(b); ok {
			b = bMap
		}
	}

	// nil and empty strings, slices or maps are considered equal for our comparisons
	if isEmptyStringSliceOrMap(a) && isEmptyStringSliceOrMap(b) {
		return true
	}

	as := efaceToString(a)
	bs := efaceToString(b)
	if as == bs {
		return true
	}

	if strings.HasPrefix(as, "arn:aws:") && strings.HasPrefix(bs, "arn:aws:") {
		// compare ARNs, ignoring empty account IDs or regions
		aa, err := arn.Parse(as)
		if err != nil {
			return false
		}
		ba, err := arn.Parse(bs)
		if err != nil {
			return false
		}
		if aa.Region == "" || ba.Region == "" {
			aa.Region, ba.Region = "", ""
		}
		if aa.AccountID == "" || ba.AccountID == "" {
			aa.AccountID, ba.AccountID = "", ""
		}
		return aa.String() == ba.String()
	}

	return cmp.Equal(a, b, cmpopts.EquateEmpty())
}

func isStringAMap(a interface{}) (map[string]interface{}, bool) {
	if astr, ok := a.(string); ok {
		var newA map[string]interface{}
		if json.Unmarshal([]byte(astr), &newA) == nil {
			return newA, true
		}
	}
	return nil, false
}

func isEmptyStringSliceOrMap(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		return v.Len() == 0
	default:
		return false
	}
}

func registerGJsonHelpers() {
	if !gjson.ModifierExists("inverse", nil) {
		// inverse a boolean (null input returns true)
		gjson.AddModifier("inverse", func(body, arg string) string {
			if body == "false" || body == "" {
				return "true"
			}
			return "false"
		})
	}
	if !gjson.ModifierExists("iftrue", nil) {
		// if given statement is true, return the arg. otherwise return nil.
		gjson.AddModifier("iftrue", func(body, arg string) string {
			b, err := strconv.ParseBool(body)
			if err != nil {
				uq, _ := strconv.Unquote(body)
				b, _ = strconv.ParseBool(uq)
			}
			if b {
				return strconv.Quote(arg)
			}
			return ""
		})
	}
	if !gjson.ModifierExists("getbool", nil) {
		// extract the given arg as key from the given object and expect it to be a boolean. returns false if doesn't exist. returns nil if not an object.
		gjson.AddModifier("getbool", func(body, arg string) string {
			if body == "" { // nil input
				return "false"
			}
			var v map[string]interface{}
			if err := json.Unmarshal([]byte(body), &v); err != nil {
				return "" // invalid input
			}
			b, ok := v[arg]
			if !ok {
				return "false" // key not in map
			}

			var bb bool
			if bb, ok = b.(bool); !ok {
				bb, _ = strconv.ParseBool(fmt.Sprintf("%v", b))
			}
			return strconv.FormatBool(bb)
		})
	}
	if !gjson.ModifierExists("if", nil) {
		// if given statement equals something, return the arg. otherwise return the statement.
		gjson.AddModifier("if", func(body, arg string) string {
			argParts := strings.SplitN(arg, ",", 2)
			if body != strconv.Quote(argParts[0]) {
				return body
			}
			return strconv.Quote(argParts[1])
		})
	}
	if !gjson.ModifierExists("split", nil) {
		// split given string into an array using the given separator
		gjson.AddModifier("split", func(body, arg string) string {
			var v string
			if err := json.Unmarshal([]byte(body), &v); err != nil {
				return "" // invalid input
			}
			vSplit := strings.Split(v, arg)
			b, _ := json.Marshal(vSplit)
			return string(b)
		})
	}
}

// getTFNameFromMap returns the terraform attribute name for the given provider attr, by looking up the attributeMap. If not found, returns the given attr itself.
func getTFNameFromMap(attr string, attributeMap map[string]string) (string, func(interface{}) interface{}) {
	if mapped := attributeMap[attr]; mapped != "" {
		return mapped, nil
	}
	for k, tfAttr := range attributeMap {
		if strings.HasPrefix(k, attr+"|") {
			return tfAttr, func(input interface{}) interface{} {
				j, _ := json.Marshal(map[string]interface{}{
					attr: input,
				})
				gj := gjson.ParseBytes(j)
				res := gj.Get(k)
				return res.Value()
			}
		}
	}

	return attr, nil
}
