The GitHub source plugin supports two authentication methods: Personal Access Token and App authentication. Which one you use is up to and the security requirements of your organization.

Keep in mind rate limits for GitHub Apps are higher than for personal access tokens.  Review [GitHub rate limits documentation](https://docs.github.com/en/apps/creating-github-apps/registering-a-github-app/rate-limits-for-github-apps) for details. 

CloudQuery requires only *read* permissions (we will never make any changes to your GitHub account or organizations),
so, following the principle of least privilege, it's recommended to grant it read-only permissions to all the resources you wish to sync.

## Personal Access Token

Follow this [guide](https://docs.github.com/en/enterprise-server@3.4/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) on how to create a personal access token for CloudQuery.

## App authentication

For App authentication, you need to create a GitHub App and install it on your organization. Follow [this guide](https://docs.github.com/en/apps/creating-github-apps/creating-github-apps/creating-a-github-app) and install the App into your organization(s). Give it all the permissions you need (read-only is recommended).

Every organization will have a unique installation ID. You can find it by going to the organization's settings page, and clicking on the "Installed GitHub Apps" tab. The installation ID is the number in the URL of the page.

<details>
    <summary>Passing `private_key` as plaintext</summary>
    You can use `|` to pass the multi-line private key as plaintext.
    
    For example:

    ```yaml
    - org: cloudquery
      private_key: |
        -----BEGIN RSA PRIVATE KEY-----
        MIIEpQIBAAKCAQEA3eVv6PCn9P8zO+EP8K7pLMfxcA2uVrSZ2f+H3GgYIavDxWtO
        vM9tE3jAA8mOjZpdLaG5yy4QfV1LQ3R7kO49JCB6VbClwN2lNvd8Iw49JCBDid7D
        ...
        -----END RSA PRIVATE KEY-----
      app_id: your_app_id
    ```

</details>

<details>
    <summary>Referencing `private_key` as environment variable</summary>
    When referencing the `private_key` as a string from environment variables, you will need to replace all the new lines in your PEM file with `\n` otherwise the new line and indent will prevent CloudQuery from reading the variable correctly. 

    For example:

    ```yaml 
    - org: cloudquery
      private_key: "${GITHUB_PRI_KEY}"
      app_id: your_app_id
      ...
    ```
    where
    ```bash
    GITHUB_PRI_KEY="-----BEGIN RSA PRIVATE KEY-----\nMIIEpQIBAAKCAQEA3eVv6PCn9P8zO+EP8K7pLMfxcA2uVrSZ2f+H3GgYIavDxWtO\n...vM9tE3jAA8mOjZpdLaG5yy4QfV1LQ3R7kO49JCB6VbClwN2lNvd8Iw==\n-----END RSA PRIVATE KEY-----"
    ```

</details>


