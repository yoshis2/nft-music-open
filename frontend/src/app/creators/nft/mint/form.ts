import { z } from "zod";

export const formSchema = z.object({
  genre_id: z.string().min(1, { message: "ジャンルを選択してください。" }),
  name: z.string().min(1, {
    message: "名前は必須項目です。",
  }),
  price: z.coerce.number().positive({
    message: "0より大きい数字をご入力ください",
  }),
  sale: z.coerce.boolean(),
  insentive: z.coerce
    .number()
    .min(0, { message: "0以上の数値を指定してください。" })
    .max(100, {
      message: "数値は100以下を指定して下さい",
    })
    .int(),
  description: z
    .string()
    .min(1, {
      message: "説明文は必須項目です",
    })
    .max(1200, {
      message: "最大1200文字までです",
    }),
});

export type NftMintFormValues = z.input<typeof formSchema>;

export const defaultFormValues: NftMintFormValues = {
  genre_id: "",
  name: "",
  description: "",
  price: 0,
  sale: false,
  insentive: 0,
};
