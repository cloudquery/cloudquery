import { FormValues } from './formSchema';

export function generateConnectionUrl(values: FormValues): string {
  const password = values.password ? '${password}' : '';
  const db = values.database ? `${escapeSingleQuotesAndBackslashes(values.database.trim())}` : '';
  const host = values.host ? `${escapeSingleQuotesAndBackslashes(values.host.trim())}` : undefined;
  const port = values.port ?? undefined;
  const address = host ? `${host}:${port}` : undefined;
  const wrappedAddress = address ? (values.tcp ? `tcp(${address})` : address) : '';

  const base = `${values.username}:${password}@${wrappedAddress}/${db}`;

  const normalizedConnectionParams: Record<string, boolean | string> = {};
  if (values.connectionParams.tls && values.connectionParams.tlsMode) {
    normalizedConnectionParams['tlsMode'] = values.connectionParams.tlsMode;
  }

  if (values.connectionParams.parseTime) {
    normalizedConnectionParams['parseTime'] = 'True';
  }

  if (values.connectionParams.charset) {
    normalizedConnectionParams['charset'] = values.connectionParams.charset.trim();
  }

  if (values.connectionParams.loc) {
    normalizedConnectionParams['loc'] = values.connectionParams.loc.trim();
  }

  if (values.connectionParams.timeout) {
    normalizedConnectionParams['timeout'] = `${values.connectionParams.timeout.toString()}s`;
  }

  if (values.connectionParams.readTimeout) {
    normalizedConnectionParams['readTimeout'] =
      `${values.connectionParams.readTimeout.toString()}s`;
  }

  if (values.connectionParams.writeTimeout) {
    normalizedConnectionParams['writeTimeout'] =
      `${values.connectionParams.writeTimeout.toString()}s`;
  }

  if (values.connectionParams.allowNativePasswords) {
    normalizedConnectionParams['allowNativePasswords'] =
      values.connectionParams.allowNativePasswords;
  }

  const queryParams = new URLSearchParams(normalizedConnectionParams as any).toString();

  return queryParams ? `${base}?${queryParams}` : base;
}

export function escapeSingleQuotesAndBackslashes(str: string) {
  return str.replaceAll('\\', '\\\\').replaceAll("'", String.raw`\'`);
}
