import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { FormValues } from './formSchema';

// TODO: the shape of this, per the API

interface ConnectionOptions {
  originalProtocol: string;
  host: string | null;
  password?: string;
  user?: string;
  port?: string | null;
  database: string | null | undefined;
  clientEncoding?: string;
  sslMode?: 'require' | 'prefer' | 'disable' | 'allow' | 'verify-ca' | 'verify-full';
}

export function prepareInitialValues(
  initialValues: FormMessagePayload['init']['initialValues'],
): FormValues {
  const connectionObj = parse(initialValues?.spec?.connection_string);
  const spec = {
    originalProtocol: connectionObj.originalProtocol || 'postgresql',
    host: connectionObj.host || '',
    password: connectionObj.password || '',
    username: connectionObj.user || '',
    port: connectionObj.port || '',
    database: connectionObj.database || '',
    clientEncoding: connectionObj.clientEncoding || '',
    ssl: !!connectionObj.sslMode,
    sslMode: connectionObj.sslMode || ('require' as const),
  };

  return {
    ...(initialValues as any), // TODO
    spec,
  };
}

//parses a connection string
function parse(url: string): ConnectionOptions {
  const originalProtocol = url.split('://')[0];
  let str = url.replace(/^[a-z]*:\/\//, 'ftp://');
  //unix socket
  if (str.charAt(0) === '/') {
    const config = str.split(' ');

    return { host: config[0], database: config[1] } as ConnectionOptions;
  }

  // Check for empty host in URL

  const config: ConnectionOptions = {
    originalProtocol,
    host: '',
    database: '',
  };
  let result;
  let dummyHost = false;
  if (/ |%[^\da-f]|%[\da-f][^\da-f]/i.test(str)) {
    // Ensure spaces are encoded as %20
    str = encodeURI(str).replaceAll(/%25(\d\d)/g, '%$1');
  }

  try {
    result = new URL(str, 'postgresql://base');
  } catch {
    // The URL is invalid so try again with a dummy host
    result = new URL(str.replace('@/', '@___DUMMY___/'), 'postgresql://base');
    dummyHost = true;
  }

  config.user = config.user || decodeURIComponent(result.username);
  config.password = config.password || decodeURIComponent(result.password);

  if (result.protocol == 'socket:') {
    config.host = decodeURI(result.pathname);
    config.database = result.searchParams.get('db');
    config.clientEncoding = result.searchParams.get('encoding') || undefined;

    return config;
  }
  const hostname = dummyHost ? '' : result.hostname;
  if (!config.host) {
    // Only set the host if there is no equivalent query param.
    config.host = decodeURIComponent(hostname);
  } else if (hostname && /^%2f/i.test(hostname)) {
    // Only prepend the hostname to the pathname if it is not a URL encoded Unix socket host.
    result.pathname = hostname + result.pathname;
  }
  if (!config.port) {
    // Only set the port if there is no equivalent query param.
    config.port = result.port;
  }

  const pathname = result.pathname.slice(1) || null;
  config.database = pathname ? decodeURI(pathname) : null;

  const searchParams = new URLSearchParams(result.search);

  config.sslMode = searchParams.get('sslmode') as ConnectionOptions['sslMode'];

  return config;
}
