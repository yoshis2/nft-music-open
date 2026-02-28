import { NextRequest, NextResponse } from "next/server";

export async function POST(request: NextRequest) {
  try {
    const req = await request.json();
    const res = await fetch(`${process.env.BACKEND_HOST}/api/v1/ipfs/meta`, {
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
