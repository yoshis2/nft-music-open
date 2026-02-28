"use client";

import type { NextPage } from "next";
import { useRouter } from "next/navigation";
import React, { Suspense, useEffect, useState } from "react";
import { FormProvider, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { ethers } from "ethers";

import { BusinessMasterItem } from "@/types/master";
import { formSchema, type UserCreateFormValues } from "../create/form";

const UserUpdateContent = () => {
  const [businessList, setBusinessList] = useState<BusinessMasterItem[]>([]);
  const [address, setAddress] = useState("");
  const [isPageLoading, setIsPageLoading] = useState(true);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [formError, setFormError] = useState<string | null>(null);
  const router = useRouter();

  const form = useForm<UserCreateFormValues>({
    resolver: zodResolver(formSchema),
  });

  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
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

        const userResponse = await fetch(`/api/users?wallet=${walletAddress}`);

        if (userResponse.status === 404) {
          router.push("/creators/users/create?error=ユーザー情報が見つかりません。まずユーザー登録をしてください。");
          return;
        }

        if (!userResponse.ok) {
          const errorData = await userResponse.json();
          throw new Error(errorData.message || "ユーザー情報の取得に失敗しました。");
        }

        const userData = await userResponse.json();
        reset(userData);

        const businessResponse = await fetch("/api/master/businesses/");
        const businessData = await businessResponse.json();
        setBusinessList(businessData);

        setAddress(walletAddress);
        setFormError(null);
      } catch (err: unknown) {
        if (
          typeof err === "object" &&
          err !== null &&
          "code" in err &&
          (err as { code: unknown }).code === "ACTION_REJECTED"
        ) {
          setFormError(
            "ウォレットへの接続が拒否されました。編集を続けるには、ページを再読み込みして接続を許可してください。",
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
  }, [router, reset]);

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
        method: "PUT",
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
      console.error("更新処理中にエラーが発生しました:", error);
      setFormError("更新処理中にエラーが発生しました。");
    } finally {
      setIsSubmitting(false);
    }
  };

  if (isPageLoading) {
    return (
      <main className="main-container max-w-4xl mx-auto">
        <h1 className="heading1">ユーザー編集</h1>
        <p>読み込み中...</p>
      </main>
    );
  }

  return (
    <main className="main-container max-w-4xl mx-auto">
      <h1 className="heading1">ユーザー編集</h1>
      <section className="p-6 mb-6 bg-white border border-gray-200 rounded-lg shadow w-full">
        <p className="mb-4 text-gray-600 dark:text-gray-400">ユーザー情報を編集してください。</p>
        {formError && (
          <div
            className="p-4 mb-4 text-sm text-red-800 rounded-lg bg-red-50 dark:bg-gray-800 dark:text-red-400"
            role="alert"
          >
            <span className="font-medium"></span> {formError}
          </div>
        )}
        <FormProvider {...form}>
          <form onSubmit={handleSubmit(onSubmit)}>
            <input type="hidden" {...register("id")} />
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
              {isSubmitting ? "更新中..." : "更新"}
            </button>
          </form>
        </FormProvider>
      </section>
    </main>
  );
};

const UserUpdatePage: NextPage = () => (
  <Suspense fallback={<div>Loading...</div>}>
    <UserUpdateContent />
  </Suspense>
);

export default UserUpdatePage;
