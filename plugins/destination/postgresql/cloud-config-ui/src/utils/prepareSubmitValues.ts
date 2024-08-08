import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { FormValues } from './formSchema';
import { escapeSingleQuotesAndBackslashes, generateConnectionUrl } from './generateConnectionUrl';
import { generateConnectionString, parseConnectionString } from './parseConnectionString';

export function prepareSubmitValues(
  values: FormValues,
): PluginUiMessagePayload['validation_passed']['values'] {
  const envs = [] as Array<{ name: string; value: string }>;
  let { connectionString } = values;

  if (values.connectionType === 'string') {
    const { password, ...connectionStringProps } = parseConnectionString(values.connectionString);

    if (password && password !== '${password}') {
      envs.push({ name: 'password', value: password });
      connectionString = generateConnectionString({
        ...connectionStringProps,
        password: '${password}',
      });
    } else if (password && password === '${password}') {
      envs.push({ name: 'password', value: '' });
    }
  } else {
    if (values.password) {
      envs.push({
        name: 'password',
        value:
          values.password === '${password}'
            ? ''
            : escapeSingleQuotesAndBackslashes(values.password),
      });
    }
  }

  return {
    name: values.name,
    envs,
    spec: {
      connection_string:
        values.connectionType === 'string' ? connectionString : generateConnectionUrl(values),
      pgx_log_level: values.pgxLogLevel,
      batch_size: values.batchSize,
      batch_size_bytes: values.batchSizeBytes,
      batch_timeout: values.batchTimeout,
    },
    migrateMode: values.migrateMode,
    writeMode: values.writeMode,
  };
}
