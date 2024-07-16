import { Stack, Typography } from '@mui/material';
import { Sections } from '../todoGetFromShared/sections';

const SECTIONS = [
  {
    header: 'Full Visibility',
    bodies: [
      {
        text: `This will guide you through creating an IAM role that allows CloudQuery to read your infrastructure configuration.`,
      },
      {
        text: `You will be taken to AWS Console and will be prompted to create a new IAM Role with a read-only inline policy using CloudFormation.`,
      },
      {
        text: `Once the Stack is created, you will see the ARN of the new IAM role in the output. Return to this page to enter it and continue the setup.`,
      },
    ],
  },
  {
    header: 'Secure',
    bodies: [
      { text: 'Cross account IAM roles are the most secure way to connect to AWS accounts.' },
    ],
  },
  {
    header: 'Read-only Access',
    bodies: [{ text: 'CloudQuery only ingests the data based on read-only permissions.' }],
  },
  {
    header: 'AWS Partner',
    bodies: [{ text: 'Proud member of the AWS Partner Network' }],
  },
];

export function AWSConsoleOverview() {
  return (
    <Stack gap={3}>
      <Typography variant="h5">AWS connection</Typography>
      <Sections sections={SECTIONS} />
    </Stack>
  );
}
