import { DocsThemeConfig } from "nextra-theme-docs";
import { Footer } from "./components/Footer";
import { EditLink } from "./components/EditLink";
import CloudQueryLogo from "./components/logos/CloudQuery";
import { useRouter } from 'next/router'
import { useConfig } from 'nextra-theme-docs'

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
    link: "https://www.cloudquery.io/discord",
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
  head: () => {
      const { frontMatter } = useConfig()
      return (
      <>
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
        <meta property="og:locale" content="en_IE" />
        <meta property="og:site_name" content="CloudQuery" />
        {(frontMatter.title) ?
            <meta property="og:title" content={frontMatter.title} />
            : null
        }
        {(frontMatter.video) ?
          <>
              <meta property="og:video:url" content={frontMatter.video} />
              <meta property="og:video:secure_url" content={frontMatter.video} />
              <meta property="og:video:type" content="application/x-shockwave-flash"/>
              <meta property="og:video:width" content="1280" />
              <meta property="og:video:height" content="720" />
          </>
          : null
        }
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
};

export default theme;
