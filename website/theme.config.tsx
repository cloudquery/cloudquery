import { DocsThemeConfig } from "nextra-theme-docs";
import { Footer } from "./components/Footer";
import CloudQueryLogo from "./components/logos/CloudQuery";

const theme: DocsThemeConfig = {
  github: "https://github.com/cloudquery/cloudquery",
  projectLink: "https://github.com/cloudquery/cloudquery",
  docsRepositoryBase:
    "https://github.com/cloudquery/cloudquery/blob/main/docs/pages",
  titleSuffix: " | cloudquery",
  search: true,
  floatTOC: true,
  font: false,
  projectChatLink: "https://cloudquery.io/discord",
  feedbackLink: "Question? Give us feedback →",
  logo: function LogoActual() {
    return (
      <>
        <CloudQueryLogo height={32} />
        <span className="sr-only">CloudQuery</span>
      </>
    );
  },
  head: (
    <>
      <meta name="viewport" content="width=device-width, initial-scale=1.0" />
      <link
        rel="apple-touch-icon"
        sizes="180x180"
        href="/images/favicon/apple-touch-icon.png"
      />
      <link
        rel="icon"
        type="image/png"
        sizes="32x32"
        href="/images/favicon/favicon-32x32.png"
      />
      <link
        rel="icon"
        type="image/png"
        sizes="16x16"
        href="/images/favicon/favicon-16x16.png"
      />
      <link
        rel="mask-icon"
        href="/images/favicon/safari-pinned-tab.svg"
        color="#000000"
      />
      <link rel="shortcut icon" href="/images/favicon/favicon.ico" />
      <meta name="msapplication-TileColor" content="#000000" />
      <meta name="theme-color" content="#000" />
      <meta name="twitter:card" content="summary_large_image" />
      <meta name="twitter:site" content="@cloudqueryio" />
      <meta name="twitter:creator" content="@cloudqueryio" />
      <meta property="og:type" content="website" />
      {/* <meta name="og:title" content={title} />
        <meta name="og:description" content={meta.description} /> */}
      {/* <meta
          property="og:url"
          content={`https://cloudquery.io${router.asPath}`}
        /> */}
      <meta
        property="twitter:image"
        content="https://cloudquery.io/og-image.png"
      />
      <meta property="og:image" content="https://cloudquery.io/og-image.png" />
      <meta property="og:locale" content="en_IE" />
      <meta property="og:site_name" content="CloudQuery" />
    </>
  ),
  footerEditLink: "Edit this page on GitHub",
  footerText: () => {
    return <Footer />;
  },
  nextThemes: {
    defaultTheme: "dark",
  },
};

export default theme;
