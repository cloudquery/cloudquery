import { Divider, Link, Stack, Typography } from '@mui/material';
import { AWSGuideHeader } from './header';
import { Sections } from './sections';

const SERVICES_SECTION = [
  {
    bodies: [
      {
        text: 'Select the services that you want to sync the data for. This will affect the API endpoints that will be queried and the tables created in your destination database.',
      },
      {
        text: 'The most popular services are listed on top. Note that some services may take a while to sync depending on the amount of resources used.',
      },
      {
        Node: () => (
          <div>
            See the [AWS Plugin Documentation](
            <Link target="_blank" href="TODO:link">
              TODO:link
            </Link>
            ) to see the full list of tables and their schema.
          </div>
        ),
      },
    ],
  },
];

export function AWSSelectServices() {
  return (
    <Stack gap={3} p={3}>
      <AWSGuideHeader />
      <Divider />
      <Typography variant="h5">Select services</Typography>
      <Sections sections={SERVICES_SECTION} />
    </Stack>
  );
}
