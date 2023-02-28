import React from "react";
import Script from "next/script";
import { Prism } from "prism-react-renderer";

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

export default function Nextra({ Component, pageProps }) {
  return (
    <React.Fragment>
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
