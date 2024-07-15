import { Card, CardContent } from '@mui/material';
import { useFormContext } from 'react-hook-form';
import { useMemo } from 'react';
import { AWSConsoleOverview } from './awsConsoleOverviewGuide';
import { AWSManualConnect } from './awsManualGuide';
import { AWSConsoleConnect } from './awsConsoleConnectGuide';
import { AWSSelectServices } from './awsSelectServicesGuide';
import { SetupType } from '../../utils/formSchema';

interface Props {}

export function Guides({}: Props) {
  const form = useFormContext();

  const usingConsoleConnection = form.watch('connector_id');
  const setupType = form.watch('_setupType');
  const isSelectServices = form.watch('_activeIndex') === 1;

  const Content = useMemo(() => {
    if (isSelectServices) {
      return AWSSelectServices;
    } else if (setupType === SetupType.Manual) {
      return AWSManualConnect;
    } else if (setupType === SetupType.Console) {
      if (usingConsoleConnection) {
        return AWSConsoleConnect;
      } else {
        return AWSConsoleOverview;
      }
    }
    return () => <></>;
  }, [usingConsoleConnection, setupType, isSelectServices]);

  return (
    <Card>
      <CardContent>
        <Content />
      </CardContent>
    </Card>
  );
}
