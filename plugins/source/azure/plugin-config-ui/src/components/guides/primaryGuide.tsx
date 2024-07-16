import { Link, Stack, Typography } from '@mui/material';
import { Sections } from './sections';

const GENERIC_SECTIONS = [
  {
    bodies: [
      {
        text: 'Cloudquery integrates with your Azure account by granting read-only access to a Service Principal.',
      },
      {
        text: 'This guide will show you how to create a Service Principal specifically for CloudQuery and assign it Reader permissions to your subscriptions.',
      },
    ],
  },
  {
    header: 'Step 1: Install the Azure CLI',
    bodies: [
      {
        text: (
          <>
            Go to the{' '}
            <Link
              target="_blank"
              href="https://learn.microsoft.com/en-us/cli/azure/install-azure-cli"
            >
              Azure CLI download page
            </Link>{' '}
            and install the az command line tool.
          </>
        ),
      },
    ],
  },
  {
    header: 'Step 2: Log in with the Azure CLI',
    bodies: [
      { text: 'Run the following command in your local terminal:' },
      {
        code: 'az login',
      },
    ],
  },
  {
    header: 'Step 3: Create a service principal',
    bodies: [
      { text: 'Create a service principal the plugin will use to access your cloud deployment.' },
      {
        code: `az provider register --namespace 'Microsoft.Security'

az ad sp create-for-rbac --name cloudquery --role Reader`,
      },
      { text: 'The output of the final command should look like this:' },
      {
        code: `{
        "appId": "YOUR AZURE_CLIENT_ID",
        "displayName": "cloudquery",
        "password": "YOUR AZURE_CLIENT_SECRET",
        "tenant": "YOUR AZURE_TENANT_ID"
        }`,
      },
    ],
  },
  {
    header: 'Step 4: Copy the values',
    bodies: [
      {
        text: (
          <>
            Copy the values for the Azure Tenant ID, Service Principal App ID, and Service Principal
            Password into the input boxes. Then click <b>Save source</b> to save and test the
            connection.
          </>
        ),
      },
    ],
  },
];

export function PrimaryGuide() {
  return (
    <Stack gap={3}>
      <Typography variant="h5">Select services</Typography>
      <Sections sections={GENERIC_SECTIONS} />
    </Stack>
  );
}
