import { INTEGRATION_WIDTH } from "./constants";

const CloudflareLogo = ({ width = INTEGRATION_WIDTH, className }: { width?: number, className?: string }) => (
    <svg
        viewBox="0 0 412 187"
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
            d="M331 326c11-26-4-38-19-38l-148-2c-4 0-4-6 1-7l150-2c17-1 37-15 43-33 0 0 10-21 9-24-9.297-44.859-49.169-77.315-94.981-77.315-41.641 0-78.846 26.813-92.019 66.315-38-25-78 9-69 46-48 3-65 46-60 72 0 1 1 2 3 2h274c1 0 3-1 3-3Zm50-102c-4 0-6-1-7 1l-5 21c-5 16 3 30 20 31l32 2c4 0 4 6-1 7l-33 1c-36 4-46 39-46 39 0 2 0 3 2 3h113l3-2a81.015 81.015 0 0 0 3.045-22c0-44.435-36.565-81-81-81H381Z"
            fill="currentColor"
            style={{
                fillRule: "nonzero",
            }}
            transform="translate(-50.16 -142.685)"
        />
    </svg>
)

export default CloudflareLogo
