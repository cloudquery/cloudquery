import { AuthType, useFormContext, usePluginContext } from '@cloudquery/plugin-config-ui-lib';

import { AWSAdvancedGuide } from './awsAdvancedGuide';
import { AWSConsoleConnect } from './awsConsoleConnectGuide';
import { AWSConsoleOverview } from './awsConsoleOverviewGuide';
import { AWSManualConnect } from './awsManualGuide';
import { AWSSelectServices } from './awsSelectServicesGuide';

export function Guides() {
  const { initialValues } = usePluginContext();
  const form = useFormContext();
  const editMode = !!initialValues;

  const connectorId = form.watch('connectorId');
  const externalId = form.watch('externalId');
  const authType = form.watch('_authType');
  const isSelectServices = form.watch('_activeIndex') === 1;
  const isAdvanced = form.watch('_activeIndex') === 2;
  const arnTouched = !!form.formState.dirtyFields.arn;

  if (isAdvanced) {
    return <AWSAdvancedGuide />;
  } else if (isSelectServices) {
    return <AWSSelectServices editMode={editMode} />;
  } else if (authType === AuthType.OTHER) {
    return (
      <AWSManualConnect
        externalId={externalId}
        externalIdCreate={!editMode}
        externalIdEdit={editMode && arnTouched}
      />
    );
  } else if (authType === AuthType.OAUTH && connectorId) {
    return <AWSConsoleConnect />;
  } else if (authType === AuthType.OAUTH) {
    return <AWSConsoleOverview />;
  } else {
    return null;
  }
}
