import { NextRequest, NextResponse, userAgent } from "next/server";

const getOS = (request) => {
  const parsed = userAgent(request);

  if (parsed.os.name.includes("Mac OS")) {
    return "macOS";
  }
  if (parsed.os.name.includes("Windows")) {
    return "Windows";
  }

  return "Linux";
};

export function middleware(request: NextRequest) {
  if (request.url.includes("/docs/quickstart?")) {
    return;
  }
  try {
    const url = request.nextUrl.clone();
    const os = getOS(request);
    url.searchParams.set("os", os);
    return NextResponse.rewrite(url);
  } catch {
    // do nothing
  }
}

export const config = {
  matcher: "/docs/quickstart",
};
