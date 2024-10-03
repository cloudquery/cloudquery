import React from 'react';

import { AuthType, useFormContext } from '@cloudquery/plugin-config-ui-lib';

import { IAMGuide } from './iam-guide';
import { ServiceAccountGuide } from './sa-guide';

export function Guide() {
  const form = useFormContext();
  const authType = form.watch('_authType');

  return authType === AuthType.OAUTH ? <IAMGuide /> : <ServiceAccountGuide />;
}
