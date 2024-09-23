export const connectionTypeValues = ['string', 'fields'] as const;
export const sslModeValues = [
  'allow',
  'disable',
  'prefer',
  'require',
  'verify-ca',
  'verify-full',
] as const;
export const migrateModeValues = ['safe', 'forced'] as const;
export const writeModeValues = ['append', 'overwrite', 'overwrite-delete-stale'] as const;
export const pgxLogLevelValues = ['error', 'warn', 'info', 'debug', 'trace'] as const;
