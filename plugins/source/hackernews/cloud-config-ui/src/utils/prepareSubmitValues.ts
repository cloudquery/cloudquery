import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { corePrepareSubmitValues, PluginTable } from '@cloudquery/plugin-config-ui-lib';

export function prepareSubmitValues(
  values: Record<string, any>,
  tablesList?: PluginTable[],
): PluginUiMessagePayload['validation_passed']['values'] {
  const payload = corePrepareSubmitValues(values, tablesList);

  if (values.item_concurrency) {
    payload.spec.item_concurrency = Number(values.item_concurrency);
  }

  if (values._startTimeEnabled) {
    payload.spec.start_time = values.start_time.toISOString();
  }

  return payload;
}
