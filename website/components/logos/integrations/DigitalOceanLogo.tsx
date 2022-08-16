import { INTEGRATION_WIDTH } from "./constants";

const DigitalOceanLogo = ({ width = INTEGRATION_WIDTH, className }: { width?: number, className?: string }) => (
    <svg
        viewBox="0 0 64 64"
        xmlns="http://www.w3.org/2000/svg"
        xmlSpace="preserve"
        style={{
            fillRule: "evenodd",
            clipRule: "evenodd",
            strokeLinejoin: "round",
            strokeMiterlimit: 2,
        }}
        className={className || "dark:text-white text-gray-900"}
        width={width}
    >
        <path
            d="M32.036 51.6c13.12-.021 23.296-13.011 18.256-26.836a18.553 18.553 0 0 0-11.066-11.062c-13.824-5.004-26.85 5.15-26.85 18.28H0C0 11.056 20.248-5.262 42.2 1.592c9.6 3.01 17.236 10.626 20.21 20.2 6.86 21.98-9.434 42.214-30.406 42.214V51.634h.032V51.6Zm-21.842 9.534v-9.5H19.7v9.5h-9.506ZM32.036 51.6h-.032v.034H19.7v-12.33h12.336V51.6Zm-21.842.034H2.248v-7.94h7.946v7.94Z"
            fill="currentColor"
            style={{
                fillRule: "nonzero",
            }}
        />
    </svg>
)

export default DigitalOceanLogo
