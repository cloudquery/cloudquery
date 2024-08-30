const mysqlErrorDescriptions = {
  INVALID_DSN:
    'The connection string (DSN) is invalid or in an incorrect format. Please check and correct your connection details.',
  CONNECT_FAILED:
    'Failed to establish a connection to the MySQL database. This is rare and might indicate a driver issue.',
  DEFAULT_DATABASE_FAILED:
    "Unable to determine the default database. Please ensure you've specified a database name in your connection string.",
  QUERY_VERSION_FAILED:
    'Failed to retrieve the MySQL version. This might indicate restricted permissions or a connection issue.',
  UNREACHABLE: 'The MySQL server is unreachable. Check your host, port, and network settings.',
  ACCESS_DENIED: 'Access denied. The provided username or password is incorrect.',
  UNKNOWN_DATABASE:
    "The specified database does not exist. Please check your database name and ensure it's created on the server.",
  PING_FAILED:
    'Failed to ping the MySQL server. This might indicate network issues or server unavailability.',
  LIST_FAILED: 'Failed to list databases. This might be due to insufficient permissions.',
};

export function parseTestConnectionError(error: Error & { code?: string }) {
  return {
    ...error,
    message:
      mysqlErrorDescriptions[error.code as keyof typeof mysqlErrorDescriptions] ||
      error.message ||
      'Unknown error',
  };
}
