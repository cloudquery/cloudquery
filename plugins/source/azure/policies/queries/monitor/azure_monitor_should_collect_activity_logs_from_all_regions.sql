insert into azure_policy_results
SELECT 
  :'execution_time',
  :'framework',
  :'check_id',
  'Azure Monitor should collect activity logs from all regions',
  s.id,
  s.id
FROM
  azure_subscription_subscriptions s
  LEFT OUTER JOIN azure_monitor_log_profiles p
  ON s.id = '/subscriptions/' || p.subscription_id
WHERE
  p.properties -> 'locations' IS NULL
  OR NOT p.properties -> 'locations'  @> '["australiacentral", "australiacentral2", "australiaeast", "australiasoutheast", "brazilsouth", "brazilsoutheast", "canadacentral", "canadaeast", "centralindia", "centralus", "eastasia", "eastus", "eastus2", "francecentral", "francesouth", "germanynorth", "germanywestcentral", "japaneast", "japanwest", "jioindiawest", "koreacentral", "koreasouth", "northcentralus", "northeurope", "norwayeast", "norwaywest", "southafricanorth", "southafricawest", "southcentralus", "southeastasia", "southindia", "switzerlandnorth", "switzerlandwest", "uaecentral", "uaenorth", "uksouth", "ukwest", "westcentralus", "westeurope", "westindia", "westus", "westus2", "westus3", "global"]'::jsonb;
