package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"
)

var awsCIS = `
views:
  - name: "aws_log_metric_filter_and_alarm"
    query: >
      CREATE VIEW aws_log_metric_filter_and_alarm AS
      SELECT * FROM aws_cloudtrail_trails
        JOIN aws_cloudtrail_trail_event_selectors on aws_cloudtrail_trails.id = aws_cloudtrail_trail_event_selectors.trail_id
        JOIN aws_cloudwatchlogs_metric_filters on aws_cloudtrail_trails.cloud_watch_logs_log_group_name = aws_cloudwatchlogs_metric_filters.log_group_name
        JOIN aws_cloudwatch_metric_alarm_metrics on aws_cloudwatchlogs_metric_filters.filter_name = aws_cloudwatch_metric_alarm_metrics.name
        JOIN aws_cloudwatch_metric_alarms on aws_cloudwatch_metric_alarm_metrics.metric_alarm_id = aws_cloudwatch_metric_alarms.id
        JOIN aws_cloudwatch_metric_alarm_actions ON aws_cloudwatch_metric_alarm_metrics.id = aws_cloudwatch_metric_alarm_actions.metric_alarm_id
        JOIN aws_sns_subscriptions ON aws_cloudwatch_metric_alarm_actions.value = aws_sns_subscriptions.topic_arn
      WHERE is_multi_region_trail=true AND is_logging=true
            AND include_management_events=true AND read_write_type = 'All'
            AND subscription_arn LIKE 'aws:arn:%'
queries:
  - name: "AWS CIS 1.1 Avoid the use of 'root' account. Show used in last 30 days (Scored)"
    query: >
      SELECT account_id, arn, password_last_used, user_name FROM aws_iam_users
      WHERE user_name = '<root_account>' AND DATE(password_last_used) > date('now', '-30 day')
  - name: "AWS CIS 1.2 Ensure MFA is enabled for all IAM users that have a console password (Scored)"
    query: >
      SELECT account_id, arn, password_last_used, user_name, mfa_active FROM aws_iam_users
      WHERE password_enabled AND NOT mfa_active
  - name: "AWS CIS 1.3 Ensure credentials unused for 90 days or greater are disabled (Scored)"
    query: >
      SELECT account_id, arn, password_last_used, user_name, access_key_id, last_used FROM aws_iam_users
        JOIN aws_iam_user_access_keys on aws_iam_users.id = aws_iam_user_access_keys.user_id
       WHERE (password_enabled AND DATE(password_last_used) < date('now', '-90 day') OR
             (DATE(last_used) < date('now', '-90 day')))
  - name: "AWS CIS 1.4 Ensure access keys are rotated every 90 days or less"
    query: >
      SELECT account_id, arn, password_last_used, user_name, access_key_id, last_used, last_rotated FROM aws_iam_users
        JOIN aws_iam_user_access_keys on aws_iam_users.id = aws_iam_user_access_keys.user_id
       WHERE DATE(last_rotated) < date('now', '-90 day')
  - name: "AWS CIS 1.5  Ensure IAM password policy requires at least one uppercase letter"
    query: >
      SELECT account_id, require_uppercase_characters FROM aws_iam_password_policies
       WHERE require_uppercase_characters = 0
  - name: "AWS CIS 1.6  Ensure IAM password policy requires at least one lowercase letter"
    query: >
      SELECT account_id, require_lowercase_characters FROM aws_iam_password_policies
       WHERE require_lowercase_characters = 0
  - name: "AWS CIS 1.7  Ensure IAM password policy requires at least one symbol"
    query: >
      SELECT account_id, require_symbols FROM aws_iam_password_policies
       WHERE require_symbols = 0
  - name: "AWS CIS 1.8  Ensure IAM password policy requires at least one number"
    query: >
      SELECT account_id, require_numbers FROM aws_iam_password_policies
       WHERE require_numbers = 0
  - name: "AWS CIS 1.9 Ensure IAM password policy requires minimum length of 14 or greater"
    query: >
      SELECT account_id, minimum_password_length FROM aws_iam_password_policies
       WHERE minimum_password_length < 14
  - name: "AWS CIS 1.10 Ensure IAM password policy prevents password reuse"
    query: >
      SELECT account_id, password_reuse_prevention FROM aws_iam_password_policies
       WHERE password_reuse_prevention is NULL or password_reuse_prevention > 24
  - name: "AWS CIS 1.11 Ensure IAM password policy expires passwords within 90 days or less"
    query: >
      SELECT account_id, max_password_age FROM aws_iam_password_policies
       WHERE max_password_age is NULL or max_password_age < 90
  - name: "AWS CIS 1.12  Ensure no root account access key exists (Scored)"
    query: >
      select * from aws_iam_users
          JOIN aws_iam_user_access_keys aiuak on aws_iam_users.id = aiuak.user_id
      WHERE user_name = '<root>'
  - name: "AWS CIS 1.13 Ensure MFA is enabled for the 'root' account"
    query: >
      SELECT account_id, arn, password_last_used, user_name, mfa_active FROM aws_iam_users
      WHERE user_name = '<root_account>' AND NOT mfa_active
  - name: "AWS CIS 1.14 Ensure hardware MFA is enabled for the 'root' account (Scored)"
    query: >
      SELECT aws_iam_users.account_id, arn, password_last_used, user_name, mfa_active, count(*) FROM aws_iam_users
          JOIN aws_iam_virtual_mfa_devices ON aws_iam_virtual_mfa_devices.user_arn = aws_iam_users.arn
      WHERE user_name = '<root_account>' AND mfa_active
      GROUP BY aws_iam_users.user_name
      HAVING count(*) = 1
  - name: "AWS CIS 1.16 Ensure IAM policies are attached only to groups or roles (Scored)"
    query: >
      SELECT aws_iam_users.account_id, arn, user_name FROM aws_iam_users
      JOIN aws_iam_user_attached_policies aiuap on aws_iam_users.id = aiuap.user_id
  - name: "AWS CIS 2.1 Ensure CloudTrail is enabled in all regions"
    query: >
      SELECT aws_cloudtrail_trails.account_id, trail_arn, is_multi_region_trail, read_write_type, include_management_events FROM aws_cloudtrail_trails
      JOIN aws_cloudtrail_trail_event_selectors on aws_cloudtrail_trails.id = aws_cloudtrail_trail_event_selectors.trail_id
      WHERE is_multi_region_trail = FALSE OR (is_multi_region_trail = TRUE AND (read_write_type != 'All' OR include_management_events = FALSE))
  - name: "AWS CIS 2.2 Ensure CloudTrail log file validation is enabled"
    query: >
      SELECT aws_cloudtrail_trails.account_id, region, trail_arn, log_file_validation_enabled FROM aws_cloudtrail_trails
      WHERE log_file_validation_enabled = 0
  - name: "AWS CIS 2.4 Ensure CloudTrail trails are integrated with CloudWatch Logs"
    query: >
      SELECT aws_cloudtrail_trails.account_id, trail_arn, latest_cloud_watch_logs_delivery_time from aws_cloudtrail_trails
      WHERE cloud_watch_logs_log_group_arn is NULL OR DATE(latest_cloud_watch_logs_delivery_time) < date('now', '-1 day')
  - name: "AWS CIS 2.6 Ensure S3 bucket access logging is enabled on the CloudTrail S3 bucket"
    query: >
      SELECT aws_cloudtrail_trails.account_id, s3_bucket_name, trail_arn from aws_cloudtrail_trails
      JOIN aws_s3_buckets on s3_bucket_name = aws_s3_buckets.name
      WHERE logging_target_bucket is NULL OR logging_target_prefix is NULL
  - name: "AWS CIS 2.7 Ensure CloudTrail logs are encrypted at rest using KMS CMKs"
    query: >
      SELECT account_id, region, trail_arn, kms_key_id from aws_cloudtrail_trails
      WHERE kms_key_id is NULL
  - name: "AWS CIS 2.8 Ensure rotation for customer created CMKs is enabled (Scored)"
    query: >
      SELECT account_id, region, arn FROM aws_kms_keys WHERE rotation_enabled = FALSE AND manager = 'CUSTOMER'
  - name: "AWS CIS 2.9 Ensure VPC flow logging is enabled in all VPCs (Scored)"
    query: >
      SELECT aws_ec2_vpcs.account_id, aws_ec2_vpcs. region, vpc_id FROM aws_ec2_vpcs
      LEFT JOIN aws_ec2_flow_logs ON aws_ec2_vpcs.vpc_id = aws_ec2_flow_logs.resource_id WHERE aws_ec2_flow_logs.resource_id is NULL
  - name: "AWS CIS 3.1.1 Ensure a log metric filter and alarm exist for unauthorized API calls (Scored)"
    query: >
      SELECT account_id, region, trail_arn FROM aws_cloudtrail_trails
        JOIN aws_cloudtrail_trail_event_selectors actes on aws_cloudtrail_trails.id = actes.trail_id
        WHERE is_multi_region_trail=true AND is_logging=false
  - name: "AWS CIS 3.1 Ensure a log metric filter and alarm exist for unauthorized API calls (Scored)"
    invert: true
    query: >
      SELECT account_id, region, cloud_watch_logs_log_group_arn  FROM aws_log_metric_filter_and_alarm
      WHERE filter_pattern='{ ($.errorCode = "*UnauthorizedOperation") || ($.errorCode = "AccessDenied*") }'
  - name: "AWS CIS 3.2 Ensure a log metric filter and alarm exist for Management Console sign-in without MFA (Scored)"
    invert: true
    query: >
      SELECT account_id, region, cloud_watch_logs_log_group_arn  FROM aws_log_metric_filter_and_alarm
      WHERE filter_pattern='{ ($.errorCode = "ConsoleLogin") || ($.additionalEventData.MFAUsed != "Yes")  }'
  - name: "AWS CIS 3.3  Ensure a log metric filter and alarm exist for usage of 'root' account (Score)"
    invert: true
    query: >
      SELECT account_id, region, cloud_watch_logs_log_group_arn  FROM aws_log_metric_filter_and_alarm
      WHERE filter_pattern='{ $.userIdentity.type = "Root" && $.userIdentity.invokedBy NOT EXISTS && $.eventType != "AwsServiceEvent" }'
  - name: "AWS CIS 3.4 Ensure a log metric filter and alarm exist for IAM policy changes (Score)"
    invert: true
    query: >
      SELECT account_id, region, cloud_watch_logs_log_group_arn  FROM aws_log_metric_filter_and_alarm
      WHERE filter_pattern='{ ($.eventName = DeleteGroupPolicy) || ($.eventName = DeleteRolePolicy) || ($.eventName = DeleteUserPolicy) || ($.eventName = PutGroupPolicy) || ($.eventName = PutRolePolicy) || ($.eventName = PutUserPolicy) || ($.eventName = CreatePolicy) || ($.eventName = DeletePolicy) || ($.eventName=CreatePolicyVersion) || ($.eventName=DeletePolicyVersion) || ($.eventName=AttachRolePolicy) || ($.eventName=DetachRolePolicy) || ($.eventName=AttachUserPolicy) || ($.eventName = DetachUserPolicy) || ($.eventName = AttachGroupPolicy) || ($.eventName = DetachGroupPolicy)}'
  - name: "AWS CIS 3.5 Ensure a log metric filter and alarm exist for CloudTrail configuration changes (Scored)"
    invert: true
    query: >
      SELECT account_id, region, cloud_watch_logs_log_group_arn  FROM aws_log_metric_filter_and_alarm
      WHERE filter_pattern='{ ($.eventName = CreateTrail) || ($.eventName = UpdateTrail) || ($.eventName = DeleteTrail) || ($.eventName = StartLogging) || ($.eventName = StopLogging) }'
  - name: "AWS CIS 3.6 Ensure a log metric filter and alarm exist for AWS Management Console authentication failures (Scored)"
    invert: true
    query: >
      SELECT account_id, region, cloud_watch_logs_log_group_arn  FROM aws_log_metric_filter_and_alarm
      WHERE filter_pattern='{ ($.eventName = ConsoleLogin) && ($.errorMessage = "Failed authentication") }'
  - name: "AWS CIS 3.7 Ensure a log metric filter and alarm exist for disabling or scheduled deletion of customer created CMKs (Scored)"
    invert: true
    query: >
      SELECT account_id, region, cloud_watch_logs_log_group_arn  FROM aws_log_metric_filter_and_alarm
      WHERE filter_pattern='{($.eventSource = kms.amazonaws.com) && (($.eventName=DisableKey)||($.eventName=ScheduleKeyDeletion)) }"'
  - name: "AWS CIS 3.8 Ensure a log metric filter and alarm exist for S3 bucket policy changes (Scored)"
    invert: true
    query: >
      SELECT account_id, region, cloud_watch_logs_log_group_arn  FROM aws_log_metric_filter_and_alarm
      WHERE filter_pattern='{ ($.eventSource = s3.amazonaws.com) && (($.eventName = PutBucketAcl) || ($.eventName = PutBucketPolicy) || ($.eventName = PutBucketCors) || ($.eventName = PutBucketLifecycle) || ($.eventName = PutBucketReplication) || ($.eventName = DeleteBucketPolicy) || ($.eventName = DeleteBucketCors) || ($.eventName = DeleteBucketLifecycle) || ($.eventName = DeleteBucketReplication)) }'
  - name: "AWS CIS 3.9 Ensure a log metric filter and alarm exist for AWS Config configuration changes (Scored)"
    invert: true
    query: >
      SELECT account_id, region, cloud_watch_logs_log_group_arn  FROM aws_log_metric_filter_and_alarm
      WHERE filter_pattern='{ ($.eventSource = config.amazonaws.com) && (($.eventName = StopConfigurationRecorder) || ($.eventName = DeleteDeliveryChannel) || ($.eventName = PutDeliveryChannel) || ($.eventName = PutConfigurationRecorder)) }'
  - name: "AWS CIS 3.10 Ensure a log metric filter and alarm exist for security group changes (Scored)"
    invert: true
    query: >
      SELECT account_id, region, cloud_watch_logs_log_group_arn  FROM aws_log_metric_filter_and_alarm
      WHERE filter_pattern='{ ($.eventName = AuthorizeSecurityGroupIngress) || ($.eventName = AuthorizeSecurityGroupEgress) || ($.eventName = RevokeSecurityGroupIngress) || ($.eventName = RevokeSecurityGroupEgress) || ($.eventName = CreateSecurityGroup) || ($.eventName = DeleteSecurityGroup) }'
  - name: "AWS CIS 3.11 Ensure a log metric filter and alarm exist for changes to Network Access Control Lists (NACL) (Scored)"
    invert: true
    query: >
      SELECT account_id, region, cloud_watch_logs_log_group_arn  FROM aws_log_metric_filter_and_alarm
      WHERE filter_pattern='{ ($.eventName = CreateNetworkAcl) || ($.eventName = CreateNetworkAclEntry) || ($.eventName = DeleteNetworkAcl) || ($.eventName = DeleteNetworkAclEntry) || ($.eventName = ReplaceNetworkAclEntry) || ($.eventName = ReplaceNetworkAclAssociation) }'
  - name: "AWS CIS 3.12 Ensure a log metric filter and alarm exist for changes to network gateways (Scored)"
    invert: true
    query: >
      SELECT account_id, region, cloud_watch_logs_log_group_arn  FROM aws_log_metric_filter_and_alarm
      WHERE filter_pattern='{ ($.eventName = CreateCustomerGateway) || ($.eventName = DeleteCustomerGateway) || ($.eventName = AttachInternetGateway) || ($.eventName = CreateInternetGateway) || ($.eventName = DeleteInternetGateway) || ($.eventName = DetachInternetGateway) }'
  - name: "AWS CIS 3.13 Ensure a log metric filter and alarm exist for route table changes (Scored)"
    invert: true
    query: >
      SELECT account_id, region, cloud_watch_logs_log_group_arn  FROM aws_log_metric_filter_and_alarm
      WHERE filter_pattern='{ ($.eventName = CreateRoute) || ($.eventName = CreateRouteTable) || ($.eventName = ReplaceRoute) || ($.eventName = ReplaceRouteTableAssociation) || ($.eventName = DeleteRouteTable) || ($.eventName = DeleteRoute) || ($.eventName = DisassociateRouteTable) }'
  - name: "AWS CIS 3.14 Ensure a log metric filter and alarm exist for VPC changes (Scored)"
    invert: true
    query: >
      SELECT account_id, region, cloud_watch_logs_log_group_arn  FROM aws_log_metric_filter_and_alarm
      WHERE filter_pattern='{ ($.eventName = CreateVpc) || ($.eventName = DeleteVpc) || ($.eventName = ModifyVpcAttribute) || ($.eventName = AcceptVpcPeeringConnection) || ($.eventName = CreateVpcPeeringConnection) || ($.eventName = DeleteVpcPeeringConnection) || ($.eventName = RejectVpcPeeringConnection) || ($.eventName = AttachClassicLinkVpc) || ($.eventName = DetachClassicLinkVpc) || ($.eventName = DisableVpcClassicLink) || ($.eventName = EnableVpcClassicLink) }'
  - name: "AWS CIS 4.1 Ensure no security groups allow ingress from 0.0.0.0/0 to port 22 (Scored)"
    query: >
      select account_id, region, group_name, from_port, to_port, cidr_ip from aws_ec2_security_groups
          JOIN aws_ec2_security_group_ip_permissions on aws_ec2_security_groups.id = aws_ec2_security_group_ip_permissions.security_group_id
          JOIN aws_ec2_security_group_ip_ranges on aws_ec2_security_group_ip_permissions.id = aws_ec2_security_group_ip_ranges.security_group_ip_permission_id
      WHERE from_port >= 0 AND to_port <= 22 AND cidr_ip = '0.0.0.0/0'
  - name: "AWS CIS 4.2 Ensure no security groups allow ingress from 0.0.0.0/0 to port 3389 (Scored)"
    query: >
      select account_id, region, group_name, from_port, to_port, cidr_ip from aws_ec2_security_groups
          JOIN aws_ec2_security_group_ip_permissions on aws_ec2_security_groups.id = aws_ec2_security_group_ip_permissions.security_group_id
          JOIN aws_ec2_security_group_ip_ranges on aws_ec2_security_group_ip_permissions.id = aws_ec2_security_group_ip_ranges.security_group_ip_permission_id
      WHERE from_port >= 0 AND to_port <= 3389 AND cidr_ip = '0.0.0.0/0'
  - name: "AWS CIS 4.3  Ensure the default security group of every VPC restricts all traffic (Scored)"
    query: >
      select account_id, region, group_name, from_port, to_port, cidr_ip from aws_ec2_security_groups
          JOIN aws_ec2_security_group_ip_permissions on aws_ec2_security_groups.id = aws_ec2_security_group_ip_permissions.security_group_id
          JOIN aws_ec2_security_group_ip_ranges on aws_ec2_security_group_ip_permissions.id = aws_ec2_security_group_ip_ranges.security_group_ip_permission_id
      WHERE group_name='default' AND cidr_ip = '0.0.0.0/0'`

