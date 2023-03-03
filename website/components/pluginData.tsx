export type Plugin = {
  id: string;
  name: string;
  href?: string;
  description?: string;
  logo: string;
  logoDark?: string;
  kind: "official" | "partner" | "community";
};

export const SOURCE_PLUGINS: Plugin[] = [
  {
    name: "Alibaba Cloud",
    logo: "/images/logos/plugins/alicloud.svg",
    logoDark: "/images/logos/plugins/alicloud-dark.svg",
    id: "alicloud",
    kind: "official",
  },
  {
    name: "AWS",
    logo: "/images/logos/plugins/aws.svg",
    logoDark: "/images/logos/plugins/aws-dark.svg",
    id: "aws",
    kind: "official",
  },
  {
    name: "Azure",
    logo: "/images/logos/plugins/azure.svg",
    id: "azure",
    kind: "official",
  },
  {
    name: "Azure DevOps",
    logo: "/images/logos/plugins/azuredevops.svg",
    id: "azuredevops",
    kind: "official",
  },
  {
    name: "Cloudflare",
    logo: "/images/logos/plugins/cloudflare.svg",
    id: "cloudflare",
    kind: "official",
  },
  {
    name: "Datadog",
    logo: "/images/logos/plugins/datadog.svg",
    logoDark: "/images/logos/plugins/datadog-dark.svg",
    id: "datadog",
    kind: "official",
  },
  {
    name: "Digital Ocean",
    logo: "/images/logos/plugins/digitalocean.svg",
    id: "digitalocean",
    kind: "official",
  },
  {
    name: "Fastly",
    logo: "/images/logos/plugins/fastly.svg",
    id: "fastly",
    kind: "official",
  },
  {
    name: "Gandi",
    logo: "/images/logos/plugins/gandi.svg",
    logoDark: "/images/logos/plugins/gandi-dark.svg",
    id: "gandi",
    kind: "official",
  },
  {
    name: "Google Cloud Platform",
    logo: "/images/logos/plugins/gcp.svg",
    id: "gcp",
    kind: "official",
  },
  {
    name: "GitHub",
    logo: "/images/logos/plugins/github.svg",
    logoDark: "/images/logos/plugins/github-dark.svg",
    id: "github",
    kind: "official",
  },
  {
    name: "GitLab",
    logo: "/images/logos/plugins/gitlab.svg",
    id: "gitlab",
    kind: "official",
  },
  {
    name: "Hacker News",
    logo: "/images/logos/plugins/hackernews.svg",
    id: "hackernews",
    kind: "official",
  },
  {
    name: "Heroku",
    logo: "/images/logos/plugins/heroku.svg",
    id: "heroku",
    kind: "official",
  },
  {
    name: "Homebrew",
    logo: "/images/logos/plugins/homebrew.svg",
    id: "homebrew",
    kind: "official",
  },
  {
    name: "HubSpot",
    logo: "/images/logos/plugins/hubspot.svg",
    id: "hubspot",
    kind: "official",
  },
  {
    name: "Kubernetes",
    logo: "/images/logos/plugins/kubernetes.svg",
    id: "k8s",
    kind: "official",
  },
  {
    name: "LaunchDarkly",
    logo: "/images/logos/plugins/launchdarkly.svg",
    logoDark: "/images/logos/plugins/launchdarkly-dark.svg",
    id: "launchdarkly",
    kind: "official",
  },
  {
    name: "Mixpanel",
    logo: "/images/logos/plugins/mixpanel.svg",
    logoDark: "/images/logos/plugins/mixpanel-dark.svg",
    id: "mixpanel",
    kind: "official",
  },
  {
    name: "Okta",
    logo: "/images/logos/plugins/okta.svg",
    id: "okta",
    kind: "official",
  },
  {
    name: "Oracle",
    logo: "/images/logos/plugins/oracle.svg",
    id: "oracle",
    kind: "official",
  },
  {
    name: "Pagerduty",
    logo: "/images/logos/plugins/pagerduty.svg",
    id: "pagerduty",
    kind: "official",
  },
  {
    name: "Plausible Analytics",
    logo: "/images/logos/plugins/plausibleanalytics.svg",
    id: "plausible",
    kind: "official",
  },
  {
    name: "PostgreSQL",
    logo: "/images/logos/plugins/postgresql.svg",
    id: "postgresql",
    kind: "official",
  },
  {
    name: "Salesforce",
    logo: "/images/logos/plugins/salesforce.svg",
    id: "salesforce",
    kind: "official",
  },
  {
    name: "SharePoint",
    logo: "/images/logos/plugins/sharepoint.svg",
    id: "sharepoint",
    kind: "community",
  },
  {
    name: "Swetrix",
    logo: "/images/logos/plugins/swetrix.svg",
    href: "https://github.com/swetrix/cq-source-swetrix",
    id: "swetrix",
    kind: "partner",
  },
  {
    name: "Scaleway",
    logo: "/images/logos/plugins/scaleway.svg",
    href: "https://github.com/scaleway/cq-source-scaleway",
    id: "scaleway",
    kind: "partner",
  },
  {
    name: "Shopify",
    logo: "/images/logos/plugins/shopify.svg",
    id: "shopify",
    kind: "official",
  },
  {
    name: "Simple Analytics",
    logo: "/images/logos/plugins/simpleanalytics.svg",
    href: "https://github.com/simpleanalytics/cq-source-simpleanalytics",
    id: "simpleanalytics",
    kind: "partner",
  },
  {
    name: "Slack",
    logo: "/images/logos/plugins/slack.svg",
    id: "slack",
    kind: "official",
  },
  {
    name: "Snyk",
    logo: "/images/logos/plugins/snyk.svg",
    id: "snyk",
    kind: "official",
  },
  {
    name: "Stripe",
    logo: "/images/logos/plugins/stripe.svg",
    id: "stripe",
    kind: "official",
  },
  {
    name: "Tailscale",
    logo: "/images/logos/plugins/tailscale.svg",
    logoDark: "/images/logos/plugins/tailscale-dark.svg",
    id: "tailscale",
    kind: "official",
  },
  {
    name: "Terraform",
    logo: "/images/logos/plugins/terraform.svg",
    id: "terraform",
    kind: "official",
  },
  {
    name: "Vercel",
    logo: "/images/logos/plugins/vercel.svg",
    logoDark: "/images/logos/plugins/vercel-dark.svg",
    id: "vercel",
    kind: "official",
  },
  {
    name: "Yandex Cloud",
    logo: "/images/logos/plugins/yandex.svg",
    id: "yandexcloud",
    href: "https://github.com/yandex-cloud/cq-source-yandex",
    kind: "partner",
  },
];

