import glob from 'glob'
import path from 'path'
import { buildDynamicMDX, buildDynamicMeta } from 'nextra/remote'
import { promises as fs } from 'fs'

export const getStaticPropsFactory = (plugin: string) => async ({ params: { table } }) => {
  const tableFile = path.join(process.cwd(), `tables/${plugin}/${table}.md`)
  const tableContent = await fs.readFile(tableFile, 'utf8')
  return {
    props: {
      ...(await buildDynamicMDX(tableContent)),
      ...(await buildDynamicMeta())
    },
  }
}

export const getStaticPaths = () => {
  return {
    paths: [],
    fallback: 'blocking'
  }
}

export const getTablesData = () => {
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
