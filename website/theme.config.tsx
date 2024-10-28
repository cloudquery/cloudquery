import { DocsThemeConfig } from "nextra-theme-docs";
import { Footer } from "./components/Footer";
import { EditLink } from "./components/EditLink";
import CloudQueryLogo from "./components/logos/CloudQuery";
import { Badge } from "./components/Badge";
import { Callout, useConfig } from "nextra-theme-docs";
import { components } from "./utils/components";
import { getSlackAppLink } from "./utils/slack-app-link";
import { WebSite, WithContext } from "schema-dts";

const jsonLd: WithContext<WebSite> = {
  "@context": "https://schema.org",
  "@type": "WebSite",
  name: "CloudQuery",
  alternateName: ["CloudQuery Docs"],
  url: "https://docs.cloudquery.io/",
};

const theme: DocsThemeConfig = {
  project: {
    link: "https://github.com/cloudquery/cloudquery",
    icon: (
      <img
        alt="CloudQuery Github repo stars"
        src="https://img.shields.io/github/stars/cloudquery/cloudquery?style=social"
      />
    ),
  },
  sidebar: {
    defaultMenuCollapseLevel: 1,
  },
  toc: {
    float: true,
  },
  docsRepositoryBase:
    "https://github.com/cloudquery/cloudquery/blob/main/website/pages",
  useNextSeoProps: () => ({
    titleTemplate: "%s | CloudQuery",
  }),
  search: { component: null },
  chat: {
    link: "https://community.cloudquery.io",
  },
  feedback: {
    content: "Question? Give us feedback →",
  },
  logo: function LogoActual() {
    return (
      <>
        <CloudQueryLogo height={32} />
        <span className="sr-only">CloudQuery</span>
      </>
    );
  },
  logoLink: "https://www.cloudquery.io",
  head: () => {
    const { frontMatter } = useConfig();
    return (
      <>
        <script
          id="faq-schema"
          type="application/ld+json"
          dangerouslySetInnerHTML={{
            __html: JSON.stringify(jsonLd),
          }}
        />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <link
          rel="apple-touch-icon"
          sizes="180x180"
          href="/favicon/apple-touch-icon.png"
        />
        <link
          rel="icon"
          type="image/png"
          sizes="32x32"
          href="/favicon/favicon-32x32.png"
        />
        <link
          rel="icon"
          type="image/png"
          sizes="16x16"
          href="/favicon/favicon-16x16.png"
        />
        <link rel="shortcut icon" href="/favicon/favicon.ico" />
        <meta name="msapplication-TileColor" content="#111111" />
        <meta name="twitter:card" content="summary_large_image" />
        <meta name="twitter:site" content="@cloudqueryio" />
        <meta name="twitter:creator" content="@cloudqueryio" />
        <meta property="og:type" content="website" />
        <meta property="og:locale" content="en_US" />
        <meta property="og:site_name" content="CloudQuery" />
        {frontMatter.image ? (
          <meta property="og:image" content={frontMatter.image} />
        ) : null}
        {frontMatter.video ? (
          <meta property="og:video" content={frontMatter.video} />
        ) : null}
      </>
    );
  },
  footer: {
    text: <Footer />,
  },
  editLink: {
    component: (props) => <EditLink {...props} />,
  },
  nextThemes: {
    defaultTheme: "dark",
  },
  components: {
    badge: ({ text }: { text: string }) => <Badge text={text} />,
    configuration: ({ kind, name }: { kind: string; name: string }) => {
      return components[`${kind}-${name}-configuration`];
    },
    authentication: ({ kind, name }: { kind: string; name: string }) => {
      return components[`${kind}-${name}-authentication`];
    },
    callout: ({ type, children }: any) => {
      return <Callout type={type}>{children}</Callout>;
    },
    "slack-app-link": () => {
      return (
        <div style={{ marginTop: "1em" }}>
          <a target="_blank" className="btn btn-blue" href={getSlackAppLink()}>
            Install App
          </a>
        </div>
      );
    },
  },
};

export default theme;
