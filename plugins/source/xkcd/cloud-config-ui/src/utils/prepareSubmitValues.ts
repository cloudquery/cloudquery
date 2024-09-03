import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { corePrepareSubmitValues, PluginTable } from '@cloudquery/plugin-config-ui-lib';

export function prepareSubmitValues(
  values: Record<string, any>,
  tablesList?: PluginTable[],
): PluginUiMessagePayload['validation_passed']['values'] {
  const payload = corePrepareSubmitValues(values, tablesList);

  return payload;
}
