import React from "react";
import Link from "next/link";
import algoliasearch from "algoliasearch/lite";
import {
  InstantSearch,
  connectSearchBox,
  connectStateResults,
} from "react-instantsearch-dom";

const searchClient = algoliasearch(
  "0OUG5EUZ6H",
  "fa5e5811f7275c275157c74b46f11345"
);

const searchHitsClassName =
  "nextra-scrollbar nx-bg-white nx-text-gray-100 dark:nx-bg-neutral-900 nx-absolute nx-top-full nx-z-20 nx-mt-2 nx-overflow-auto nx-overscroll-contain nx-rounded-xl nx-py-2.5 nx-shadow-xl nx-max-h-[min(calc(50vh-11rem-env(safe-area-inset-bottom)),400px)] md:nx-max-h-[min(calc(100vh-5rem-env(safe-area-inset-bottom)),400px)] nx-inset-x-0 ltr:md:nx-left-auto rtl:md:nx-right-auto contrast-more:nx-border contrast-more:nx-border-gray-900 contrast-more:dark:nx-border-gray-50 nx-w-screen nx-min-h-[100px] nx-max-w-[min(calc(100vw-2rem),calc(100%+20rem))]";

const SearchBox = ({ refine }) => {
  return (
    <input
      className="nx-relative nx-flex nx-items-center nx-text-gray-900 contrast-more:nx-text-gray-800 dark:nx-text-gray-300 contrast-more:dark:nx-text-gray-300"
      type="search"
      placeholder="Discover resources..."
      onChange={(e) => refine(e.currentTarget.value)}
    />
  );
};
const ConnectedSearchBox = connectSearchBox(SearchBox);

const SearchHits = ({ searchState, searchResults }) => {
  const validQuery = searchState.query?.length >= 3;

  if (!validQuery) {
    return null;
  }
  const hits = searchResults?.hits || [];
  if (hits.length === 0) {
    return (
      <div className={searchHitsClassName}>
        <div>No results found!</div>
      </div>
    );
  }

  return (
    <ul className={searchHitsClassName}>
      {hits.map((hit) => (
        <li
          key={hit.objectID}
          className="nx-mx-2.5 nx-break-words nx-rounded-md contrast-more:nx-border nx-bg-primary-500/10 nx-text-primary-500 contrast-more:nx-border-primary-500"
        >
          <Link href={hit.docs_link}>
            <a
              className="nx-block nx-scroll-m-12 nx-px-2.5 nx-py-2"
              target="_blank"
              rel="noopener noreferrer"
            >
              {hit.name}
            </a>
          </Link>
        </li>
      ))}
    </ul>
  );
};

const ConnectedSearchHits = connectStateResults(SearchHits);

export const TableSearch = () => {
  return (
    <div className="nextra-search nx-relative md:nx-w-64 nx-hidden md:nx-inline-block mx-min-w-[200px]">
      <InstantSearch searchClient={searchClient} indexName="all-sources-tables">
        <ConnectedSearchBox />
        <ConnectedSearchHits />
      </InstantSearch>
    </div>
  );
};
