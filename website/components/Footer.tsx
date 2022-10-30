import { useRouter } from "next/router";
import Link from "next/link";
import { useState, ReactNode, FormEvent } from "react";
import CloudQueryLogo from "./logos/CloudQuery";

function FooterLink({ href, children }: { href: string; children: ReactNode }) {
  const classes =
    "text-sm text-gray-500 no-underline betterhover:hover:text-gray-700 betterhover:hover:dark:text-white transition";
  if (href.startsWith("http")) {
    return (
      <a href={href} className={classes}>
        {children}
      </a>
    );
  }
  return (
    <Link href={href}>
      <a className={classes}>{children}</a>
    </Link>
  );
}

function FooterHeader({ children }: { children: ReactNode }) {
  return <h3 className="text-sm text-gray-900 dark:text-white">{children}</h3>;
}

const navigation = {
  general: [
    { name: "Documentation", href: "/docs" },
    { name: "CLI reference", href: "/docs/reference/cli/cloudquery" },
    { name: "Blog", href: "/blog" },
    {
      name: "Releases",
      href: "https://github.com/cloudquery/cloudquery/releases",
    },
    { name: "FAQ", href: "/docs/faq" },
  ],
  support: [
    {
      name: "GitHub",
      href: "https://github.com/cloudquery/cloudquery",
    },
    {
      name: "Discord",
      href: "https://www.cloudquery.io/discord",
    },
  ],
  company: [
    { name: "GitHub", href: "https://github.com/cloudquery/cloudquery" },
    { name: "Twitter", href: "https://twitter.com/cloudqueryio" },
    { name: "Brand Logo", href: "/logo" },
  ],
  legal: [
    { name: "Privacy Policy", href: "/privacy" },
    { name: "Terms of Use", href: "/terms" },
  ],
};

export function Footer() {
  return (
    <footer className="" aria-labelledby="footer-heading">
      <h2 id="footer-heading" className="sr-only">
        Footer
      </h2>
      <div className="py-8 mx-auto max-w-7xl">
        <div className="xl:grid xl:grid-cols-3 xl:gap-8">
          <div className="grid grid-cols-2 gap-8 xl:col-span-2">
            <div className="md:grid md:grid-cols-2 md:gap-8">
              <div>
                <FooterHeader>Solutions</FooterHeader>
                <ul role="list" className="mt-4 space-y-1.5 list-none ml-0">
                  {navigation.general.map((item) => (
                    <li key={item.name}>
                      <FooterLink href={item.href}>{item.name}</FooterLink>
                    </li>
                  ))}
                </ul>
              </div>
              <div className="mt-12 md:!mt-0">
                <FooterHeader>Support</FooterHeader>
                <ul role="list" className="mt-4 space-y-1.5 list-none ml-0">
                  {navigation.support.map((item) => (
                    <li key={item.name}>
                      <FooterLink href={item.href}>{item.name}</FooterLink>
                    </li>
                  ))}
                </ul>
              </div>
            </div>
            <div className="md:grid md:grid-cols-2 md:gap-8">
              <div>
                <FooterHeader>Company</FooterHeader>
                <ul role="list" className="mt-4 space-y-1.5 list-none ml-0">
                  {navigation.company.map((item) => (
                    <li key={item.name}>
                      <FooterLink href={item.href}>{item.name}</FooterLink>
                    </li>
                  ))}
                </ul>
              </div>
              <div className="mt-12 md:!mt-0">
                <FooterHeader>Legal</FooterHeader>
                <ul role="list" className="mt-4 space-y-1.5 list-none ml-0">
                  {navigation.legal.map((item) => (
                    <li key={item.name}>
                      <FooterLink href={item.href}>{item.name}</FooterLink>
                    </li>
                  ))}
                </ul>
              </div>
            </div>
          </div>
          <div className="mt-12 xl:!mt-0">
            <FooterHeader>Subscribe to our newsletter</FooterHeader>
            <p className="mt-4 text-sm text-gray-600 dark:text-gray-500">
              Join the CloudQuery newsletter and stay updated on new releases
              and features, guides, and case studies.
            </p>
            <SubmitForm />
          </div>
        </div>

        <div className="pt-8 mt-8 md:flex md:items-center md:justify-between">
          <div>
            <a
              className="text-current"
              target="_blank"
              rel="noopener noreferrer"
              href="https://www.cloudquery.io"
            >
              <CloudQueryLogo />
            </a>
            <p className="mt-4 text-xs text-gray-500 ">
              &copy; {new Date().getFullYear()} CloudQuery, Inc. All rights
              reserved.
            </p>
          </div>
        </div>
      </div>
    </footer>
  );
}

function SubmitForm() {
  const [email, setEmail] = useState("");
  const router = useRouter();

  const subscribe = async (e: FormEvent) => {
    e.preventDefault();

    const params = new URLSearchParams();
    params.append("EMAIL", email);
    params.append("u", "5a433922a8db867a38c8a4a8e");
    params.append("id", "4a9aca3509");
    params.append("c", "?");

    await fetch(
      `https://cloudquery.us1.list-manage.com/subscribe/post-json?${params.toString()}`,
      { mode: "no-cors" }
    );

    router.push("/confirm");
  };

  return (
    <form
      className="mt-4 sm:flex sm:max-w-md"
      autoComplete="off"
      onSubmit={subscribe}
    >
      <label htmlFor="email-address" className="sr-only">
        Email address
      </label>
      <div>
        <input
          type="email"
          name="email-address"
          id="email-address"
          autoComplete="off"
          required
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          className="w-full min-w-0 px-4 py-2 text-base text-gray-900 placeholder-gray-500 bg-white border border-transparent rounded-md appearance-none dark:text-white sm:text-sm dark:border-gray-700 dark:bg-transparent focus:outline-none focus:ring-2 focus:ring-gray-800 dark:focus:border-white focus:placeholder-gray-400"
          placeholder="you@domain.com"
        />
      </div>
      <div className="mt-3 rounded-md sm:mt-0 sm:ml-3 sm:flex-shrink-0">
        <button
          type="submit"
          className="flex items-center justify-center w-full px-4 py-2 text-base font-medium text-white !bg-black dark:!bg-white dark:text-black border border-transparent rounded-md sm:text-sm betterhover:hover:!bg-gray-600 focus:outline-none focus:ring-2 focus:ring-gray-800 dark:focus:ring-white dark:betterhover:hover:!bg-gray-300"
        >
          Subscribe
        </button>
      </div>
    </form>
  );
}
