import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';

import { connectionTypeValues, formValidationSchema, FormValues } from './formSchema';
import { connectionStringToFields } from './connectionStringtoFields';

const defaultValues = formValidationSchema.getDefault();

export function prepareInitialValues(
  initialValues: FormMessagePayload['init']['initialValues'],
): FormValues {
  const url = initialValues?.spec?.connection_string || '';

  const connectionObj: Record<string, any> = connectionStringToFields(url);

  const spec = {
    connectionType: connectionTypeValues[1],
    connectionString: url,
    host: connectionObj.host || defaultValues.host,
    password: connectionObj.password || defaultValues.password,
    username: connectionObj.username || defaultValues.username,
    port: connectionObj.port || defaultValues.port,
    database: connectionObj.database || defaultValues.database,
    tcp: connectionObj.tcp || defaultValues.tcp,
    connectionParams: connectionObj.connectionParams || defaultValues.connectionParams,
    batchSize: initialValues?.spec?.batch_size ?? defaultValues.batchSize,
    batchSizeBytes: initialValues?.spec?.batch_size_bytes ?? defaultValues.batchSizeBytes,
  };

  return {
    ...spec,
    name: initialValues?.name || '',
    envs: initialValues?.envs || [],
    migrateMode: initialValues?.migrateMode || defaultValues.migrateMode,
    writeMode: initialValues?.writeMode || defaultValues.writeMode,
  };
}
