import { Transporter } from '@/config/nodemailer';
import { NextRequest, NextResponse } from 'next/server';

export async function POST(request: NextRequest) {
  const { NEXT_PUBLIC_OWNER_MAIL } = process.env;
  try {
    const data = await request.formData();
    const info = await Transporter.sendMail({
      from: '"NFTミュージックお問合せ" <' + NEXT_PUBLIC_OWNER_MAIL + '>',
      to: NEXT_PUBLIC_OWNER_MAIL + ', ' + data.get('email'),
      subject:
        'お問合せいただきありがとうございました - (' +
        String(data.get('subject')) +
        ')', // Subject line
      text: contents(data), // plain text body
    });

    return new NextResponse(JSON.stringify({ info }), { status: 200 });
  } catch (e) {
    console.log(e);
    return new NextResponse(
      JSON.stringify({ error: 'Internal Server Error' }),
      { status: 500 },
    );
  }
}

const contents = (data: FormData) => {
  return (
    String(data.get('name')) +
    '様\n\n' +
    'お問合せいただきありがとうございます \n' +
    'スリーネクストNFTミュージックです。\n\n' +
    'お問合せ内容について３営業日以内に回答しますので今しばらくお待ちください\n\n' +
    '件名 : ' +
    String(data.get('subject')) +
    '\n\n' +
    'お問合せ内容 \n' +
    String(data.get('contents'))
  );
};
