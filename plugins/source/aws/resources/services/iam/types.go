package iam

import "github.com/aws/aws-sdk-go-v2/service/iam/types"

type ReportUser struct {
	User                  string `csv:"user"`
	ARN                   string `csv:"arn"`
	PasswordStatus        string `csv:"password_enabled"`
	PasswordLastChanged   string `csv:"password_last_changed"`
	PasswordNextRotation  string `csv:"password_next_rotation"`
	MfaActive             bool   `csv:"mfa_active"`
	AccessKey1Active      bool   `csv:"access_key_1_active"`
	AccessKey2Active      bool   `csv:"access_key_2_active"`
	AccessKey1LastRotated string `csv:"access_key_1_last_rotated"`
	AccessKey2LastRotated string `csv:"access_key_2_last_rotated"`
	Cert1Active           bool   `csv:"cert_1_active"`
	Cert2Active           bool   `csv:"cert_2_active"`
	Cert1LastRotated      string `csv:"cert_1_last_rotated"`
	Cert2LastRotated      string `csv:"cert_2_last_rotated"`
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
