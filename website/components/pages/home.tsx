import { DocumentDuplicateIcon } from "@heroicons/react/24/outline";
import copy from "copy-to-clipboard";
import Head from "next/head";
import Link from "next/link";
import toast, { Toaster } from "react-hot-toast";
import { Container } from "../Container";
import Features from "../Features";
import { Integrations } from "../Integrations";
import { LogosBlock } from "../clients/LogosBlock";
import { QueriesExamples } from "../QueriesExamples";

const HERO_IMAGE_DATA = [
  {
    src: '/images/hero/section1.svg',
    title: 'Extract from sources',
    className: 'rounded-t-2xl ',
  },
  {
    src: '/images/hero/section2.svg',
    title: 'Load to destination',
  },
  {
    src: '/images/hero/section3.svg',
    title: 'Transform',
    description: 'Run SQL policies and create views',
  },
  {
    src: '/images/hero/section4.svg',
    title: 'Visualize',
    description: 'Connect to your BI stack (Grafana, Preset)',
    className: 'rounded-b-2xl ',
  },
]

export default function Home() {
  const onClick = (code: string) => {
    copy(code);
    toast.success("Copied to clipboard");
  };

  return <>
    <Head>
      <title>CloudQuery</title>
      <meta
        name="og:description"
        content="CloudQuery is an open source high performance data integration platform built for developers."
      />
    </Head>

    <div className="flex flex-col md:flex-row justify-between px-4 pt-16 pb-8 mx-auto sm:pt-24 lg:px-8 w-auto lg:max-w-7xl">
      <div className="flex flex-col justify-between md:mr-4">
        <div>
          <h1 className="max-w-5xl mx-auto nx-text-6xl font-extrabold tracking-tighter leading-[1.1] sm:text-7xl lg:nx-text-8xl xl:nx-text-8xl">
             Frustratingly
             <br className="hidden lg:block" />
             <span className="pr-1 inline-block text-transparent bg-clip-text bg-gradient-to-r from-green-500 to-blue-500 ">
             reliable
             </span>
             &nbsp;
             ELT
          </h1>
          <p className="max-w-2xl mt-6 nx-text-xl font-medium leading-tight text-gray-400 sm:nx-text-2xl md:nx-text-3xl lg:nx-text-4xl">
            Sync any source to any destination, transform and visualize.
          </p>
          <p className="max-w-2xl mt-6 nx-text-xl font-medium leading-tight text-gray-400 sm:nx-text-2xl md:nx-text-3xl lg:nx-text-4xl">
          CloudQuery is an open source high performance data integration platform built for developers.
          </p>
        </div>
        <div className="flex fix-flex-col h-32 mt-4 md:mt-0 mb-6 md:mb-0 items-center gap-3 md:flex-row xl:flex-row">
          <div className="rounded-md fix-width-auto xl:w-auto">
            <Link
              href="/docs/quickstart"
              className="flex items-center justify-center w-full px-8 py-3 text-base font-medium text-white no-underline bg-black border border-transparent rounded-md dark:bg-white dark:text-black betterhover:dark:hover:bg-gray-300 betterhover:hover:bg-gray-700 md:py-3 md:text-lg md:px-10 md:leading-6">
              
                Get Started →
              
            </Link>
          </div>
          <div className="relative rounded-md fix-width-auto xl:w-auto">
            <button
              onClick={() => onClick('brew install cloudquery/tap/cloudquery')}
              className="flex items-center justify-center w-full px-8 py-3 font-mono text-sm font-medium text-gray-600 bg-black border border-transparent border-gray-200 rounded-md bg-opacity-5 dark:bg-white dark:text-gray-300 dark:border-gray-700 dark:bg-opacity-5 betterhover:hover:bg-gray-50 betterhover:dark:hover:bg-gray-900 md:py-3 md:text-base md:leading-6 md:px-10"
            >
              brew install cloudquery/tap/cloudquery
              <DocumentDuplicateIcon className="w-6 h-6 ml-2 -mr-3 text-gray-400" />
            </button>
          </div>
        </div>
      </div>
      <div className="w-full max-w-[436px] flex flex-col gap-1.5 m-auto md:justify-start">
        {HERO_IMAGE_DATA.map(({ title, src, className, description }) => (
          <div key={title} className={`hero-image-bg h-[80px] md:h-[106px] overflow-hidden flex items-center justify-between ${className ? className : ''}`}>
            <div className="pr-1 text-white font-semibold text-base md:text-lg ml-6 leading-tight">
              {title}
              {description &&
                <div className="mt-1 font-normal nx-text-xs leading-none">
                  {description}
                </div>
              }
            </div>
            <img src={src} height='inherit' alt={title} className="h-[inherit]" />
          </div>
        ))}
      </div>
    </div>

    <div className="py-16">
      <div className="mx-auto">
        <p className="pb-8 text-sm font-semibold tracking-wide text-center text-gray-400 uppercase dark:text-gray-500">
          Trusted by teams from around the world
        </p>
        <LogosBlock />
      </div>
    </div>

    <div className="relative from-gray-50 to-gray-100">
      <div className="px-4 py-16 mx-auto sm:pt-20 sm:pb-24 lg:max-w-7xl lg:pt-24">
        <h2 className="nx-text-4xl font-extrabold tracking-tight lg:nx-text-5xl xl:nx-text-6xl lg:text-center dark:text-white">
          Regain access to your data
        </h2>
        <p className="mx-auto mt-4 text-lg font-medium text-gray-400 lg:max-w-3xl lg:text-xl lg:text-center">
          Apply the best practices in data engineering to solve infrastructure security, compliance, cost and search
          use cases.
        </p>
        <Features />
      </div>
    </div>

    <div className="relative from-gray-50 to-gray-100">
      <div className="px-4 py-16 mx-auto sm:pt-20 sm:pb-24 lg:max-w-7xl lg:pt-24">
        <h2 className="nx-text-4xl font-extrabold tracking-tight lg:nx-text-5xl xl:nx-text-6xl lg:text-center dark:text-white">
            Built for any team
        </h2>
        <p className="mx-auto mt-4 text-lg font-medium text-gray-400 lg:max-w-3xl lg:text-xl lg:text-center">
          Engineering, Infrastructure, Security, Marketing, Sales, Support, FinOps
        </p>
        <QueriesExamples onClick={onClick} />
      </div>
    </div >

    <div className="sm:py-20 lg:py-24">
      <div className="max-w-4xl px-4 pb-12 mx-auto lg:px-8 ">
        <h2 className="nx-text-4xl font-extrabold leading-tight tracking-tight lg:nx-text-5xl xl:nx-text-6xl text-center dark:text-white">
          Integrations
        </h2>
        <p className="mx-auto mt-4 font-medium text-gray-400 lg:max-w-3xl lg:nx-text-xl text-center">
          Integrate with a growing list of <a className="dark:text-white" href="/docs/plugins/sources/overview">30+ cloud providers and SaaS apps</a> with more than 1,000 unique tables. Sync to your <a href="/docs/plugins/destinations/overview" className="dark:text-white">favorite database, data warehouse or data lake</a>.
        </p>
        <Integrations />
      </div>
      <Container>
        <div className="px-4 py-16 mx-auto mt-10 sm:max-w-none sm:flex sm:justify-center">
          <div className="space-y-4 sm:space-y-0 sm:mx-auto ">
            <Link
              href="/docs"
              className="flex items-center justify-center w-full px-8 py-3 text-base font-medium text-white no-underline bg-black border border-transparent rounded-md dark:bg-white dark:text-black betterhover:dark:hover:bg-gray-300 betterhover:hover:bg-gray-700 md:py-3 md:text-lg md:px-10 md:leading-6">
              
                Start Syncing →
              
            </Link>
          </div>
        </div>
      </Container>
    </div>
    <Toaster position="bottom-right" />
  </>;
}
