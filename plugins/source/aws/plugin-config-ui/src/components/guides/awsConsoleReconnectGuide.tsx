import { Divider, Stack, Typography } from '@mui/material';
import { AWSGuideHeader } from './header';
import { Sections } from './sections';

const SECTIONS = [
  {
    header: 'Step 1: Create stack',
    bodies: [
      {
        text: 'In the tab, review the parameters of the CloudFormation Stack and acknowledge that AWS might create IAM resources and click the Create Stack button.',
      },
      {
        image: '/screenshots/cloudFormation1.png',
        text: 'Create stack',
      },
    ],
  },
  {
    header: 'Step 2: Find ARN and paste it into the input field',
    bodies: [
      {
        text: `Once the stack creation is complete, find the ARN of the newly created IAM role in the output of the CloudFormation stack and enter it in the field on the left.`,
      },
      {
        image: '/screenshots/cloudFormation2.png',
        text: 'Copy ARN',
      },
    ],
  },
  {
    header: 'Step 3: Test the conneciton',
    bodies: [
      {
        text: `Click the Test Connection button to check if CloudQuery can connect and to continue to select the data to sync.`,
      },
      {
        text: 'Click the Save Source button if you wish to continue later.',
      },
    ],
  },
];

export function AWSConsoleReconnect() {
  return (
    <Stack gap={3} p={3}>
      <AWSGuideHeader />
      <Divider />
      <Typography variant="h5">Create the CloudFormation Stack</Typography>
      <Sections sections={SECTIONS} />
    </Stack>
  );
}
