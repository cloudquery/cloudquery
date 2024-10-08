import { Link, RenderGuide } from '@cloudquery/plugin-config-ui-lib';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';

import { pluginUiMessageHandler } from '../../utils/messageHandler';

const GENERIC_SECTIONS = [
  {
    bodies: [
      {
        text: 'This will guide you through creating an IAM role that allows CloudQuery to read your infrastructure configuration.',
      },

      {
        text: 'Once the Stack is created, you will see the ARN of the new IAM role in the output. Return to this page to enter it and continue the setup.',
      },
    ],
  },
];

const getAwsManualSections = ({
  externalId,
  externalIdCreate,
  externalIdEdit,
}: {
  externalId: string;
  externalIdCreate?: boolean;
  externalIdEdit?: boolean;
}) => {
  const showExternalIdStep = externalIdCreate || externalIdEdit;

  return [
    {
      header: 'Step 1: Enter the name of the S3 bucket',
      bodies: [
        {
          text: 'Enter the name of the S3 bucket CloudQuery will sync to.',
        },
      ],
    },
    {
      header: 'Step 2: Create IAM role',
      bodies: [
        {
          text: (
            <>
              Open the{' '}
              <Link
                href="https://us-east-1.console.aws.amazon.com/iam/home#/roles"
                pluginUiMessageHandler={pluginUiMessageHandler}
              >
                AWS IAM Console
              </Link>
              , navigate to Access management - Roles. Then click Create Role.
            </>
          ),
        },
        {
          text: 'In the Trusted Entity Type, select the “Custom Trust Policy”.',
        },
        {
          image: `images/createIAM1a.webp`,
          text: 'Create IAM role',
        },
        {
          text: 'In the policy editor, insert the following JSON:',
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
			        "sts:ExternalId": "${externalId}"
			    }
			}
		}
	]
}`,
        },
        {
          image: `images/createIAM1b.webp`,
          text: 'Select trusted entity',
        },
      ],
    },
    {
      header: 'Step 3: Click Next to specify the permissions',
      bodies: [
        {
          text: `In the list of permission policies, filter by type “AWS managed - job function” and check 
the box next to the policy with the name “ReadOnlyAccess”.`,
        },
        {
          image: `images/createIAM2.webp`,
          text: 'Add permissions',
        },
      ],
    },
    {
      header: 'Step 4: Click Next',
      bodies: [
        {
          text: `Set the role name to “CloudQueryIntegrationRoleForAWSSource”`,
        },
        {
          image: `images/createIAM3.webp`,
          text: 'Name, review, and create',
        },
      ],
    },
    {
      header: 'Step 5: Scroll down to click the “Create role”',
      bodies: [
        {
          image: `images/createIAM4.webp`,
          text: 'Create role',
        },
      ],
    },
    {
      header: 'Step 6: Select the role that you just created',
      bodies: [
        {
          image: `images/createIAM5.webp`,
          text: 'Select role',
        },
      ],
    },
    {
      header: 'Step 7: Get the ARN',
      bodies: [
        {
          text: 'In the Summary section, copy the ARN value and enter it in the input on the left ',
        },
        {
          image: `images/createIAM6.webp`,
          text: 'Copy ARN',
        },
      ],
    },
    showExternalIdStep && {
      header: 'Step 8: The External ID',
      bodies: [
        {
          text: (
            <>
              At times, you need to give a third party access to your AWS resources (delegate
              access). One important aspect of this scenario is the{' '}
              <Link
                href="https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html"
                pluginUiMessageHandler={pluginUiMessageHandler}
              >
                External ID
              </Link>
              , optional information that you can use in an IAM role trust policy to designate who
              can assume the role.
            </>
          ),
        },
        externalIdCreate && {
          text: 'The pre-populated External ID matches to the Trust Policy json above.',
        },
        {
          text: 'If you want to use a custom External ID, you can overwrite it on the left. Ensure that the new value matches to your Trust Policy in AWS.',
        },
        {
          image: `images/createIAM7.webp`,
          text: 'Check External ID',
        },
      ].filter(Boolean),
    },
    {
      header: `Step ${showExternalIdStep ? 9 : 8}: Proceed to the next page`,
      bodies: [
        {
          text: `Click the Continue button on the left to select the data to sync. `,
        },
      ],
    },
  ].filter(Boolean);
};

export function AWSManualConnect({
  externalId = '',
  externalIdCreate,
  externalIdEdit,
}: {
  externalId?: string;
  externalIdCreate?: boolean;
  externalIdEdit?: boolean;
}) {
  const AWS_MANUAL_SECTIONS = getAwsManualSections({
    externalId,
    externalIdCreate,
    externalIdEdit,
  });

  return (
    <Stack gap={3}>
      <RenderGuide pluginUiMessageHandler={pluginUiMessageHandler} sections={GENERIC_SECTIONS} />
      <Typography variant="h5">Connect Account via the AWS IAM Console</Typography>
      <RenderGuide
        pluginUiMessageHandler={pluginUiMessageHandler}
        sections={AWS_MANUAL_SECTIONS as any}
      />
    </Stack>
  );
}
