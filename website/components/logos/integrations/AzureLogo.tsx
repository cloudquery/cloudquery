import {INTEGRATION_HEIGHT} from "./constants";

const AzureLogo = ({height = INTEGRATION_HEIGHT}) => (
    <svg
        viewBox="0 0 162 129"
        xmlns="http://www.w3.org/2000/svg"
        xmlSpace="preserve"
        style={{
            fillRule: "evenodd",
            clipRule: "evenodd",
            strokeLinejoin: "round",
            strokeMiterlimit: 2,
        }}
        className="dark:text-white text-gray-900"
        height={height}
    >
        <path
            d="M88.33 16.33 40.67 57.66 0 130.66h36.67L88.33 16.33ZM94.67 26 74.33 83.33l39 49-75.66 13h124L94.67 26Z"
            fill="currentColor"
            style={{
                fillRule: "nonzero",
            }}
            transform="translate(0 -16.33)"
        />
    </svg>
)

export default AzureLogo
