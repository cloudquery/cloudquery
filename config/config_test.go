package config

import "testing"

const test1 =
`
{"providers": 
	[
		{
			"name": "aws",
			"regions": ["us-east-1"],
			"resources": 
				[
					{
					  "name": "ec2.images"
					},
					{
					  "name": "ec2.instances"
					},
					{
					  "name": "s3.buckets"
					}
				]
		}
	]
}
`
const test2 =
`
{"providers": 
	[
		{
			"name": true,
			"regions": ["us-east-1"],
			"resources": 
				[
					{
					  "name": "ec2.images"
					},
					{
					  "name": "ec2.instances"
					},
					{
					  "name": "s3.buckets"
					}
				]
		}
	]
}
`

func TestConfig_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "basicUnmarshal",
			args: args{[]byte(test1)},
			wantErr: false,
		},
		{
			name: "basicUnmarshalFail",
			args: args{[]byte(test2)},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{}
			var err error
			if err = c.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				// probably a better way to test for proper parsing but this will do for now
				if len(c.Providers) != 1 {
					t.Errorf("Did not parse any providers")
				}
				if c.Providers[0].Rest == nil {
					t.Errorf("Did not parse Rest")
				}
				if _, ok := c.Providers[0].Rest["resources"]; !ok {
					t.Errorf("resources key is not in Rest")
				}
			}
		})
	}
}
