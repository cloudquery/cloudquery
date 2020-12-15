package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"
)

var policyHeaderConfig = `queries:`

var awsCIS = `
  - name: "AWS CIS check 1.1. Avoid the use of 'root' account. Show used in last 30 days."
    query: >
      SELECT account_id, arn, password_last_used, user_name FROM aws_iam_users
      WHERE user_name = '<root_account>' AND DATE(password_last_used) > date('now', '-30 day')
  - name: "AWS CIS check 1.2. Ensure MFA is enabled for all IAM users that have a console password"
    query: >
      SELECT account_id, arn, password_last_used, user_name, mfa_active FROM aws_iam_users
      WHERE password_enabled AND NOT mfa_active
  - name: "AWS CIS check 1.3. Ensure credentials unused for 90 days or greater are disabled"
    query: >
      SELECT account_id, arn, password_last_used, user_name, access_key_id, last_used FROM aws_iam_users
        JOIN aws_iam_user_access_keys on aws_iam_users.id = aws_iam_user_access_keys.user_id
       WHERE (password_enabled AND DATE(password_last_used) < date('now', '-90 day') OR
             (DATE(last_used) < date('now', '-90 day')))
  - name: "AWS CIS check 1.4. Ensure access keys are rotated every 90 days or less"
    query: >
      SELECT account_id, arn, password_last_used, user_name, access_key_id, last_used, last_rotated FROM aws_iam_users
        JOIN aws_iam_user_access_keys on aws_iam_users.id = aws_iam_user_access_keys.user_id
       WHERE DATE(last_rotated) < date('now', '-90 day')
  - name: "AWS CIS check 1.5.  Ensure IAM password policy requires at least one uppercase letter"
    query: >
      SELECT account_id, require_uppercase_characters FROM aws_iam_password_policies
       WHERE require_uppercase_characters = 0
  - name: "AWS CIS check 1.6.  Ensure IAM password policy requires at least one lowercase letter"
    query: >
      SELECT account_id, require_lowercase_characters FROM aws_iam_password_policies
       WHERE require_lowercase_characters = 0
  - name: "AWS CIS check 1.7.  Ensure IAM password policy requires at least one symbol"
    query: >
      SELECT account_id, require_symbols FROM aws_iam_password_policies
       WHERE require_symbols = 0
  - name: "AWS CIS check 1.8.  Ensure IAM password policy requires at least one number"
    query: >
      SELECT account_id, require_numbers FROM aws_iam_password_policies
       WHERE require_numbers = 0
  - name: "AWS CIS check 1.9. Ensure IAM password policy requires minimum length of 14 or greater"
    query: >
      SELECT account_id, minimum_password_length FROM aws_iam_password_policies
       WHERE minimum_password_length < 14
  - name: "AWS CIS check 1.10. Ensure IAM password policy requires minimum length of 14 or greater"
    query: >
      SELECT account_id, password_reuse_prevention FROM aws_iam_password_policies
       WHERE password_reuse_prevention is NULL or password_reuse_prevention != 1
  - name: "AWS CIS check 1.11. Ensure IAM password policy expires passwords within 90 days or less"
    query: >
      SELECT account_id, max_password_age FROM aws_iam_password_policies
       WHERE max_password_age is NULL or max_password_age < 90
  - name: "AWS CIS check 1.13. Ensure MFA is enabled for the 'root' account"
    query: >
      SELECT account_id, arn, password_last_used, user_name, mfa_active FROM aws_iam_users
      WHERE user_name = '<root_account>' AND NOT mfa_active
  - name: "AWS CIS check 2.2. Ensure CloudTrail log file validation is enabled"
    query: >
      SELECT account_id, region, trail_arn, log_file_validation_enabled FROM aws_cloudtrail_trails
      WHERE log_file_validation_enabled = 0`

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
		_, err = s.WriteString(policyHeaderConfig)
		if err != nil {
			return err
		}
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
