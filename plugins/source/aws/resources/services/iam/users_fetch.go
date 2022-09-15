package iam

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/gocarina/gocsv"
	"github.com/spf13/cast"
)

type AccessKeyWrapper struct {
	types.AccessKeyMetadata
	LastRotated time.Time
}

type UserWrapper struct {
	types.User
	UserCreationTime      time.Time
	PasswordStatus        string
	PasswordLastChanged   string
	PasswordNextRotation  string
	MfaActive             bool
	AccessKey1Active      bool
	AccessKey2Active      bool
	AccessKey1LastRotated string
	AccessKey2LastRotated string
	Cert1Active           bool
	Cert2Active           bool
	Cert1LastRotated      string
	Cert2LastRotated      string

	// internal
	reportUser *reportUser
	isRoot     bool
}

type reportUser struct {
	User                  string    `csv:"user"`
	ARN                   string    `csv:"arn"`
	UserCreationTime      time.Time `csv:"user_creation_time"`
	PasswordStatus        string    `csv:"password_enabled"`
	PasswordLastChanged   string    `csv:"password_last_changed"`
	PasswordNextRotation  string    `csv:"password_next_rotation"`
	MfaActive             bool      `csv:"mfa_active"`
	AccessKey1Active      bool      `csv:"access_key_1_active"`
	AccessKey2Active      bool      `csv:"access_key_2_active"`
	AccessKey1LastRotated string    `csv:"access_key_1_last_rotated"`
	AccessKey2LastRotated string    `csv:"access_key_2_last_rotated"`
	Cert1Active           bool      `csv:"cert_1_active"`
	Cert2Active           bool      `csv:"cert_2_active"`
	Cert1LastRotated      string    `csv:"cert_1_last_rotated"`
	Cert2LastRotated      string    `csv:"cert_2_last_rotated"`
}

type reportUsers []*reportUser

const rootName = "<root_account>"

func fetchIamUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return client.ListAndDetailResolver(ctx, meta, res, listUsers, userDetail)
}

func listUsers(ctx context.Context, meta schema.ClientMeta, detailChan chan<- interface{}) error {
	report, err := getCredentialReport(ctx, meta)
	if err != nil {
		return err
	}

	for _, user := range report {
		detailChan <- user
	}
	return nil
}

func userDetail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, listInfo interface{}) {
	c := meta.(*client.Client)

	reportUser := listInfo.(*reportUser)
	if reportUser.ARN == c.AccountGlobalARN(client.IamService, "root") {
		resultsChan <- UserWrapper{
			User: types.User{
				Arn:        aws.String(reportUser.ARN),
				CreateDate: aws.Time(reportUser.UserCreationTime),
				UserId:     aws.String("root"),
				UserName:   aws.String(reportUser.User),
			},
			UserCreationTime:      reportUser.UserCreationTime,
			PasswordStatus:        reportUser.PasswordStatus,
			PasswordLastChanged:   reportUser.PasswordLastChanged,
			PasswordNextRotation:  reportUser.PasswordNextRotation,
			MfaActive:             reportUser.MfaActive,
			AccessKey1Active:      reportUser.AccessKey1Active,
			AccessKey2Active:      reportUser.AccessKey2Active,
			AccessKey1LastRotated: reportUser.AccessKey1LastRotated,
			AccessKey2LastRotated: reportUser.AccessKey2LastRotated,
			Cert1Active:           reportUser.Cert1Active,
			Cert2Active:           reportUser.Cert2Active,
			Cert1LastRotated:      reportUser.Cert1LastRotated,
			Cert2LastRotated:      reportUser.Cert2LastRotated,
			reportUser:            reportUser,
			isRoot:                true,
		}
		return
	}
	svc := meta.(*client.Client).Services().IAM
	userDetail, err := svc.GetUser(ctx, &iam.GetUserInput{
		UserName: aws.String(reportUser.User),
	})
	if err != nil {
		if c.IsNotFoundError(err) {
			return
		}
		errorChan <- err
		return
	}
	resultsChan <- UserWrapper{
		User:                  *userDetail.User,
		UserCreationTime:      reportUser.UserCreationTime,
		PasswordStatus:        reportUser.PasswordStatus,
		PasswordLastChanged:   reportUser.PasswordLastChanged,
		PasswordNextRotation:  reportUser.PasswordNextRotation,
		MfaActive:             reportUser.MfaActive,
		AccessKey1Active:      reportUser.AccessKey1Active,
		AccessKey2Active:      reportUser.AccessKey2Active,
		AccessKey1LastRotated: reportUser.AccessKey1LastRotated,
		AccessKey2LastRotated: reportUser.AccessKey2LastRotated,
		Cert1Active:           reportUser.Cert1Active,
		Cert2Active:           reportUser.Cert2Active,
		Cert1LastRotated:      reportUser.Cert1LastRotated,
		Cert2LastRotated:      reportUser.Cert2LastRotated,
		reportUser:            reportUser,
		isRoot:                false,
	}
}

