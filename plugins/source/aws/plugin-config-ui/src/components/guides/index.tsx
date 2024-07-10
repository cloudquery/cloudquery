import { Card, CardContent } from '@mui/material';
import { useFormContext } from 'react-hook-form';
import { useMemo } from 'react';
import { AWSConsoleConnection } from './awsConsoleGuide';
import { AWSManualConnect } from './awsManualGuide';
import { AWSConsoleReconnect } from './awsConsoleReconnectGuide';
import { AWSSelectServices } from './awsSelectServicesGuide';
import { SetupType } from '../../utils/formSchema';

interface Props {}

export function Guides({}: Props) {
  const form = useFormContext();

  const hasARN = !!form.watch('arn');
  const setupType = form.watch('_setupType');

  const Content = useMemo(() => {
    if (setupType === SetupType.Manual) {
      return AWSManualConnect;
    } else if (setupType === SetupType.Console) {
      if (hasARN) {
        return AWSConsoleReconnect;
      } else {
        return AWSConsoleConnection;
      }
    } else {
      return AWSSelectServices;
    }
  }, [hasARN, setupType]);

  return (
    <Card>
      <CardContent>
        <Content />
      </CardContent>
    </Card>
  );
}
