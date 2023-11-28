In Order for CloudQuery to sync resources from your Notion setup, you will need to create a notion integration key and export the Token in `NOTION_SECRET_KEY` environment variable.
How to create the notion integration key? [see here](https://developers.notion.com/docs/create-a-notion-integration#create-your-integration-in-notion). Make sure to give proper **Content Capabilities** and **User Capabilities** from capabilities settings. Also give your integration page permissions [see here](https://developers.notion.com/docs/create-a-notion-integration#give-your-integration-page-permissions). Only pages and databases with permission will able to sync. 

```bash
export NOTION_SECRET_KEY=<your_notion_integration_key>
```