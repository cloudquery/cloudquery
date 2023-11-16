import {ALL_PLUGINS, DESTINATION_CATEGORIES, SOURCE_CATEGORIES} from "./pluginData";

import {LogoContainer} from "./LogoContainer";
import {CategoryTitle} from "./Category";
import {useState} from "react";
import {KINDS, KindTitle} from "./Kind";

const sourceCategories = [{
  id: "",
  name: "All",
}, ...SOURCE_CATEGORIES.map(cat => {
  return {
    id: cat,
    name: CategoryTitle(cat),
  }
})];

const destinationCategories = [{
  id: "",
  name: "All",
}, ...DESTINATION_CATEGORIES.map(cat => {
  return {
    id: cat,
    name: CategoryTitle(cat),
  }
})];

const kinds = KINDS.map(k => {
  return {
    id: k,
    name: KindTitle(k),
  }
});

const filters = [
  {
    id: 'availability',
    name: 'Availability',
    options: [
      { value: 'free', label: 'Free', checked: true },
      { value: 'premium', label: 'Premium', checked: true },
      { value: 'unpublished', label: 'Pre-order', checked: false },
    ],
  },
]

export function Plugins() {
  const [search, setSearch] = useState('');
  const [category, setCategory] = useState('');
  const [availability, setAvailability] = useState(['free', 'premium']);
  const [kind, setKind] = useState('source');
  const selectedPlugins = ALL_PLUGINS.filter(plugin => {
    let match = plugin.kind === kind;
    if (category !== '') {
      match = match && plugin.category === category;
    }
    if (search !== '') {
      let matchesName = plugin.name.toLowerCase().includes(search.toLowerCase());
      let matchesCategory = CategoryTitle(plugin.category).toLowerCase().includes(search.toLowerCase());
      match = match && (matchesName || matchesCategory);
    }
    let pa = plugin.availability;
    switch (pa) {
      case 'community':
        pa = 'free';
        break;
      case 'partner':
        pa = 'free';
        break;
    }
    match = match && availability.includes(pa);
    return match;
  });
  const categories = kind === 'source' ? sourceCategories : destinationCategories;
  return (
        <div>
          <main className="mx-auto max-w-8xl px-4 sm:px-6 lg:px-8">
            <section>
              <div className="relative max-w-xl my-6 mx-auto">
                <div className="absolute inset-y-0 flex items-center pl-3 pointer-events-none">
                  <svg className="w-4 h-4 text-gray-500 dark:text-gray-400" aria-hidden="true"
                       xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20">
                    <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2"
                          d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"/>
                  </svg>
                </div>
                <input type="search" id="default-search"
                       className="block w-full p-2 pl-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                       placeholder="Search" onChange={(e) => {
                          setSearch(e.target.value);
                }} />
              </div>
            </section>
            <section aria-labelledby="filters-heading" className="pb-24 pt-6">
              <div className="grid grid-cols-1 gap-x-8 gap-y-10 lg:grid-cols-4">
                <form className="hidden lg:block" onSubmit={(e) => {e.preventDefault();}}>
                  <h3 className="nx-text-2xl font-extrabold leading-tight tracking-tight lg:nx-text-3xl xl:nx-text-4xl dark:text-white pb-4">Filters</h3>
                  {filters.map((section) => (
                      <div key={section.id} className="space-y-4 pb-6 border-b border-gray-200">
                        {section.options.map((option, optionIdx) => (
                            <div key={option.value} className="flex items-center">
                              <input
                                  id={`filter-${section.id}-${optionIdx}`}
                                  name={`${section.id}[]`}
                                  defaultValue={option.value}
                                  type="checkbox"
                                  defaultChecked={option.checked}
                                  className="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                                  onChange={(e) => {
                                    if (section.id === 'availability') {
                                      let av = [...availability];
                                      if (e.target.checked) {
                                          av.push(e.target.value);
                                      } else {
                                          av = av.filter(item => item !== e.target.value);
                                      }
                                      setAvailability(av);
                                    }
                                  }}
                              />
                              <label
                                  htmlFor={`filter-${section.id}-${optionIdx}`}
                                  className="ml-3 text-sm text-gray-500 dark:text-gray-100"
                              >
                                {option.label}
                              </label>
                            </div>
                        ))}
                      </div>
                  ))}
                  <ul role="list" className="space-y-1 py-6 text-sm pb-6 border-b border-gray-200">
                    {kinds.map((k) => (
                        <li key={k.id} className={"my-0 " + (kind == k.id ? "bg-neutral-100 dark:bg-neutral-600" : "")}>
                          <a href={"#" + k} className={"block p-3 " + (kind == k.id ? "font-bold" : "font-medium")} onClick={(e) => {
                            setKind(k.id);
                            e.preventDefault();
                          }
                          }>{k.name}</a>
                        </li>
                    ))}
                  </ul>
                  <ul role="list" className="space-y-1 py-6 text-sm font-medium">
                    {categories.map((c) => (
                        <li key={c.id} className={"my-0 " + (category == c.id ? "bg-neutral-100 dark:bg-neutral-600" : "")}>
                          <a href={"#" + c.id} className={"block p-3 " + (category == c.id ? "font-bold" : "font-medium")} onClick={(e) => {
                              setCategory(c.id);
                              e.preventDefault();
                            }
                          }>{c.name}</a>
                        </li>
                    ))}
                  </ul>
                </form>

                <div className="lg:col-span-3">
                  <div className="flex justify-center items-center flex-wrap gap-9 sm:mt-4">
                    {
                      selectedPlugins.length > 0 ?
                      selectedPlugins.map((plugin) => {
                        return <LogoContainer
                            title={plugin.name}
                            href={(plugin.kind === 'source') ? (
                                plugin.availability === 'premium'
                                    ? `https://hub.cloudquery.io/plugins/source/cloudquery/${plugin.id}`
                                    : `/integrations/${plugin.id}`
                            ) : `/docs/plugins/destinations/${plugin.id}/overview`}
                            key={plugin.id}
                            external={false}
                            logo={plugin.logo}
                            logoDark={plugin.logoDark}
                            name={plugin.name}
                            availability={plugin.availability}
                            category={plugin.category}
                        ></LogoContainer>
                      }) :
                      <>
                        <p className="text-center font-light text-gray-500 sm:text-xl dark:text-gray-400">
                          No plugins matched the search criteria.<br/>Try adjusting the search and filters or <a href="https://github.com/cloudquery/cloudquery/issues/new/choose" className="font-bold dark:text-white">request a new plugin</a>.
                        </p>
                      </>
                    }
                  </div>
                </div>
              </div>
            </section>
          </main>
        </div>
  );
}
