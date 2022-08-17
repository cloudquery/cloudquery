import { DuplicateIcon } from "@heroicons/react/outline";
import copy from "copy-to-clipboard";
import Head from "next/head";
import Link from "next/link";
import toast, { Toaster } from "react-hot-toast";
import { Container } from "../Container";
import Features from "../Features";
import { Integrations } from "../Integrations";
import { LogosBlock } from "../clients/LogosBlock";

export default function Home() {
  const onClick = () => {
    copy("brew install cloudquery");
    toast.success("Copied to clipboard");
  };

  return (
    <>
      <Head>
        <title>CloudQuery</title>
        <meta
          name="og:description"
          content="CloudQuery is an open source high performance data integration platform designed for security and infrastructure teams"
        />
      </Head>
      <div className="w-auto px-4 pt-16 pb-8 mx-auto sm:pt-24 lg:px-8">
        <h1 className="max-w-5xl text-center mx-auto text-6xl font-extrabold tracking-tighter leading-[1.1] sm:text-7xl lg:text-8xl xl:text-8xl">
          Data integration
          <br className="hidden lg:block" />
          <span className="inline-block text-transparent bg-clip-text bg-gradient-to-r from-green-500 to-blue-500 ">
            that works.
          </span>{" "}
        </h1>
        <p className="max-w-lg mx-auto mt-6 text-xl font-medium leading-tight text-center text-gray-400 sm:max-w-4xl sm:text-2xl md:text-3xl lg:text-4xl">
          CloudQuery is an open source high performance data integration
          platform designed for security and infrastructure teams.
        </p>
        <div className="max-w-xl mx-auto mt-5 sm:flex sm:justify-center md:mt-8">
          <div className="rounded-md ">
            <Link href="/docs">
              <a className="flex items-center justify-center w-full px-8 py-3 text-base font-medium text-white no-underline bg-black border border-transparent rounded-md dark:bg-white dark:text-black betterhover:dark:hover:bg-gray-300 betterhover:hover:bg-gray-700 md:py-3 md:text-lg md:px-10 md:leading-6">
                Get Started →
              </a>
            </Link>
          </div>
          <div className="relative mt-3 rounded-md sm:mt-0 sm:ml-3">
            <button
              onClick={onClick}
              className="flex items-center justify-center w-full px-8 py-3 font-mono text-sm font-medium text-gray-600 bg-black border border-transparent border-gray-200 rounded-md bg-opacity-5 dark:bg-white dark:text-gray-300 dark:border-gray-700 dark:bg-opacity-5 betterhover:hover:bg-gray-50 betterhover:dark:hover:bg-gray-900 md:py-3 md:text-base md:leading-6 md:px-10"
            >
              brew install cloudquery
              <DuplicateIcon className="w-6 h-6 ml-2 -mr-3 text-gray-400" />
            </button>
          </div>
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
          <h2 className="text-4xl font-extrabold tracking-tight lg:text-5xl xl:text-6xl lg:text-center dark:text-white">
            Regain access to your data
          </h2>
          <p className="mx-auto mt-4 text-lg font-medium text-gray-400 lg:max-w-3xl lg:text-xl lg:text-center">
            Apply the best practices in data engineering to solve infrastructure security, compliance, cost and search
            use cases.
          </p>
          <Features />
        </div>
      </div>
      <div className="sm:py-20 lg:py-24">
        <div className="max-w-4xl px-4 pb-12 mx-auto lg:px-8 ">
          <h2 className="text-4xl font-extrabold leading-tight tracking-tight lg:text-5xl xl:text-6xl text-center dark:text-white">
            Integrations
          </h2>
          <p className="mx-auto mt-4 font-medium text-gray-400 lg:max-w-3xl lg:text-xl text-center">
            Integrate with 10+ cloud providers and SaaS apps with more than 1,000 unique tables.
          </p>
          <Integrations />
        </div>
        <Container>
          <div className="px-4 py-16 mx-auto mt-10 sm:max-w-none sm:flex sm:justify-center">
            <div className="space-y-4 sm:space-y-0 sm:mx-auto ">
              <Link href="/docs">
                <a className="flex items-center justify-center w-full px-8 py-3 text-base font-medium text-white no-underline bg-black border border-transparent rounded-md dark:bg-white dark:text-black betterhover:dark:hover:bg-gray-300 betterhover:hover:bg-gray-700 md:py-3 md:text-lg md:px-10 md:leading-6">
                  Start Building →
                </a>
              </Link>
            </div>
          </div>
        </Container>
      </div>
      <Toaster position="bottom-right" />
    </>
  );
}
