import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { FormValues } from './formSchema';

export function prepareSubmitValues(
  values: FormValues,
): PluginUiMessagePayload['validation_passed']['values'] {
  const envs = [] as Array<{ name: string; value: string }>;

  const spec = {
    start_time: values.spec.startTimeEnabled ? values.spec.startTime.toISOString() : undefined,
    item_concurrency: values.spec.itemConcurrency,
  };

  return {
    name: values.name,
    envs,
    tables: Object.keys(values.tables).filter(
      (key) => values.tables[key as keyof typeof values.tables],
    ),
    spec,
  };
}
