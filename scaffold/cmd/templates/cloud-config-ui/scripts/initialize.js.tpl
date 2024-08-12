require('child_process').fork('scripts/set_environment.js');

const path = require('path');
if (require('path').basename(path.join(__dirname, '../../..')) === 'source') {
  require('child_process').fork('scripts/set_tables.js');
}
