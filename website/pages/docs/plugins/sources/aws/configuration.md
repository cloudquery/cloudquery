# AWS Source Plugin Configuration Reference

## Simple Example

This example connects a single AWS account in one region to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

```yaml
kind: source
spec:
  # Source spec section
  name: aws
  path: cloudquery/aws
  version: "VERSION_SOURCE_AWS"
  tables: ["*"]
  destinations: ["postgresql"]
  spec: 
    # AWS Spec section described below
    regions: 
      - us-east-1
    accounts:
      - id: "account1"
        local_profile: "account1"
    aws_debug: false
```

### AWS Organization Example


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


CloudQuery supports discovery of AWS Accounts via AWS Organizations. This means that as Accounts get added or removed from your organization CloudQuery will be able to handle new or removed accounts without any configuration changes.

Prerequisites for using AWS Org functionality:
1. Have a role (or user) in an Admin account with the following access:

  - `organizations:ListAccounts`
  - `organizations:ListAccountsForParent`
  - `organizations:ListChildren`

2. Have a role in each member account that has a trust policy with a single principal. The default profile name is `OrganizationAccountAccessRole`. More information can be found [here](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_access.html#orgs_manage_accounts_create-cross-account-role), including how to create the role if it doesn't already exist in your account.






Using AWS Organization:
1. Specify member role name:

```yaml
    org:
      member_role_name: OrganizationAccountAccessRole
```

2. Getting credentials that have  the necessary `organizations` permissions:

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




## AWS Spec

This is the (nested) spec used by the AWS source plugin.

- `regions` ([]string) (default: Empty. Will use all enabled regions)

  Regions to use.

- `accounts` ([][account](#accounts)) (default: current account)

  List of all accounts to fetch information from

- `org` ([org](#org)) (default: not used)

  In AWS organization mode, CloudQuery will source all accounts underneath automatically

- `aws_debug` (bool) (default: false)

  If true, will log AWS debug logs, including retries and other request/response metadata

## accounts

This is used to specify one or more accounts to extract information from. Note that it should be an array of objects, each with the following fields:

- `id` (string) (**required**)

  Will be used as an alias in the source plugin and in the logs

- `local_profile` (string) (default: will use current credentials)

  [Local profile](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html) to use to authenticate this account with.
  Please note this should be set to the name of the profile. For example, with the following credentials file:

  ```ini
  [default]
  aws_access_key_id=xxxx
  aws_secret_access_key=xxxx

  [user1]
  aws_access_key_id=xxxx
  aws_secret_access_key=xxxx
  ```

  `local_profile` should be set to either `default` or `user1`.

- `role_arn` (string)

  If specified will use this to assume role

- `role_session_name` (string)

  If specified will use this session name when assume role to `role_arn`

- `external_id` (string)

  If specified will use this when assuming role to `role_arn`

- `default_region` (string) (default: `us-east-1`)

  If specified, this region will be used as the default region for the account.

- `regions` (string)

  Regions to use for this account. Defaults to global `regions` setting.


## org

- `organization_units` ([]string)

  List of Organizational Units that CloudQuery should use to source accounts from. If you specify an OU, CloudQuery will not traverse nested OUs

- `admin_account` ([Account](#account))

  Configuration for how to grab credentials from an Admin account

- `member_trusted_principal` ([Account](#account))

  Configuration for how to specify the principle to use in order to assume a role in the member accounts

- `member_role_name` (string) (**required**)

  Role name that CloudQuery should use to assume a role in the member account from the admin account. Note: This is not a full ARN, it is just the name

- `member_role_session_name` (string)

  Override the default Session name.

- `member_external_id` (string)

  Specify an ExternalID for use in the trust policy

- `member_regions` ([]string)

  Limit fetching resources within this specific account to only these regions. This will override any regions specified in the provider block. You can specify all regions by using the `*` character as the only argument in the array
