import { resetYupDefaultErrorMessages } from '@cloudquery/cloud-ui';
import * as yup from 'yup';

resetYupDefaultErrorMessages(yup);

export const formValidationSchema = yup.object({
  // defaults
  envs: yup
    .array()
    .of(
      yup.object({
        name: yup.string().default('').required(),
        value: yup.string().default(''),
      }),
    )
    .default([]),
  // user entry
  name: yup.string().default('').required(),
  services: yup.array().of(yup.string().required()).default([]), // => tables

  // form state
  _activeIndex: yup.number().default(0),
});

export type FormValues = yup.InferType<typeof formValidationSchema>;
