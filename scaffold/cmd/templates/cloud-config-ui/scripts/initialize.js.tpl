var argv = require('minimist')(process.argv.slice(2));

require('child_process').fork('scripts/set_environment.js');
require('child_process').fork(`scripts/gen_tables.js`, ['--f', argv.f]);