var policies = map[string]string{
	"aws_cis": awsCIS,
}

var policyValidArgs = []string{"aws_cis"}
var policyPath string

var policyCmd = &cobra.Command{
	Use:       fmt.Sprintf("policy [choose one or more of: %s]", strings.Join(policyValidArgs, ",")),
	Short:     "Generate initial policy.yml for query command",
	ValidArgs: policyValidArgs,
	Args:      cobra.RangeArgs(1, len(policyValidArgs)),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cobra.OnlyValidArgs(cmd, args)
		if err != nil {
			return fmt.Errorf("invalid argument %s for cloudquery gen policy. choose from %v", args[0], validArgs)
		}
		var s strings.Builder
		for _, provider := range args {
			s.WriteString(policies[provider])
		}
		s.WriteString("\n")
		if _, err := os.Stat(policyPath); err == nil {
			return fmt.Errorf("file %s already exists. Either delete it or specify other path via --path flag", policyPath)
		} else if os.IsNotExist(err) {
			return ioutil.WriteFile(policyPath, []byte(s.String()), 0644)
		} else {
			return err
		}
	},
}

func init() {
	genCmd.AddCommand(policyCmd)
	policyCmd.Flags().StringVar(&policyPath, "path", "./policy.yml", "path to output generated policy file")
}
