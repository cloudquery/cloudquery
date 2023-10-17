The GitHub source plugin supports two authentication methods: Personal Access Token and App authentication. Which one you use is up to and the security requirements of your organization.

Keep in mind rate limits for GitHub Apps are higher than for personal access tokens.  Review [GitHub rate limits documentation](https://docs.github.com/en/apps/creating-github-apps/registering-a-github-app/rate-limits-for-github-apps) for details. 

CloudQuery requires only *read* permissions (we will never make any changes to your GitHub account or organizations),
so, following the principle of least privilege, it's recommended to grant it read-only permissions to all the resources you wish to sync.

## Personal Access Token

Follow this [guide](https://docs.github.com/en/enterprise-server@3.4/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) on how to create a personal access token for CloudQuery.

## App authentication

For App authentication, you need to create a GitHub App and install it on your organization. Follow [this guide](https://docs.github.com/en/apps/creating-github-apps/creating-github-apps/creating-a-github-app) and install the App into your organization(s). Give it all the permissions you need (read-only is recommended).

Every organization will have a unique installation ID. You can find it by going to the organization's settings page, and clicking on the "Installed GitHub Apps" tab. The installation ID is the number in the URL of the page.
