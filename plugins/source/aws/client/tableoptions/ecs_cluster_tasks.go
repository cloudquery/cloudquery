package tableoptions

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/cloudquery/plugin-sdk/v3/caser"
)

type ECSTaskAPIs struct {
	ListTasksOpts []CustomListTasksOpts `json:"list_tasks,omitempty"`
}

type CustomListTasksOpts struct {
	ecs.ListTasksInput
}

// UnmarshalJSON implements the json.Unmarshaler interface for the CustomGetFindingsOpts type.
// It is the same as default, but allows the use of underscore in the JSON field names.
func (s *CustomListTasksOpts) UnmarshalJSON(data []byte) error {
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

func (s *ECSTaskAPIs) validateListTasks() error {
	for _, opt := range s.ListTasksOpts {
		if aws.ToString(opt.NextToken) != "" {
			return errors.New("invalid input: cannot set NextToken in ListTasks")
		}

		if aws.ToString(opt.Cluster) != "" {
			return errors.New("invalid input: cannot set Cluster in ListTasks")
		}

		// As per https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ListTasks.html
		if aws.ToInt32(opt.MaxResults) < 0 || aws.ToInt32(opt.MaxResults) > 100 {
			return errors.New("invalid range: MaxResults must be within range [1-100]")
		}
	}
	return nil
}

func (s *ECSTaskAPIs) setDefaults() {
	for i := 0; i < len(s.ListTasksOpts); i++ {
		if aws.ToInt32(s.ListTasksOpts[i].MaxResults) == 0 {
			s.ListTasksOpts[i].MaxResults = aws.Int32(100)
		}
	}
}

func (s *ECSTaskAPIs) Validate() error {
	s.setDefaults()
	return s.validateListTasks()
}
