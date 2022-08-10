# Troubleshooting

## Help Channels

### Discord

First things first - feel free to join our new [Discord](https://discord.gg/2mPfFYyAtQ)!

### GitHub Issues

There are couple of ways to get help for any CloudQuery-related issues or questions.

1. Check out previous issues at <https://github.com/cloudquery> and open a new one if no previous one has been opened or resolved.

### Intercom Chat

Our Intercom chat is available on all our public sites, so feel free to drop us a line there.

## Debugging

### Verbose Logging

Usually the first step that will be needed to debug/resolve an issue is to run `cloudquery` with `-v` to enable verbose logging.

### I am trying to run a policy with "cloudquery policy run", but am getting a "Failed to run policies…Invalid value for path parameter" error

Are you correctly using `//` when specifying the policy you want to run? The `//` separates the "path to the
root policy" from the "path to the subpolicy" - it must appear **right after the root policy**.

For instance, to run the `foundational_security/ec2` subpolicy of the `aws` root policy, you must use:

```bash
cloudquery policy run aws//foundational_security/ec2
```

[Read more in the FAQ](faq#what-is-double-slash).
