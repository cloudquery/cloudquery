# V1 Migration Guide Generation Script

This is a once-off script to generate v0 to v1 schema migration guides, such as [this one for AWS](../../plugins/source/aws/docs/v1-migration.md).

## Running Locally

Generating these require a bit of manual setup, but the process can be replicated using something like this:

```shell
git worktree add ../v0
git worktree add ../v1
cd ../v0 && git checkout ab65d1e33 # or any v0 commit
cd ../v1 && make build
cd ..
```

Then run this script from outside the git root:

```shell
plugins="aws gcp azure cloudflare digitalocean github okta k8s" 

mkdir -p docs/tables-v0
mkdir -p docs/tables-v1
for p in $plugins
do
	echo "Generating migration guide for $p"
	cp -r v0/plugins/source/$p/docs/tables/* docs/tables-v0
	./v1/bin/$p doc docs/tables-v1
	./v1/bin/v1-migration -o cloudquery/plugins/source/$p/docs/v1-migration.md -v1 docs/tables-v0 -v2 docs/tables-v1
done
```