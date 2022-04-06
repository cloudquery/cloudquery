package drift

import (
	"bytes"
	"encoding/gob"
	"sort"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
)

func TestApplyWildProvider(t *testing.T) {
	t.Parallel()

	table := []struct {
		Name     string
		Wild     *ProviderConfig
		Dst      *ProviderConfig
		Expected *ProviderConfig
	}{
		{
			Name: "Apply to empty",
			Wild: &ProviderConfig{
				IgnoreResources: ResourceSelectors{
					{
						ID: aws.String("foo"),
					},
				},
				CheckResources: ResourceSelectors{
					{
						ID: aws.String("bar"),
					},
				},
				AccountIDs: []string{"baz"},
			},
			Dst: &ProviderConfig{},
			Expected: &ProviderConfig{
				IgnoreResources: ResourceSelectors{
					{
						ID: aws.String("foo"),
					},
				},
				CheckResources: ResourceSelectors{
					{
						ID: aws.String("bar"),
					},
				},
				AccountIDs: []string{"baz"},
			},
		},
		{
			Name: "Apply to partially empty",
			Wild: &ProviderConfig{
				IgnoreResources: ResourceSelectors{
					{
						ID: aws.String("foo"),
					},
				},
				CheckResources: ResourceSelectors{
					{
						ID: aws.String("bar"),
					},
				},
				AccountIDs: []string{"baz"},
			},
			Dst: &ProviderConfig{
				CheckResources: ResourceSelectors{
					{
						ID: aws.String("qux"),
					},
				},
				AccountIDs: []string{"quux"},
			},
			Expected: &ProviderConfig{
				IgnoreResources: ResourceSelectors{
					{
						ID: aws.String("foo"),
					},
				},
				CheckResources: ResourceSelectors{
					{
						ID: aws.String("qux"),
					},
				},
				AccountIDs: []string{"quux"},
			},
		},
		{
			Name: "Apply to not empty",
			Wild: &ProviderConfig{
				IgnoreResources: ResourceSelectors{
					{
						ID: aws.String("foo"),
					},
				},
				CheckResources: ResourceSelectors{
					{
						ID: aws.String("bar"),
					},
				},
				AccountIDs: []string{"baz"},
			},
			Dst: &ProviderConfig{
				IgnoreResources: ResourceSelectors{
					{
						ID: aws.String("qux"),
					},
				},
				CheckResources: ResourceSelectors{
					{
						ID: aws.String("quux"),
					},
				},
				AccountIDs: []string{"quuz"},
			},
			Expected: &ProviderConfig{
				IgnoreResources: ResourceSelectors{
					{
						ID: aws.String("qux"),
					},
				},
				CheckResources: ResourceSelectors{
					{
						ID: aws.String("quux"),
					},
				},
				AccountIDs: []string{"quuz"},
			},
		},
	}

	for i := range table {
		t.Run(table[i].Name, func(t *testing.T) {
			table[i].Dst.applyWildProvider(table[i].Wild)
			assert.EqualValues(t, table[i].Expected, table[i].Dst)
		})
	}
}

