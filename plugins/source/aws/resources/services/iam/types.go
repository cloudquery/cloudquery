package iam

import "github.com/aws/aws-sdk-go-v2/service/iam/types"

type ReportUser struct {
	User string `csv:"user"`
	// The Amazon Resource Name (ARN) that identifies the user
	ARN string `csv:"arn"`
	// When the user has a password, this value is TRUE. Otherwise it is FALSE.The value for the AWS account root user is always not_supported
	PasswordStatus string `csv:"password_enabled"`
	// The date and time when the user's password was last set, in ISO 8601 date-time format. If the user does not have a password, the value in this field is N/A (not applicable). The value for the AWS account (root) is always NULL
	PasswordLastChanged string `csv:"password_last_changed"`
	// When the account has a password policy that requires password rotation, this field contains the date and time, in ISO 8601 date-time format, when the user is required to set a new password. The value for the AWS account (root) is always NULL
	PasswordNextRotation string `csv:"password_next_rotation"`
	// When a multi-factor authentication (MFA) device has been enabled for the user, this value is TRUE. Otherwise it is FALSE
	MfaActive bool `csv:"mfa_active"`
	// When the user has an access key and the access key's status is Active, this value is TRUE. Otherwise it is FALSE
	AccessKey1Active bool `csv:"access_key_1_active"`
	// When the user has an access key and the access key's status is Active, this value is TRUE. Otherwise it is FALSE
	AccessKey2Active bool `csv:"access_key_2_active"`
	// The date and time, in ISO 8601 date-time format, when the user's access key was created or last changed
	AccessKey1LastRotated string `csv:"access_key_1_last_rotated"`
	// The date and time, in ISO 8601 date-time format, when the user's access key was created or last changed
	AccessKey2LastRotated string `csv:"access_key_2_last_rotated"`
	// When the user has an X.509 signing certificate and that certificate's status is Active, this value is TRUE. Otherwise it is FALSE
	Cert1Active bool `csv:"cert_1_active"`
	// When the user has an X.509 signing certificate and that certificate's status is Active, this value is TRUE. Otherwise it is FALSE
	Cert2Active bool `csv:"cert_2_active"`
	// The date and time, in ISO 8601 date-time format, when the user's signing certificate was created or last changed
	Cert1LastRotated string `csv:"cert_1_last_rotated"`
	// The date and time, in ISO 8601 date-time format, when the user's signing certificate was created or last changed
	Cert2LastRotated string `csv:"cert_2_last_rotated"`
}

type ReportUsers []*ReportUser

type wrappedKey struct {
	types.AccessKeyMetadata
	types.AccessKeyLastUsed
}
type WrappedUser struct {
	types.User
	*ReportUser
}
