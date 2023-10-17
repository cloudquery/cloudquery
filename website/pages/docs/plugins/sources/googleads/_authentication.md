[Google Ads API](https://developers.google.com/google-ads/api/)
authentication is based on [OAuth 2.0 authorization](https://developers.google.com/identity/protocols/oauth2)
along with [Developer Token](https://developers.google.com/google-ads/api/docs/get-started/dev-token).

Two methods are supported: OAuth 2.0 and Application Default Credentials.

**Note**: See the official [documentation](https://developers.google.com/google-ads/api/docs/oauth/overview)
describing different authorization options.

### OAuth 2.0

The following options are available when using OAuth:

- **Using an existing access token**

  This token should be authorized for `https://www.googleapis.com/auth/adwords` scope (e.g. by using [OAuth 2.0 Playground](https://developers.google.com/oauthplayground/)).

- **Using OAuth client ID & client secret**

  You can get your own OAuth credentials using [this guide](https://developers.google.com/identity/protocols/oauth2#1.-obtain-oauth-2.0-credentials-from-the-dynamic_data.setvar.console_name-.).

### Application Default Credentials

See the official [Application Default Credentials guide](https://cloud.google.com/sdk/gcloud/reference/auth/application-default).

**Note**: You will still need to authorize these credentials for `https://www.googleapis.com/auth/adwords` scope.

Available options are all the same options described [here](https://cloud.google.com/docs/authentication/provide-credentials-adc) in detail.
