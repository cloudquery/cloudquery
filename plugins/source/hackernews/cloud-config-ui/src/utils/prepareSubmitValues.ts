import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import {
  corePrepareSubmitValues,
  PluginConfig,
  PluginTable,
} from '@cloudquery/plugin-config-ui-lib';

export function prepareSubmitValues(
  config: PluginConfig,
  values: Record<string, any>,
  tablesList?: PluginTable[],
): PluginUiMessagePayload['validation_passed']['values'] {
  const payload = corePrepareSubmitValues(config, values, tablesList);

  if (values.item_concurrency) {
    payload.spec.item_concurrency = Number(values.item_concurrency);
  }

  if (values.start_time) {
    payload.spec.start_time = values.start_time.toISOString();
  }

  return payload;
}
