import { resetYupDefaultErrorMessages } from '@cloudquery/cloud-ui';
import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { generateName } from '@cloudquery/plugin-config-ui-lib';
import { default as dayjs } from '../utils/date';
import * as yup from 'yup';

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
    name: yup
      .string()
      .default(generateName('hackernews'))
      .matches(
        /^[a-z](-?[\da-z]+)+$/,
        'Name must consist of a lower case letter, followed by alphanumeric segments separated by single dashes',
      )
      .max(255)
      .required(),
    envs: yup
      .array()
      .of(
        yup.object({
          name: yup.string().default('').required(),
          value: yup.string().default(''),
        }),
      )
      .default([]),

    spec: yup.object({
      itemConcurrency: yup
        .number()
        .default(initialValues?.spec?.item_concurrency ?? 100)
        .required(),
      startTimeEnabled: yup.boolean().default(!!initialValues?.spec?.startTime).required(),
      startTime: yup
        .mixed<dayjs.Dayjs>()
        .default(getDefaultStartTime(initialValues?.spec?.startTime)),
    }),
  });
}

export type FormValues = yup.InferType<ReturnType<typeof getFormValidationSchema>>;
