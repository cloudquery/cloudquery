export const connectionTypeValues = ['string', 'fields'] as const;
export const migrateModeValues = ['forced', 'safe'] as const;
export const writeModeValues = ['append', 'overwrite', 'overwrite-delete-stale'] as const;
export const compressValues = ['none', 'zstd', 'lz4', 'gzip', 'deflate', 'br'] as const;
export const connectionOpenStrategyValues = ['random', 'round_robin', 'in_order'] as const;
