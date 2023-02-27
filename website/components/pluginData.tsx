export type Plugin = {
  id: string;
  name: string;
  href?: string;
  description?: string;
  logo: string;
  logoDark?: string;
  kind: "source" | "destination";
};

export const SOURCE_PLUGINS: Plugin[] = [
  {
    name: "Alibaba Cloud",
    logo: "/images/logos/plugins/alicloud.svg",
    logoDark: "/images/logos/plugins/alicloud-dark.svg",
    id: "alicloud",
    kind: "source",
  },
  {
    name: "AWS",
    logo: "/images/logos/plugins/aws.svg",
    logoDark: "/images/logos/plugins/aws-dark.svg",
    id: "aws",
    kind: "source",
  },
  {
    name: "Azure",
    logo: "/images/logos/plugins/azure.svg",
    id: "azure",
    kind: "source",
  },
  {
    name: "Azure DevOps",
    logo: "/images/logos/plugins/azuredevops.svg",
    id: "azuredevops",
    kind: "source",
  },
  {
    name: "Cloudflare",
    logo: "/images/logos/plugins/cloudflare.svg",
    id: "cloudflare",
    kind: "source",
  },
  {
    name: "Datadog",
    logo: "/images/logos/plugins/datadog.svg",
    logoDark: "/images/logos/plugins/datadog-dark.svg",
    id: "datadog",
    kind: "source",
  },
  {
    name: "Digital Ocean",
    logo: "/images/logos/plugins/digitalocean.svg",
    id: "digitalocean",
    kind: "source",
  },
  {
    name: "Fastly",
    logo: "/images/logos/plugins/fastly.svg",
    id: "fastly",
    kind: "source",
  },
  {
    name: "Gandi",
    logo: "/images/logos/plugins/gandi.svg",
    logoDark: "/images/logos/plugins/gandi-dark.svg",
    id: "gandi",
    kind: "source",
  },
  {
    name: "Google Cloud Platform",
    logo: "/images/logos/plugins/gcp.svg",
    id: "gcp",
    kind: "source",
  },
  {
    name: "GitHub",
    logo: "/images/logos/plugins/github.svg",
    logoDark: "/images/logos/plugins/github-dark.svg",
    id: "github",
    kind: "source",
  },
  {
    name: "GitLab",
    logo: "/images/logos/plugins/gitlab.svg",
    id: "gitlab",
    kind: "source",
  },
  {
    name: "Hacker News",
    logo: "/images/logos/plugins/hackernews.svg",
    id: "hackernews",
    kind: "source",
  },
  {
    name: "Heroku",
    logo: "/images/logos/plugins/heroku.svg",
    id: "heroku",
    kind: "source",
  },
  {
    name: "Homebrew",
    logo: "/images/logos/plugins/homebrew.svg",
    id: "homebrew",
    kind: "source",
  },
  {
    name: "HubSpot",
    logo: "/images/logos/plugins/hubspot.svg",
    id: "hubspot",
    kind: "source",
  },
  {
    name: "Kubernetes",
    logo: "/images/logos/plugins/kubernetes.svg",
    id: "k8s",
    kind: "source",
  },
  {
    name: "LaunchDarkly",
    logo: "/images/logos/plugins/launchdarkly.svg",
    logoDark: "/images/logos/plugins/launchdarkly-dark.svg",
    id: "launchdarkly",
    kind: "source",
  },
  {
    name: "Mixpanel",
    logo: "/images/logos/plugins/mixpanel.svg",
    logoDark: "/images/logos/plugins/mixpanel-dark.svg",
    id: "mixpanel",
    kind: "source",
  },
  {
    name: "Okta",
    logo: "/images/logos/plugins/okta.svg",
    id: "okta",
    kind: "source",
  },
  {
    name: "Oracle",
    logo: "/images/logos/plugins/oracle.svg",
    id: "oracle",
    kind: "source",
  },
  {
    name: "Pagerduty",
    logo: "/images/logos/plugins/pagerduty.svg",
    id: "pagerduty",
    kind: "source",
  },
  {
    name: "Plausible Analytics",
    logo: "/images/logos/plugins/plausibleanalytics.svg",
    id: "plausible",
    kind: "source",
  },
  {
    name: "PostgreSQL",
    logo: "/images/logos/plugins/postgresql.svg",
    id: "postgresql",
    kind: "source",
  },
  {
    name: "Salesforce",
    logo: "/images/logos/plugins/salesforce.svg",
    id: "salesforce",
    kind: "source",
  },
  {
    name: "Swetrix",
    logo: "/images/logos/plugins/swetrix.svg",
    href: "https://github.com/swetrix/cq-source-swetrix",
    id: "swetrix",
    kind: "source",
  },
  {
    name: "Scaleway",
    logo: "/images/logos/plugins/scaleway.svg",
    href: "https://github.com/scaleway/cq-source-scaleway",
    id: "scaleway",
    kind: "source",
  },
  {
    name: "Shopify",
    logo: "/images/logos/plugins/shopify.svg",
    id: "shopify",
    kind: "source",
  },
  {
    name: "Simple Analytics",
    logo: "/images/logos/plugins/simpleanalytics.svg",
    href: "https://github.com/simpleanalytics/cq-source-simpleanalytics",
    id: "simpleanalytics",
    kind: "source",
  },
  {
    name: "Slack",
    logo: "/images/logos/plugins/slack.svg",
    id: "slack",
    kind: "source",
  },
  {
    name: "Snyk",
    logo: "/images/logos/plugins/snyk.svg",
    id: "snyk",
    kind: "source",
  },
  {
    name: "Stripe",
    logo: "/images/logos/plugins/stripe.svg",
    id: "stripe",
    kind: "source",
  },
  {
    name: "Tailscale",
    logo: "/images/logos/plugins/tailscale.svg",
    logoDark: "/images/logos/plugins/tailscale-dark.svg",
    id: "tailscale",
    kind: "source",
  },
  {
    name: "Terraform",
    logo: "/images/logos/plugins/terraform.svg",
    id: "terraform",
    kind: "source",
  },
  {
    name: "Vercel",
    logo: "/images/logos/plugins/vercel.svg",
    logoDark: "/images/logos/plugins/vercel-dark.svg",
    id: "vercel",
    kind: "source",
  },
  {
    name: "Yandex Cloud",
    logo: "/images/logos/plugins/yandex.svg",
    id: "yandexcloud",
    href: "https://github.com/yandex-cloud/cq-source-yandex",
    kind: "source",
  },
  {
    name: "All Plugins",
    logo: "/images/logos/plugins/more.svg",
    logoDark: "/images/logos/plugins/more-dark.svg",
    id: "more",
    href: "/docs/plugins/overview",
    kind: "source",
  },
];

