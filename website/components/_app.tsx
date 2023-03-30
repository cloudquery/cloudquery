import React from "react";
import Script from "next/script";
import { Prism } from "prism-react-renderer";
import {DefaultSeo} from "next-seo";
import {useRouter} from "next/router";

(typeof global !== "undefined" ? global : window).Prism = Prism;

require("prismjs/components/prism-powershell");
require("prismjs/components/prism-ini");
require("prismjs/components/prism-docker");
require("prismjs/components/prism-toml");

// Shim requestIdleCallback in Safari
if (typeof window !== "undefined" && !("requestIdleCallback" in window)) {
  // @ts-expect-error
  window.requestIdleCallback = (fn) => setTimeout(fn, 1);
  // @ts-expect-error
  window.cancelIdleCallback = (e) => clearTimeout(e);
}

declare global {
  interface Window {
    sa_event: any;
  }
}

const getCanonicalUrl = (path: string) => {
  if (process.env.NEXT_PUBLIC_VERCEL_ENV !== "production") {
    return
  }

  return (`https://www.cloudquery.io` + (path === "/" ? "": path)).split("?")[0]
}

export default function Nextra({ Component, pageProps }) {
    const router = useRouter();
    const canonicalUrl = getCanonicalUrl(router.asPath);

    return (
    <React.Fragment>
      <DefaultSeo
          canonical={canonicalUrl}
      />
      <Component {...pageProps} />
      <Script>
        {typeof window !== "undefined" &&
          (window.sa_event =
            window.sa_event ||
            function () {
              var a = [].slice.call(arguments);
              window.sa_event.q
                ? window.sa_event.q.push(a)
                : (window.sa_event.q = [a]);
            })}
      </Script>
      <Script defer data-domain="cloudquery.io" src="https://plausible.io/js/script.js"></Script>
      <Script src="https://scripts.simpleanalyticscdn.com/latest.js" />
      <noscript>
        {/* eslint-disable @next/next/no-img-element */}
        <img
          src="https://queue.simpleanalyticscdn.com/noscript.gif"
          alt=""
          referrerPolicy="no-referrer-when-downgrade"
        />
      </noscript>
    </React.Fragment>
  );
}
