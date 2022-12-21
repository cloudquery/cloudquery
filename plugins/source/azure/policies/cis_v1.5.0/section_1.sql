\echo "Executing CIS V1.5.0 Section 1: Identity and Access Management"

-- Section 1.1 Security Defaults
\set check_id '1.1.1'
\echo "Executing check 1.1.1"
\echo "Ensure Security Defaults is enabled on Azure Active Directory must be done manually"

\set check_id '1.1.2'
\echo "Executing check 1.1.2"
\echo "Ensure that `Multi-Factor Auth Status` is `Enabled` for all Privileged Users must be done manually"

\set check_id '1.1.3'
\echo "Executing check 1.1.3"
\echo "Ensure that `Multi-Factor Auth Status` is `Enabled` for all Non-Privileged Users must be done manually"

\set check_id '1.1.4'
\echo "Executing check 1.1.4"
\echo "Ensure that `Restore multi-factor authentication on all remembered devices` is Enabled must be done manually"

-- Section 1.2 Conditional Access
\set check_id '1.2.1'
\echo "Executing check 1.2.1"
\echo "Ensure Trusted Locations Are Defined must be done manually"

\set check_id '1.2.2'
\echo "Executing check 1.2.2"
\echo "Ensure that an exclusionary Georgraphic Access Policy is considered must be done manually"

\set check_id '1.2.3'
\echo "Executing check 1.2.3"
\echo "Ensure that a Multi-factor Authentication Policy Exists for Administrative Groups must be done manually"

\set check_id '1.2.4'
\echo "Executing check 1.2.4"
\echo "Ensure that a Multi-factor Authentication Policy Exists for All Users must be done manually"

\set check_id '1.2.5'
\echo "Executing check 1.2.5"
\echo "Ensure Multi-factor Authentication is Required for Risky Sign-ins must be done manually"

\set check_id `1.2.6`
\echo "Executing check 1.2.6"
\echo "Ensure Multi-factor Authentication is Required for Azure Management must be done manually"

\set check_id `1.3`
\echo "Executing check 1.3"
\echo "Ensure Access Review is Set Up for External Users in Azure AD Privileged Identity Management must be done manually"

\set check_id `1.4`
\echo "Executing check 1.4"
\echo "Ensure Guest Users Are Reviewed on a Regular Basis must be done manually"

\set check_id `1.5`
\echo "Executing check 1.5"
\echo "Ensure that 'Allow users to remember multi-factor authentication on devices they trust' is 'Disabled' must be done manually"

\set check_id `1.6`
\echo "Executing check 1.6"
\echo "Ensure that 'Number of methods required to reset' is set to '2' must be done manually"

\set check_id `1.7`
\echo "Executing check 1.7"
\echo "Ensure that a Custom Bad Password List is set to 'Enforce' for your Organization must be done manually"

\set check_id `1.8`
\echo "Executing check 1.8"
\echo "Ensure that Multi-factor Authentication is Required for Azure Management must be done manually"

\set check_id `1.9`
\echo "Executing check 1.9"
\echo "Ensure that 'Number of days before users are asked to re- confirm their authentication information' is not set to '0' must be done manually"

\set check_id `1.10`
\echo "Executing check 1.10"
\echo "Ensure that 'Notify all admins when other admins reset their password?' is set to 'Yes' must be done manually"

\set check_id `1.11`
\echo "Executing check 1.11"
\echo "Ensure that ‘Users Can Consent to Apps Accessing Company Data on Their Behalf’ Is Set To ‘Allow for Verified Publishers’ must be done manually"

\set check_id `1.12`
\echo "Executing check 1.12"
\echo "Ensure that 'Users can consent to apps accessing company data on their behalf' is set to 'No' must be done manually"

\set check_id `1.13`
\echo "Executing check 1.13"
\echo "Ensure that 'Users can add gallery apps to My Apps' is set to 'No' must be done manually"

\set check_id `1.14`
\echo "Executing check 1.14"
\echo "Ensure that ‘Users Can Register Applications’ Is Set to ‘No’ must be done manually"

\set check_id `1.15`
\echo "Executing check 1.15"
\echo "Ensure that 'Guest users access restrictions' is set to 'Guest user access is restricted to properties and memberships of their own directory objects' must be done manually"

\set check_id `1.16`
\echo "Executing check 1.16"
\echo "Ensure that Ensure that 'Guest invite restrictions' is set to `Only users assigned to specific admin roles can invite guest users` must be done manually"

\set check_id `1.17`
\echo "Executing check 1.17"
\echo "Ensure That 'Restrict access to Azure AD administration portal' is Set to 'Yes' must be done manually"

\set check_id `1.18`
\echo "Executing check 1.18"
\echo "Ensure that 'Restrict user ability to access groups features in the Access Pane' is Set to 'Yes' must be done manually"

\set check_id `1.19`
\echo "Executing check 1.19"
\echo "Ensure that 'Users can create security groups in Azure portals, API or PowerShell' is set to 'No' must be done manually"

\set check_id `1.20`
\echo "Executing check 1.20"
\echo "Ensure that 'Owners can manage group membership requests in the Access Panel' is set to 'No' must be done manually"

\set check_id `1.21`
\echo "Executing check 1.21"
\echo "Ensure that 'Users can create Microsoft 365 groups in Azure portals, API or PowerShell' is set to 'No' must be done manually"

\set check_id `1.22`
\echo "Executing check 1.22"
\echo "Ensure that 'Require Multi-Factor Authentication to register or join devices with Azure AD' is set to 'Yes' must be done manually"

\set check_id `1.23`
\echo "Executing check 1.23"
\echo "Ensure that No Custom Subscription Owner Roles Are Created"
\ir ../queries/iam/custom_subscription_owner_roles.sql
--TODO: Need to validate this check

\set check_id `1.24`
\echo "Executing check 1.24"
\echo "Ensure a Custom Role is Assigned Permissions for Administering Resource Locks must be done manually"

\set check_id `1.25`
\echo "Executing check 1.25"
\echo "Ensure that Subscription Entering AAD Directory’ and ‘Subscription Leaving AAD Directory’ Is Set To ‘Permit No One’ must be done manually"