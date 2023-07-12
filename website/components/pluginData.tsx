import {Category} from "./Category";
import {Kind} from "./Kind";

export type Availability = "free" | "premium" | "partner" | "community" | "unpublished";

export type Plugin = {
  id: string;
  name: string;
  href?: string; // external link to plugin
  website?: string; // related website for the API, if any
  description?: string;
  logo?: string;
  logoDark?: string;
  kind: Kind;
  availability: Availability;
  category: Category;
  buyLinks?: any;
};

export const ALL_PLUGINS: Plugin[] = [
  {
    name: "Alibaba Cloud",
    logo: "/images/logos/plugins/alicloud.svg",
    logoDark: "/images/logos/plugins/alicloud-dark.svg",
    id: "alicloud",
    kind: "source",
    availability: "free",
    category: "cloud-infrastructure",
  },
  {
    name: "Atlassian Jira",
    id: "atlassian-jira",
    kind: "source",
    availability: "premium",
    category: "engineering-analytics",
    logo: "/images/logos/plugins/jira.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/cN29AE5HX8px5ig4i7',
      'extended': 'https://buy.stripe.com/14k0048U935ddOMcOE'
    },
  },
  {
    name: "AWS",
    logo: "/images/logos/plugins/aws.svg",
    logoDark: "/images/logos/plugins/aws-dark.svg",
    id: "aws",
    kind: "source",
    availability: "free",
    category: "cloud-infrastructure",
  },
  {
    name: "AWS Pricing",
    logo: "/images/logos/plugins/aws.svg",
    logoDark: "/images/logos/plugins/aws-dark.svg",
    id: "awspricing",
    kind: "source",
    availability: "free",
    category: "cloud-finops",
    website: "https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/price-changes.html"
  },
  {
    name: "Azure",
    logo: "/images/logos/plugins/azure.svg",
    id: "azure",
    kind: "source",
    availability: "free",
    category: "cloud-infrastructure",
  },
  {
    name: "Azure DevOps",
    logo: "/images/logos/plugins/azuredevops.svg",
    id: "azuredevops",
    kind: "source",
    availability: "free",
    category: "cloud-infrastructure",
  },
  {
    name: "Cloud Firestore",
    logo: "/images/logos/plugins/firestore-light.svg",
    logoDark: "/images/logos/plugins/firestore-dark.svg",
    id: "firestore",
    kind: "source",
    availability: "free",
    category: "databases",
  },
  {
    name: "Cloudflare",
    logo: "/images/logos/plugins/cloudflare.svg",
    id: "cloudflare",
    kind: "source",
    availability: "free",
    category: "cloud-infrastructure",
  },
  {
    name: "Coinpaprika",
    logo: "/images/logos/plugins/coinpaprika.svg",
    id: "coinpaprika",
    href: "https://github.com/coinpaprika/cq-source-coinpaprika",
    kind: "source",
    availability: "partner",
    category: "other",
  },
  {
    name: "Crowdstrike",
    logo: "/images/logos/plugins/crowdstrike.svg",
    id: "crowdstrike",
    href: "https://github.com/justmiles/cq-source-crowdstrike",
    kind: "source",
    availability: "community",
    category: "security",
  },
  {
    name: "Datadog",
    logo: "/images/logos/plugins/datadog.svg",
    logoDark: "/images/logos/plugins/datadog-dark.svg",
    id: "datadog",
    kind: "source",
    availability: "free",
    category: "engineering-analytics",
  },
  {
    name: "Digital Ocean",
    logo: "/images/logos/plugins/digitalocean.svg",
    id: "digitalocean",
    kind: "source",
    availability: "free",
    category: "cloud-infrastructure",
  },
  {
    name: "Facebook Marketing",
    logo: "/images/logos/plugins/meta.svg",
    id: "facebookmarketing",
    kind: "source",
    availability: "free",
    category: "marketing-analytics",
  },
  {
    name: "Fastly",
    logo: "/images/logos/plugins/fastly.svg",
    id: "fastly",
    kind: "source",
    availability: "free",
    category: "cloud-infrastructure",
  },
  {
    name: "Gandi",
    logo: "/images/logos/plugins/gandi.svg",
    logoDark: "/images/logos/plugins/gandi-dark.svg",
    id: "gandi",
    kind: "source",
    availability: "premium",
    category: "cloud-infrastructure",
    buyLinks: {
      standard: "https://buy.stripe.com/00g6oseeteNVbGE03A",
      extended: "https://buy.stripe.com/eVa9AE2vL2199yw9E9",
    }
  },
  {
    name: "GitHub",
    logo: "/images/logos/plugins/github.svg",
    logoDark: "/images/logos/plugins/github-dark.svg",
    id: "github",
    kind: "source",
    availability: "free",
    category: "engineering-analytics",
  },
  {
    name: "GitLab",
    logo: "/images/logos/plugins/gitlab.svg",
    id: "gitlab",
    kind: "source",
    availability: "free",
    category: "engineering-analytics",
  },
  {
    name: "Google Ads",
    id: "googleads",
    kind: "source",
    availability: "premium",
    category: "marketing-analytics",
    logo: "/images/logos/plugins/ga2.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/6oE3cg9Yd6hp3a801T',
      'extended': 'https://buy.stripe.com/4gweUY8U99tB3a8g0S'
    },
  },
  {
    name: "Google Analytics",
    logo: "/images/logos/plugins/ga.svg",
    id: "googleanalytics",
    kind: "source",
    availability: "free",
    category: "marketing-analytics",
  },
  {
    name: "Google Cloud Platform",
    logo: "/images/logos/plugins/gcp.svg",
    id: "gcp",
    kind: "source",
    availability: "free",
    category: "cloud-infrastructure",
  },
  {
    name: "Hacker News API",
    logo: "/images/logos/plugins/hackernews.svg",
    website: "https://github.com/HackerNews/API",
    id: "hackernews",
    kind: "source",
    availability: "free",
    category: "other",
  },
  {
    name: "Heroku",
    logo: "/images/logos/plugins/heroku.svg",
    id: "heroku",
    kind: "source",
    availability: "premium",
    category: "cloud-infrastructure",
    buyLinks: {
      'standard': 'https://buy.stripe.com/6oE004eetgW3cKIeYf',
      'extended': 'https://buy.stripe.com/eVa1485HX49h2646rK'
    },
  },
  {
    name: "Homebrew",
    logo: "/images/logos/plugins/homebrew.svg",
    id: "homebrew",
    kind: "source",
    availability: "free",
    category: "marketing-analytics",
  },
  {
    name: "HubSpot",
    logo: "/images/logos/plugins/hubspot.svg",
    id: "hubspot",
    kind: "source",
    availability: "free",
    category: "marketing-analytics",
  },
  {
    name: "Kubernetes",
    logo: "/images/logos/plugins/kubernetes.svg",
    id: "k8s",
    kind: "source",
    availability: "free",
    category: "cloud-infrastructure",
  },
  {
    name: "LaunchDarkly",
    logo: "/images/logos/plugins/launchdarkly.svg",
    logoDark: "/images/logos/plugins/launchdarkly-dark.svg",
    id: "launchdarkly",
    kind: "source",
    availability: "premium",
    category: "engineering-analytics",
    buyLinks: {
      'standard': 'https://buy.stripe.com/3cs1480nD219fWUaI1',
      'extended': 'https://buy.stripe.com/28o4gkgmBgW3cKI7vQ'
    },
  },
  {
    name: "Mixpanel",
    logo: "/images/logos/plugins/mixpanel.svg",
    logoDark: "/images/logos/plugins/mixpanel-dark.svg",
    id: "mixpanel",
    kind: "source",
    availability: "premium",
    category: "marketing-analytics",
    buyLinks: {
      'standard': 'https://buy.stripe.com/4gwfZ29Yd7ltaCA3fB',
      'extended': 'https://buy.stripe.com/4gw28cb2haxF8usaI4'
    },
  },
  {
    name: "MySQL",
    logo: "/images/logos/plugins/mysql.svg",
    logoDark: "/images/logos/plugins/mysql-dark.svg",
    id: "mysql",
    kind: "source",
    availability: "free",
    category: "databases",
  },
  {
    name: "Okta",
    logo: "/images/logos/plugins/okta.svg",
    id: "okta",
    kind: "source",
    availability: "free",
    category: "cloud-infrastructure",
  },
  {
    name: "Onfleet",
    logo: "/images/logos/plugins/onfleet.svg",
    id: "onfleet",
    kind: "source",
    href: "https://github.com/onfleet/cq-source-onfleet",
    availability: "partner",
    category: "fleet-management",
  },
  {
    name: "Oracle",
    logo: "/images/logos/plugins/oracle.svg",
    id: "oracle",
    kind: "source",
    availability: "free",
    category: "cloud-infrastructure",
  },
  {
    name: "Oracle Database",
    logo: "/images/logos/plugins/oracle.svg",
    id: "oracledb",
    kind: "source",
    availability: "free",
    category: "databases",
  },
  {
    name: "Pagerduty",
    logo: "/images/logos/plugins/pagerduty.svg",
    id: "pagerduty",
    kind: "source",
    availability: "free",
    category: "engineering-analytics",
  },
  {
    name: "Plausible Analytics",
    logo: "/images/logos/plugins/plausibleanalytics.svg",
    id: "plausible",
    kind: "source",
    availability: "premium",
    category: "marketing-analytics",
    buyLinks: {
      'standard': 'https://buy.stripe.com/cN28wA9YdfRZ9ywcQj',
      'extended': 'https://buy.stripe.com/00g28cb2hfRZ6mkg2w'
    },
  },
  {
    name: "PostgreSQL",
    logo: "/images/logos/plugins/postgresql.svg",
    id: "postgresql",
    kind: "source",
    availability: "free",
    category: "databases",
  },
  {
    name: "Salesforce",
    logo: "/images/logos/plugins/salesforce.svg",
    id: "salesforce",
    kind: "source",
    availability: "free",
    category: "marketing-analytics",
  },
  {
    name: "SharePoint",
    logo: "/images/logos/plugins/sharepoint.svg",
    href: "https://github.com/koltyakov/cq-source-sharepoint",
    id: "sharepoint",
    kind: "source",
    availability: "community",
    category: "cloud-infrastructure",
  },
  {
    name: "Swetrix",
    logo: "/images/logos/plugins/swetrix.svg",
    href: "https://github.com/swetrix/cq-source-swetrix",
    id: "swetrix",
    kind: "source",
    availability: "partner",
    category: "cloud-infrastructure",
  },
  {
    name: "Scaleway",
    logo: "/images/logos/plugins/scaleway.svg",
    href: "https://github.com/scaleway/cq-source-scaleway",
    id: "scaleway",
    kind: "source",
    availability: "partner",
    category: "cloud-infrastructure",
  },
  {
    name: "Shopify",
    logo: "/images/logos/plugins/shopify.svg",
    id: "shopify",
    kind: "source",
    availability: "free",
    category: "marketing-analytics",
  },
  {
    name: "Simple Analytics",
    logo: "/images/logos/plugins/simpleanalytics.svg",
    href: "https://github.com/simpleanalytics/cq-source-simpleanalytics",
    id: "simpleanalytics",
    kind: "source",
    availability: "partner",
    category: "marketing-analytics",
  },
  {
    name: "Slack",
    logo: "/images/logos/plugins/slack.svg",
    id: "slack",
    kind: "source",
    availability: "premium",
    category: "engineering-analytics",
    buyLinks: {
      'standard': 'https://buy.stripe.com/6oE28c1rHfRZaCA03r',
      'extended': 'https://buy.stripe.com/7sIaEI8U935dfWU8zY'
    },
  },
  {
    name: "Snyk",
    logo: "/images/logos/plugins/snyk.svg",
    id: "snyk",
    kind: "source",
    availability: "free",
    category: "security",
  },
  {
    name: "Stripe",
    logo: "/images/logos/plugins/stripe.svg",
    id: "stripe",
    kind: "source",
    availability: "free",
    category: "cloud-finops",
  },
  {
    name: "Tailscale",
    logo: "/images/logos/plugins/tailscale.svg",
    logoDark: "/images/logos/plugins/tailscale-dark.svg",
    id: "tailscale",
    kind: "source",
    availability: "free",
    category: "cloud-infrastructure",
    buyLinks: {
      'standard': 'https://buy.stripe.com/14k8wAdap49hbGE6rR',
      'extended': 'https://buy.stripe.com/aEU1482vL6hpcKI17y'
    },
  },
  {
    name: "Terraform",
    logo: "/images/logos/plugins/terraform.svg",
    id: "terraform",
    kind: "source",
    availability: "free",
    category: "cloud-infrastructure",
  },
  {
    name: "Vercel",
    logo: "/images/logos/plugins/vercel.svg",
    logoDark: "/images/logos/plugins/vercel-dark.svg",
    id: "vercel",
    kind: "source",
    availability: "free",
    category: "cloud-infrastructure",
    buyLinks: {
      'standard': 'https://buy.stripe.com/bIY7swc6lcFN9yw5nP',
      'extended': 'https://buy.stripe.com/5kA6oseet0X55ig8A2'
    },
  },
  {
    name: "Yandex Cloud",
    logo: "/images/logos/plugins/yandex.svg",
    id: "yandexcloud",
    kind: "source",
    href: "https://github.com/yandex-cloud/cq-source-yandex",
    availability: "partner",
    category: "cloud-infrastructure",
  },
  {
    name: "Azure Blob Storage",
    logo: "/images/logos/plugins/azblob.svg",
    id: "azblob",
    kind: "destination",
    availability: "free",
    category: "data-warehouses-lakes",
  },
  {
    name: "BigQuery",
    logo: "/images/logos/plugins/bigquery.svg",
    id: "bigquery",
    kind: "destination",
    availability: "free",
    category: "data-warehouses-lakes",
  },
  {
    name: "ClickHouse",
    logo: "/images/logos/plugins/clickhouse.svg",
    id: "clickhouse",
    kind: "destination",
    availability: "free",
    category: "data-warehouses-lakes",
  },
  {
    name: "DuckDB",
    logo: "/images/logos/plugins/duckdb.svg",
    logoDark: "/images/logos/plugins/duckdb-dark.svg",
    id: "duckdb",
    kind: "destination",
    availability: "free",
    category: "data-warehouses-lakes",
  },
  {
    name: "Elasticsearch",
    logo: "/images/logos/plugins/elasticsearch.svg",
    id: "elasticsearch",
    kind: "destination",
    availability: "free",
    category: "databases",
  },
  {
    name: "File",
    logo: "/images/logos/plugins/file.svg",
    id: "file",
    kind: "destination",
    availability: "free",
    category: "data-warehouses-lakes",
  },
  {
    name: "Firehose",
    logo: "/images/logos/plugins/s3.svg",
    logoDark: "/images/logos/plugins/s3-dark.svg",
    id: "firehose",
    kind: "destination",
    availability: "premium",
    category: "data-warehouses-lakes",
    buyLinks: {
      'standard': 'https://buy.stripe.com/9AQ4gk1rH8pxbGE4jz',
      'extended': 'https://buy.stripe.com/cN26os5HX6hpcKI7vM'
    },
  },
  {
    name: "GCS",
    logo: "/images/logos/plugins/gcs.svg",
    id: "gcs",
    kind: "destination",
    availability: "free",
    category: "data-warehouses-lakes",
  },
  {
    name: "Gremlin",
    logo: "/images/logos/plugins/gremlin.svg",
    id: "gremlin",
    kind: "destination",
    availability: "free",
    category: "databases",
  },
  {
    name: "Kafka",
    logo: "/images/logos/plugins/kafka.svg",
    logoDark: "/images/logos/plugins/kafka-dark.svg",
    id: "kafka",
    kind: "destination",
    availability: "free",
    category: "databases",
  },
  {
    name: "Meilisearch",
    logo: "/images/logos/plugins/meilisearch.svg",
    id: "meilisearch",
    kind: "destination",
    availability: "free",
    category: "databases",
  },
  {
    name: "Microsoft SQL Server",
    logo: "/images/logos/plugins/mssql.svg",
    id: "mssql",
    kind: "destination",
    availability: "free",
    category: "databases",
  },
  {
    name: "MySQL",
    logo: "/images/logos/plugins/mysql.svg",
    logoDark: "/images/logos/plugins/mysql-dark.svg",
    id: "mysql",
    kind: "destination",
    availability: "free",
    category: "databases",
  },
  {
    name: "MongoDB",
    logo: "/images/logos/plugins/mongodb.svg",
    id: "mongodb",
    kind: "destination",
    availability: "free",
    category: "databases",
  },
  {
    name: "Neo4j",
    logo: "/images/logos/plugins/neo4j.svg",
    id: "neo4j",
    kind: "destination",
    availability: "free",
    category: "databases",
  },
  {
    name: "PostgreSQL",
    logo: "/images/logos/plugins/postgresql.svg",
    id: "postgresql",
    kind: "destination",
    availability: "free",
    category: "databases",
  },
  {
    name: "S3",
    logo: "/images/logos/plugins/s3.svg",
    id: "s3",
    kind: "destination",
    availability: "free",
    category: "data-warehouses-lakes",
  },
  {
    name: "Snowflake",
    logo: "/images/logos/plugins/snowflake.svg",
    id: "snowflake",
    kind: "destination",
    availability: "free",
    category: "data-warehouses-lakes",
  },
  {
    name: "SQLite",
    logo: "/images/logos/plugins/sqlite.svg",
    id: "sqlite",
    kind: "destination",
    availability: "free",
    category: "databases",
  },
  {
    name: "XKCD",
    logo: "/images/logos/plugins/xkcd.svg",
    logoDark: "/images/logos/plugins/xkcd-dark.svg",
    id: "xkcd",
    href: "https://github.com/hermanschaaf/cq-source-xkcd",
    kind: "source",
    availability: "community",
    category: "other",
  },
  {
    name: "Airtable",
    id: "airtable",
    href: "https://github.com/cloudquery/cloudquery/issues/8869",
    kind: "source",
    availability: "unpublished",
    category: "project-management",
    logo: "/images/logos/plugins/airtable.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/14kbIMeetbBJ6mk14r',
      'extended': 'https://buy.stripe.com/dR61481rHbBJ2646s0'
    },
  },
  {
    name: "AfterShip",
    id: "aftership",
    href: "https://github.com/cloudquery/cloudquery/issues/9047",
    kind: "source",
    availability: "unpublished",
    category: "shipment-tracking",
    logo: "/images/logos/plugins/aftership.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/6oE6os9Yd9tBdOMcNd',
      'extended': 'https://buy.stripe.com/7sIaEI8U9gW39yw17H'
    },
  },
  {
    name: "Amazon Ads",
    id: "amazonads",
    href: "https://github.com/cloudquery/cloudquery/issues/8424",
    kind: "source",
    availability: "unpublished",
    category: "marketing-analytics",
    logo: "/images/logos/plugins/amazon.svg",
    logoDark: "/images/logos/plugins/amazon-dark.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/00g28cgmBfRZfWU9B2',
      'extended': 'https://buy.stripe.com/cN2cMQ0nDbBJ7qo8wZ'
    },
  },
  {
    name: "Amplitude",
    id: "amplitude",
    href: "https://github.com/cloudquery/cloudquery/issues/8423",
    kind: "source",
    availability: "unpublished",
    category: "product-analytics",
    logo: "/images/logos/plugins/amplitude.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/fZe28c3zPaxF26428E',
      'extended': 'https://buy.stripe.com/00g4gk2vL0X53a8dUr'
    },
  },
  {
    name: "Bamboo HR",
    id: "bamboo-hr",
    href: "https://github.com/cloudquery/cloudquery/issues/8426",
    kind: "source",
    availability: "unpublished",
    category: "other",
    logo: "/images/logos/plugins/bamboohr.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/3cs4gk9Yd0X53a8dRn',
      'extended': 'https://buy.stripe.com/bIYbIM1rH5dlfWUbJg'
    },
  },
  {
    name: "Baremetrics",
    id: "baremetrics",
    href: "https://github.com/cloudquery/cloudquery/issues/9045",
    kind: "source",
    availability: "unpublished",
    category: "marketing-analytics",
    logo: "/images/logos/plugins/baremetrics.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/6oEaEI4DTaxF3a89B9',
      'extended': 'https://buy.stripe.com/dR66osgmB6hp4eccNm'
    },
  },
  {
    name: "Bing Ads",
    id: "bing-ads",
    href: "https://github.com/cloudquery/cloudquery/issues/8425",
    kind: "source",
    availability: "unpublished",
    category: "marketing-analytics",
    logo: "/images/logos/plugins/bingads.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/5kA5koc6l8pxdOM5kV',
      'extended': 'https://buy.stripe.com/5kA28cc6lgW34ec28K'
    },
  },
  {
    name: "BitBucket",
    id: "bitbucket",
    href: "https://github.com/cloudquery/cloudquery/issues/5510",
    kind: "source",
    availability: "unpublished",
    category: "engineering-analytics",
    logo: "/images/logos/plugins/bitbucket.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/28o8wA1rHdJR9yw3cP',
      'extended': 'https://buy.stripe.com/bIY9AE3zP2193a83cQ'
    },
  },
  {
    name: "Chargebee",
    id: "chargebee",
    href: "https://github.com/cloudquery/cloudquery/issues/9048",
    kind: "source",
    availability: "unpublished",
    category: "cloud-finops",
    logo: "/images/logos/plugins/chargebee.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/9AQ4gkfix49hh0Y00F',
      'extended': 'https://buy.stripe.com/5kA8wA7Q50X58usaFk'
    },
  },
  {
    name: "Cloudinary",
    id: "cloudinary",
    href: "https://github.com/cloudquery/cloudquery/issues/9044",
    kind: "source",
    availability: "unpublished",
    category: "cloud-infrastructure",
    logo: "/images/logos/plugins/cloudinary.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/5kA9AE4DTgW3dOM7t9',
      'extended': 'https://buy.stripe.com/00g8wA4DTfRZ8us9Bi'
    },
  },
  {
    name: "Detrack",
    id: "detrack",
    href: "https://github.com/cloudquery/cloudquery/issues/9054",
    kind: "source",
    availability: "unpublished",
    category: "shipment-tracking",
    buyLinks: {
      'standard': 'https://buy.stripe.com/28obIMb2heNV5ig3cV',
      'extended': 'https://buy.stripe.com/8wM28cc6lcFN9yw14O'
    },
  },
  {
    name: "Gmail",
    id: "gmail",
    href: "https://github.com/cloudquery/cloudquery/issues/8135",
    kind: "source",
    availability: "unpublished",
    category: "marketing-analytics",
    buyLinks: {
      'standard': 'https://buy.stripe.com/7sI7sw7Q52199yw7td',
      'extended': 'https://buy.stripe.com/aEUcMQfixfRZ264cNy'
    },
  },
  {
    name: "Google Sheets",
    id: "google-sheets",
    href: "https://github.com/cloudquery/cloudquery/issues/5190",
    kind: "source",
    availability: "unpublished",
    category: "other",
    logo: "/images/logos/plugins/google-sheets.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/14kfZ25HXeNVbGEaFr',
      'extended': 'https://buy.stripe.com/4gw148eet7lt6mkbJw'
    },
  },
  {
    name: "Hashicorp Vault",
    id: "hashicorp-vault",
    href: "https://github.com/cloudquery/cloudquery/issues/6738",
    kind: "source",
    availability: "unpublished",
    category: "cloud-infrastructure",
    logo: "/images/logos/plugins/hashicorp-vault.svg",
    logoDark: "/images/logos/plugins/hashicorp-vault-dark.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/bIY5koeetfRZh0Y14T',
      'extended': 'https://buy.stripe.com/7sIcMQc6lbBJ6mkcNC'
    },
  },
  {
    name: "Infoblox",
    id: "infoblox",
    href: "https://github.com/cloudquery/cloudquery/issues/8383",
    kind: "source",
    availability: "unpublished",
    category: "cloud-infrastructure",
    logo: "/images/logos/plugins/infoblox.svg",
    logoDark: "/images/logos/plugins/infoblox-dark.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/28oaEIb2hcFN5ig28Z',
      'extended': 'https://buy.stripe.com/4gwaEIdap9tB5ig14W'
    },
  },
  {
    name: "Intercom",
    id: "intercom",
    href: "https://github.com/cloudquery/cloudquery/issues/9041",
    kind: "source",
    availability: "unpublished",
    category: "cloud-infrastructure",
    logo: "/images/logos/plugins/intercom.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/4gw6os4DT5dl9yw9Bt',
      'extended': 'https://buy.stripe.com/3cseUYeet7ltaCAeVO'
    },
  },
  {
    name: "LinkedIn Ads",
    id: "linkedin-ads",
    href: "https://github.com/cloudquery/cloudquery/issues/9033",
    kind: "source",
    availability: "unpublished",
    category: "marketing-analytics",
    logo: "/images/logos/plugins/linkedin.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/3cs1480nDdJReSQ293',
      'extended': 'https://buy.stripe.com/3csfZ25HXcFN6mk4hc'
    },
  },
  {
    name: "Mailchimp",
    id: "mailchimp",
    href: "https://github.com/cloudquery/cloudquery/issues/8430",
    kind: "source",
    availability: "unpublished",
    category: "marketing-analytics",
    logo: "/images/logos/plugins/mailchimp.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/cN2eUY8U9axF120dRN',
      'extended': 'https://buy.stripe.com/4gw3cgc6lfRZh0Y296'
    },
  },
  {
    name: "Mailgun",
    id: "mailgun",
    href: "https://github.com/cloudquery/cloudquery/issues/9050",
    kind: "source",
    availability: "unpublished",
    category: "marketing-analytics",
    logo: "/images/logos/plugins/mailgun.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/4gw28c7Q535d264cNL',
      'extended': 'https://buy.stripe.com/cN24gkdapbBJ264fZY'
    },
  },
  {
    name: "Marketo",
    id: "marketo",
    href: "https://github.com/cloudquery/cloudquery/issues/8428",
    kind: "source",
    availability: "unpublished",
    category: "marketing-analytics",
    logo: "/images/logos/plugins/marketo.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/cN29AE0nD49h9yweVV',
      'extended': 'https://buy.stripe.com/4gwcMQ2vLeNVcKI156'
    },
  },
  {
    name: "Monday",
    id: "monday",
    href: "https://github.com/cloudquery/cloudquery/issues/8431",
    kind: "source",
    availability: "unpublished",
    category: "project-management",
    buyLinks: {
      'standard': 'https://buy.stripe.com/5kA1483zPcFNfWUg01',
      'extended': 'https://buy.stripe.com/9AQ0041rHbBJdOMg02'
    },
  },
  {
    name: "MongoDB Atlas",
    id: "mongodb-atlas",
    href: "https://github.com/cloudquery/cloudquery/issues/9134",
    kind: "source",
    availability: "unpublished",
    category: "databases",
    logo: "/images/logos/plugins/mongodb.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/3cs28c1rH6hpaCA4hl',
      'extended': 'https://buy.stripe.com/28oeUYeet219bGE7ty'
    },
  },
  {
    name: "Microsoft SQL Server",
    id: "mssql",
    href: "https://github.com/cloudquery/cloudquery/issues/8861",
    kind: "source",
    availability: "unpublished",
    category: "databases",
    logo: "/images/logos/plugins/mssql.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/bIYbIMb2h7lteSQ8xD',
      'extended': 'https://buy.stripe.com/14k3cg6M135d4ecg06'
    },
  },
  {
    name: "New Relic",
    id: "new-relic",
    href: "https://github.com/cloudquery/cloudquery/issues/9040",
    kind: "source",
    availability: "unpublished",
    category: "product-analytics",
    buyLinks: {
      'standard': 'https://buy.stripe.com/fZe7sw0nDaxF8us7tB',
      'extended': 'https://buy.stripe.com/4gw6os7Q55dlbGE3dm'
    },
  },
  {
    name: "Paypal",
    id: "paypal",
    href: "https://github.com/cloudquery/cloudquery/issues/6781",
    kind: "source",
    availability: "unpublished",
    category: "cloud-finops",
    logo: "/images/logos/plugins/paypal.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/4gwcMQfixeNV6mkaFP',
      'extended': 'https://buy.stripe.com/fZebIM4DT219h0YcNY'
    },
  },
  {
    name: "Pendo",
    id: "pendo",
    href: "https://github.com/cloudquery/cloudquery/issues/9360",
    kind: "source",
    availability: "unpublished",
    category: "marketing-analytics",
    buyLinks: {
      'standard': 'https://buy.stripe.com/dR6cMQ2vLgW3cKI6pB',
      'extended': 'https://buy.stripe.com/aEUbIM9Yd9tBh0Y29m'
    },
  },
  {
    name: "Prisma",
    id: "prisma",
    href: "https://github.com/cloudquery/cloudquery/issues/6582",
    kind: "source",
    availability: "unpublished",
    category: "cloud-infrastructure",
    logo: "/images/logos/plugins/prisma.svg",
    logoDark: "/images/logos/plugins/prisma-dark.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/9AQ9AE4DT6hp5ig5lz',
      'extended': 'https://buy.stripe.com/eVafZ24DT0X526415k'
    },
  },
  {
    name: "Render",
    id: "render",
    website: "https://render.com/",
    kind: "source",
    availability: "premium",
    category: "cloud-infrastructure",
    buyLinks: {
      'standard': 'https://buy.stripe.com/8wMdQUeet5dlcKI16x',
      'extended': 'https://buy.stripe.com/eVa004dapdJR7qo5mO'
    },
  },
  {
    name: "Reddit Advertising",
    id: "reddit-advertising",
    href: "https://github.com/cloudquery/cloudquery/issues/8684",
    kind: "source",
    availability: "unpublished",
    category: "marketing-analytics",
    logo: "/images/logos/plugins/reddit.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/6oEbIMb2hcFNaCA9BT',
      'extended': 'https://buy.stripe.com/eVaaEI2vLgW3120cO6'
    },
  },
  {
    name: "Samsara",
    id: "samsara",
    href: "https://github.com/cloudquery/cloudquery/issues/9052",
    kind: "source",
    availability: "unpublished",
    category: "fleet-management",
    buyLinks: {
      'standard': 'https://buy.stripe.com/aEU3cgdap219bGEbK3',
      'extended': 'https://buy.stripe.com/bIY7sweetaxF2649BW'
    },
  },
  {
    name: "SentinelOne",
    id: "sentinel-one",
    href: "https://github.com/cloudquery/cloudquery/issues/9136",
    kind: "source",
    availability: "unpublished",
    category: "security",
    buyLinks: {
      'standard': 'https://buy.stripe.com/00gaEIgmB9tBdOM29v',
      'extended': 'https://buy.stripe.com/cN2cMQ5HXeNVfWUg0m'
    },
  },
  {
    name: "Shippo",
    id: "shippo",
    href: "https://github.com/cloudquery/cloudquery/issues/9038",
    kind: "source",
    availability: "unpublished",
    category: "shipment-tracking",
    logo: "/images/logos/plugins/shippo.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/5kA0043zP49hdOM7tR',
      'extended': 'https://buy.stripe.com/cN2cMQb2h6hp9yweWk'
    },
  },
  {
    name: "Shipup",
    id: "shipup",
    href: "https://github.com/cloudquery/cloudquery/issues/9055",
    kind: "source",
    availability: "unpublished",
    category: "shipment-tracking",
    buyLinks: {
      'standard': 'https://buy.stripe.com/bIYaEIfix9tB8useWl',
      'extended': 'https://buy.stripe.com/6oE8wA5HXcFN6mk29A'
    },
  },
  {
    name: "S3 Bucket",
    id: "s3",
    href: "https://github.com/cloudquery/cloudquery/issues/8320",
    kind: "source",
    availability: "unpublished",
    category: "data-warehouses-lakes",
    logo: "/images/logos/plugins/s3.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/00g3cg8U9cFN5igcOf',
      'extended': 'https://buy.stripe.com/cN24gk0nD8px2649C4'
    },
  },
  {
    name: "Snowflake Configuration",
    id: "snowflake-config",
    href: "https://github.com/cloudquery/cloudquery/issues/9135",
    kind: "source",
    availability: "unpublished",
    category: "databases",
    logo: "/images/logos/plugins/snowflake.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/28o8wA2vL219aCA5lP',
      'extended': 'https://buy.stripe.com/fZe3cg0nDaxF6mkdSm'
    },
  },
  {
    name: "Sonatype Nexus",
    id: "sonatype-nexus",
    href: "https://github.com/cloudquery/cloudquery/issues/7640",
    kind: "source",
    availability: "unpublished",
    category: "cloud-infrastructure",
    logo: "/images/logos/plugins/sonatype.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/cN2bIM2vL6hp4eceWr',
      'extended': 'https://buy.stripe.com/fZeaEIeet9tB8us8y4'
    },
  },
  {
    name: "Sophos",
    id: "sophos",
    href: "https://github.com/cloudquery/cloudquery/issues/9131",
    kind: "source",
    availability: "unpublished",
    category: "security",
    buyLinks: {
      'standard': 'https://buy.stripe.com/9AQ9AE7Q549h5ig15D',
      'extended': 'https://buy.stripe.com/9AQfZ26M1dJRaCA01A'
    },
  },
  {
    name: "Square",
    id: "square",
    href: "https://github.com/cloudquery/cloudquery/issues/9037",
    kind: "source",
    availability: "unpublished",
    category: "cloud-finops",
    logo: "/images/logos/plugins/square.svg",
    logoDark: "/images/logos/plugins/square-dark.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/5kA6osdapaxFcKI6pZ',
      'extended': 'https://buy.stripe.com/4gwcMQfixcFNeSQ4hS'
    },
  },
  {
    name: "Tenable",
    id: "tenable",
    href: "https://github.com/cloudquery/cloudquery/issues/9132",
    kind: "source",
    availability: "unpublished",
    category: "security",
    buyLinks: {
      'standard': 'https://buy.stripe.com/00g6os6M17ltaCA7u5',
      'extended': 'https://buy.stripe.com/cN2bIMfix6hp120cOq'
    },
  },
  {
    name: "Twilio",
    id: "twilio",
    href: "https://github.com/cloudquery/cloudquery/issues/9035",
    kind: "source",
    availability: "unpublished",
    category: "cloud-infrastructure",
    logo: "/images/logos/plugins/twilio.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/7sI5ko8U9dJRaCAaGj',
      'extended': 'https://buy.stripe.com/bIYaEI0nD49hh0Y7u8'
    },
  },
  {
    name: "Twilio Sendgrid",
    id: "sendgrid",
    href: "https://github.com/cloudquery/cloudquery/issues/9039",
    kind: "source",
    availability: "unpublished",
    category: "marketing-analytics",
    logo: "/images/logos/plugins/sendgrid.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/fZe28c5HXeNVfWUcOt',
      'extended': 'https://buy.stripe.com/eVa1486M1gW37qog0G'
    },
  },
  {
    name: "Typeform",
    id: "typeform",
    href: "https://github.com/cloudquery/cloudquery/issues/9034",
    kind: "source",
    availability: "unpublished",
    category: "marketing-analytics",
    logo: "/images/logos/plugins/typeform.svg",
    logoDark: "/images/logos/plugins/typeform-dark.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/28o3cg5HXbBJdOM3dV',
      'extended': 'https://buy.stripe.com/6oEeUYfixcFN8us01K'
    },
  },
  {
    name: "Zoho Campaign",
    id: "zoho-campaign",
    website: "https://www.zoho.com",
    href: "https://github.com/cloudquery/cloudquery/issues/9028",
    kind: "source",
    availability: "unpublished",
    category: "product-analytics",
    logo: "/images/logos/plugins/zoho.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/bIYcMQ0nD7lt1203dX',
      'extended': 'https://buy.stripe.com/cN2dQU9Yd8px1205m6'
    },
  },
  {
    name: "Zoho CRM",
    id: "zoho-crm",
    website: "https://www.zoho.com",
    href: "https://github.com/cloudquery/cloudquery/issues/9029",
    kind: "source",
    availability: "unpublished",
    category: "other",
    logo: "/images/logos/plugins/zoho.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/dR6dQUb2h5dldOM4i3',
      'extended': 'https://buy.stripe.com/8wM5ko6M18px8us8yk'
    },
  },
  {
    name: "Zoom",
    id: "zoom",
    href: "https://github.com/cloudquery/cloudquery/issues/1507",
    kind: "source",
    availability: "unpublished",
    category: "cloud-infrastructure",
    logo: "/images/logos/plugins/zoom.svg",
    buyLinks: {
      'standard': 'https://buy.stripe.com/00g5ko4DT8pxcKI9Cp',
      'extended': 'https://buy.stripe.com/00g3cgfix35d1205ma'
    },
  },
].sort((a, b) => a.name.localeCompare(b.name)) as Plugin[];

