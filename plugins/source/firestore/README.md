# Firestore source Plugin

The CloudQuery Firestore source plugin replicates a Firestore database to any supported CloudQuery destination.

## Links

- [User Guide](https://www.cloudquery.io/docs/plugins/sources/firestore/overview)

## Test

```bash
docker run \
  --env "FIRESTORE_PROJECT_ID=cqtest-project" \
  --publish 8080:8080 \
  mtlynch/firestore-emulator-docker
```
