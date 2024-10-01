import { escapeSingleQuotesAndBackslashes } from '@cloudquery/plugin-config-ui-lib';

export function generateConnectionStringURI(values: any): string {
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

/**
 * Generates a connection string using key value pairs
 *
 * The connection string is in the format of "dbtype='postgresql' user='user' password='pass' host='myhost' port='1234' dbname='db'",
 * which allows special characters to be used in the values. In addition, this format is more readily parsed back into a key value
 * object when the plugin configuration is in editing mode.
 *
 * @param values
 * @returns {string} connection string in key value format
 */
export function generateConnectionStringKeyValue(values: any): string {
  const password = values.password ? '${password}' : '';

  let finalUrl = `dbtype='postgresql' user='${values.user}' password='${password}' host='${escapeSingleQuotesAndBackslashes(values.host)}' dbname='${escapeSingleQuotesAndBackslashes(values.database)}'`;

  if (values.port) {
    finalUrl += ` port='${values.port}'`;
  }

  if (values.connectionParams?.ssl) {
    finalUrl += ` sslmode='${values.connectionParams?.sslmode}'`;
  }

  if (values.connectionParams?.search_path) {
    finalUrl += ` search_path='${values.connectionParams?.search_path}'`;
  }

  return finalUrl;
}
