const glob = require('glob')
const path = require('path')

const getTablesData = () => {
  const tables = glob.globSync('tables/**/*.md')
  const tablesData = tables.map((file) => {
    const withoutDir = file.replace('tables/', '')
    const [plugin, tableFile] = withoutDir.split('/')
    const table = path.basename(tableFile, '.md')
    return {
        plugin,
        table,
    }
  })

  tablesData.sort((a, b) => a.table.localeCompare(b.table));
  return tablesData
};

module.exports = {
    getTablesData,
};
