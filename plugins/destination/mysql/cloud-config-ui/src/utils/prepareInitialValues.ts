import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';

import { connectionTypeValues, formValidationSchema, FormValues } from './formSchema';

const defaultValues = formValidationSchema.getDefault();

const connectionStringToFields = (connectionString: string) => {
  // unit test TODO
  // const example =
  //   'user:password@localhost:3306/dbname?timeout=30s\u0026readTimeout=1s\u0026writeTimeout=1s';

  const username = connectionString.split(':')[0] ?? '';
  const password = connectionString.split(':')[1]?.split('@')[0] ?? '';
  const address = connectionString.split('@')[1]?.split('/')[0] ?? '';
  const tcp = address.startsWith('tcp(');
  const host = (tcp ? address.split('(')[1]?.split(':')[0] : address.split(':')[0]) ?? '';
  const port = (tcp ? address.split(':')[1]?.split(')')[0] : address.split(':')[1]) ?? '';
  const database = connectionString.split('/')[1]?.split('?')[0] ?? '';
  // const params = connectionString.split('?')[1]; // TODO

  // const parts = connectionString.match(/([^:@/()?]+)/g);
  return {
    username,
    password,
    host,
    port,
    database,
    tcp,
    connectionParams: {
      // TODO decodeURI
    },
  };
};

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
