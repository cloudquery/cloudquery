import { resetYupDefaultErrorMessages } from '@cloudquery/cloud-ui';
import * as yup from 'yup';

resetYupDefaultErrorMessages(yup);

export const formValidationSchema = yup.object({
  token: yup.string().default('').required(),
});

export type FormValues = yup.InferType<typeof formValidationSchema>;
