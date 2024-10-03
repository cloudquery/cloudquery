import { RenderGuide, useFormContext, GCPConnect } from '@cloudquery/plugin-config-ui-lib';

import connectionBodyContent from './connectionFields';

import { pluginUiMessageHandler } from '../../utils/messageHandler';

export function IAMGuide() {
  const form = useFormContext();

  // this is the usage that necessitates ejection from the useConfig json
  const serviceAccount = form.watch('_serviceAccount');

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
          header: 'Step 1: Authorize CloudQuery',
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
            { image: 'images/iam-2.webp' },
            {
              text: (
                <>
                  3. In the main section, under the <b>Permissions</b> for your project, click the{' '}
                  <b>Grant Access</b>.
                </>
              ),
            },
            { image: 'images/iam-3.webp' },

            {
              text: (
                <>
                  4. In the <b>Add Principals</b> section, copy and add the below principal:
                </>
              ),
            },
            { image: 'images/iam-4.webp' },

            {
              code: serviceAccount,
            },
            {
              text: (
                <>
                  5. In the <b>Assign Roles</b> section, click <b>Basic</b> and add the{' '}
                  <b>Viewer</b> role.
                </>
              ),
            },
            { image: 'images/iam-5.webp' },

            {
              text: (
                <>
                  6. Click <b>Save</b> to grant the access.
                </>
              ),
            },
            {
              text: (
                <>
                  7. Repeat these steps for each project that you want to sync with CloudQuery. When
                  ready, click <b>Continue</b> to select the services to sync with CloudQuery.
                </>
              ),
            },
          ],
        },
        {
          header: 'Step 2: Fill in Connection Options',
          bodies: connectionBodyContent,
        },
      ]}
    />
  );
}
