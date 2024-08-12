const fs = require('fs');

let tables = [];
if (fs.existsSync('../docs/__tables.json')) {
  tables = JSON.parse(fs.readFileSync('../docs/__tables.json', 'utf8'));
}

if (!fs.existsSync('./src/data')) {
  fs.mkdirSync('./src/data');
}

fs.writeFileSync('./src/data/tables.json', JSON.stringify(tables, null, 2));
