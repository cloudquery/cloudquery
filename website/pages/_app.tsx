import "../styles.css";
import "nextra-theme-docs/style.css";
import "../custom.css";
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
  window.requestIdleCallback = (fn) => setTimeout(fn, 1);
  window.cancelIdleCallback = (e) => clearTimeout(e);
}

export default function Nextra({ Component, pageProps }) {
  return (
    <React.Fragment>
      <Component {...pageProps} />
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
