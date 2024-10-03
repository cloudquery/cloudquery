import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { AuthType, corePrepareSubmitValues, PluginConfig } from '@cloudquery/plugin-config-ui-lib';

export function prepareSubmitValues(
  config: PluginConfig,
  values: Record<string, any>,
): PluginUiMessagePayload['validation_passed']['values'] {
  const payload = corePrepareSubmitValues(config, values);

  payload.migrateMode = values.migrateMode;
  payload.writeMode = 'append'; // it is only option
  payload.spec.project_id = values.project_id;
  payload.spec.dataset_id = values.dataset_id;

  if (values.dataset_location) {
    payload.spec.dataset_location = values.dataset_location;
  }

  if (values.time_partitioning) {
    payload.spec.time_partitioning = values.time_partitioning;
  }

  if (values.batch_size) {
    payload.spec.batch_size = Number(values.batch_size);
  }

  if (values.batch_size_bytes) {
    payload.spec.batch_size_bytes = Number(values.batch_size_bytes);
  }

  if (values._authType === AuthType.OTHER) {
    payload.spec['service_account_key_json'] = '${service_account_key_json}';
    payload.envs.push({
      name: 'service_account_key_json',
      value: values.service_account_key_json,
    });
  }

  return payload;
}