export const ALL_SOURCE_PLUGINS: Plugin[] = ALL_PLUGINS.filter((p) => p.kind === "source");
export const ALL_DESTINATION_PLUGINS: Plugin[] = ALL_PLUGINS.filter((p) => p.kind === "destination");

export const PUBLISHED_SOURCE_PLUGINS: Plugin[] = ALL_SOURCE_PLUGINS.filter(
    (p) => p.availability !== "unpublished"
);
export const UNPUBLISHED_SOURCE_PLUGINS: Plugin[] = ALL_SOURCE_PLUGINS.filter(
  (p) => p.availability === "unpublished"
);
export const PUBLISHED_DESTINATION_PLUGINS: Plugin[] = ALL_DESTINATION_PLUGINS.filter(
    (p) => p.availability !== "unpublished"
);
export const UNPUBLISHED_DESTINATION_PLUGINS: Plugin[] = ALL_DESTINATION_PLUGINS.filter(
    (p) => p.availability === "unpublished"
);

export const SOURCE_CATEGORIES: Category[] = ALL_SOURCE_PLUGINS.map((p) => p.category).filter(
    (c, pos) => ALL_SOURCE_PLUGINS.map((p) => p.category).indexOf(c) === pos
).sort() as Category[];

export const DESTINATION_CATEGORIES: Category[] = ALL_DESTINATION_PLUGINS.map((p) => p.category).filter(
    (c, pos) => ALL_DESTINATION_PLUGINS.map((p) => p.category).indexOf(c) === pos
).sort() as Category[];

export function getPlugin(kind: string, id: string): Plugin {
  if (kind === "destination") {
    return ALL_DESTINATION_PLUGINS.find((p) => p.id === id);
  }
  return ALL_SOURCE_PLUGINS.find((p) => p.id === id);
}
