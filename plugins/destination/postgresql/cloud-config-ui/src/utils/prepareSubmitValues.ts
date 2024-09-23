import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { corePrepareSubmitValues } from '@cloudquery/plugin-config-ui-lib';
import { generateConnectionUrl } from './generateConnectionUrl';
import { parseConnectionString } from './parseConnectionString';

export function prepareSubmitValues(
  values: Record<string, any>,
): PluginUiMessagePayload['validation_passed']['values'] {
  const payload =
    values._connectionType === 'string'
      ? prepareSubmitValuesFromConnectionString(values)
      : prepareSubmitValuesFromFields(values);

  payload.migrateMode = values.migrateMode;
  payload.writeMode = values.writeMode;

  if (values.pgx_log_level) {
    payload.spec.pgx_log_level = values.pgx_log_level;
  }
  if (values.batch_size) {
    payload.spec.batch_size = values.batch_size;
  }
  if (values.batch_size_bytes) {
    payload.spec.batch_size_bytes = values.batch_size_bytes;
  }
  if (values.batch_timeout) {
    payload.spec.batch_timeout = values.batch_timeout;
  }

  return payload;
}

function prepareSubmitValuesFromFields(
  values: Record<string, any>,
): PluginUiMessagePayload['validation_passed']['values'] {
  const payload = corePrepareSubmitValues(values);
  payload.spec.connection_string = generateConnectionUrl(values);
  delete payload.spec.password;
  payload.envs = payload.envs.filter(({ name }) => name !== 'connection_string');

  return payload;
}

function prepareSubmitValuesFromConnectionString(
  values: Record<string, any>,
): PluginUiMessagePayload['validation_passed']['values'] {
  const connectionFields = parseConnectionString(values.connection_string);

  return prepareSubmitValuesFromFields({ ...values, ...connectionFields });
}
