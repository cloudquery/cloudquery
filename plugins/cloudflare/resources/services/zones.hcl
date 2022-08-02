service          = "cloudflare"
output_directory = "."
add_generate     = true

resource "cloudflare" "" "zones" {
  path = "github.com/cloudflare/cloudflare-go/.Zone"

  multiplex "CFAccount" {
    path   = "github.com/cloudquery/cq-provider-cloudflare/client.AccountMultiplex"
  }

  options {
    primary_keys = [
      "id"
    ]
  }

  userDefinedColumn "account_id" {
    description = "The Account ID of the resource."
    type        = "string"
    resolver "resolveCFAccount" {
      path = "github.com/cloudquery/cq-provider-cloudflare/client.ResolveAccountId"
    }
  }

  column "id" {
    description = "The unique universal identifier for a Cloudflare zone."
  }

  column "name" {
    description = "Cloudflare zone name."
  }

  column "dev_mode" {
    description = "DevMode contains the time in seconds until development expires (if positive) or since it expired (if negative)"
  }

  column "original_n_s" {
    rename      = "original_ns"
    description = "Creation timestamp of the account."
  }

  column "original_registrar" {
    description = "Cloudflare zone original name servers."
  }

  column "original_dns_host" {
    description = "Cloudflare zone original registrar."
  }

  column "created_on" {
    description = "Cloudflare zone original dns host."
  }

  column "modified_on" {
    description = "Zone created on date and time."
  }

  column "name_servers" {
    description = "Zone last modified date and time."
  }

  column "owner_id" {
    description = "Zone owner id."
  }

  column "owner_email" {
    description = "Zone owner email address."
  }

  column "owner_name" {
    description = "Zone owner name."
  }

  column "owner_type" {
    description = "Zone owner type."
  }

  column "permissions" {
    description = "Zone permissions."
  }

  column "status" {
    description = "Zone status."
  }

  column "paused" {
    description = "True if zone is paused."
  }

  column "type" {
    description = "Zone type."
  }

  column "host" {
    skip        = true
    description = "Zone host."
  }

  userDefinedColumn "host_name" {
    description       = "Zone host name."
    type              = "string"
    generate_resolver = false
  }

  userDefinedColumn "host_website" {
    description       = "Zone host website."
    type              = "string"
    generate_resolver = false
  }

  column "vanity_n_s" {
    rename      = "vanity_ns"
    description = "Zone vanity name servers."
  }

  column "betas" {
    description = "Zone betas."
  }

  column "deact_reason" {
    rename      = "deactivation_reason"
    description = "Zone deactivation reason."
  }

  column "meta" {
    skip_prefix = true
  }

  column "page_rule_quota" {
    description = "custom_certificate_quota is broken - sometimes it's a string, sometimes a number! CustCertQuota     int    `json:\"custom_certificate_quota\"`"
  }

  column "plan_zone_plan_common_id" {
    rename = "plan_id"
    description = "The unique universal identifier for a Cloudflare zone plan."
  }

  column "plan_zone_plan_common_name" {
    rename = "plan_name"
    description = "Cloudflare zone plan name."
  }

  column "plan_zone_plan_common_price" {
    rename = "plan_price"
    description = "Cloudflare zone plan price."
  }

  column "plan_zone_plan_common_currency" {
    rename = "plan_currency"
    description = "Cloudflare zone plan currency."
  }

  column "plan_zone_plan_common_frequency" {
    rename = "plan_frequency"
    description = "Cloudflare zone plan frequency."
  }

  column "plan_legacy_id" {
    description = "True if zone plan is subscribed."
  }

  column "plan_is_subscribed" {
    description = "True if zone plan can subscribe."
  }

  column "plan_can_subscribe" {
    description = "Cloudflare zone plan legacy id."
  }

  column "plan_legacy_discount" {
    description = "True if zone plan has legacy discount."
  }

  column "plan_externally_managed" {
    description = "True if zone plan is externally managed."
  }

  column "plan_pending_zone_plan_common_id" {
    rename = "plan_pending_id"
    description = "The unique universal identifier for a Cloudflare zone plan."
  }

  column "plan_pending_zone_plan_common_name" {
    rename = "plan_pending_name"
    description = "Cloudflare zone plan name."
  }

  column "plan_pending_zone_plan_common_price" {
    rename = "plan_pending_price"
    description = "Cloudflare zone plan price."
  }

  column "plan_pending_zone_plan_common_currency" {
    rename = "plan_pending_currency"
    description = "Cloudflare zone plan currency."
  }

  column "plan_pending_zone_plan_common_frequency" {
    rename = "plan_pending_frequency"
    description = "Cloudflare zone plan frequency."
  }

  column "plan_pending_legacy_id" {
    description = "True if zone plan is subscribed."
  }

  column "plan_pending_is_subscribed" {
    description = "True if zone plan can subscribe."
  }

  column "plan_pending_can_subscribe" {
    description = "Cloudflare zone plan legacy id."
  }

  column "plan_pending_legacy_discount" {
    description = "True if zone plan has legacy discount."
  }

  column "plan_pending_externally_managed" {
    description = "True if zone plan is externally managed."
  }

  column "account" {
    skip        = true
    description = "Zone account id."
  }

  column "verification_key" {
    description = "Zone verification key."
  }
}
