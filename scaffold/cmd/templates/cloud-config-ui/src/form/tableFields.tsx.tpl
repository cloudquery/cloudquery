import { PluginTable } from '@cloudquery/plugin-config-ui-lib';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';

import { PluginTableSelector } from '../components/tableSelector';

interface Props {
  tablesList: PluginTable[];
}
export function TableFields({ tablesList }: Props) {
  return (
    <Card>
      <CardContent>
        <Stack gap={2}>
          <Stack gap={1}>
            <Typography variant="h6">Table Selection</Typography>
          </Stack>

          <PluginTableSelector pluginTables={tablesList} />
        </Stack>
      </CardContent>
    </Card>
  );
}
