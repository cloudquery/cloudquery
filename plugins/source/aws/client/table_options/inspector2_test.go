package table_options

import (
	"reflect"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client/table_option_inputs/inspector2_input"
	"github.com/stretchr/testify/assert"

	"github.com/cloudquery/plugin-sdk/faker"
)

func findNilOrDefaultFields(v reflect.Value, nilOrDefFields []string) []string {
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := v.Type().Field(i)

		switch fieldValue.Kind() {
		case reflect.Ptr:
			if fieldValue.IsNil() {
				nilOrDefFields = append(nilOrDefFields, fieldType.Name)
			} else {
				if fieldValue.Elem().Kind() == reflect.Struct {
					if fieldValue.Elem().Type() == reflect.TypeOf(time.Time{}) {
						if fieldValue.Elem().Interface().(time.Time).IsZero() {
							nilOrDefFields = append(nilOrDefFields, fieldType.Name)
						}
					} else {
						nilOrDefFields = findNilOrDefaultFields(fieldValue, nilOrDefFields)
					}
				} else {
					zeroValue := reflect.Zero(fieldType.Type).Interface()
					if reflect.DeepEqual(fieldValue.Interface(), zeroValue) {
						nilOrDefFields = append(nilOrDefFields, fieldType.Name)
					}
				}
			}
		case reflect.Struct:
			if fieldType.Type == reflect.TypeOf(time.Time{}) {
				if fieldValue.Interface().(time.Time).IsZero() {
					nilOrDefFields = append(nilOrDefFields, fieldType.Name)
				}
			} else {
				nilOrDefFields = findNilOrDefaultFields(fieldValue, nilOrDefFields)
			}

		default:
			zeroValue := reflect.Zero(fieldType.Type).Interface()
			if reflect.DeepEqual(fieldValue.Interface(), zeroValue) {
				nilOrDefFields = append(nilOrDefFields, fieldType.Name)
			}
		}
	}

	return nilOrDefFields
}

func TestInspector2ListFindings(t *testing.T) {
	u := inspector2_input.ListFindingsInput{}
	if err := faker.FakeObject(&u); err != nil {
		t.Fatal(err)
	}

	api := Inspector2APIs{
		ListFindingOpts: u,
	}
	// Ensure that the validation works as expected
	_, err := api.ListFindings()
	assert.EqualError(t, err, "invalid input: cannot set NextToken in ListFindings")

	// Ensure that as soon as the validation passes that there are no unexpected empty or nil fields
	api.ListFindingOpts.NextToken = nil
	input, err := api.ListFindings()
	nilFields := findNilOrDefaultFields(reflect.ValueOf(*input), []string{})

	assert.Equal(t, nilFields, []string{"NextToken"}, "These are the only fields that should have a default value")
	assert.Nil(t, err)
}
