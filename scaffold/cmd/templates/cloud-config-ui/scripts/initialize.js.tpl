var argv = require('minimist')(process.argv.slice(2));

require('child_process').fork('scripts/set_environment.js');

const path = require('path');
if (require('path').basename(path.join(__dirname, '../../..')) === 'source') {
  require('child_process').fork(`scripts/gen_tables.js`, ['--f', argv.f]);
}
