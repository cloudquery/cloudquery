import { Card, CardContent } from '@mui/material';
import { useFormContext } from 'react-hook-form';
import { useMemo } from 'react';
import { AWSConsoleConnection } from './awsConsoleGuide';
import { AWSManualConnect } from './awsManualGuide';
import { AWSConsoleReconnect } from './awsConsoleReconnectGuide';
import { AWSSelectServices } from './awsSelectServicesGuide';

interface Props {}

export function Guides({}: Props) {
  const form = useFormContext();

  const hasARN = !!form.watch('arn');
  const setupType = form.watch('_setupType');

  const Content = useMemo(() => {
    if (setupType === 'manual') {
      return AWSManualConnect;
    } else if (setupType === 'console') {
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
