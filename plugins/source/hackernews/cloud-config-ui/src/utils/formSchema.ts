import { resetYupDefaultErrorMessages } from '@cloudquery/cloud-ui';
import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { generateDisplayName } from '@cloudquery/plugin-config-ui-lib';

import * as yup from 'yup';

import { default as dayjs } from '../utils/date';

resetYupDefaultErrorMessages(yup);

export function getDefaultStartTime(startTime?: string): dayjs.Dayjs {
  if (!startTime) {
    return dayjs().subtract(1, 'month');
  }

  return dayjs(startTime);
}

export function getFormValidationSchema(
  initialValues?: FormMessagePayload['init']['initialValues'],
) {
  return yup.object({
    displayName: yup
      .string()
      .default(initialValues?.displayName || generateDisplayName('Hacker News'))
      .max(255)
      .required(),
    name: yup
      .string()
      .default(initialValues?.name || '')
      .max(255),
    envs: yup
      .array()
      .of(
        yup.object({
          name: yup.string().default('').required(),
          value: yup.string().default(''),
        }),
      )
      .default(initialValues?.envs || []),

    spec: yup.object({
      itemConcurrency: yup
        .number()
        .default(initialValues?.spec?.item_concurrency ?? 100)
        .required(),
      startTimeEnabled: yup
        .boolean()
        .default(!initialValues || !!initialValues?.spec?.start_time)
        .required(),
      startTime: yup
        .mixed<dayjs.Dayjs>()
        .default(getDefaultStartTime(initialValues?.spec?.start_time)),
    }),

    tables: yup
      .object()
      .test('valid tables', function (value: Record<string, true>) {
        if (Object.keys(value || {}).filter((key) => value[key]).length === 0) {
          return this.createError({
            message: 'At least one table must be selected',
            path: 'tables',
          });
        }

        return true;
      })
      .default(
        initialValues?.tables
          ? Object.fromEntries(initialValues.tables.map((table) => [table, true]))
          : { hackernews_items: true },
      ),
  });
}

export type FormValues = yup.InferType<ReturnType<typeof getFormValidationSchema>>;
