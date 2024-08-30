import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { corePrepareSubmitValues } from '@cloudquery/plugin-config-ui-lib';

import { convertConnectionStringToFields } from './convertConnectionStringToFields';
import { generateConnectionUrl } from './generateConnectionUrl';

export function prepareSubmitValues(
  values: Record<string, any>,
): PluginUiMessagePayload['validation_passed']['values'] {
  const mutatedValues = { ...values };
  if (values._connectionType === 'string') {
    const { password } = convertConnectionStringToFields(values.connectionString);
    mutatedValues.password = password;
  }

  const payload = corePrepareSubmitValues(mutatedValues);

  payload.migrateMode = values.migrateMode;
  payload.writeMode = values.writeMode;

  payload.spec.connection_string = generateConnectionUrl(mutatedValues);

  if (values.batch_size) {
    payload.spec.batch_size = Number(values.batch_size);
  }

  if (values.batch_size_bytes) {
    payload.spec.batch_size_bytes = Number(values.batch_size_bytes);
  }

  console.log({ payload });

  return payload;
}
