
# Table: aws_iam_virtual_mfa_devices
Contains information about a virtual MFA device.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|serial_number|text|The serial number associated with VirtualMFADevice.|
|base32_string_seed|bytea|The base32 seed defined as specified in RFC3548 (https://tools.ietf.org/html/rfc3548.txt). The Base32StringSeed is base64-encoded. |
|enable_date|timestamp without time zone|The date and time on which the virtual MFA device was enabled. |
|qr_code_png|bytea|A QR code PNG image that encodes otpauth://totp/$virtualMFADeviceName@$AccountName?secret=$Base32String where $virtualMFADeviceName is one of the create call arguments. AccountName is the user name if set (otherwise, the account ID otherwise), and Base32String is the seed in base32 format. The Base32String value is base64-encoded. |
|tags|jsonb|A list of tags that are attached to the virtual MFA device. For more information about tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User Guide. |
|user_arn|text|The Amazon Resource Name (ARN) that identifies the user. For more information about ARNs and how to use ARNs in policies, see IAM Identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.|
|user_create_date|timestamp without time zone|The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the user was created.|
|user_path|text|The path to the user. For more information about paths, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide. The ARN of the policy used to set the permissions boundary for the user.|
|user_id|text|The stable and unique string identifying the user. For more information about IDs, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.|
|user_name|text|The friendly name identifying the user.|
|user_password_last_used|timestamp without time zone|The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the user's password was last used to sign in to an AWS website. For a list of AWS websites that capture a user's last sign-in time, see the Credential reports (https://docs.aws.amazon.com/IAM/latest/UserGuide/credential-reports.html) topic in the IAM User Guide. If a password is used more than once in a five-minute span, only the first use is returned in this field. If the field is null (no value), then it indicates that they never signed in with a password. This can be because:|
|user_permissions_boundary_permissions_boundary_arn|text|The ARN of the policy used to set the permissions boundary for the user or role. |
|user_permissions_boundary_permissions_boundary_type|text|The permissions boundary usage type that indicates what type of IAM resource is used as the permissions boundary for an entity. This data type can only have a value of Policy. |
|user_tags|jsonb|A list of tags that are associated with the user. For more information about tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User Guide. |
