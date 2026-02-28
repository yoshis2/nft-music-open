import { NextRequest, NextResponse } from "next/server";

export async function GET() {
  try {
    const res = await fetch(`${process.env.BACKEND_HOST}/api/v1/nfts`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });
    const data = await res.json();
    return new NextResponse(JSON.stringify(data), { status: 200 });
  } catch (e) {
    console.log(e);
    const errorMessage = e instanceof Error ? e.message : String(e);
    return new NextResponse(JSON.stringify({ error: "エラー内容 : " + errorMessage }), {
      status: 500,
    });
  }
}

export async function POST(request: NextRequest) {
  try {
    const req = await request.json();
    console.log(JSON.stringify(req));
    const res = await fetch(`${process.env.BACKEND_HOST}/api/v1/nfts`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(req),
    });
    const data = await res.json();
    return new NextResponse(JSON.stringify(data), { status: 200 });
  } catch (e) {
    console.log(e);
    const errorMessage = e instanceof Error ? e.message : String(e);
    return new NextResponse(JSON.stringify({ error: "エラー内容 : " + errorMessage }), {
      status: 500,
    });
  }
}
