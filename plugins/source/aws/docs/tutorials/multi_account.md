# Multi Account Configuration Tutorial

### AWS Organizations:

The plugin supports discovery of AWS Accounts via AWS Organizations. This means that as Accounts get added or removed from your organization, it will be able to handle new or removed accounts without any configuration changes.

Prerequisites for using AWS Org functionality:

1. Have a role (or user) in an Admin account with the following access:

- `organizations:ListAccounts`
- `organizations:ListAccountsForParent`
- `organizations:ListChildren`

2. Have a role in each child account that has a trust policy with a single principal. The default profile name is `OrganizationAccountAccessRole`. More information can be found [here](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_access.html#orgs_manage_accounts_create-cross-account-role), including how to create the role if it doesn't already exist in your account.

Using AWS Organization:

1. Specify member role name:

```yaml
org:
  member_role_name: OrganizationAccountAccessRole
```

2. Getting credentials that have the necessary `organizations` permissions:

   1. Sourcing Credentials from the default credential tool chain:

   ```yaml
   org:
     member_role_name: OrganizationAccountAccessRole
   ```

   2. Sourcing credentials from a named profile in the shared configuration or credentials file

   ```yaml
   org:
     member_role_name: OrganizationAccountAccessRole
     admin_account:
       local_profile: <Named-Profile>
   ```

   3. Assuming a role in admin account using credentials in the shared configuration or credentials file:

   ```yaml
       org:
         member_role_name: OrganizationAccountAccessRole
         admin_account:
           local_profile: <Named-Profile>
           role_arn: arn:aws:iam::<ACCOUNT_ID>:role/<ROLE_NAME>

           // Optional. Specify the name of the session
           // role_session_name: ""

           // Optional. Specify the ExternalID if required for trust policy
           // external_id: ""
   ```

3. Optional. If the trust policy configured for the member accounts requires different credentials than you configured in the previous step, then you can specify the credentials to use in the `member_trusted_principal` block

```yaml
org:
  member_role_name: OrganizationAccountAccessRole
  admin_account:
    local_profile: <Named-Profile-Admin>
  member_trusted_principal:
    local_profile: <Named-Profile-Member>
  organization_units:
    - ou-<ID-1>
    - ou-<ID-2>
```

4. Optional. If you want to specify specific Organizational Units to fetch from you can add them to the `organization_units` list.

```yaml
org:
  member_role_name: OrganizationAccountAccessRole
  admin_account:
    local_profile: <Named-Profile-Admin>
  organization_units:
    - ou-<ID-1>
    - ou-<ID-2>
```

**_note: If you specify an OU, CloudQuery will not traverse child OUs_**

#### Arguments for Org block

- `organization_units` **(Optional)** - List of Organizational Units that CloudQuery should use to source accounts from
- `admin_account` **(Optional)** - Configuration on how to grab credentials from an Admin account
- `member_trusted_principal` **(Optional)** - Configuration on how to specify the principle to use in order to assume a role in the member accounts
- `member_role_name` **(Required)** - Role name that CloudQuery should use to assume a role in the member account from the admin account. Note: This is not a full ARN, it is just the name
- `member_role_session_name` **(Optional)** - Override the default Session name.
- `member_external_id` **(Optional)** - Specify an ExternalID for use in the trust policy
- `member_regions` **(Optional)** - Limit fetching resources within this specific account to only these regions. This will override any regions specified in the provider block. You can specify all regions by using the `*` character as the only argument in the array

## Multi Account- Specific Accounts

CloudQuery can fetch from multiple accounts in parallel by using AssumeRole (You will need to use credentials that can AssumeRole to all other specified accounts). Below is an example configuration:

```yaml
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

#### Arguments for Accounts block

- `role_arn` **(Optional)** - The role that CloudQuery will use to perform the fetch
- `local_profile` **(Optional)** - Local Profile is the named profile in your shared configuration file (usually `~/.aws/config`) that you want to use for the account
- `external_id` **(Optional)** - The unique identifier used by non-AWS entities to assume a role in an AWS account
- `role_session_name` **(Optional)** - Override the default Session name.
- `default_region` **(Optional)** - this sets the Default Region for the AWS SDK. If you are assuming a role in a partition other than the AWS commercial region, it is important that this attribute is set
- `regions` **(Optional)** - Limit fetching resources within this specific account to only these regions. This will override any regions specified in the provider block. You can specify all regions by using the `*` character as the only argument in the array

### Assume Role with MFA

In order to assume role with MFA, you need to request temporary credentials using STS "get-session-token".

```bash
aws sts get-session-token --serial-number <YOUR_MFA_SERIAL_NUMBER> --token-code <YOUR_MFA_TOKEN_CODE> --duration-seconds 3600
```

Then export the temporary credentials to your environment variables.

```bash
export AWS_ACCESS_KEY_ID=<YOUR_ACCESS_KEY_ID>
export AWS_SECRET_ACCESS_KEY=<YOUR_SECRET_ACCESS_KEY>
export AWS_SESSION_TOKEN=<YOUR_SESSION_TOKEN>
```
