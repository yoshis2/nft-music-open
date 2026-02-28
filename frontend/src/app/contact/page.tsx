"use client";

import type { NextPage } from "next"; // Metadataは不要そうなので削除しても良いかもしれません
import React, { useState } from "react";
import { useForm, SubmitHandler, Resolver, FieldErrors } from "react-hook-form";
import MessageDialog from "@/components/layouts/contact-mail-modal";

// フォームデータの型を定義
interface FormValues {
  name: string;
  email: string;
  subject: string;
  contents: string;
}

// フォーム全体のバリデーションを行うリゾルバ関数
const formResolver: Resolver<FormValues> = async (data) => {
  const errors: FieldErrors<FormValues> = {};

  // Name validation
  if (!data.name) {
    errors.name = { type: "required", message: "氏名は必須です" };
  } else if (data.name.length < 2) {
    errors.name = {
      type: "minLength",
      message: "氏名は2文字以上で入力してください",
    };
  }

  // Email validation
  if (!data.email) {
    errors.email = { type: "required", message: "メールアドレスは必須です" };
  } else if (!/^\S+@\S+\.\S+$/.test(data.email)) {
    // より一般的なメール形式の正規表現
    errors.email = {
      type: "pattern",
      message: "有効なメールアドレスを入力してください",
    };
  }

  // Subject validation
  if (!data.subject) {
    errors.subject = { type: "required", message: "件名は必須です" };
  }

  // Contents validation
  if (!data.contents) {
    errors.contents = { type: "required", message: "お問合せ内容は必須です" };
  }

  return {
    values: Object.keys(errors).length === 0 ? data : {},
    errors,
  };
};

// メタデータ (もしこのページで個別に設定したい場合)
// export const metadata: Metadata = {
//   title: "お問い合わせ - NFT Music",
//   description: "NFT Musicへのお問い合わせはこちらから。",
// };

const Contact: NextPage = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [message, setMessage] = useState("");

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormValues>({
    defaultValues: {
      name: "",
      email: "",
      subject: "",
      contents: "", // contentsのデフォルト値も追加
    },
    resolver: formResolver, // 作成したリゾルバ関数を指定
    mode: "onSubmit", // resolverを使用する場合、onSubmitが一般的ですが、onChangeやonBlurも設定可能です
    // onChangeでリアルタイムバリデーションしたい場合は、resolverの挙動と合わせて検討してください
  });

  const onSubmit: SubmitHandler<FormValues> = async (data) => {
    setIsLoading(true);
    setMessage("");
    try {
      const apiResult = await sendMail(data); // sendMailは成功時にAPIからのJSONレスポンスを返す
      console.log("APIからのレスポンス:", apiResult); // このログでAPIの実際の応答構造を確認してください

      // sendMailがエラーをスローしなかった場合、HTTPリクエストは成功しています。
      // APIが返すJSONボディにメッセージが含まれていればそれを使用し、
      // そうでなければデフォルトの成功メッセージを表示します。
      if (apiResult && typeof apiResult.message === "string") {
        setMessage(apiResult.message);
      } else {
        // APIが特定のメッセージを返さない場合のデフォルト成功メッセージ
        setMessage("お問い合わせ内容が送信されました。");
      }
      setIsOpen(true); // モーダルを開く
      // フォームをリセットする場合は、ここでreset()を呼び出す
      // reset(); // useFormから取得したreset関数を使用
    } catch (error) {
      // sendMail内でスローされたエラー (HTTPエラーやネットワークエラーなど)
      if (error instanceof Error) {
        setMessage(error.message || "エラーが発生しました。");
      } else {
        setMessage("予期せぬエラーが発生しました。");
      }
    } finally {
      setIsLoading(false);
    }
  };

  const sendMail = async (formData: FormValues) => {
    try {
      const apiFormData = new FormData(); // HTMLのFormDataオブジェクトと区別するため別名に
      apiFormData.set("name", formData.name);
      apiFormData.set("email", formData.email);
      apiFormData.set("subject", formData.subject);
      apiFormData.set("contents", formData.contents);

      const response = await fetch("/api/mail/send", {
        method: "POST",
        body: apiFormData,
      });

      if (!response.ok) {
        // サーバーからのエラーレスポンスを処理
        const errorData = await response.json().catch(() => ({ message: `HTTPエラー: ${response.status}` }));
        throw new Error(errorData.message || `サーバーエラーが発生しました: ${response.status}`);
      }

      const resJson = await response.json();
      return resJson;
    } catch (err) {
      console.error(err);
      throw err; // エラーを呼び出し元に伝播させる
    }
  };

  return (
    <main className="main-container">
      <h1 className="heading1">お問合せ</h1>
      <p className="text-center text-lg text-gray-600 mb-12">ご質問やご相談など、お気軽にお問い合わせください。</p>
      <div className="max-w-4xl w-full mx-auto">
        <form onSubmit={handleSubmit(onSubmit)} className="bg-white p-8 w-full rounded-lg shadow-md space-y-6">
          <div>
            <label htmlFor="name" className="block text-sm font-medium text-gray-700">
              氏名 <span className="text-red-500">*</span>
            </label>
            <div className="mt-1">
              <input
                type="text"
                id="name"
                {...register("name")}
                className="block w-full px-3 py-2 bg-white border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              />
            </div>
            {errors.name && <p className="mt-2 text-sm text-red-600">{errors.name.message}</p>}
          </div>

          <div>
            <label htmlFor="email" className="block text-sm font-medium text-gray-700">
              メールアドレス <span className="text-red-500">*</span>
            </label>
            <div className="mt-1">
              <input
                type="email"
                id="email"
                {...register("email")}
                className="block w-full px-3 py-2 bg-white border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              />
            </div>
            {errors.email && <p className="mt-2 text-sm text-red-600">{errors.email.message}</p>}
          </div>

          <div>
            <label htmlFor="subject" className="block text-sm font-medium text-gray-700">
              件名 <span className="text-red-500">*</span>
            </label>
            <div className="mt-1">
              <input
                type="text"
                id="subject"
                {...register("subject")}
                className="block w-full px-3 py-2 bg-white border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              />
            </div>
            {errors.subject && <p className="mt-2 text-sm text-red-600">{errors.subject.message}</p>}
          </div>

          <div>
            <label htmlFor="contents" className="block text-sm font-medium text-gray-700">
              お問合せ内容 <span className="text-red-500">*</span>
            </label>
            <div className="mt-1">
              <textarea
                rows={5}
                id="contents"
                {...register("contents")}
                className="block w-full px-3 py-2 bg-white border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              />
            </div>
            {errors.contents && <p className="mt-2 text-sm text-red-600">{errors.contents.message}</p>}
          </div>

          <div>
            <button
              type="submit"
              disabled={isLoading}
              className="w-full flex justify-center py-3 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:bg-indigo-400 disabled:cursor-not-allowed transition-colors"
            >
              {isLoading ? "送信中..." : "送信"}
            </button>
          </div>
          {message && (
            <p
              className={`text-center text-sm ${message.includes("失敗") || message.includes("エラー") ? "text-red-500" : "text-green-500"}`}
            >
              {message}
            </p>
          )}
        </form>
      </div>
      <MessageDialog open={isOpen} onCancel={() => setIsOpen(false)} onOk={() => setIsOpen(false)} />
    </main>
  );
};
export default Contact;
