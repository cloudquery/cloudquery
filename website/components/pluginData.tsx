export type Plugin = {
  id: string;
  name: string;
  href?: string;
  description?: string;
  logo: string;
  logoDark?: string;
};

export const SOURCE_PLUGINS: Plugin[] = [
  {
    name: "Alibaba Cloud",
    logo: "/images/logos/plugins/alicloud.svg",
    logoDark: "/images/logos/plugins/alicloud-dark.svg",
    id: "alicloud",
  },
  {
    name: "AWS",
    logo: "/images/logos/plugins/aws.svg",
    logoDark: "/images/logos/plugins/aws-dark.svg",
    id: "aws",
  },
  {
    name: "Azure",
    logo: "/images/logos/plugins/azure.svg",
    id: "azure",
  },
  {
    name: "Azure DevOps",
    logo: "/images/logos/plugins/azuredevops.svg",
    id: "azuredevops",
  },
  {
    name: "Cloudflare",
    logo: "/images/logos/plugins/cloudflare.svg",
    id: "cloudflare",
  },
  {
    name: "Datadog",
    logo: "/images/logos/plugins/datadog.svg",
    logoDark: "/images/logos/plugins/datadog-dark.svg",
    id: "datadog",
  },
  {
    name: "Digital Ocean",
    logo: "/images/logos/plugins/digitalocean.svg",
    id: "digitalocean",
  },
  {
    name: "Fastly",
    logo: "/images/logos/plugins/fastly.svg",
    id: "fastly",
  },
  {
    name: "Gandi",
    logo: "/images/logos/plugins/gandi.svg",
    logoDark: "/images/logos/plugins/gandi-dark.svg",
    id: "gandi",
  },
  {
    name: "Google Cloud Platform",
    logo: "/images/logos/plugins/gcp.svg",
    id: "gcp",
  },
  {
    name: "GitHub",
    logo: "/images/logos/plugins/github.svg",
    logoDark: "/images/logos/plugins/github-dark.svg",
    id: "github",
  },
  {
    name: "GitLab",
    logo: "/images/logos/plugins/gitlab.svg",
    id: "gitlab",
  },
  {
    name: "Hacker News",
    logo: "/images/logos/plugins/hackernews.svg",
    id: "hackernews",
  },
  {
    name: "Heroku",
    logo: "/images/logos/plugins/heroku.svg",
    id: "heroku",
  },
  {
    name: "Homebrew",
    logo: "/images/logos/plugins/homebrew.svg",
    id: "homebrew",
  },
  {
    name: "HubSpot",
    logo: "/images/logos/plugins/hubspot.svg",
    id: "hubspot",
  },
  {
    name: "Kubernetes",
    logo: "/images/logos/plugins/kubernetes.svg",
    id: "k8s",
  },
  {
    name: "LaunchDarkly",
    logo: "/images/logos/plugins/launchdarkly.svg",
    logoDark: "/images/logos/plugins/launchdarkly-dark.svg",
    id: "launchdarkly",
  },
  {
    name: "Mixpanel",
    logo: "/images/logos/plugins/mixpanel.svg",
    logoDark: "/images/logos/plugins/mixpanel-dark.svg",
    id: "mixpanel",
  },
  {
    name: "Okta",
    logo: "/images/logos/plugins/okta.svg",
    id: "okta",
  },
  {
    name: "Oracle",
    logo: "/images/logos/plugins/oracle.svg",
    id: "oracle",
  },
  {
    name: "Pagerduty",
    logo: "/images/logos/plugins/pagerduty.svg",
    id: "pagerduty",
  },
  {
    name: "Plausible Analytics",
    logo: "/images/logos/plugins/plausibleanalytics.svg",
    id: "plausible",
  },
  {
    name: "PostgreSQL",
    logo: "/images/logos/plugins/postgresql.svg",
    id: "postgresql",
  },
  {
    name: "Salesforce",
    logo: "/images/logos/plugins/salesforce.svg",
    id: "salesforce",
  },
  {
    name: "Swetrix",
    logo: "/images/logos/plugins/swetrix.svg",
    href: "https://github.com/swetrix/cq-source-swetrix",
    id: "swetrix",
  },
  {
    name: "Scaleway",
    logo: "/images/logos/plugins/scaleway.svg",
    href: "https://github.com/scaleway/cq-source-scaleway",
    id: "scaleway",
  },
  {
    name: "Shopify",
    logo: "/images/logos/plugins/shopify.svg",
    id: "shopify",
  },
  {
    name: "Simple Analytics",
    logo: "/images/logos/plugins/simpleanalytics.svg",
    href: "https://github.com/simpleanalytics/cq-source-simpleanalytics",
    id: "simpleanalytics",
  },
  {
    name: "Slack",
    logo: "/images/logos/plugins/slack.svg",
    id: "slack",
  },
  {
    name: "Snyk",
    logo: "/images/logos/plugins/snyk.svg",
    id: "snyk",
  },
  {
    name: "Stripe",
    logo: "/images/logos/plugins/stripe.svg",
    id: "stripe",
  },
  {
    name: "Tailscale",
    logo: "/images/logos/plugins/tailscale.svg",
    logoDark: "/images/logos/plugins/tailscale-dark.svg",
    id: "tailscale",
  },
  {
    name: "Terraform",
    logo: "/images/logos/plugins/terraform.svg",
    id: "terraform",
  },
  {
    name: "Vercel",
    logo: "/images/logos/plugins/vercel.svg",
    logoDark: "/images/logos/plugins/vercel-dark.svg",
    id: "vercel",
  },
  {
    name: "Yandex Cloud",
    logo: "/images/logos/plugins/yandex.svg",
    id: "yandexcloud",
    href: "https://github.com/yandex-cloud/cq-source-yandex",
  },
  {
    name: "All Plugins",
    logo: "/images/logos/plugins/more.svg",
    logoDark: "/images/logos/plugins/more-dark.svg",
    id: "more",
    href: "/docs/plugins/overview",
  },
];