func TestApplyWildResource(t *testing.T) {
	t.Parallel()

	table := []struct {
		Name     string
		Wild     *ResourceConfig
		Dst      *ResourceConfig
		Expected *ResourceConfig
	}{
		{
			Name: "Apply to empty",
			Wild: &ResourceConfig{
				Identifiers:       []string{"id_a", "id_b", "id_c"},
				IgnoreIdentifiers: []string{"ig_a", "ig_b", "ig_c"},
				Attributes:        []string{"at_a", "at_b", "at_c"},
				IgnoreAttributes:  []string{"igat_a", "igat_b", "igat_c"},
				Deep:              aws.Bool(true),
				Filters:           []string{"f_a", "f_b"},
				Sets:              []string{"at_a"},
				IAC: map[iacProvider]*IACConfig{
					iacTerraform: {
						Type:         "test",
						Path:         "path",
						Identifiers:  []string{"iac_a"},
						AttributeMap: []string{"at_b=iac_b"},
					},
				},
			},
			Dst: &ResourceConfig{},
			Expected: &ResourceConfig{
				Identifiers:       []string{"id_a", "id_b", "id_c"},
				IgnoreIdentifiers: []string{"ig_a", "ig_b", "ig_c"},
				Attributes:        []string{"at_a", "at_b", "at_c"},
				IgnoreAttributes:  []string{"igat_a", "igat_b", "igat_c"},
				Deep:              aws.Bool(true),
				Filters:           []string{"f_a", "f_b"},
				Sets:              []string{"at_a"},
				IAC: map[iacProvider]*IACConfig{
					iacTerraform: {
						Type:         "test",
						Path:         "path",
						Identifiers:  []string{"iac_a"},
						AttributeMap: []string{"at_b=iac_b"},
					},
				},
			},
		},
		{
			Name: "Apply to partially empty",
			Wild: &ResourceConfig{
				Identifiers:       []string{"id_a", "id_b", "id_c"},
				IgnoreIdentifiers: []string{"ig_a", "ig_b", "ig_c"},
				Attributes:        []string{"at_a", "at_b", "at_c"},
				IgnoreAttributes:  []string{"igat_a", "igat_b", "igat_c"},
				Deep:              aws.Bool(true),
				Filters:           []string{"f_a", "f_b"},
				Sets:              []string{"at_a"},
				IAC: map[iacProvider]*IACConfig{
					iacTerraform: {
						Type:         "test",
						Path:         "path",
						Identifiers:  []string{"iac_a"},
						AttributeMap: []string{"at_b=iac_b"},
					},
				},
			},
			Dst: &ResourceConfig{
				Identifiers: []string{"did_a"},
				Sets:        []string{"dat_a"},
			},
			Expected: &ResourceConfig{
				Identifiers:       []string{"did_a"},
				IgnoreIdentifiers: []string{"ig_a", "ig_b", "ig_c"},
				Attributes:        []string{"at_a", "at_b", "at_c"},
				IgnoreAttributes:  []string{"igat_a", "igat_b", "igat_c"},
				Deep:              aws.Bool(true),
				Filters:           []string{"f_a", "f_b"},
				Sets:              []string{"at_a", "dat_a"},
				IAC: map[iacProvider]*IACConfig{
					iacTerraform: {
						Type:         "test",
						Path:         "path",
						Identifiers:  []string{"iac_a"},
						AttributeMap: []string{"at_b=iac_b"},
					},
				},
			},
		},
		{
			Name: "Apply to not empty",
			Wild: &ResourceConfig{
				Identifiers:       []string{"id_a", "id_b", "id_c"},
				IgnoreIdentifiers: []string{"ig_a", "ig_b", "ig_c"},
				Attributes:        []string{"at_a", "at_b", "at_c"},
				IgnoreAttributes:  []string{"igat_a", "igat_b", "igat_c"},
				Deep:              aws.Bool(true),
				Filters:           []string{"f_a", "f_b"},
				Sets:              []string{"at_a"},
				IAC: map[iacProvider]*IACConfig{
					iacTerraform: {
						Type:         "test",
						Path:         "path",
						Identifiers:  []string{"iac_a"},
						AttributeMap: []string{"at_b=iac_b"},
					},
				},
			},
			Dst: &ResourceConfig{
				Identifiers:       []string{"did_a"},
				IgnoreIdentifiers: []string{"dig_a", "dig_b", "dig_c"},
				Attributes:        []string{"dat_a", "dat_b", "dat_c"},
				IgnoreAttributes:  []string{"digat_a"},
				Deep:              aws.Bool(false),
				Filters:           []string{"df_a", "f_a"},
				Sets:              []string{"dat_a"},
				IAC: map[iacProvider]*IACConfig{
					iacTerraform: {
						Type:         "test",
						Path:         "path",
						Identifiers:  []string{"iac_a"},
						AttributeMap: []string{"at_b=iac_b"},
					},
				},
			},
			Expected: &ResourceConfig{
				Identifiers:       []string{"did_a"},
				IgnoreIdentifiers: []string{"dig_a", "dig_b", "dig_c", "ig_a", "ig_b", "ig_c"},
				Attributes:        []string{"dat_a", "dat_b", "dat_c"},
				IgnoreAttributes:  []string{"digat_a", "igat_a", "igat_b", "igat_c"},
				Deep:              aws.Bool(false),
				Filters:           []string{"df_a", "f_a", "f_b"},
				Sets:              []string{"at_a", "dat_a"},
				IAC: map[iacProvider]*IACConfig{
					iacTerraform: {
						Type:         "test",
						Path:         "path",
						Identifiers:  []string{"iac_a"},
						AttributeMap: []string{"at_b=iac_b"},
					},
				},
			},
		},
	}

	for i := range table {
		t.Run(table[i].Name, func(t *testing.T) {
			table[i].Dst.applyWildResource(table[i].Wild)
			assert.EqualValues(t, table[i].Expected, table[i].Dst)
		})
	}
}

