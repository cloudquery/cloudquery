interface ConnectionOptions {
  originalProtocol: string;
  host: string;
  password?: string;
  user?: string;
  port: string;
  database: string | null | undefined;
  clientEncoding?: string;
  sslMode?: 'require' | 'prefer' | 'disable' | 'allow' | 'verify-ca' | 'verify-full';
  schema?: string;
}

//parses a connection string
export function parseConnectionString(url: string): ConnectionOptions {
  try {
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
      port: '',
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
  } catch {
    return {
      originalProtocol: '',
      host: '',
      password: '',
      user: '',
      port: '',
      database: '',
      clientEncoding: '',
    };
  }
}

export function generateConnectionString(options: ConnectionOptions): string {
  const url = new URL('https://cloudquery.io');
  url.protocol = options.originalProtocol;
  url.hostname = options.host;
  url.port = options.port;
  url.pathname = encodeURIComponent(options.database || '');

  if (options.sslMode) {
    url.searchParams.set('sslmode', options.sslMode);
  }
  if (options.schema) {
    url.searchParams.set('search_path', options.schema);
  }

  const userInfo = [options.user, options.password].filter(Boolean).join(':');

  let finalUrl = url.toString().replace(/^[a-z]*:\/\//, `${options.originalProtocol}://`);

  if (userInfo) {
    finalUrl = finalUrl.replace(
      `${options.originalProtocol}://`,
      `${options.originalProtocol}://${userInfo}@`,
    );
  }

  return finalUrl;
}
