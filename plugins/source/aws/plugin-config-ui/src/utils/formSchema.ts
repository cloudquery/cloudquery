import { resetYupDefaultErrorMessages } from '@cloudquery/cloud-ui';
import * as yup from 'yup';

export enum SetupType {
  Manual = 'manual',
  Console = 'console',
}
export const setupTypes = Object.values(SetupType);

resetYupDefaultErrorMessages(yup);

// TODO: the shape of this, per the API
export const formValidationSchema = yup.object({
  // defaults
  migrateMode: yup
    .string()
    .oneOf(['forced', 'safe'] as const)
    .default('safe'),
  writeMode: yup
    .string()
    .oneOf(['append', 'overwrite', 'overwrite-delete-stale'] as const)
    .default('append'),
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
  arn: yup.string().default('').required(),
  services: yup.array().of(yup.string().required()).default([]), // => tables

  // derived
  spec: yup
    .object({
      regions: yup.array().of(yup.string().required()).default([]),
      accounts: yup
        .array()
        .of(yup.object({ role_arn: yup.string().required() }).required())
        .default([]),
    })
    .required(),
  // form state
  _setupType: yup.string().oneOf(setupTypes).default(SetupType.Console).strip(true),
});

export type FormValues = yup.InferType<typeof formValidationSchema>;
