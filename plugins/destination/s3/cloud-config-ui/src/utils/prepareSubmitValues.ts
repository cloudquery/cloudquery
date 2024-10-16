import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { corePrepareSubmitValues, PluginConfig } from '@cloudquery/plugin-config-ui-lib';

export function prepareSubmitValues(
  config: PluginConfig,
  values: Record<string, any>,
): PluginUiMessagePayload['validation_passed']['values'] {
  const payload = corePrepareSubmitValues(config, values);

  payload.connectorId = values.connectorId;

  delete payload.spec.arn;
  delete payload.spec.externalId;

  payload.envs = payload.envs.filter((env) => env.name !== 'arn' && env.name !== 'externalId');

  payload.spec.bucket = values.bucket;
  payload.spec.region = values.region;
  payload.spec.path = values.path;
  payload.spec.format = values.format;
  payload.spec.format_spec = {};

  if (values.format === 'csv') {
    payload.spec.format_spec.delimiter = values.format_spec_csv_delimiter;
    payload.spec.format_spec.skip_header = values.format_spec_csv_skip_header;
  } else if (values.format === 'parquet') {
    payload.spec.format_spec.version = values.format_spec_parquet_version;
    payload.spec.format_spec.root_repetition = values.format_spec_parquet_root_repetition;
  }

  if (values.server_side_encryption_configuration_enabled) {
    payload.spec.format_spec.server_side_encryption_configuration = {
      sse_kms_key_id: values.server_side_encryption_configuration_sse_kms_key_id,
      server_side_encryption: values.server_side_encryption_configuration_server_side_encryption,
    };
  }

  if (values.compression && values.format !== 'parquet') {
    payload.spec.compression = values.compression;
  }

  if (values.no_rotate) {
    payload.spec.no_rotate = values.no_rotate;
  }

  if (values.athena) {
    payload.spec.athena = values.athena;
  }

  if (values.format === 'parquet' && values.write_empty_objects_for_empty_tables) {
    payload.spec.write_empty_objects_for_empty_tables = values.write_empty_objects_for_empty_tables;
  }

  if (values.test_write) {
    payload.spec.test_write = values.test_write;
  }

  if (values.endpoint) {
    payload.spec.endpoint = values.endpoint;
  }

  if (values.acl) {
    payload.spec.acl = values.acl;
  }

  if (values.endpoint_skip_tls_verify) {
    payload.spec.endpoint_skip_tls_verify = values.endpoint_skip_tls_verify;
  }

  if (values.use_path_style) {
    payload.spec.use_path_style = values.use_path_style;
  }

  if (!values.no_rotate) {
    if (values.batch_size) {
      payload.spec.batch_size = Number.parseInt(values.batch_size);
    }

    if (values.batch_size_bytes) {
      payload.spec.batch_size_bytes = Number.parseInt(values.batch_size_bytes);
    }

    if (values.batch_timeout) {
      payload.spec.batch_timeout = values.batch_timeout;
    }
  }

  return payload;
}
