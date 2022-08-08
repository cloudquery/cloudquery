package client

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	stsTypes "github.com/aws/aws-sdk-go-v2/service/sts/types"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

type mockAssumeRole func(ctx context.Context, params *sts.AssumeRoleInput, optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error)

// emptyInterfaceFieldNames looks at value s, which should be a struct (or a pointer to a struct),
// and returns the list of its field names which represent interface values but have nil value.
func emptyInterfaceFieldNames(s interface{}) []string {
	if s == nil {
		return nil
	}
	v := reflect.ValueOf(s)
	if v.Type().Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = reflect.Indirect(v)
	}
	if v.Type().Kind() != reflect.Struct {
		return nil
	}
	var empty []string
	for i := 0; i < v.Type().NumField(); i++ {
		field := v.Field(i)
		if t := field.Type(); t == nil || t.Kind() != reflect.Interface {
			continue
		}
		if field.IsNil() {
			empty = append(empty, v.Type().Field(i).Name)
		}
	}
	return empty
}

// emptyInterfaceFieldNames is a test helper but it is not trivial and uses reflection. So let's test it too.
func Test_emptyInterfaceFieldNames(t *testing.T) {
	// nested structs are ok here to simplify the test matrix
	// nolint:revive
	tests := []struct {
		s    interface{}
		want []string
	}{
		{nil, nil},
		{
			struct {
				x int
				y *string
			}{}, nil,
		},
		{
			struct {
				x interface{}
				y interface{}
			}{0, "test"}, nil,
		},
		{
			struct {
				x interface{}
				y interface{}
			}{},
			[]string{"x", "y"},
		},
		{
			struct {
				x interface{}
				y interface{}
			}{nil, 1},
			[]string{"x"},
		},
		{
			struct {
				x interface{}
				y interface{}
			}{1, nil},
			[]string{"y"},
		},
		{
			&struct { // test that pointer to a struct works too
				x interface{}
				y interface{}
			}{1, nil},
			[]string{"y"},
		},
	}
	for _, tt := range tests {
		got := emptyInterfaceFieldNames(tt.s)
		results := cmp.Diff(got, tt.want)
		if results != "" {
			t.Errorf(results)
		}
	}
}

func Test_initServices_NoNilValues(t *testing.T) {
	// the purpose of this test is to call initServices and check that returned Services struct
	// has no nil values in its fields.
	empty := emptyInterfaceFieldNames(initServices("us-east-1", aws.Config{}))
	for _, name := range empty {
		t.Errorf("initServices().%s == nil", name)
	}
}

func Test_obfuscateAccountId(t *testing.T) {
	assert.Equal(t, "1111xxxxxxxx", obfuscateAccountId("1111111111"))
	assert.Equal(t, "11", obfuscateAccountId("11"))
}

func Test_isValidRegions(t *testing.T) {
	tests := []struct {
		regions []string
		want    error
	}{
		{
			regions: []string{"us-east-1"},
			want:    nil,
		}, {
			regions: []string{"us-east-1", "*"},
			want:    errInvalidRegion,
		}, {
			regions: []string{"*"},
			want:    nil,
		}, {
			regions: []string{"*", "us-east-1"},
			want:    errInvalidRegion,
		},
		{
			regions: []string{"us-easta-1"},
			want:    errUnknownRegion("us-easta-1"),
		},
		{
			regions: []string{"*", "us-easta-1"},
			want:    errInvalidRegion,
		},
		{
			regions: []string{"us-easta-1", "*"},
			want:    errUnknownRegion("us-easta-1"),
		},
	}
	for i, tt := range tests {
		err := verifyRegions(tt.regions)
		if err != nil {
			assert.EqualErrorf(t, err, tt.want.Error(), "Case-%d: Error should be: %v, got: %v", i, tt.want.Error(), err)
		}
	}
}
func Test_isAllRegions(t *testing.T) {
	tests := []struct {
		regions []string
		want    bool
	}{
		{
			regions: []string{"us-east-1"},
			want:    false,
		}, {
			regions: []string{"us-east-1", "*"},
			want:    false,
		}, {
			regions: []string{"*"},
			want:    true,
		}, {
			regions: []string{"*", "us-east-1"},
			want:    false,
		},
	}
	for i, tt := range tests {
		err := isAllRegions(tt.regions)
		results := cmp.Diff(err, tt.want)
		if results != "" {
			t.Errorf("Case-%d failed: %s", i, results)
		}
	}
}

func (m mockAssumeRole) AssumeRole(ctx context.Context, params *sts.AssumeRoleInput, optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error) {
	return m(ctx, params, optFns...)
}

