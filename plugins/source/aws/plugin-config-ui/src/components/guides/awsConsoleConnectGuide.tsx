import { Stack, Typography } from '@mui/material';
import { Sections } from '../todoGetFromShared/sections';

const SECTIONS = [
  {
    header: 'Step 1: Create stack',
    bodies: [
      {
        text: 'In the tab, review the parameters of the CloudFormation Stack and acknowledge that AWS might create IAM resources and click the Create Stack button.',
      },
      {
        image: '/screenshots/cloudFormation1.webp',
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
        image: '/screenshots/cloudFormation2.webp',
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

export function AWSConsoleConnect() {
  return (
    <Stack gap={3}>
      <Typography variant="h5">Create the CloudFormation Stack</Typography>
      <Sections sections={SECTIONS} />
    </Stack>
  );
}