export const DESTINATION_PLUGINS: Plugin[] = [
  {
    name: "Azure Blob Storage",
    logo: "/images/logos/plugins/azblob.svg",
    id: "azblob",
  },
  {
    name: "BigQuery",
    logo: "/images/logos/plugins/bigquery.svg",
    id: "bigquery",
  },
  {
    name: "ClickHouse",
    logo: "/images/logos/plugins/clickhouse.svg",
    id: "clickhouse",
  },
  {
    name: "Elasticsearch",
    logo: "/images/logos/plugins/elasticsearch.svg",
    id: "elasticsearch",
  },
  {
    name: "File",
    logo: "/images/logos/plugins/file.svg",
    id: "file",
  },
  {
    name: "GCS",
    logo: "/images/logos/plugins/gcs.svg",
    id: "gcs",
  },
  {
    name: "Kafka",
    logo: "/images/logos/plugins/kafka.svg",
    logoDark: "/images/logos/plugins/kafka-dark.svg",
    id: "kafka",
  },
  {
    name: "Microsoft SQL Server",
    logo: "/images/logos/plugins/mssql.svg",
    id: "mssql",
  },
  {
    name: "MySQL",
    logo: "/images/logos/plugins/mysql.svg",
    logoDark: "/images/logos/plugins/mysql-dark.svg",
    id: "mysql",
  },
  {
    name: "MongoDB",
    logo: "/images/logos/plugins/mongodb.svg",
    id: "mongodb",
  },
  {
    name: "Neo4j",
    logo: "/images/logos/plugins/neo4j.svg",
    id: "neo4j",
  },
  {
    name: "PostgreSQL",
    logo: "/images/logos/plugins/postgresql.svg",
    id: "postgresql",
  },
  {
    name: "S3",
    logo: "/images/logos/plugins/s3.svg",
    id: "s3",
  },
  {
    name: "Snowflake",
    logo: "/images/logos/plugins/snowflake.svg",
    id: "snowflake",
  },
  {
    name: "SQLite",
    logo: "/images/logos/plugins/sqlite.svg",
    id: "sqlite",
  },
];

export function getPlugin(kind: string, id: string): Plugin {
  if (kind === "destination") {
    return DESTINATION_PLUGINS.find((p) => p.id === id);
  }
  return SOURCE_PLUGINS.find((p) => p.id === id);
}