package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemovePII(t *testing.T) {
	cases := []struct {
		Input    string
		Expected string
	}{
		{
			"msg AccountID 123456789 something",
			"msg AccountID xxxx something",
		},
		{
			"msg 123456789 something",
			"msg 1234xxxxxxxx something",
		},
		{
			"operation error S3: GetBucketLogging, https response error StatusCode: 404, RequestID: 3PQRRTJ1BAB82DWH, HostID: MZfZa61jiI+47KWhZjIl1N92GVKOMhslal/A2dcWMJwk7rZazsIflh1LNa3yIDqjrRpF1fF17/k=, api error NoSuchBucket: The specified bucket does not exist",
			"operation error S3: GetBucketLogging, https response error StatusCode: 404, RequestID: xxxx, HostID: xxxx, api error NoSuchBucket: The specified bucket does not exist",
		},
		{
			"AccessDenied: User: arn:aws:sts::123456789:assumed-role/some-role/i-012304405c679abcd is not authorized to perform: sts:AssumeRole on resource: arn:aws:iam::123456789:role/other-role\n\tstatus code: 403, request id: d2f12332-d1f2-12c5-1234-abc12345d123",
			"AccessDenied: User: arn:aws:xxxx is not authorized to perform: sts:AssumeRole on resource: arn:aws:xxxx code: 403, request id: xxxx",
		},
		{
			"IAM: GetUserPolicy - User: arn:aws:sts::123456789:assumed-role/some-role/i-012304405c679abcd is not authorized to perform: iam:GetUserPolicy on resource: user some_user",
			"IAM: GetUserPolicy - User: arn:aws:xxxx is not authorized to perform: iam:GetUserPolicy on resource: user xxxx",
		},
		{
			"operation error CloudWatch Logs: DescribeMetricFilters, exceeded maximum number of attempts, 10, https response error StatusCode: 0, RequestID: , request send failed, Post \"https://logs.eu-central-1.amazonaws.com/\": dial tcp: lookup logs.eu-central-1.amazonaws.com on 192.168.1.1:53: read udp 192.168.1.2:5353->192.168.1.1:53: i/o timeout",
			"operation error CloudWatch Logs: DescribeMetricFilters, exceeded maximum number of attempts, 10, https response error StatusCode: 0, RequestID: , request send failed, Post \"https://xxxx\": dial tcp: lookup xxxx on xxxx:xx: read udp xxxx:xx->xxxx:xx: i/o timeout",
		},
		{
			"EC2: DescribeImageAttribute - You are not authorized to perform this operation. Encoded authorization failure message: SOMEENCODEDMESSAGEWITHNUMBERS1234567ANDDASHANDUNDERSCORES-ABCDE_123123123_EXAMPLEMESSAGE",
			"EC2: DescribeImageAttribute - You are not authorized to perform this operation. Encoded authorization failure message: xxxx",
		},
		{
			"operation error Elastic Beanstalk: DescribeConfigurationOptions, https response error StatusCode: 400, RequestID: 3PQRRTJ1BAB82DWH, api error InvalidParameterValue: Access Denied: S3Bucket=some-bucket-1, S3Key=object_path/some_key.ext (Service: Amazon S3; Status Code: 403; Error Code: AccessDenied; Request ID: 3PQRRTJ1BAB82DWH; Proxy: null)",
			"operation error Elastic Beanstalk: DescribeConfigurationOptions, https response error StatusCode: 400, RequestID: xxxx, api error InvalidParameterValue: Access Denied: S3Bucket=xxxx, S3Key=xxxx (Service: Amazon S3; Status Code: 403; Error Code: AccessDenied; Request ID: xxxx; Proxy: null)",
		},
		{
			"operation error Direct Connect: DescribeVirtualInterfaces, exceeded maximum number of attempts, 10, https response error StatusCode: 0, RequestID: , request send failed, Post \"https://logs.eu-central-1.amazonaws.com/\": dial tcp 177.72.244.112:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.",
			"operation error Direct Connect: DescribeVirtualInterfaces, exceeded maximum number of attempts, 10, https response error StatusCode: 0, RequestID: , request send failed, Post \"https://xxxx\": dial tcp xxxx:xx: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.",
		},
		{
			`operation error Cognito Identity Provider: ListUserPools, exceeded maximum number of attempts, 10, https response error StatusCode: 0, RequestID: , request send failed, Post "https://cognito-idp.us-west-2.amazonaws.com/": dial tcp [2600:1f14:917:5700:4845:5c16:891b:7127]:443: connect: network is unreachable`,
			`operation error Cognito Identity Provider: ListUserPools, exceeded maximum number of attempts, 10, https response error StatusCode: 0, RequestID: , request send failed, Post "https://xxxx": dial tcp xxxx:xx: connect: network is unreachable`,
		},
		{
			"operation error EC2: DescribeSnapshotAttribute, https response error StatusCode: 400, RequestID: xxxx, api error InvalidSnapshot.NotFound: The snapshot 'snap-11111111111111111' does not exist.",
			"operation error EC2: DescribeSnapshotAttribute, https response error StatusCode: 400, RequestID: xxxx, api error InvalidSnapshot.NotFound: The snapshot 'xxxx' does not exist.",
		},
		{
			"ResourceType name not found - Could not find example request type named 'resource-dev-1111'",
			"ResourceType name not found - Could not find example request type named 'xxxx'",
		},
		{
			`qldb.ledgers: failed to resolve table "aws_qldb_ledgers": error at github.com/cloudquery/cq-provider-aws/resources/services/qldb.fetchQldbLedgers[ledgers.go:264] operation error QLDB: ListLedgers, exceeded maximum number of attempts, 10, https response error StatusCode: 0, RequestID: , request send failed, Get "https://qldb.ap-southeast-1.amazonaws.com/ledgers": dial tcp: lookup qldb.ap-southeast-1.amazonaws.com on 172.20.0.10:53: no such host`,
			`qldb.ledgers: failed to resolve table "aws_qldb_ledgers": error at github.com/cloudquery/cq-provider-aws/resources/services/qldb.fetchQldbLedgers[ledgers.go:264] operation error QLDB: ListLedgers, exceeded maximum number of attempts, 10, https response error StatusCode: 0, RequestID: , request send failed, Get "https://xxxx": dial tcp: lookup xxxx on xxxx:xx: no such host`,
		},
		{
			`operation error EC2: DescribeImageAttribute, https response error StatusCode: 400, RequestID: 3PQRRTJ1BAB82DWH, api error InvalidAMIID.Unavailable: The image ID 'ami-01964cde3b8020132' is no longer available`,
			`operation error EC2: DescribeImageAttribute, https response error StatusCode: 400, RequestID: xxxx, api error InvalidAMIID.Unavailable: The image ID 'xxxx' is no longer available`,
		},
		{
			`operation error Auto Scaling: DescribePolicies, https response error StatusCode: 400, RequestID: 3PQRRTJ1BAB82DWH, api error ValidationError: Group group-name not found`,
			`operation error Auto Scaling: DescribePolicies, https response error StatusCode: 400, RequestID: xxxx, api error ValidationError: Group xxxx not found`,
		},
	}
	for i, tc := range cases {
		res := removePII([]string{"123456789"}, tc.Input)
		assert.Equalf(t, tc.Expected, res, "Case #%d", i+1)
	}
}
