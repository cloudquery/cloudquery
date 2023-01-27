# Slack Source Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("source", "slack")}/>

The CloudQuery Slack plugin extracts information from your Slack organization(s) and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

## Installation

### Step 1

Install a custom Slack app with read-only permissions into your workspace by clicking the button below and following the instructions:

import { getSlackAppLink } from "../../../../../utils/slack-app-link";

<div style={{marginTop: '1em'}}>
    <a target="_blank" href={getSlackAppLink()} class="btn btn-blue">Install App</a>
</div>

Once installed into the workspace, go to **Install App** (under **Settings**) and **copy the Bot User OAuth Token**. You will need this in your source config in the next step.

### Step 2

Set up a CloudQuery source config. See [Configuration](/docs/plugins/sources/slack/configuration) for an example and available options.

If this is your first time running CloudQuery, check out the [Quickstart](/docs/quickstart) page for help on setting up a destination and running your first sync. 

That's it, happy querying!

## Syncing message histories

The Slack source plugin supports syncing of message histories, but **only for channels that the bot is added to**. If you would like to sync the messages and threads in a channel, add the CloudQuery bot that you installed in Step 1 to the channel before running a sync, and make sure that the `slack_conversation_histories` table is included in the tables list in your Slack plugin source config.

## Example Queries

### List all active users in the Slack workspace

```sql
select id, name from slack_users where deleted is not true;
```

### Rank users by number of messages sent in public channels

```sql
select 
    u.name, count(h.user) 
from slack_conversation_histories h 
    join slack_conversation_replies r on h.ts = r.conversation_history_ts 
    join slack_users u on u.id = h.user 
group by u.name order by count desc;
```

### List all bookmarks

```sql
select title, link from slack_conversation_bookmarks;
```


### List URLs to all uploaded files

```sql
select title, url_private from slack_files;
```

### List all external files

```sql
select title, url_private from slack_files where is_external is true;
``` 