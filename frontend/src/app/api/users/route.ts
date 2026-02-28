import { NextRequest, NextResponse } from "next/server";

// wallet addressからユーザー情報を取得
export async function GET(request: NextRequest): Promise<NextResponse> {
  try {
    const url = new URL(request.url);
    const walletAddress = url.searchParams.get("wallet"); // 変数名をより明確に
    const res = await fetch(`${process.env.BACKEND_HOST}/api/v1/users/wallet/${walletAddress}`, {
      // 新しいエンドポイントを呼び出す
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

export async function POST(request: NextRequest): Promise<NextResponse> {
  try {
    const req = await request.json();
    const res = await fetch(`${process.env.BACKEND_HOST}/api/v1/users`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(req),
    });
    const data = await res.json();
    console.log(data);
    if (typeof data === "string" && data.includes("error_type")) {
      return new NextResponse(JSON.stringify({ error: "エラー内容 : " + data }), {
        status: 500,
      });
    }
    return new NextResponse(JSON.stringify(data), { status: 200 });
  } catch (e) {
    console.error("エラー内容" + e);
    const errorMessage = e instanceof Error ? e.message : String(e);
    return new NextResponse(JSON.stringify({ error: "エラー内容 : " + errorMessage }), {
      status: 500,
    });
  }
}

export async function PUT(request: NextRequest): Promise<NextResponse> {
  try {
    const req = await request.json();
    console.log("req : ");
    console.log(req);

    const res = await fetch(`${process.env.BACKEND_HOST}/api/v1/users/${req.id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(req),
    });
    const data = await res.json();
    console.log(data);

    return new NextResponse(JSON.stringify(data), { status: 200 });
  } catch (e) {
    console.error("エラー内容" + e);
    const errorMessage = e instanceof Error ? e.message : String(e);
    return new NextResponse(JSON.stringify({ error: "エラー内容 : " + errorMessage }), {
      status: 500,
    });
  }
}
