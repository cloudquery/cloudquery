
# Table: github_organization_member_membership
Membership represents the status of a user's membership in an organization or team.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|organization_member_cq_id|uuid|Unique CloudQuery ID of github_organization_members table (FK)|
|url|text||
|state|text|State is the user's status within the organization or team. Possible values are: "active", "pending"|
|role|text|Role identifies the user's role within the organization or team. Possible values for organization membership:     member - non-owner organization member     admin - organization owner  Possible values for team membership are:     member - a normal member of the team     maintainer - a team maintainer|
|organization_url|text|For organization membership, the API URL of the organization.|
