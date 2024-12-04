import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { corePrepareSubmitValues, PluginConfig } from '@cloudquery/plugin-config-ui-lib';

/* eslint-disable unicorn/no-abusive-eslint-disable */
/* eslint-disable */

import { convertConnectionStringToFields } from './convertConnectionStringToFields';
import { generateConnectionStringURI } from './generateConnectionString';

export function prepareSubmitValues(
  config: PluginConfig,
  values: Record<string, any>,
): PluginUiMessagePayload['validation_passed']['values'] {
  const payload =
    values._connectionType === 'string'
      ? prepareSubmitValuesFromConnectionString(config, values)
      : prepareSubmitValuesFromFields(config, values);

  payload.migrateMode = values.migrateMode;
  payload.writeMode = values.writeMode;

  if (values.batch_size) {
    payload.spec.batch_size = Number(values.batch_size);
  }

  if (values.batch_size_bytes) {
    payload.spec.batch_size_bytes = Number(values.batch_size_bytes);
  }

  return payload;
}

function prepareSubmitValuesFromFields(
  config: PluginConfig,
  values: Record<string, any>,
): PluginUiMessagePayload['validation_passed']['values'] {
  const payload = corePrepareSubmitValues(config, values);

  const { connection_string, envs } = generateConnectionStringURI(values);
  payload.spec.connection_string = connection_string;

  payload.envs = envs;

  delete payload.spec.password;
  delete payload.spec.username;
  delete payload.spec.database;

  return payload;
}

function prepareSubmitValuesFromConnectionString(
  config: PluginConfig,
  values: Record<string, any>,
): PluginUiMessagePayload['validation_passed']['values'] {
  const connectionFields = convertConnectionStringToFields(values.connection_string);
  return prepareSubmitValuesFromFields(config, { ...values, ...connectionFields });
}
