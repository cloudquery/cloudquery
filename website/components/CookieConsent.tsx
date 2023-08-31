import React from "react";
import CookieConsent, {
  Cookies,
  getCookieConsentValue as getValue,
  resetCookieConsentValue,
} from "react-cookie-consent";

type Properties = {
  onAccept: () => void;
  onDecline: () => void;
};

const CQ_COOKIE_CONSENT_NAME = "cq-cookie-consent";

export const getCookieConsentValue = () => {
  return getValue(CQ_COOKIE_CONSENT_NAME) === "true";
};

export const optIn = () => {
  resetCookieConsentValue(CQ_COOKIE_CONSENT_NAME);
  window.location.reload();
};

export const optOut = () => {
  Cookies.set(CQ_COOKIE_CONSENT_NAME, "false");
  window.location.reload();
};

const buttonComponent = (props) => {
  const isAccept = props.children === "Accept";
  const className = isAccept
    ? "font-medium text-white no-underline bg-black border border-transparent rounded-md dark:bg-green-400 dark:text-black betterhover:dark:hover:bg-green-500 betterhover:hover:bg-gray-700 md:py-3 md:text-lg md:px-10 md:leading-6"
    : "font-medium text-white no-underline bg-black border border-transparent rounded-md dark:bg-white dark:text-black betterhover:dark:hover:bg-gray-300 betterhover:hover:bg-gray-700 md:py-3 md:text-lg md:px-10 md:leading-6";
  return (
    <button
      id={props.id}
      aria-label={props["aria-label"]}
      onClick={props.onClick}
      className={className}
      style={{ margin: "10px 5px" }}
    >
      {props.children}
    </button>
  );
};

export const CQCookieConsent: React.FC<Properties> = ({
  onAccept,
  onDecline,
}) => {
  return (
    <CookieConsent
      enableDeclineButton
      ButtonComponent={buttonComponent}
      buttonText="Accept"
      declineButtonText="Decline"
      onAccept={onAccept}
      onDecline={onDecline}
      cookieName={CQ_COOKIE_CONSENT_NAME}
    >
      We use tracking cookies to understand how you use the product and help us
      improve it. Please accept cookies to help us improve. You can always opt
      out later via the link in the footer.
    </CookieConsent>
  );
};
