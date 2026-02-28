import { NextRequest, NextResponse } from "next/server";

export async function GET(request: NextRequest) {
  try {
    const url = new URL(request.url);
    const params = url.searchParams.get("params"); // 変数名をより明確に
    const res = await fetch(`${process.env.BACKEND_HOST}/api/v1/nfts/search?${params}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });
    const data = await res.json();
    return new NextResponse(JSON.stringify(data), { status: 200 });
  } catch (e) {
    console.error("エラー内容 : ", e);
    const errorMessage = e instanceof Error ? e.message : String(e);
    return new NextResponse(JSON.stringify({ error: "エラー内容 : " + errorMessage }), {
      status: 500,
    });
  }
}
