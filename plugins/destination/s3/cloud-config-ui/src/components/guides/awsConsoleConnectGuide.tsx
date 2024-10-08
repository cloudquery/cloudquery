import { RenderGuide } from '@cloudquery/plugin-config-ui-lib';

import { pluginUiMessageHandler } from '../../utils/messageHandler';

const SECTIONS = [
  {
    header: 'Step 1: Navigate to the AWS Console',
    bodies: [
      {
        text: `Click the Reconnect CloudQuery via AWS Console button`,
      },
      {
        text: 'You will be taken to AWS Console and will be prompted to create a new IAM Role with a read-only inline policy using CloudFormation.',
      },
    ],
  },
  {
    header: 'Step 2: Create stack',
    bodies: [
      {
        text: 'In the tab, review the parameters of the CloudFormation Stack and acknowledge that AWS might create IAM resources and click the Create Stack button.',
      },
      {
        image: `images/cloudFormation1.webp`,
        text: 'Create stack',
      },
    ],
  },
  {
    header: 'Step 3: Find ARN and paste it into the input field',
    bodies: [
      {
        text: `Once the stack creation is complete, find the ARN of the newly created IAM role in the output of the CloudFormation stack and enter it in the field on the left.`,
      },
      {
        image: `images/cloudFormation2.webp`,
        text: 'Copy ARN',
      },
    ],
  },
  {
    header: 'Step 4: Proceed to the next page',
    bodies: [
      {
        text: `Click the Continue button on the left to select the data to sync. `,
      },
    ],
  },
];

export function AWSConsoleConnect() {
  return <RenderGuide pluginUiMessageHandler={pluginUiMessageHandler} sections={SECTIONS} />;
}
