import { resetYupDefaultErrorMessages } from '@cloudquery/cloud-ui';
import { generateDisplayName } from '@cloudquery/plugin-config-ui-lib';
import * as yup from 'yup';

resetYupDefaultErrorMessages(yup);

export const formValidationSchema = yup.object({
  displayName: yup
    .string()
    .default(generateDisplayName('XKCD'))
    .max(255)
    .required(),
  name: yup
    .string()
    .default('')
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
  spec: yup.object({}),
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
    .default({ xkcd_comics: true }),
});

export type FormValues = yup.InferType<typeof formValidationSchema>;
