# Multi Account Configuration Tutorial

## AWS Organizations

The plugin supports discovery of AWS Accounts via AWS Organizations. This means that as Accounts get added or removed from your organization, it will be able to handle new or removed accounts without any configuration changes.

```yaml
kind: source
spec:
  name: aws
  registry: github
  path: cloudquery/aws
  version: "VERSION_SOURCE_AWS"
  tables: ['*']
  destinations: ["postgresql"]
  spec:
    aws_debug: false
    org:
      admin_account:
        local_profile: "<NAMED_PROFILE>"
      member_role_name: OrganizationAccountAccessRole
    regions:
      - '*'
  ```

Prerequisites for using AWS Org functionality:

1. Have a role (or user) in an Admin account with the following access:

  - `organizations:ListAccounts`
  - `organizations:ListAccountsForParent`
  - `organizations:ListChildren`

2. Have a role in each member account that has a trust policy with a single principal. The default profile name is `OrganizationAccountAccessRole`. More information can be found [here](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_access.html#orgs_manage_accounts_create-cross-account-role), including how to create the role if it doesn't already exist in your account.

Configuring AWS Organization:

1. It is always necessary to specify a member role name:

   ```yaml copy
       org:
         member_role_name: OrganizationAccountAccessRole
   ```

2. Sourcing credentials that have the necessary `organizations` permissions can be done in any of the following ways:

    1. Source credentials from the default credential tool chain:
       ```yaml copy
           org:
             member_role_name: OrganizationAccountAccessRole
       ```

    2. Source credentials from a named profile in the shared configuration or credentials file
       ```yaml copy
           org:
             member_role_name: OrganizationAccountAccessRole
             admin_account:
               local_profile: <Named-Profile>
       ```

    3. Assume a role in admin account using credentials in the shared configuration or credentials file:

       ```yaml copy
           org:
             member_role_name: OrganizationAccountAccessRole
             admin_account:
               local_profile: <Named-Profile>
               role_arn: arn:aws:iam::<ACCOUNT_ID>:role/<ROLE_NAME>
               
               # Optional. Specify the name of the session 
               # role_session_name: ""
   
               # Optional. Specify the ExternalID if required for trust policy 
               # external_id: ""
       ```

3. Optional. If the trust policy configured for the member accounts requires different credentials than you configured in the previous step, then you can specify the credentials to use in the `member_trusted_principal` block: 

   ```yaml copy
       org:
         member_role_name: OrganizationAccountAccessRole
         member_trusted_principal:
           local_profile: <Named-Profile-Member>
   ```

4. Optional. If you want to specify specific Organizational Units to fetch from you can add them to the `organization_units` list. 

   ```yaml copy
       org:
         member_role_name: OrganizationAccountAccessRole
         organization_units:
           - ou-<ID-1>
           - ou-<ID-2>
   ```
   
   Child OUs will also be included. To skip a child OU or account, use the `skip_organization_units` or `skip_member_accounts` options respectively:

   ```yaml copy
       org:
         member_role_name: OrganizationAccountAccessRole
         organization_units:
           - ou-<ID-1>
           - ou-<ID-2>
         skip_organization_units:
           - ou-<ID-3>
         skip_member_accounts:
           - <ACCOUNT_ID>
   ```

import { Callout } from 'nextra-theme-docs'

<Callout type="info">
Note that in AWS plugin versions before v9.0.0, child OUs were not traversed when specifying an OU, and `skip_organization_units` and `skip_member_accounts` were not supported. These options are only available in v9.0.0 and above, and child OUs are now traversed by default.
</Callout>

### Arguments for Org block

See [AWS org configuration](/docs/plugins/sources/aws/configuration#org) for more information on all the arguments in the `org` block.

## Multi Account: Specific Accounts

CloudQuery can fetch from multiple accounts in parallel by using AssumeRole (You will need to use credentials that can AssumeRole to all other specified accounts). Below is an example configuration:

```yaml copy
accounts:
  - id: <AccountID_Alias_1>
    role_arn: <YOUR_ROLE_ARN_1>
    # Optional. Local Profile is the named profile in your shared configuration file (usually `~/.aws/config`) that you want to use for this specific account
    local_profile: <NAMED_PROFILE>
    # Optional. Specify the Role Session name
    role_session_name: ""
  - id: <AccountID_Alias_2>
    local_profile: provider
    # Optional. Role ARN we want to assume when accessing this account
    role_arn: <YOUR_ROLE_ARN_2>
```

### Arguments for Accounts block

See [AWS accounts configuration](/docs/plugins/sources/aws/configuration#accounts) for more information on all the arguments in the `accounts` block.
