package tableoptions

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/invopop/jsonschema"
)

type ECSTasks struct {
	ListTasksOpts []CustomECSListTasksInput `json:"list_tasks,omitempty"`
}

type CustomECSListTasksInput struct {
	ecs.ListTasksInput
}

// UnmarshalJSON implements the json.Unmarshaler interface for the CustomECSListTasksInput type.
// It is the same as default, but allows the use of underscore in the JSON field names.
func (s *CustomECSListTasksInput) UnmarshalJSON(data []byte) error {
	m := map[string]any{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	csr := caser.New()
	changeCaseForObject(m, csr.ToPascal)
	b, _ := json.Marshal(m)
	return json.Unmarshal(b, &s.ListTasksInput)
}

func (CustomECSListTasksInput) JSONSchemaExtend(sc *jsonschema.Schema) {
	// The following properties are prohibited in spec
	sc.Properties.Delete("NextToken")
	sc.Properties.Delete("Cluster")
	// The following properties have additional constraints
	propertyMaxResults := sc.Properties.Value("MaxResults")
	if len(propertyMaxResults.OneOf) == 2 {
		propertyMaxResults = propertyMaxResults.OneOf[0] // 0 = value, 1 = null
	}
	propertyMaxResults.Minimum = json.Number("1")
	propertyMaxResults.Maximum = json.Number("100")
	propertyMaxResults.Default = 100
}

func (s *ECSTasks) validateListTasks() error {
	for _, opt := range s.ListTasksOpts {
		if opt.NextToken != nil {
			return errors.New("invalid input: cannot set NextToken in ListTasks")
		}
		if opt.Cluster != nil {
			return errors.New("invalid input: cannot set Cluster in ListTasks")
		}
		if aws.ToInt32(opt.MaxResults) < 1 || aws.ToInt32(opt.MaxResults) > 100 {
			return errors.New("invalid range: MaxResults must be within range [1-100]")
		}
	}
	return nil
}

func (s *ECSTasks) sanitized() *ECSTasks {
	var result ECSTasks
	if s != nil {
		result = *s
	}

	if len(result.ListTasksOpts) == 0 {
		result.ListTasksOpts = []CustomECSListTasksInput{{ListTasksInput: ecs.ListTasksInput{}}}
	}
	for i, opt := range result.ListTasksOpts {
		if aws.ToInt32(opt.MaxResults) == 0 {
			result.ListTasksOpts[i].MaxResults = aws.Int32(100)
		}
	}
	return &result
}

func (s *ECSTasks) Validate() error {
	return s.sanitized().validateListTasks()
}

func (s *ECSTasks) Filters() []CustomECSListTasksInput {
	if s != nil && s.ListTasksOpts != nil {
		return s.ListTasksOpts
	}
	return s.sanitized().ListTasksOpts
}
