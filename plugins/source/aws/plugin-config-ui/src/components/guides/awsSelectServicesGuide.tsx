import { Link, Stack, Typography } from '@mui/material';
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
        text: (
          <>
            See the [
            <Link
              target="_blank"
              href="https://hub.cloudquery.io/plugins/source/cloudquery/aws/latest/docs"
            >
              AWS Plugin Documentation
            </Link>
            ] to see the full list of tables and their schema.
          </>
        ),
      },
    ],
  },
];

export function AWSSelectServices() {
  return (
    <Stack gap={3}>
      <Typography variant="h5">Select services</Typography>
      <Sections sections={SERVICES_SECTION} />
    </Stack>
  );
}
