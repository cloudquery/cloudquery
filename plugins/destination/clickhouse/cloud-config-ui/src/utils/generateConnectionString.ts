import { escapeSingleQuotesAndBackslashes } from '@cloudquery/plugin-config-ui-lib';

export function generateConnectionStringURI(values: any): string {
  const password = values.password ? '${password}' : '';
  const username = values.username ? '${username}' : '';
  const database = values.database ? '${database}' : '';
  const credentials = `${username}:${password}@`;
  const address = values.hosts
    ? values.hosts
        .map((host: string) => `${escapeSingleQuotesAndBackslashes(host.trim())}`)
        .join(',')
    : '';

  const base = `clickhouse://${credentials}${address}/${database}`;

  const normalizedConnectionParams: Record<string, boolean | string> = values.connectionParams;
  // handle postfixes
  if (values.connectionParams.dial_timeout) {
    normalizedConnectionParams['dial_timeout'] = `${values.connectionParams.dial_timeout}ms`;
  }
  if (values.connectionParams.read_timeout) {
    normalizedConnectionParams['read_timeout'] = `${values.connectionParams.read_timeout}s`;
  }

  const queryParams = new URLSearchParams(values.connectionParams as any).toString();

  return queryParams ? `${base}?${queryParams}` : base;
}
