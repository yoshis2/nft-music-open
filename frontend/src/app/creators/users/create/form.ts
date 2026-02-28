import { z } from "zod";

export const formSchema = z.object({
  id: z.string(),
  name: z.string().min(1, {
    message: "名前は必須項目です。",
  }),
  email: z.string().email({ message: "メールアドレスの形式でご記入ください" }).min(1, {
    message: "メールアドレスは必須項目です",
  }),
  wallet: z.string(), // ウォレットアドレスはuseEffectで設定するため、ここでは必須としない
  address: z.string().optional(), // 任意入力の場合
  business_id: z.string().min(1, {
    // 必須入力の場合
    message: "職種を選択してください。",
  }),
  website: z.string().url({ message: "有効なURLを入力してください。" }).optional().or(z.literal("")), // URL形式、任意入力、または空文字列を許容
  profile: z.string().optional(), // 任意入力の場合
});

// フォームの型をエクスポートしておくと便利です
export type UserCreateFormValues = z.infer<typeof formSchema>;

// フォームのデフォルト値をエクスポート
export const defaultFormValues: UserCreateFormValues = {
  id: "",
  name: "",
  email: "",
  wallet: "", // walletはuseEffectで設定されるため、初期値は空で良い
  address: "",
  business_id: "",
  website: "",
  profile: "",
};
