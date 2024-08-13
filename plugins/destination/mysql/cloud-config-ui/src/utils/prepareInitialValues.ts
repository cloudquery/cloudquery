import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';

import { formValidationSchema, FormValues } from './formSchema';

const defaultValues = formValidationSchema.getDefault();

export function prepareInitialValues(
  initialValues: FormMessagePayload['init']['initialValues'],
): FormValues {
  const url = initialValues?.spec?.connection_string || '';
  const connectionType = url.startsWith('postgresql://')
    ? ('string' as const)
    : ('fields' as const);
  const connectionObj: Record<string, any> = {};

  if (connectionType === 'fields') {
    const params = url.split(' ');

    for (const param of params) {
      const [key, value] = param.split('=');
      connectionObj[key] = value?.replaceAll("'", ''); // Remove the single quotes around the value
    }
  }

  const spec = {
    connectionType,
    connectionString: connectionType === 'string' ? url : '',
    host: connectionObj.host || defaultValues.host,
    password: connectionObj.password || defaultValues.password,
    username: connectionObj.user || defaultValues.username,
    port: connectionObj.port || defaultValues.port,
    database: connectionObj.dbname || defaultValues.database,
    ssl: !!connectionObj.sslmode,
    // sslMode: connectionObj.sslmode || defaultValues.sslMode,
    // schemaName: connectionObj.search_path || defaultValues.schemaName,
    batchSize: initialValues?.spec?.batch_size ?? defaultValues.batchSize,
    batchSizeBytes: initialValues?.spec?.batch_size_bytes ?? defaultValues.batchSizeBytes,
  };

  return {
    ...spec,
    connectionParams: {},
    name: initialValues?.name || '',
    envs: initialValues?.envs || [],
    migrateMode: initialValues?.migrateMode || defaultValues.migrateMode,
    writeMode: initialValues?.writeMode || defaultValues.writeMode,
  };
}
