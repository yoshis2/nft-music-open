import { describe, it, expect, vi, beforeEach } from "vitest";
import { POST } from "./route";
import { NextRequest } from "next/server";
import { Transporter } from "@/config/nodemailer";

vi.mock("@/config/nodemailer", () => ({
  Transporter: {
    sendMail: vi.fn(),
  },
}));

describe("api/mail/send", () => {
  beforeEach(() => {
    vi.resetAllMocks();
  });

  it("POST: メールを送信できること", async () => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    vi.mocked(Transporter.sendMail).mockResolvedValueOnce({ messageId: "123" } as any);

    // multipart/form-data形式のbodyを手動で作成
    const boundary = "----WebKitFormBoundary7MA4YWxkTrZu0gW";
    const body =
      `--${boundary}\r\n` +
      `Content-Disposition: form-data; name="name"\r\n\r\nTest User\r\n` +
      `--${boundary}\r\n` +
      `Content-Disposition: form-data; name="email"\r\n\r\ntest@example.com\r\n` +
      `--${boundary}\r\n` +
      `Content-Disposition: form-data; name="subject"\r\n\r\nTest Subject\r\n` +
      `--${boundary}\r\n` +
      `Content-Disposition: form-data; name="contents"\r\n\r\nTest Contents\r\n` +
      `--${boundary}--\r\n`;

    const request = new NextRequest("http://localhost/api/mail/send", {
      method: "POST",
      body: Buffer.from(body),
      headers: {
        "Content-Type": `multipart/form-data; boundary=${boundary}`,
      },
    });

    const response = await POST(request);
    const data = await response.json();

    expect(response.status).toBe(200);
    expect(data.info.messageId).toBe("123");
    expect(Transporter.sendMail).toHaveBeenCalled();
  });
});
