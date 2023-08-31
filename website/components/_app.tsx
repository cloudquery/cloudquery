import React, { useState, useEffect } from "react";
import Script from "next/script";
import { Prism } from "prism-react-renderer";
import { DefaultSeo } from "next-seo";
import { useRouter } from "next/router";
import { CQCookieConsent, getCookieConsentValue } from "./CookieConsent";

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
    sa_event?: any;
  }
}

const getCanonicalUrl = (path: string) => {
  if (process.env.NEXT_PUBLIC_VERCEL_ENV !== "production") {
    return;
  }

  return (`https://www.cloudquery.io` + (path === "/" ? "" : path)).split(
    "?",
  )[0];
};

const Analytics = () => (
  <>
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
    <Script src="https://scripts.simpleanalyticscdn.com/latest.js" />
    <Script
      src={`https://www.googletagmanager.com/gtag/js?id=AW-11077035012`}
    />
    <Script id="google-tag-manager">
      {`window.dataLayer = window.dataLayer || [];
          function gtag(){dataLayer.push(arguments);}
          gtag('js', new Date());
            
          gtag('config', 'AW-11077035012');`}
    </Script>
    <Script>{`window.faitracker=window.faitracker||function(){this.q=[];var t=new CustomEvent("FAITRACKER_QUEUED_EVENT");return this.init=function(t,e,a){this.TOKEN=t,this.INIT_PARAMS=e,this.INIT_CALLBACK=a,window.dispatchEvent(new CustomEvent("FAITRACKER_INIT_EVENT"))},this.call=function(){var e={k:"",a:[]};if(arguments&&arguments.length>=1){for(var a=1;a<arguments.length;a++)e.a.push(arguments[a]);e.k=arguments[0]}this.q.push(e),window.dispatchEvent(t)},this.message=function(){window.addEventListener("message",function(t){"faitracker"===t.data.origin&&this.call("message",t.data.type,t.data.message)})},this.message(),this.init("gm3eyhbta0tee4d4mmzghwgpl4wa0vsf",{host:"https://api.dyh8ken8pc.com"}),this}(),function(){var t=document.createElement("script");t.type="text/javascript",t.src="https://asset.dyh8ken8pc.com/dyh8ken8pc.js",t.async=!0,(d=document.getElementsByTagName("script")[0]).parentNode.insertBefore(t,d)}();`}</Script>
  </>
);

export default function Nextra({ Component, pageProps }) {
  const router = useRouter();
  const canonicalUrl = getCanonicalUrl(router.asPath);
  const [consent, setConsent] = useState(getCookieConsentValue());

  return (
    <React.Fragment>
      <DefaultSeo canonical={canonicalUrl} />
      <Component {...pageProps} />
      {consent && <Analytics />}
      <noscript>
        {/* eslint-disable @next/next/no-img-element */}
        <img
          src="https://queue.simpleanalyticscdn.com/noscript.gif"
          alt=""
          referrerPolicy="no-referrer-when-downgrade"
        />
      </noscript>
      <CQCookieConsent
        onAccept={() => setConsent(true)}
        onDecline={() => setConsent(false)}
      />
    </React.Fragment>
  );
}
