"use client";

import type { NextPage } from "next";
import { useRouter } from "next/navigation"; // Import useRouter

import { useState } from "react";
import { FormProvider, useForm } from "react-hook-form";

import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";

import ConfirmDialog from "@/components/modals/master";
const formSchema = z.object({
  name: z.string().min(1, {
    message: "名前は必須項目です。",
  }),
});

const GenreCreate: NextPage = () => {
  const router = useRouter(); // Initialize router
  const [isOpen, setIsOpen] = useState(false);
  const form = useForm<z.infer<typeof formSchema>>({
    // Add type argument to useForm
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "",
    },
  });
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = form;

  // フォームデータはこの関数内で使用しないため、引数を削除
  const onSubmit = async () => {
    setIsOpen(true);
  };

  const handleOk = async () => {
    const data = form.getValues();
    try {
      await fetch("/api/master/genres", {
        method: "POST",
        body: JSON.stringify(data),
      });
      setIsOpen(false);
      router.push("/master/genres/list"); // Use router.push for client-side navigation
    } catch (error) {
      console.error("ジャンル作成エラー:", error);
      alert("作成処理中にエラーが発生しました。");
    }
  };

  return (
    <main className="main-container">
      <h1 className="heading1">ジャンル作成</h1>
      <FormProvider {...form}>
        <form onSubmit={handleSubmit(onSubmit)} className="form-contents">
          <label htmlFor="text">ジャンル名</label>
          <input type="text" {...register("name")} />
          {errors.name?.message && <p className="error-text">{errors.name?.message}</p>}
          <button type="submit" className="submit-button">
            送信
          </button>
        </form>
      </FormProvider>
      <ConfirmDialog
        open={isOpen}
        kind="作成"
        name={form.getValues("name")}
        onCancel={() => setIsOpen(false)}
        onOk={handleOk}
      />
    </main>
  );
};

export default GenreCreate;