func TestInterpolatedResourceMap(t *testing.T) {
	t.Parallel()

	prov := &ProviderConfig{
		IgnoreResources: ResourceSelectors{
			{
				ID: aws.String("foo"),
			},
		},
		CheckResources: ResourceSelectors{
			{
				ID: aws.String("bar"),
			},
		},
		AccountIDs: []string{"baz"},
		Resources: map[string]*ResourceConfig{
			"test": {
				Identifiers:       []string{"id1", "id2"},
				IgnoreIdentifiers: []string{"id2"},
				Attributes:        []string{"a1", "a2"},
				IgnoreAttributes:  []string{"a1"},
				IAC: map[iacProvider]*IACConfig{
					iacTerraform: {
						Type: "tf_test",
					},
				},
			},
			"test2": {
				Identifiers: []string{"id3"},
				IAC: map[iacProvider]*IACConfig{
					iacCloudformation: {
						Type: "cf_test",
					},
				},
			},
		},
	}

	ret := prov.interpolatedResourceMap(iacTerraform, hclog.NewNullLogger())

	assert.EqualValues(t, map[string]*ResourceConfig{
		"test": {
			Identifiers:       []string{"id1"},
			IgnoreIdentifiers: []string{"id2"},
			Attributes:        []string{"a2"},
			IgnoreAttributes:  []string{"a1"},
			IAC: map[iacProvider]*IACConfig{
				iacTerraform: {
					Type: "tf_test",
				},
			},
		},
	}, ret)
}

func copyMap(in, out interface{}) {
	buf := new(bytes.Buffer)
	_ = gob.NewEncoder(buf).Encode(in)
	_ = gob.NewDecoder(buf).Decode(out)
}

