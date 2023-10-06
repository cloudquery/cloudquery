# CloudQuery notion Source Plugin

A notion source plugin for CloudQuery that loads data from notion to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more.

## Links

 - [CloudQuery Quickstart Guide](https://www.cloudquery.io/docs/quickstart)
 - [Supported Tables](docs/tables/README.md)

## Authentication

In Order for CloudQuery to sync resources from your Notion setup, you will need to create a notion integration key and export the Token in NOTION_SECRET_KEY environment variable.
How to create the notion integration key? [see here](https://developers.notion.com/docs/create-a-notion-integration#create-your-integration-in-notion). Make sure to give proper **Content Capabilities** and **User Capabilities** from capabilities settings. Also give your integration page permissions [see here](https://developers.notion.com/docs/create-a-notion-integration#give-your-integration-page-permissions). Only pages and databases with permission will able to sync. 

```bash
export NOTION_SECRET_KEY=<your_notion_integration_key>
```

## Configuration

The following source configuration file will sync to a PostgreSQL database. See [the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more information on how to configure the source and destination.

```yaml
kind: source
spec:
  name: "notion"
  path: "cloudquery/notion"
  version: "${VERSION}"
  destinations:
    - "postgresql"
  spec:
    bearer_token: "${NOTION_SECRET_KEY}"
```


