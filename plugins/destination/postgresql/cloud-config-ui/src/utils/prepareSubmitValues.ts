import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { FormValues } from './formSchema';
import { escapeSingleQuotesAndBackslashes, generateConnectionUrl } from './generateConnectionUrl';

export function prepareSubmitValues(
  values: FormValues,
): PluginUiMessagePayload['validation_passed']['values'] {
  const envs = [] as Array<{ name: string; value: string }>;

  if (values.spec.username) {
    envs.push({
      name: 'username',
      value:
        values.spec.username === '${username}'
          ? ''
          : escapeSingleQuotesAndBackslashes(values.spec.username),
    });
  }
  if (values.spec.password) {
    envs.push({
      name: 'password',
      value:
        values.spec.password === '${password}'
          ? ''
          : escapeSingleQuotesAndBackslashes(values.spec.password),
    });
  }

  return {
    name: values.name,
    envs,
    spec: {
      connection_string: generateConnectionUrl(values, true),
      pgx_log_level: values.spec.pgxLogLevel,
      batch_size: values.spec.batchSize,
      batch_size_bytes: values.spec.batchSizeBytes,
      batch_timeout: values.spec.batchTimeout,
    },
    migrateMode: values.migrateMode,
    writeMode: values.writeMode,
  };
}
