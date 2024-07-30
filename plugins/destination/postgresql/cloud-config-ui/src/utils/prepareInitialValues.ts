import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { FormValues } from './formSchema';

export function prepareInitialValues(
  initialValues: FormMessagePayload['init']['initialValues'],
): FormValues {
  const url = initialValues?.spec?.connection_string || '';
  const connectionObj: Record<string, any> = {};
  const params = url.split(' ');

  for (const param of params) {
    const [key, value] = param.split('=');
    connectionObj[key] = value.replaceAll("'", ''); // Remove the single quotes around the value
  }

  const spec = {
    host: connectionObj.host || '',
    password: connectionObj.password || '',
    username: connectionObj.user || '',
    port: connectionObj.port || '',
    database: connectionObj.dbname || '',
    clientEncoding: connectionObj.client_encoding || '',
    ssl: !!connectionObj.sslmode,
    sslMode: connectionObj.sslmode || ('require' as const),
    pgxLogLevel: initialValues?.spec?.pgx_log_level,
    batchSize: initialValues?.spec?.batch_size,
    batchSizeBytes: initialValues?.spec?.batch_size_bytes,
    batchTimeout: initialValues?.spec?.batch_timeout,
  };

  return {
    connectionType: 'multipleFields',
    name: initialValues?.name || '',
    envs: initialValues?.envs || [],
    spec,
    migrateMode: initialValues?.migrateMode || 'safe',
    writeMode: initialValues?.writeMode || 'overwrite-delete-stale',
  };
}
