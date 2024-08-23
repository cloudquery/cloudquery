import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { convertStringToSlug } from '@cloudquery/plugin-config-ui-lib';

import { FormValues } from './formSchema';

export function prepareSubmitValues(
  values: FormValues,
): PluginUiMessagePayload['validation_passed']['values'] {
  const envs = [] as Array<{ name: string; value: string }>;

  return {
    displayName: values.displayName,
    name: values.name || convertStringToSlug(values.displayName),
    envs,
    tables: Object.keys(values.tables).filter(
      (key) => values.tables[key as keyof typeof values.tables],
    ),
    spec: {},
  };
}
