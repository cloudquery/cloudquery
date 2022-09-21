import { INTEGRATION_WIDTH } from "./constants";

const HerokuLogo = ({
  width = INTEGRATION_WIDTH,
  className,
}: {
  width?: number;
  className?: string;
}) => (
  <svg
    xmlSpace="preserve"
    width={width}
    className={className || "dark:text-white text-gray-900"}
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 5.12 5.12"
  >
    <path d="M4.12 5.12H.968a.49.49 0 0 1-.488-.488V.488A.49.49 0 0 1 .968 0H4.12a.49.49 0 0 1 .488.488v4.144a.49.49 0 0 1-.488.488z" />
    <path
      d="M3.068 4.415V2.382s.132-.487-1.63.2C1.436 2.6 1.436.7 1.436.7L2.01.697v1.2s1.61-.635 1.61.48v2.026h-.555zm.328-2.986h-.6c.22-.27.42-.73.42-.73h.63s-.108.3-.44.73zm-1.95 2.982V3.254l.58.58-.58.58z"
      fill="currentColor"
    />
  </svg>
);

export default HerokuLogo;
