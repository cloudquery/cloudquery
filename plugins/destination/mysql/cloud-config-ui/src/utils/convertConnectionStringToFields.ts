export const convertConnectionStringToFields = (connectionString: string) => {
  const connectionParams: Record<string, any> = {};

  const username = connectionString.split(':')[0] ?? '';
  const password = connectionString.split(':')[1]?.split('@')[0] ?? '';
  const address = connectionString.split('@')[1]?.split('/')[0] ?? '';
  const tcp = address.startsWith('tcp(');
  const host = (tcp ? address.split('(')[1]?.split(':')[0] : address.split(':')[0]) ?? '';
  const port = (tcp ? address.split(':')[1]?.split(')')[0] : address.split(':')[1]) ?? '';
  const database = connectionString.split('/')[1]?.split('?')[0] ?? '';

  const params = decodeURI(connectionString).split('?')[1].split('&'); // TODO
  for (const param of params) {
    const [key, value] = param.split('=');
    if (key.toLowerCase().includes('timeout')) {
      connectionParams[key] = Number(value.replace('s', ''));
    } else if (key === 'tlsMode') {
      connectionParams.tls = true;
      connectionParams.tlsMode = value;
    } else if (['True', 'False', 'true', 'false'].includes(value)) {
      connectionParams[key] = ['True', 'true'].includes(value);
    } else {
      connectionParams[key] = value;
    }
  }

  return {
    username,
    password,
    host,
    port,
    database,
    tcp,
    connectionParams,
  };
};
