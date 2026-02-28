"use client";

import type { NextPage } from "next";
import { useSearchParams, useRouter } from "next/navigation";
import React, { Suspense, useEffect, useState } from "react";
import { FormProvider, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { ethers } from "ethers";

import { BusinessMasterItem } from "@/types/master";
import { defaultFormValues, formSchema, type UserCreateFormValues } from "./form";

const UserCreateContent = () => {
  const [businessList, setBisneessList] = useState<BusinessMasterItem[]>([]);
  const [address, setAddress] = useState("");
  const [isPageLoading, setIsPageLoading] = useState(true);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [formError, setFormError] = useState<string | null>(null);
  const router = useRouter();

  // URLのクエリパラメータを取得
  const searchParams = useSearchParams();
  const queryError = searchParams.get("error");

  // フォームの初期化
  const form = useForm({
    resolver: zodResolver(formSchema),
    defaultValues: defaultFormValues, // インポートしたデフォルト値を使用
  });

  const {
    register,
    handleSubmit,
    formState: { errors },
    setValue,
  } = form;

  useEffect(() => {
    const fetchData = async () => {
      if (!window.ethereum) {
        setFormError("ウォレットがインストールされていません。");
        setIsPageLoading(false);
        return;
      }
      try {
        const provider = new ethers.BrowserProvider(window.ethereum);
        await provider.send("eth_requestAccounts", []);
        const signer = await provider.getSigner();
        const walletAddress = await signer.getAddress();

        // ユーザーが既に登録済みか確認
        const userResponse = await fetch(`/api/users?wallet=${walletAddress}`, {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
        });
        if (userResponse.status === 404) {
          // 404 (Not Found) は正常ケースとして扱い、アカウント作成フローを継続
        } else if (userResponse.ok) {
          const userData = await userResponse.json();
          // ユーザーが存在する場合（IDで判断）、クリエイターページにリダイレクト
          if (userData && userData.id) {
            router.push("/creators");
            return; // これ以降の処理は不要
          }
        } else {
          const errorData = await userResponse.json();
          throw new Error(errorData.message || "ユーザー情報の確認に失敗しました。");
        }

        // 未登録の場合、フォームの準備を続ける
        const businessResponse = await fetch("/api/master/businesses/", {
          method: "GET",
        });
        const businessData = await businessResponse.json();
        setBisneessList(businessData);

        setAddress(walletAddress);
        setValue("wallet", walletAddress, { shouldValidate: true });
        setFormError(null);
      } catch (err: unknown) {
        if (
          typeof err === "object" &&
          err !== null &&
          "code" in err &&
          (err as { code: unknown }).code === "ACTION_REJECTED"
        ) {
          setFormError(
            "ウォレットへの接続が拒否されました。登録を続けるには、ページを再読み込みして接続を許可してください。",
          );
        } else if (err instanceof Error) {
          console.error("データ取得中にエラーが発生しました:", err);
          setFormError(`エラー: ${err.message}`);
        } else {
          console.error("不明なエラー:", err);
          setFormError("不明なエラーが発生しました。");
        }
      } finally {
        setIsPageLoading(false);
      }
    };
    fetchData();
  }, [router, setValue]);

  const businesses = businessList.map((business) => (
    <option key={business.id} value={business.id}>
      {business.name}
    </option>
  ));

  const onSubmit = async (data: UserCreateFormValues) => {
    setIsSubmitting(true);
    setFormError(null);
    try {
      const response = await fetch("/api/users", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });

      if (!response.ok) {
        const resJson = await response.json();
        setFormError(resJson.message ?? "不明なエラーが発生しました。");
        return;
      }

      router.push("/creators");
    } catch (error) {
      console.error("登録処理中にエラーが発生しました:", error);
      setFormError("登録処理中にエラーが発生しました。");
    } finally {
      setIsSubmitting(false);
    }
  };

  if (isPageLoading) {
    return (
      <main className="main-container max-w-4xl mx-auto">
        <h1 className="heading1">ユーザー登録</h1>
        <p>読み込み中...</p>
      </main>
    );
  }

  return (
    <main className="main-container max-w-4xl mx-auto">
      <h1 className="heading1">ユーザー登録</h1>
      <section className="p-6 mb-6 bg-white border border-gray-200 rounded-lg shadow w-full">
        <p className="mb-4 text-gray-600 dark:text-gray-400">
          クリエイターとして活動するために、ユーザー情報を登録してください。
          <br />
          ウォレットアドレスは自動で取得されます。
        </p>
        {(queryError || formError) && (
          <div
            className="p-4 mb-4 text-sm text-red-800 rounded-lg bg-red-50 dark:bg-gray-800 dark:text-red-400"
            role="alert"
          >
            <span className="font-medium"></span> {queryError || formError}
          </div>
        )}
        <FormProvider {...form}>
          <form onSubmit={handleSubmit(onSubmit)}>
            <label htmlFor="name">名前</label>
            <input type="text" id="name" {...register("name")} />
            {errors.name?.message && <p className="error-text">{errors.name?.message}</p>}
            <hr />
            <label htmlFor="email">メールアドレス</label>
            <input type="text" id="email" {...register("email")} />
            {errors.email?.message && <p className="error-text">{errors.email?.message}</p>}
            <hr />
            <label htmlFor="wallet">ウォレットアドレス</label>
            <input type="text" id="wallet" {...register("wallet")} value={address} disabled />
            {errors.wallet?.message && <p className="error-text">{errors.wallet?.message}</p>}
            <hr />
            <label htmlFor="address">住所</label>
            <input type="text" id="address" {...register("address")} />
            {errors.address?.message && <p className="error-text">{errors.address?.message}</p>} <hr />
            <label htmlFor="business_id">職種</label>
            <select className="input-text" id="business_id" {...register("business_id")}>
              <option>選択してください</option>
              {businesses}
            </select>
            {errors.business_id?.message && <p className="error-text">{errors.business_id?.message}</p>} <hr />
            <label htmlFor="website">ウェブサイト</label>
            <input type="text" id="website" {...register("website")} />
            {errors.website?.message && <p className="error-text">{errors.website?.message}</p>} <hr />
            <label htmlFor="profile">プロフィール文</label>
            <textarea {...register("profile")} className="input-text"></textarea>
            {errors.profile?.message && <p className="error-text">{errors.profile?.message}</p>} <hr />
            <button type="submit" className="submit-button" disabled={isSubmitting}>
              {isSubmitting ? (
                <>
                  <svg
                    className="animate-spin -ml-1 mr-3 h-5 w-5 text-white inline"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                  >
                    <circle
                      className="opacity-25"
                      cx="12"
                      cy="12"
                      r="10"
                      stroke="currentColor"
                      strokeWidth="4"
                    ></circle>
                    <path
                      className="opacity-75"
                      fill="currentColor"
                      d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                    ></path>
                  </svg>
                  送信中...
                </>
              ) : (
                "送信"
              )}
            </button>
          </form>
        </FormProvider>
      </section>
    </main>
  );
};

// Suspenseでラップしてエクスポート
const NftCreatePage: NextPage = () => (
  <Suspense fallback={<div>Loading...</div>}>
    <UserCreateContent />
  </Suspense>
);

export default NftCreatePage;