export const DESTINATION_PLUGINS: Plugin[] = [
  {
    name: "Azure Blob Storage",
    logo: "/images/logos/plugins/azblob.svg",
    id: "azblob",
    kind: "destination",
  },
  {
    name: "BigQuery",
    logo: "/images/logos/plugins/bigquery.svg",
    id: "bigquery",
    kind: "destination",
  },
  {
    name: "ClickHouse",
    logo: "/images/logos/plugins/clickhouse.svg",
    id: "clickhouse",
    kind: "destination",
  },
  {
    name: "Elasticsearch",
    logo: "/images/logos/plugins/elasticsearch.svg",
    id: "elasticsearch",
    kind: "destination",
  },
  {
    name: "File",
    logo: "/images/logos/plugins/file.svg",
    id: "file",
    kind: "destination",
  },
  {
    name: "GCS",
    logo: "/images/logos/plugins/gcs.svg",
    id: "gcs",
    kind: "destination",
  },
  {
    name: "Kafka",
    logo: "/images/logos/plugins/kafka.svg",
    id: "kafka",
    kind: "destination",
  },
  {
    name: "Microsoft SQL Server",
    logo: "/images/logos/plugins/mssql.svg",
    id: "mssql",
    kind: "destination",
  },
  {
    name: "MongoDB",
    logo: "/images/logos/plugins/mongodb.svg",
    id: "mongodb",
    kind: "destination",
  },
  {
    name: "Neo4j",
    logo: "/images/logos/plugins/neo4j.svg",
    id: "neo4j",
    kind: "destination",
  },
  {
    name: "PostgreSQL",
    logo: "/images/logos/plugins/postgresql.svg",
    id: "postgresql",
    kind: "destination",
  },
  {
    name: "S3",
    logo: "/images/logos/plugins/s3.svg",
    id: "s3",
    kind: "destination",
  },
  {
    name: "Snowflake",
    logo: "/images/logos/plugins/snowflake.svg",
    id: "snowflake",
    kind: "destination",
  },
  {
    name: "SQLite",
    logo: "/images/logos/plugins/sqlite.svg",
    id: "sqlite",
    kind: "destination",
  },
];

export const ALL_PLUGINS: Plugin[] = [...SOURCE_PLUGINS, ...DESTINATION_PLUGINS];

export function getPlugin(kind: string, id: string): Plugin {
  if (kind === "destination") {
    return DESTINATION_PLUGINS.find((p) => p.id === id);
  }
  return SOURCE_PLUGINS.find((p) => p.id === id);
}