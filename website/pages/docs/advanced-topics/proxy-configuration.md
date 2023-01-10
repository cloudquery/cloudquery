import { Callout } from 'nextra-theme-docs'

# Using CloudQuery CLI with a Proxy Server

If you run the CloudQuery CLI in an environment that requires a proxy server for outgoing traffic, you'll need to set it via environment variables. To configure a proxy server for HTTPS traffic set the `HTTPS_PROXY` environment variable. For HTTP traffic set the `HTTP_PROXY` environment variable.

Example:

```bash copy
export HTTPS_PROXY=http://example.com:3128/proxy
export HTTP_PROXY=http://example.com:3128/proxy
cloudquery sync [files or directories]
```

<Callout type="info">

`HTTPS` in `HTTPS_PROXY` variable name means that it is a proxy for HTTPS requests, and not the protocol of the proxy, so its value can start with `http://`.

</Callout>
