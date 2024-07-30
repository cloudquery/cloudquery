import { FormValues } from './formSchema';

export function generateConnectionUrl(values: FormValues, replaceWithEnvs: boolean): string {
  let { username } = values.spec;

  if (replaceWithEnvs) {
    username = replaceWithEnvs ? (values.spec.username ? '${username}' : '') : values.spec.username;
  }

  const password = values.spec.password ? '${password}' : '';

  let finalUrl = `dbtype='postgresql' user='${username}' password='${password}' host='${escapeSingleQuotesAndBackslashes(values.spec.host)}' dbname='${escapeSingleQuotesAndBackslashes(values.spec.database)}'`;

  if (values.spec.port) {
    finalUrl += ` port='${values.spec.port}'`;
  }

  if (values.spec.ssl) {
    finalUrl += ` sslmode='${values.spec.sslMode}'`;
  }

  if (values.spec.clientEncoding) {
    finalUrl += ` client_encoding='${escapeSingleQuotesAndBackslashes(values.spec.clientEncoding)}'`;
  }

  return finalUrl;
}

export function escapeSingleQuotesAndBackslashes(str: string) {
  return str.replaceAll('\\', '\\\\').replaceAll("'", String.raw`\'`);
}
