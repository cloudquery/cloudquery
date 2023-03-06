import path from 'path'
import { buildDynamicMDX, buildDynamicMeta } from 'nextra/remote'
import { promises as fs } from 'fs'
import { getTablesData as getTablesDataJS } from './tables-data'

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

export const getTablesData = getTablesDataJS;