func TestApplyProvider(t *testing.T) {
	prov := ProviderConfig{
		Name:    "aws",
		Version: ">=1.5.0",
		Resources: map[string]*ResourceConfig{
			"test1": {
				IAC: map[iacProvider]*IACConfig{
					iacTerraform: {
						Type: "tf_test1",
					},
				},
			},
			"test2": {
				IAC: map[iacProvider]*IACConfig{
					iacTerraform: {
						Type: "tf_test2",
					},
				},
			},
		},
	}
	var err error
	prov.versionConstraints, err = version.NewConstraint(prov.Version)
	assert.NoError(t, err)

	modify := func(p ProviderConfig, fn func(*ProviderConfig)) ProviderConfig {
		var resCopy map[string]*ResourceConfig
		copyMap(p.Resources, &resCopy)
		p.Resources = resCopy
		fn(&p)
		return p
	}

	d := &Drift{
		logger: hclog.NewNullLogger(),
	}

	table := []struct {
		name              string
		cfg               ProviderConfig
		schema            *cqproto.GetProviderSchemaResponse
		expectedResult    bool
		expectedError     bool
		checkResourceList bool
		expectedResources []string
	}{
		{
			name: "Name mismatch",
			cfg: modify(prov, func(p *ProviderConfig) {
				p.Name = "not_aws"
			}),
			schema: &cqproto.GetProviderSchemaResponse{
				Name:    "aws",
				Version: "1.6.0",
			},
			expectedResult: false,
			expectedError:  false,
		},
		{
			name: "Old version",
			cfg:  prov,
			schema: &cqproto.GetProviderSchemaResponse{
				Name:    "aws",
				Version: "1.0.0",
			},
			expectedResult: false,
			expectedError:  false,
		},
		{
			name: "New version",
			cfg:  prov,
			schema: &cqproto.GetProviderSchemaResponse{
				Name:    "aws",
				Version: "1.6.0",
				ResourceTables: map[string]*schema.Table{
					"test1": {
						Name: "aws_test1",
					},
					"test2": {
						Name: "aws_test2",
					},
				},
			},
			expectedResult: true,
			expectedError:  false,
		},
		{
			name: "Ignore some resources",
			cfg: modify(prov, func(p *ProviderConfig) {
				p.IgnoreResources = ResourceSelectors{
					{
						Type: "test2",
						ID:   aws.String("*"),
					},
				}
			}),
			schema: &cqproto.GetProviderSchemaResponse{
				Name:    "aws",
				Version: "1.6.0",
				ResourceTables: map[string]*schema.Table{
					"test1": {
						Name: "aws_test1",
					},
					"test2": {
						Name: "aws_test2",
					},
				},
			},
			expectedResult:    true,
			expectedError:     false,
			checkResourceList: true,
			expectedResources: []string{"test1"},
		},
		{
			name: "Allow some resources",
			cfg: modify(prov, func(p *ProviderConfig) {
				p.CheckResources = ResourceSelectors{
					{
						Type: "test2",
						ID:   aws.String("*"),
					},
				}
			}),
			schema: &cqproto.GetProviderSchemaResponse{
				Name:    "aws",
				Version: "1.6.0",
				ResourceTables: map[string]*schema.Table{
					"test1": {
						Name: "aws_test1",
					},
					"test2": {
						Name: "aws_test2",
					},
				},
			},
			expectedResult:    true,
			expectedError:     false,
			checkResourceList: true,
			expectedResources: []string{"test2"},
		},
		{
			name: "Allow some resources with hash",
			cfg: modify(prov, func(p *ProviderConfig) {
				p.Resources["test3#hash1"] = &ResourceConfig{
					IAC: map[iacProvider]*IACConfig{
						iacTerraform: {
							Type: "tf_test3a",
						},
					},
				}
				p.Resources["test3#hash2"] = &ResourceConfig{
					IAC: map[iacProvider]*IACConfig{
						iacTerraform: {
							Type: "tf_test3b",
						},
					},
				}
				p.CheckResources = ResourceSelectors{
					{
						Type: "test3",
						ID:   aws.String("*"),
					},
				}
			}),
			schema: &cqproto.GetProviderSchemaResponse{
				Name:    "aws",
				Version: "1.6.0",
				ResourceTables: map[string]*schema.Table{
					"test1": {
						Name: "aws_test1",
					},
					"test3": {
						Name: "aws_test3",
					},
				},
			},
			expectedResult:    true,
			expectedError:     false,
			checkResourceList: true,
			expectedResources: []string{"test3#hash1", "test3#hash2"},
		},
		{
			name: "Tag filter (all resources)",
			cfg: modify(prov, func(p *ProviderConfig) {
				tags := map[string]string{"key": "value"}
				p.CheckResources = ResourceSelectors{
					{
						Type: "*",
						Tags: &tags,
					},
				}
			}),
			schema: &cqproto.GetProviderSchemaResponse{
				Name:    "aws",
				Version: "1.6.0",
				ResourceTables: map[string]*schema.Table{
					"test1": {
						Name: "aws_test1",
					},
					"test2": {
						Name: "aws_test2",
					},
				},
			},
			expectedResult:    true,
			expectedError:     false,
			checkResourceList: true,
			expectedResources: []string{"test1", "test2"}, // Tag filters should not filter out any resources at this point
		},
	}

	for i := range table {
		t.Run(table[i].name, func(t *testing.T) {
			p := table[i].cfg

			d.tableMap = nil // reinit every time
			res, diags := d.applyProvider(&p, table[i].schema)

			assert.Equal(t, table[i].expectedError, diags.HasErrors(), "diags: %s", diags.Error())
			assert.Equal(t, table[i].expectedResult, res, "unexpected result")
			if table[i].checkResourceList {
				sort.Strings(table[i].expectedResources)
				assert.Equal(t, table[i].expectedResources, p.resourceKeys())
			}
		})
	}
}

