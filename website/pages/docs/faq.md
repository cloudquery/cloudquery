# FAQ

## Does CloudQuery access any application data in my cloud?

No. CloudQuery only accesses metadata and configuration data. It never pulls data from your application databases or cloud storage files.

### What happens when I run two (or more) fetches? Will the second fetch remove resources that no longer exist from the database?

After a `fetch`, your PostgreSQL database will always mirror your cloud environment. This means that a `fetch` will add to the database
any resources that were added to the cloud during the time from the previous `fetch`, update any resources that were changed in the cloud,
and remove from the database any resources that were removed from the cloud.

The only caveat is that in the case of configuration changes, subsequent fetches won't always remove stale resources. For instance:

- If an account was removed from `cloudquery.yml`, subsequent fetches won't process that account (and won't remove that account's resources).
- If a resource type was removed from `cloudquery.yml`, that resource type won't get fetched (and database entries for this resources also won't get removed).
- When using [AWS Organizations](https://hub.cloudquery.io/providers/cloudquery/aws/latest#aws-organizations), if an account is deleted or removed from the org then the database entries for that account will be untouched.
- When using [GCP Folders or Organizations](https://hub.cloudquery.io/providers/cloudquery/gcp/latest#configuration), if a project is deleted or removed from the folder, the database entries for that account will be untouched.

#### What if I want to remove resources from accounts/projects that I'm no longer fetching from?

You can use the `cloudquery provider purge [provider]` to delete old resources. For example, you can remove
resources that weren't updated in the last three days with `cloudquery provider purge [provider] --last-update=72h --dry-run=false`.
You can find the full details [here](https://docs.cloudquery.io/docs/cli/commands/provider_purge).

### What is "//"? How is it different from "/" and where should it go? {#what-is-double-slash}

The `//` indicator is used when specifying a subpolicy/subquery for `cloudquery policy run`. It is an idiosyncrasy of the way the `cloudquery` CLI works - it separates the "path to the policy" from the "path to the subpolicy (in the policy)". It must always appear **right after the root policy**.

So, if the policy I'd like to run is `./my_policy`, and I would like
to run the `old-stopped-ec2-instances` query in the `ec2` subpolicy, I would use:

```bash
cloudquery policy run ./my_policy//ec2/old-stopped-ec2-instances
```

It is worth mentioning here that the `cloudquery` CLI also supports running policies from
our official github. The `//` separator serves the same function
described above - separating the "path to the policy" from the "path inside the policy". In case of running
policies from github, it helps the `cloudquery` CLI to know which repository to clone.

So, to run the `foundational_security` **subpolicy** in the `aws` **policy**, we run.

```bash
cloudquery policy run aws//foundational_security
```
