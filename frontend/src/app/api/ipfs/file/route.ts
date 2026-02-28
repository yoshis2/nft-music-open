import { NextRequest, NextResponse } from "next/server";

export async function POST(request: NextRequest) {
  try {
    const data = await request.formData();
    const res = await fetch(`${process.env.BACKEND_HOST}/api/v1/ipfs`, {
      method: "POST",
      body: data,
    });

    const response = await res.json();
    return new NextResponse(JSON.stringify(response), { status: 200 });
  } catch (e) {
    console.log(e);
    const errorMessage = e instanceof Error ? e.message : String(e);
    return new NextResponse(JSON.stringify({ error: "Internal Server Error: " + errorMessage }), { status: 500 });
  }
}
