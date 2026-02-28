import { type NextRequest, NextResponse } from "next/server";

export async function GET(
  request: NextRequest,
  { params }: { params: Promise<{ transaction_id: string }> },
): Promise<NextResponse> {
  const { transaction_id } = await params;
  try {
    const res = await fetch(`${process.env.BACKEND_HOST}/api/v1/nfts/detail/${transaction_id}`, {
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
