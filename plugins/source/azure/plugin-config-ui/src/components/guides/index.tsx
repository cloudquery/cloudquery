import { Box, Card, CardContent, CardHeader, Divider, Stack } from '@mui/material';
import { useFormContext } from 'react-hook-form';
import { useMemo } from 'react';
import { PrimaryGuide } from './primaryGuide';
import { AzureSelectServices } from './selectServicesGuide';
import { AzureGuideHeader } from './header';

interface Props {}

export function Guides({}: Props) {
  const form = useFormContext();

  const isSelectServices = form.watch('_activeIndex') === 1;

  const Content = useMemo(() => {
    if (isSelectServices) {
      return AzureSelectServices;
    } else {
      return PrimaryGuide;
    }
  }, [isSelectServices]);

  return (
    <Card>
      <CardHeader></CardHeader>
      <CardContent sx={{ pt: 0 }}>
        <Stack gap={3} p={3} pt={0}>
          <AzureGuideHeader />
          <Divider />
          <Box height="calc(100vh - 300px)" sx={{ overflowY: 'auto' }}>
            <Content />
          </Box>
        </Stack>
      </CardContent>
    </Card>
  );
}
