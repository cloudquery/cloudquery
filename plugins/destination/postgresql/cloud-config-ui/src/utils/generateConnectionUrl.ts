import { FormValues } from './formSchema';

export function generateConnectionUrl(values: FormValues, replaceWithEnvs: boolean): string {
  let { username } = values;

  if (replaceWithEnvs) {
    username = replaceWithEnvs ? (values.username ? '${username}' : '') : values.username;
  }

  const password = values.password ? '${password}' : '';

  let finalUrl = `dbtype='postgresql' user='${username}' password='${password}' host='${escapeSingleQuotesAndBackslashes(values.host)}' dbname='${escapeSingleQuotesAndBackslashes(values.database)}'`;

  if (values.port) {
    finalUrl += ` port='${values.port}'`;
  }

  if (values.ssl) {
    finalUrl += ` sslmode='${values.sslMode}'`;
  }

  if (values.schemaName) {
    finalUrl += ` search_path='${values.schemaName}'`;
  }

  return finalUrl;
}

export function escapeSingleQuotesAndBackslashes(str: string) {
  return str.replaceAll('\\', '\\\\').replaceAll("'", String.raw`\'`);
}