func Test_Configure(t *testing.T) {
	ctx := context.Background()
	logger := hclog.New(&hclog.LoggerOptions{})
	f, err := ioutil.TempFile("", "")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	defer os.Remove(f.Name())
	data := []byte(`[test]
	aws_access_key_id = <YOUR_TEMP_ACCESS_KEY_ID>
	aws_secret_access_key = <YOUR_TEMP_SECRET_ACCESS_KEY>
	aws_session_token = <YOUR_SESSION_TOKEN>
	[default]
	aws_access_key_id = <DEFAULT>
	aws_secret_access_key = <YOUR_TEMP_SECRET_ACCESS_KEY>
	aws_session_token = <YOUR_SESSION_TOKEN>
	`)
	if _, err := f.Write(data); err != nil {
		log.Fatal(err)
	}

	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", f.Name())
	os.Unsetenv("AWS_ACCESS_KEY_ID")
	os.Unsetenv("AWS_SECRET_ACCESS_KEY")
	os.Unsetenv("AWS_SESSION_TOKEN")

	// nested structs are ok here to simplify the test matrix
	// nolint:revive
	tests := []struct {
		stsclient    func(t *testing.T) AssumeRoleAPIClient
		account      Account
		awsConfig    *Config
		keyId        string
		envVariables []struct {
			key string
			val string
		}
	}{
		{
			stsclient: func(t *testing.T) AssumeRoleAPIClient {
				return mockAssumeRole(func(ctx context.Context, params *sts.AssumeRoleInput, optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error) {
					t.Helper()
					return &sts.AssumeRoleOutput{}, nil
				})
			},
			account: Account{
				LocalProfile: "test",
			},
			awsConfig: &Config{},
			keyId:     "<YOUR_TEMP_ACCESS_KEY_ID>",
		}, {
			stsclient: func(t *testing.T) AssumeRoleAPIClient {
				return mockAssumeRole(func(ctx context.Context, params *sts.AssumeRoleInput, optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error) {
					t.Helper()
					return &sts.AssumeRoleOutput{}, nil
				})
			},
			account:   Account{},
			awsConfig: &Config{},
			keyId:     "<DEFAULT>",
		},
		{
			stsclient: func(t *testing.T) AssumeRoleAPIClient {
				return mockAssumeRole(func(ctx context.Context, params *sts.AssumeRoleInput, optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error) {
					t.Helper()
					return &sts.AssumeRoleOutput{
						Credentials: &stsTypes.Credentials{
							AccessKeyId:     aws.String("<AssumedRoleKeyId>"),
							Expiration:      aws.Time(time.Now()),
							SecretAccessKey: aws.String("<AssumedRoleKeySecret>"),
							SessionToken:    aws.String("<AssumedRoleSessionToken>"),
						},
					}, nil
				})
			},

			account: Account{
				LocalProfile: "test",
				RoleARN:      "arn:aws:iam::123456789012:role/demo",
			},
			awsConfig: &Config{},
			keyId:     "<AssumedRoleKeyId>",
		}, {
			stsclient: func(t *testing.T) AssumeRoleAPIClient {
				return mockAssumeRole(func(ctx context.Context, params *sts.AssumeRoleInput, optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error) {
					t.Helper()
					return &sts.AssumeRoleOutput{
						Credentials: &stsTypes.Credentials{
							AccessKeyId:     aws.String("<AssumedRoleKeyId>"),
							Expiration:      aws.Time(time.Now()),
							SecretAccessKey: aws.String("<AssumedRoleKeySecret>"),
							SessionToken:    aws.String("<AssumedRoleKeySecret>"),
						},
					}, nil
				})
			},

			account: Account{
				LocalProfile: "test",
				RoleARN:      "arn:aws:iam::123456789012:role/demo",
				AccountID:    "asdfasdf",
			},
			awsConfig: &Config{},
			keyId:     "<AssumedRoleKeyId>",
		},
	}

	for i, tt := range tests {
		stsClient := tt.stsclient(t)
		awsClient, err := configureAwsClient(ctx, logger, tt.awsConfig, tt.account, stsClient)
		if err != nil {
			t.Errorf("Case-%d failed: %+v", i, err)
		}
		a, err := awsClient.Credentials.Retrieve(ctx)
		if err != nil {
			t.Errorf("Case-%d failed: %+v", i, err)
		}
		if a.AccessKeyID != tt.keyId {
			t.Errorf("Case-%d failed: %+v", i, err)
		}
	}
}
