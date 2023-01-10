import { NextRequest, NextResponse, userAgent } from "next/server";

const getOS = (request) => {
  const parsed = userAgent(request);

  if (parsed.os.name.includes("Mac OS")) {
    return "macOS";
  }
  if (parsed.os.name.includes("Windows")) {
    return "windows";
  }

  return "linux";
};

export function middleware(request: NextRequest) {
  if (request.url.includes("/docs/quickstart/")) {
    // Don't redirect if the URL has the OS already to avoid infinite redirects
    return;
  }
  try {
    const url = request.nextUrl.clone();
    const os = getOS(request);
    const nextUrl = new URL(`/docs/quickstart/${os}`, url);
    return NextResponse.redirect(nextUrl);
  } catch {
    // do nothing
  }
}

export const config = {
  matcher: "/docs/quickstart",
};
