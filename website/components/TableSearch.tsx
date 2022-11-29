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

const SearchBox = ({ refine }) => {
  return (
    <input
      className="search-box"
      type="search"
      placeholder="Search..."
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
      <div className={"search-hits"}>
        <div>No results found!</div>
      </div>
    );
  }

  console.log(searchState, searchResults, validQuery, hits);
  return (
    <div className={"search-hits"}>
      {hits.map((hit) => (
        <div key={hit.objectID}>
          <Link href={hit.docs_link}>
            <a
              className="hover:underline"
              target="_blank"
              rel="noopener noreferrer"
            >
              {hit.name}
            </a>
          </Link>
        </div>
      ))}
    </div>
  );
};

const ConnectedSearchHits = connectStateResults(SearchHits);

export const TableSearch = () => {
  return (
    <div className={"algolia-search"}>
      <InstantSearch searchClient={searchClient} indexName="all-sources-tables">
        <ConnectedSearchBox />
        <ConnectedSearchHits />
      </InstantSearch>
    </div>
  );
};
