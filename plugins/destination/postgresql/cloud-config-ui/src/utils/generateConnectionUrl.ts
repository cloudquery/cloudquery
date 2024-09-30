import { escapeSingleQuotesAndBackslashes } from '@cloudquery/plugin-config-ui-lib';

export function generateConnectionUrl(values: any): string {
  const password = values.password ? '${password}' : '';
  const credentials = values.user ? `${values.user}:${password}@` : '';
  const database = values.database
    ? `${escapeSingleQuotesAndBackslashes(values.database.trim())}`
    : '';
  const host = values.host ? `${escapeSingleQuotesAndBackslashes(values.host.trim())}` : '';
  const port = values.port ?? '';
  const address = port ? `${host}:${port}` : `${host}`;
  const wrappedAddress = address ? (values.tcp ? `tcp(${address})` : address) : '';

  const base = `postgres://${credentials}${wrappedAddress}/${database}`;

  const normalizedConnectionParams: Record<string, boolean | string> = {};
  if (values.connectionParams.ssl && values.connectionParams.sslmode) {
    normalizedConnectionParams['sslmode'] = values.connectionParams.sslmode;
  }
  if (values.connectionParams.search_path) {
    normalizedConnectionParams['search_path'] = values.connectionParams.search_path;
  }

  const queryParams = new URLSearchParams(normalizedConnectionParams as any).toString();

  return queryParams ? `${base}?${queryParams}` : base;
}
