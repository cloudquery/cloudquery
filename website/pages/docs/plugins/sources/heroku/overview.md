# Heroku Plugin

The CloudQuery Heroku plugin extracts your Heroku data and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

## Authentication

The CloudQuery Heroku plugin requires an OAuth token. After creating it, you
will need to copy the token into your `heroku.yml` file (See example below):

```yaml
  ...
  token: <Token HERE>
```

### Option 1: Generate a token with the Heroku CLI

A token can be generated using the Heroku CLI.

 1. **Install the Heroku CLI**: Follow the [official instructions](https://devcenter.heroku.com/articles/heroku-cli) to install the Heroku CLI.
 2. **Generate an OAuth token**: With the Heroku CLI installed, use your terminal to run:
    ```bash
    heroku authorizations:create --short --description="CloudQuery token" --scope="read,identity"
    ```

    (For additional options for this command, such as expiry, see the [Heroku CLI commands documentation](https://devcenter.heroku.com/articles/heroku-cli-commands#heroku-authorizations-create))

### Option 2: Generate a token with the Heroku API

It is also possible to manage OAuth tokens directly through the Heroku API. For more information, see the [Heroku documentation](https://devcenter.heroku.com/articles/platform-api-reference#oauth-authorization-create). 

### A Note about OAuth Scopes

CloudQuery needs to be authenticated with your Heroku account's token in order to fetch information about your Heroku resources.
CloudQuery requires only **read** permissions (we will never make any changes to your Heroku account or apps).
Following the principle of least privilege, it is recommended to grant it read-only permissions. The `--scope="read,identity"`
parameter suggested above achieves this. 

However, certain Heroku resources require a `global` scope, even for reading. At the time of writing, these resources are:
 - `app_webhook_deliveries`
 - `app_webhook_events`
 - `app_webhooks`
 - `credits`
 - `invoices`
 - `keys`
 - `oauth_authorizations`
 - `oauth_clients`
 - `permission_entities`
 - `team_features`
 - `team_invitations`
 - `team_invoices`
 - `team_members`
 - `team_spaces`

If you are interested in fetching any of these resources, a `global` scope will be necessary. See the Heroku documentation for [more information about OAuth scopes on Heroku](https://devcenter.heroku.com/articles/oauth#scopes).
