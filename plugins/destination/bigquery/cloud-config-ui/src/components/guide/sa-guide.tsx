import { Link, RenderGuide, GCPConnect } from '@cloudquery/plugin-config-ui-lib';

import { connectionFields } from './connectionFields';

import { pluginUiMessageHandler } from '../../utils/messageHandler';

export function ServiceAccountGuide() {
  return (
    <RenderGuide
      pluginUiMessageHandler={pluginUiMessageHandler}
      sections={[
        {
          bodies: [
            {
              text: (
                <>
                  CloudQuery integrates with your GCP account using a <b>service account</b>. During
                  the setup, you will grant this service account read access to your GCP project.
                  CloudQuery will have visibility into the cloud inventory data only and nothing
                  else.
                </>
              ),
            },
          ],
        },
        {
          header: 'Step 1: Enable Resource Manager API',
          bodies: [
            {
              text: (
                <>
                  1. Review the{' '}
                  <Link
                    pluginUiMessageHandler={pluginUiMessageHandler}
                    href="https://cloud.google.com/iam/docs/create-service-agents#before-you-begin"
                  >
                    GCP Service Agent Documentation
                  </Link>
                </>
              ),
            },
            {
              text: (
                <>
                  2. Click <b>Enable the API</b>
                </>
              ),
            },
            { image: 'images/sa-1-2.webp' },
            {
              text: (
                <>
                  3. On the top, make sure you select the project you want to grant CloudQuery
                  access to. Then, click <b>Next</b>.
                </>
              ),
            },
            { image: 'images/sa-1-3.webp' },
            {
              text: (
                <>
                  4. Click <b>Enable</b>
                </>
              ),
            },
            { image: 'images/sa-1-4.webp' },
          ],
        },
        {
          header: 'Step 2: Authorize CloudQuery',
          bodies: [
            {
              text: (
                <>
                  1. Open the{' '}
                  <GCPConnect variant="link" pluginUiMessageHandler={pluginUiMessageHandler} />.
                </>
              ),
            },
            {
              text: '2. On the top, make sure you select the project you want to grant CloudQuery access to.',
            },
            { image: 'images/sa-2.webp' },
            {
              text: (
                <>
                  3. From the left, select <b>Service Accounts</b>.
                </>
              ),
            },
            { image: 'images/sa-3.webp' },
            {
              text: (
                <>
                  4. In the main section, click <b>Create Service Account</b>.
                </>
              ),
            },
            { image: 'images/sa-4.webp' },
            {
              text: (
                <>
                  5. Choose a display name and a service account ID, then click{' '}
                  <b>Create and Continue</b>.
                </>
              ),
            },
            { image: 'images/sa-5.webp' },
            {
              text: (
                <>
                  6. In the <b>Grant this service account access to the project</b> section, click{' '}
                  <b>Basic</b> add the <b>Viewer</b> role. Click <b>Done</b>.
                </>
              ),
            },
            { image: 'images/sa-6.webp' },
            {
              text: (
                <>
                  7. Back in the list of service accounts, click the kebab menu button on the right
                  (the three dots) and select <b>Manage Keys</b>.
                </>
              ),
            },
            { image: 'images/sa-7.webp' },
            {
              text: (
                <>
                  8. Using the <b>Add Key</b> dropdown, select <b>Create new key</b>. Use{' '}
                  <b>JSON</b> for the key type and click <b>Create</b>.
                </>
              ),
            },
            { image: 'images/sa-8.webp' },
            {
              text: '9. A file with the key will download automatically. Upload this file using the upload button on the left.',
            },
            {
              text: (
                <>
                  10. You can grant access to other projects to this service account. When ready,
                  click <b>Continue</b> to select the services to sync with CloudQuery.
                </>
              ),
            },
          ],
        },
        {
          header: 'Step 3: Fill in Connection Options',
          bodies: connectionFields,
        },
      ]}
    />
  );
}
