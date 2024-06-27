import { yup } from '../../utils/validation';

export const formValidationSchema = yup.object({
  migrateMode: yup.string(),
  writeMode: yup.string(),
  secrets: yup
    .array()
    .of(
      yup.object({
        name: yup.string().default('').required(),
        value: yup.string().default('').required(),
      }),
    )
    .default([]),
  tables: yup.array().of(yup.string()),
  spec: yup
    .object({
      email: yup.string().default('').email().required(),
      name: yup.string().default('').required(),
    })
    .default({ email: '', name: '' }),
});
