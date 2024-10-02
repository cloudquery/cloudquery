import { AuthType, useFormContext } from '@cloudquery/plugin-config-ui-lib';

import { IAMGuide } from './iam-guide';
import { ServiceAccountGuide } from './sa-guide';
import { ServicesGuide } from './services-guide';

export function Guide() {
  const form = useFormContext();
  const step = form.watch('_step');
  const authType = form.watch('_authType');

  if (step === 1) {
    return <ServicesGuide />;
  } else {
    return authType === AuthType.OAUTH ? <IAMGuide /> : <ServiceAccountGuide />;
  }
}