export const DESTINATION_PLUGINS: Plugin[] = [
  {
    name: "Azure Blob Storage",
    logo: "/images/logos/plugins/azblob.svg",
    id: "azblob",
    kind: "official",
  },
  {
    name: "BigQuery",
    logo: "/images/logos/plugins/bigquery.svg",
    id: "bigquery",
    kind: "official",
  },
  {
    name: "ClickHouse",
    logo: "/images/logos/plugins/clickhouse.svg",
    id: "clickhouse",
    kind: "official",
  },
  {
    name: "Elasticsearch",
    logo: "/images/logos/plugins/elasticsearch.svg",
    id: "elasticsearch",
    kind: "official",
  },
  {
    name: "File",
    logo: "/images/logos/plugins/file.svg",
    id: "file",
    kind: "official",
  },
  {
    name: "GCS",
    logo: "/images/logos/plugins/gcs.svg",
    id: "gcs",
    kind: "official",
  },
  {
    name: "Kafka",
    logo: "/images/logos/plugins/kafka.svg",
    logoDark: "/images/logos/plugins/kafka-dark.svg",
    id: "kafka",
    kind: "official",
  },
  {
    name: "Microsoft SQL Server",
    logo: "/images/logos/plugins/mssql.svg",
    id: "mssql",
    kind: "official",
  },
  {
    name: "MySQL",
    logo: "/images/logos/plugins/mysql.svg",
    logoDark: "/images/logos/plugins/mysql-dark.svg",
    id: "mysql",
    kind: "official",
  },
  {
    name: "MongoDB",
    logo: "/images/logos/plugins/mongodb.svg",
    id: "mongodb",
    kind: "official",
  },
  {
    name: "Neo4j",
    logo: "/images/logos/plugins/neo4j.svg",
    id: "neo4j",
    kind: "official",
  },
  {
    name: "PostgreSQL",
    logo: "/images/logos/plugins/postgresql.svg",
    id: "postgresql",
    kind: "official",
  },
  {
    name: "S3",
    logo: "/images/logos/plugins/s3.svg",
    id: "s3",
    kind: "official",
  },
  {
    name: "Snowflake",
    logo: "/images/logos/plugins/snowflake.svg",
    id: "snowflake",
    kind: "official",
  },
  {
    name: "SQLite",
    logo: "/images/logos/plugins/sqlite.svg",
    id: "sqlite",
    kind: "official",
  },
];

export function getPlugin(kind: string, id: string): Plugin {
  if (kind === "destination") {
    return DESTINATION_PLUGINS.find((p) => p.id === id);
  }
  return SOURCE_PLUGINS.find((p) => p.id === id);
}
