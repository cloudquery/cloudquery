import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { FormValues } from './formSchema';

export function prepareSubmitValues(
  values: FormValues,
): PluginUiMessagePayload['validation_passed']['values'] {
  const envs = [] as Array<{ name: string; value: string }>;

  return {
    name: values.name,
    envs,
    tables: Object.keys(values.tables).filter(
      (key) => values.tables[key as keyof typeof values.tables],
    ),
    spec: {},
  };
}