func TestSubResourceLookup(t *testing.T) {
	prov := ProviderConfig{
		Name:    "aws",
		Version: ">=1.5.0",
		Resources: map[string]*ResourceConfig{
			"test1": {
				IAC: map[iacProvider]*IACConfig{
					iacTerraform: {
						Type: "tf_test1",
					},
				},
			},
			"test2": {
				IAC: map[iacProvider]*IACConfig{
					iacTerraform: {
						Type: "tf_test2",
					},
				},
			},
			"aws_test1_2": {
				Identifiers:      []string{"${" + string(placeholderResourceOptsPrimaryKeys) + "}"},
				Attributes:       []string{"${" + string(placeholderResourceColumnNames) + "}"},
				IgnoreAttributes: []string{"data3_ignored"},
				IAC: map[iacProvider]*IACConfig{
					iacTerraform: {
						Type: "tf_test1_2",
					},
				},
			},
			"aws_test1_3": {
				IAC: map[iacProvider]*IACConfig{
					iacTerraform: {
						Type: "tf_test1_3",
					},
				},
			},
		},
	}
	var err error
	prov.versionConstraints, err = version.NewConstraint(prov.Version)
	assert.NoError(t, err)

	d := &Drift{
		logger: hclog.NewNullLogger(),
	}

	sch := &cqproto.GetProviderSchemaResponse{
		Name:    "aws",
		Version: "1.6.0",
		ResourceTables: map[string]*schema.Table{
			"test1": {
				Name: "aws_test1",
				Relations: []*schema.Table{
					{
						Name:    "aws_test1_2",
						Options: schema.TableCreationOptions{PrimaryKeys: []string{"parent_cq_id", "data1"}},
						Columns: []schema.Column{
							schema.SetColumnMeta(schema.Column{
								Name: "parent_cq_id",
								Type: schema.TypeString,
							}, &schema.ColumnMeta{
								Resolver: &schema.ResolverMeta{
									Name:    "schema.ParentIdResolver",
									Builtin: true,
								},
							}),
							schema.SetColumnMeta(schema.Column{
								Name: "account_id",
								Type: schema.TypeString,
							}, &schema.ColumnMeta{
								Resolver: &schema.ResolverMeta{
									Name:    "github.com/cloudquery/cq-provider-aws/client.ResolveAWSAccount",
									Builtin: false,
								},
							}),
							{
								Name: "data1",
								Type: schema.TypeString,
							},
							{
								Name: "data2",
								Type: schema.TypeString,
							},
							{
								Name: "data3_ignored",
								Type: schema.TypeString,
							},
						},
						Relations: []*schema.Table{
							{
								Name: "aws_test1_3",
								Relations: []*schema.Table{
									{
										Name: "aws_test1_4", // This won't be processed because IAC config doesn't have an entry for it
									},
								},
							},
						},
					},
				},
			},
			"test2": {
				Name: "aws_test2",
			},
		},
	}

	res, diags := d.applyProvider(&prov, sch)
	assert.False(t, diags.HasErrors(), "diags: %s", diags.Error())
	assert.True(t, res, "unexpected result")

	assert.Equal(t, []string{"aws_test1_2", "aws_test1_3", "test1", "test2"}, prov.resourceKeys())

	ret := prov.interpolatedResourceMap(iacTerraform, d.logger)

	assert.Equal(t, []string{"data1"}, ret["aws_test1_2"].Identifiers)
	assert.Equal(t, []string{"data1", "data2"}, ret["aws_test1_2"].Attributes)

	tbl := d.lookupResource("aws_test1_2", sch)
	assert.NotNil(t, tbl)
	assert.Equal(t, []string{"data1", "data2", "data3_ignored"}, tbl.NonCQColumns())
	assert.Equal(t, []string{"data1"}, tbl.NonCQPrimaryKeys())
	assert.Equal(t, "account_id", tbl.AccountIDColumn())
}
