import nodemailer from 'nodemailer';

const {
  NEXT_PUBLIC_SMTP,
  NEXT_PUBLIC_OWNER_MAIL,
  NEXT_PUBLIC_OWNER_MAIL_PASS,
} = process.env;

export const Transporter = nodemailer.createTransport({
  host: NEXT_PUBLIC_SMTP,
  port: 465,
  secure: true, // Use `true` for port 465, `false` for all other ports
  auth: {
    user: NEXT_PUBLIC_OWNER_MAIL,
    pass: NEXT_PUBLIC_OWNER_MAIL_PASS,
  },
});
