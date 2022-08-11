import { Callout } from 'nextra-theme-docs'

# Telemetry and Crash Reporting

CloudQuery collects anonymized usage statistics and crash reports. Crash reports help us iron out any issues.

## What is Stored

- Which operation (fetch, policy, etc.) was executed
- How long the operation took and whether it succeeded
- Whether you're running CloudQuery in a CI environment
- CloudQuery version used, build-date and commit id
- Installed CloudQuery providers on your system along with their versions
- Version and build date of the `cloudquery` binary you're using
- Basic data about which part of the operation took (fetch duration, resource fetch duration, etc')

### For vague identification purposes

- A randomly generated UUID, persisted across sessions
- A SHA1 hash of your: IP address, MAC addresses, Configuration and hostname
- Your Postgres database instance id, this is a unique generated ID created by your database.
- Your vague geographic location based on the IP address

This does not allow us to track individual users but does enable us to accurately measure user counts vs. invocation counts.

<Callout type="info">

The random ID is stored in the `.cq/telemetry-random-id` file. If you wish to anonymize your requests further but still keep sending us usage statistics, you could remove this file before every invocation (or create a directory with the same name, which also stops the file from getting created)

</Callout>

### Collector

- We use the [rudder](https://www.rudderstack.com/) collector to send usage analytics.

## What is NOT Stored

- We don't store your IP address directly
- We don't store any of the command arguments or options (as they might contain sensitive information)
- We don't store any credentials
- We don't store any logs

## Inspecting Telemetry Contents

To check what kind of data is collected, invoke the `cloudquery` executable with the `--inspect-telemetry` option. This will stop sending telemetry information and write it data to the log file, by default it should be in `cloudquery.log`. You can then inspect analytics log event contents and decide whether to opt-out or not.

```bash
# Invoke CloudQuery
cloudquery --inspect-telemetry [operation] [arguments]
```

## Opting out of Telemetry

To opt out of telemetry, simply invoke `cloudquery` executable with the `--no-telemetry` option, or set the `CQ_NO_TELEMETRY` environment variable.

```bash
# Set the environment variable
export CQ_NO_TELEMETRY=1

# Invoke CloudQuery. No telemetry information or crash reports will be sent
cloudquery [operation] [arguments]
```
