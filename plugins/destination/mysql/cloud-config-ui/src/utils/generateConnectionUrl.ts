import { FormValues } from './formSchema';

export function generateConnectionUrl(values: FormValues): string {
  const password = values.password ? '${password}' : '';
  const db = values.database ? `${escapeSingleQuotesAndBackslashes(values.database)}` : '';
  const host = values.host ? `${escapeSingleQuotesAndBackslashes(values.host)}` : undefined;
  const port = values.port ?? undefined;
  const address = host ? `${host}:${port}` : undefined;
  const wrappedAddress = address ? (values.tcp ? `tcp(${address})` : address) : '';

  let base = `${values.username}:${password}@${wrappedAddress}/${db}`;

  const normalizedConnectionParams: Record<string, boolean | string> = {};
  if (values.connectionParams.tls && values.connectionParams.tlsMode) {
    normalizedConnectionParams['tlsMode'] = values.connectionParams.tlsMode;
  }

  if (values.connectionParams.parseTime) {
    normalizedConnectionParams['parseTime'] = 'True';
  }

  if (values.connectionParams.charset) {
    normalizedConnectionParams['charset'] = values.connectionParams.charset;
  }

  if (values.connectionParams.loc) {
    normalizedConnectionParams['loc'] = values.connectionParams.loc;
  }

  const queryParams = new URLSearchParams(normalizedConnectionParams as any).toString();

  return queryParams ? `${base}?${queryParams}` : base;
}

export function escapeSingleQuotesAndBackslashes(str: string) {
  return str.replaceAll('\\', '\\\\').replaceAll("'", String.raw`\'`);
}
