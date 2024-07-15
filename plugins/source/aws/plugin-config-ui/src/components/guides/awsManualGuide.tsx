import { Link, Stack, Typography } from '@mui/material';
import { Sections } from './sections';

const GENERIC_SECTIONS = [
  {
    bodies: [
      {
        text: 'Follow the instructions below to create a policy and an IAM role that allows CloudQuery to read your infrastructure configuration.',
      },
    ],
  },
  {
    header: 'Full Visibility',
    bodies: [
      {
        text: `Once you connect your AWS account, you will be able to run a sync to a destination of your choice to get a full visibility into your cloud infrastructure`,
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
    bodies: [
      { text: 'Proud member of the AWS Partner Network' },
      {
        text: 'This will guide you through creating an IAM role that allows CloudQuery to read your infrastructure configuration.',
      },
      {
        text: 'You will be taken to AWS Console and will be prompted to create a new IAM Role with a read-only inline policy using CloudFormation.',
      },
      {
        text: 'Once the Stack is created, you will see the ARN of the new IAM role in the output. Return to this page to enter it and continue the setup.',
      },
    ],
  },
];

const AWS_MANUAL_SECTIONS = [
  {
    header: 'Step 1: Create IAM role',
    bodies: [
      {
        Node: () => (
          <div>
            Open the [AWS IAM Console](
            <Link target="_blank" href="https://console.aws.amazon.com/iam/home#home">
              https://console.aws.amazon.com/iam/home#home
            </Link>
            ), navigate to Access management - Roles. Then click Create Role.
          </div>
        ),
      },
      {
        image: '/screenshots/createIAM1a.webp',
        text: 'Create IAM role',
      },
      {
        code: `{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Sid": "AllowAssumeRole",
			"Effect": "Allow",
			"Principal": {
			    "AWS": "arn:aws:iam::767397982801:role/managed-syncs-source-aws-role-for-cross-account-access"
			},
			"Action": "sts:AssumeRole",
			"Condition": {
			    "StringEquals": {
			        "sts:ExternalId": "fa17225d-70f6-4a2f-b001-bf63c49869a3"
			    }
			}
		}
	]
}`,
      },
      {
        image: '/screenshots/createIAM1b.webp',
        text: 'Select trusted entity',
      },
    ],
  },
  {
    header: 'Step 2: Click Next to specify the permissions',
    bodies: [
      {
        text: `In the list of permission policies, filter by type “AWS managed - job function” and check 
the box next to the policy with the name “ReadOnlyAccess”.`,
      },
      {
        image: '/screenshots/createIAM2.webp',
        text: 'Add permissions',
      },
    ],
  },
  {
    header: 'Step 3: Click Next',
    bodies: [
      {
        text: `Set the role name to “CloudQueryIntegrationRoleForAWSSource”`,
      },
      {
        image: '/screenshots/createIAM3.webp',
        text: 'Name, review, and create',
      },
    ],
  },
  {
    header: 'Step 4: Scroll down to click the “Create role”',
    bodies: [
      {
        image: '/screenshots/createIAM4.webp',
        text: 'Create role',
      },
    ],
  },
  {
    header: 'Step 5: Select the role that you just created.',
    bodies: [
      {
        image: '/screenshots/createIAM5.webp',
        text: 'Select role',
      },
    ],
  },
  {
    header: 'Step 6: In the Summary section, copy the ARN value and enter it in the input below ',
    bodies: [
      {
        image: '/screenshots/createIAM6.webp',
        text: 'Copy ARN',
      },
    ],
  },
];

export function AWSManualConnect() {
  return (
    <Stack gap={3}>
      <Typography variant="h5">AWS connection</Typography>
      <Sections sections={GENERIC_SECTIONS} />
      <Typography variant="h5">Connect Account via the AWS IAM Console</Typography>
      <Sections sections={AWS_MANUAL_SECTIONS} />
    </Stack>
  );
}
