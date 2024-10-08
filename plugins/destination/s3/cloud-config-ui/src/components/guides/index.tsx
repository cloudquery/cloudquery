import { AuthType, useFormContext, usePluginContext } from '@cloudquery/plugin-config-ui-lib';

import { AWSConsoleConnect } from './awsConsoleConnectGuide';
import { AWSConsoleOverview } from './awsConsoleOverviewGuide';
import { AWSManualConnect } from './awsManualGuide';
import { BucketConfigurationGuide } from './bucketConfigurationGuide';

export function Guides() {
  const { initialValues } = usePluginContext();
  const form = useFormContext();
  const editMode = !!initialValues;

  const connectorId = form.watch('connectorId');
  const externalId = form.watch('externalId');
  const authType = form.watch('_authType');
  const step = form.watch('_step');
  const arnTouched = !!form.formState.dirtyFields.arn;

  if (step === 1) {
    return <BucketConfigurationGuide />;
  } else if (authType === AuthType.OTHER) {
    return (
      <AWSManualConnect
        externalId={externalId}
        externalIdCreate={!editMode}
        externalIdEdit={editMode && arnTouched}
      />
    );
  } else if (authType === AuthType.OAUTH && connectorId) {
    return <AWSConsoleOverview />;
  } else if (authType === AuthType.OAUTH) {
    return <AWSConsoleConnect />;
  } else {
    return null;
  }
}
