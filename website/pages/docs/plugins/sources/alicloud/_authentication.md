Authentication is done via access keys. See [Obtain an AccessKey pair](https://www.alibabacloud.com/help/en/basics-for-beginners/latest/obtain-an-accesskey-pair) for more information.

It is highly recommended that you use environment variable expansion to store your access key pair. For example, you can set the following environment variables:

  * `ALICLOUD_ACCESS_KEY` - The access key ID.
  * `ALICLOUD_SECRET_KEY` - The access key secret.

And then use them in the configuration as follows:

```yaml
access_key: ${ALICLOUD_ACCESS_KEY}
secret_key: ${ALICLOUD_SECRET_KEY}
```