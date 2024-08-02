import dayjs from 'dayjs';

import relativeTimePlugin from 'dayjs/plugin/relativeTime';
import utcPlugin from 'dayjs/plugin/utc';

dayjs.extend(utcPlugin);
dayjs.extend(relativeTimePlugin);

// eslint-disable-next-line unicorn/prefer-export-from
export default dayjs;
