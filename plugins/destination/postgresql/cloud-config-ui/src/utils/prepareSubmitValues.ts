import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { FormValues } from './formSchema';
import { escapeSingleQuotesAndBackslashes, generateConnectionUrl } from './generateConnectionUrl';

export function prepareSubmitValues(
  values: FormValues,
): PluginUiMessagePayload['validation_passed']['values'] {
  const envs = [] as Array<{ name: string; value: string }>;

  if (values.username) {
    envs.push({
      name: 'username',
      value:
        values.username === '${username}' ? '' : escapeSingleQuotesAndBackslashes(values.username),
    });
  }
  if (values.password) {
    envs.push({
      name: 'password',
      value:
        values.password === '${password}' ? '' : escapeSingleQuotesAndBackslashes(values.password),
    });
  }

  return {
    name: values.name,
    envs,
    spec: {
      connection_string: generateConnectionUrl(values, true),
      pgx_log_level: values.pgxLogLevel,
      batch_size: values.batchSize,
      batch_size_bytes: values.batchSizeBytes,
      batch_timeout: values.batchTimeout,
    },
    migrateMode: values.migrateMode,
    writeMode: values.writeMode,
  };
}
