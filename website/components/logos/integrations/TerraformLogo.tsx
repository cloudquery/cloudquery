import { INTEGRATION_WIDTH } from "./constants";

const TerraformLogo = ({ width = INTEGRATION_WIDTH, className }: { width?: number, className?: string }) => (
    <svg
        viewBox="0 0 32 32"
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
            d="m11.041 5.688 9.912 5.031v10.073l-9.912-5.037V5.688Zm11 5.031v10.073l9.917-5.037V5.688l-9.917 5.031ZM.047.068v10.068l9.912 5.036V5.104L.047.068Zm10.994 26.853 9.912 5.037V21.895l-9.912-5.036v10.062Z"
            fill="currentColor"
            style={{
                fillRule: "nonzero",
            }}
        />
    </svg>
)

export default TerraformLogo
