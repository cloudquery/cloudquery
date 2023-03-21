import React from "react";
import PricingCard from "../PricingCard";
import {PricingTier} from "../PricingCard";

const tiers: PricingTier[] = [
    {
        name: "Community",
        description: "Open source community support",
        attributes: {
            Price: "Free",
            PriceMonthly: false,
            Communication: "Discord public channels, GitHub issues",
            Contract: "None needed",
        }
    },
    {
        name: "Business",
        description: "For businesses that need prioritized support",
        attributes: {
            Price: "$4k",
            PriceMonthly: true,
            PriceDetails: "Annual contract. 10% discount if billed annually",
            Communication: "Private channel in any of Discord, Slack, or Teams",
            SLA: "24 hour response time on support tickets. 48 hours for a P0 bug fix",
            TAM: "Up to 2 hours / month",
            Contract: "Standard Support Contract (minor modifications allowed)",
        }
    },
    {
        name: "Enterprise",
        description: "Prioritized support and custom development",
        attributes: {
            Price: "$25k",
            PriceMonthly: true,
            PriceDetails: "Annual contract. 10% discount if billed annually",
            Communication: "Private channel in any of Discord, Slack, or Teams",
            SLA: "24 hour response time on support tickets. 24 hours for a P0 bug fix",
            TAM: "Up to 4 hours / month",
            Features: "Up to 5 resources / month; more on a best-effort basis.",
            Contract: "Standard Support Contract (minor modifications allowed)",
        }
    },
    {
        name: "Enterprise Custom",
        description: "Customized contract to fit unique requirements",
        attributes: {
            Price: "Custom",
            PriceMonthly: false,
            Communication: "Private channel in any of Discord, Slack, or Teams",
            SLA: "24 hour response time on support tickets. 24 hours for a P0 bug fix",
            TAM: "Customizable hours / month",
            Features: "Custom feature, resource and plugin development",
            Contract: "Custom Support Contract",
        }
    },
];

export default function Home() {
  return <>
      <section>
          <div className="py-8 px-4 mx-auto max-w-screen-xl lg:py-16 lg:px-6">
              <div className="mx-auto max-w-screen-md text-center mb-8 lg:mb-12">
                  <h2 className="mb-4 text-4xl tracking-tight font-extrabold text-gray-900 dark:text-white">Support Plans</h2>
                  <p className="mb-5 font-light text-gray-500 sm:text-xl dark:text-gray-400">We offer flexible support plans that range from free open source community support to custom enterprise-level support.</p>
              </div>
              <div className="space-y-8 lg:grid lg:grid-cols-4 sm:gap-6 xl:gap-10 lg:space-y-0">
                  {tiers.map((tier) => (
                        <PricingCard key={tier.name} tier={tier} />
                    ))}
              </div>
          </div>
      </section>
  </>;
}