func postIamUserResolver(_ context.Context, _ schema.ClientMeta, resource *schema.Resource) error {
	r := resource.Item.(UserWrapper)
	if r.reportUser == nil {
		return nil
	}

	location, err := time.LoadLocation("UTC")
	if err != nil {
		return err
	}

	// Only set if cast is successful
	if enabled, err := cast.ToBoolE(r.PasswordStatus); err == nil {
		if err := resource.Set("password_enabled", enabled); err != nil {
			return err
		}
	}

	if r.reportUser.ARN == "" {
		if err := resource.Set("password_next_rotation", nil); err != nil {
			return err
		}
		if err := resource.Set("password_last_changed", nil); err != nil {
			return err
		}
		if err := resource.Set("cert_1_last_rotated", nil); err != nil {
			return err
		}
		if err := resource.Set("cert_2_last_rotated", nil); err != nil {
			return err
		}
		if err := resource.Set("access_key_1_last_rotated", nil); err != nil {
			return err
		}

		return resource.Set("access_key_2_last_rotated", nil)
	}

	if r.PasswordNextRotation == "N/A" || r.PasswordNextRotation == "not_supported" {
		if err := resource.Set("password_next_rotation", nil); err != nil {
			return err
		}
	} else {
		passwordNextRotation, err := time.ParseInLocation(time.RFC3339, r.PasswordNextRotation, location)
		if err != nil {
			return err
		}
		if err := resource.Set("password_next_rotation", passwordNextRotation); err != nil {
			return err
		}
	}

	if r.PasswordLastChanged == "N/A" || r.PasswordLastChanged == "not_supported" {
		if err := resource.Set("password_last_changed", nil); err != nil {
			return err
		}
	} else {
		passwordLastChanged, err := time.ParseInLocation(time.RFC3339, r.PasswordLastChanged, location)
		if err != nil {
			return err
		}
		if err := resource.Set("password_last_changed", passwordLastChanged); err != nil {
			return err
		}
	}

	if r.Cert1LastRotated == "N/A" || r.Cert1LastRotated == "not_supported" {
		if err := resource.Set("cert_1_last_rotated", nil); err != nil {
			return err
		}
	} else {
		cert1LastRotated, err := time.ParseInLocation(time.RFC3339, r.Cert1LastRotated, location)
		if err != nil {
			return err
		}
		if err := resource.Set("cert_1_last_rotated", cert1LastRotated); err != nil {
			return err
		}
	}

	if r.Cert2LastRotated == "N/A" || r.Cert2LastRotated == "not_supported" {
		if err := resource.Set("cert_2_last_rotated", nil); err != nil {
			return err
		}
	} else {
		cert2LastRotated, err := time.ParseInLocation(time.RFC3339, r.Cert2LastRotated, location)
		if err != nil {
			return err
		}
		if err := resource.Set("cert_2_last_rotated", cert2LastRotated); err != nil {
			return err
		}
	}

	if r.AccessKey1LastRotated == "N/A" || r.AccessKey1LastRotated == "not_supported" {
		if err := resource.Set("access_key_1_last_rotated", nil); err != nil {
			return err
		}
	} else {
		accessKey1LastRotated, err := time.ParseInLocation(time.RFC3339, r.AccessKey1LastRotated, location)
		if err != nil {
			return err
		}
		if err := resource.Set("access_key_1_last_rotated", accessKey1LastRotated); err != nil {
			return err
		}
	}

	if r.AccessKey2LastRotated == "N/A" || r.AccessKey2LastRotated == "not_supported" {
		if err := resource.Set("access_key_2_last_rotated", nil); err != nil {
			return err
		}
	} else {
		accessKey2LastRotated, err := time.ParseInLocation(time.RFC3339, r.AccessKey2LastRotated, location)
		if err != nil {
			return err
		}
		if err := resource.Set("access_key_2_last_rotated", accessKey2LastRotated); err != nil {
			return err
		}
	}

	return nil
}

func fetchIamUserGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListGroupsForUserInput
	p := parent.Item.(UserWrapper)
	if aws.ToString(p.UserName) == rootName {
		return nil
	}
	svc := meta.(*client.Client).Services().IAM
	config.UserName = p.UserName
	for {
		output, err := svc.ListGroupsForUser(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.Groups
		if output.Marker == nil {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}

func fetchIamUserAccessKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListAccessKeysInput
	p := parent.Item.(UserWrapper)
	svc := meta.(*client.Client).Services().IAM
	if aws.ToString(p.UserName) == rootName {
		return nil
	}
	config.UserName = p.UserName
	for {
		output, err := svc.ListAccessKeys(ctx, &config)
		if err != nil {
			return err
		}

		keys := make([]AccessKeyWrapper, len(output.AccessKeyMetadata))
		for i, key := range output.AccessKeyMetadata {
			switch i {
			case 0:
				rotated := parent.Get("access_key_1_last_rotated")
				if rotated != nil {
					keys[i] = AccessKeyWrapper{key, rotated.(time.Time)}
				} else {
					keys[i] = AccessKeyWrapper{key, *key.CreateDate}
				}
			case 1:
				rotated := parent.Get("access_key_2_last_rotated")
				if rotated != nil {
					keys[i] = AccessKeyWrapper{key, rotated.(time.Time)}
				} else {
					keys[i] = AccessKeyWrapper{key, *key.CreateDate}
				}
			default:
				keys[i] = AccessKeyWrapper{key, time.Time{}}
			}
		}
		res <- keys
		if output.Marker == nil {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}

func postIamUserAccessKeyResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	r := resource.Item.(AccessKeyWrapper)
	if r.AccessKeyId == nil {
		return nil
	}
	svc := meta.(*client.Client).Services().IAM
	output, err := svc.GetAccessKeyLastUsed(ctx, &iam.GetAccessKeyLastUsedInput{AccessKeyId: r.AccessKeyId})
	if err != nil {
		return err
	}
	if output.AccessKeyLastUsed != nil {
		if err := resource.Set("last_used", output.AccessKeyLastUsed.LastUsedDate); err != nil {
			return err
		}
		if err := resource.Set("last_used_service_name", output.AccessKeyLastUsed.ServiceName); err != nil {
			return err
		}
	}
	return nil
}

func fetchIamUserAttachedPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListAttachedUserPoliciesInput
	p := parent.Item.(UserWrapper)
	if aws.ToString(p.UserName) == rootName {
		return nil
	}
	svc := meta.(*client.Client).Services().IAM
	config.UserName = p.UserName
	for {
		output, err := svc.ListAttachedUserPolicies(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.AttachedPolicies
		if output.Marker == nil {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}

func (r reportUsers) GetUser(arn string) *reportUser {
	for _, u := range r {
		if u.ARN == arn {
			return u
		}
	}
	return nil
}

func getCredentialReport(ctx context.Context, meta schema.ClientMeta) (reportUsers, error) {
	var err error
	var apiErr smithy.APIError
	var reportOutput *iam.GetCredentialReportOutput
	svc := meta.(*client.Client).Services().IAM
	for {
		reportOutput, err = svc.GetCredentialReport(ctx, &iam.GetCredentialReportInput{})
		if err == nil && reportOutput != nil {
			var users reportUsers
			err = gocsv.UnmarshalBytes(reportOutput.Content, &users)
			if err != nil {
				return nil, err
			}
			return users, nil
		}
		if !errors.As(err, &apiErr) {
			return nil, err
		}
		switch apiErr.ErrorCode() {
		case "ReportNotPresent", "ReportExpired":
			_, err := svc.GenerateCredentialReport(ctx, &iam.GenerateCredentialReportInput{})
			if err != nil {
				var serviceError smithy.APIError
				if !errors.As(err, &serviceError) {
					return nil, err
				}
				// LimitExceeded is the only specific error that should not stop processing
				// If Limit Exceeded is returned we should try and see if there is a credential report
				// already generated so we want to sleep for 5 seconds then continue
				if serviceError.ErrorCode() != "LimitExceeded" {
					return nil, err
				}
				if err := client.Sleep(ctx, 5*time.Second); err != nil {
					return nil, err
				}
			}
		case "ReportInProgress":
			meta.Logger().Debug().Msg("Waiting for credential report to be generated")
			if err := client.Sleep(ctx, 5*time.Second); err != nil {
				return nil, err
			}
		default:
			return nil, err
		}
	}
}
