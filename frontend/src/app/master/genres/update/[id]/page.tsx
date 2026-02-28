"use client";

import type { NextPage } from "next";
import { useParams, useRouter } from "next/navigation";

import { useState, useEffect } from "react";
import { FormProvider, useForm } from "react-hook-form";

import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";

import ConfirmDialog from "@/components/modals/master";
const formSchema = z.object({
  name: z.string().min(1, {
    message: "名前は必須項目です。",
  }),
});

const GenreUpdate: NextPage = () => {
  const pathname = useParams();
  const router = useRouter();
  const id = pathname?.id as string;
  const [isOpen, setIsOpen] = useState(false);
  const form = useForm({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "",
    },
  });
  const {
    register,
    handleSubmit,
    setValue,
    formState: { errors },
  } = form;

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(`/api/master/genres/${id}`);
        if (!response.ok) throw new Error("ジャンルのデータ取得に失敗しました。");
        const resJson = await response.json();
        setValue("name", resJson.name);
      } catch (error) {
        console.error("データ取得エラー:", error);
        // ここでユーザーにエラーを通知することもできます
      }
    };
    fetchData();
  }, [id, setValue]);

  const onSubmit = async () => {
    setIsOpen(true);
  };

  const handleOk = async () => {
    const data = form.getValues();
    try {
      await fetch(`/api/master/genres/${id}`, {
        method: "PUT",
        body: JSON.stringify(data),
      });
      setIsOpen(false);
      router.push("/master/genres/list");
    } catch (error) {
      console.error("ジャンル更新エラー:", error);
      alert("更新処理中にエラーが発生しました。");
    }
  };

  return (
    <main className="main-container">
      <h1 className="heading1">ジャンル更新</h1>
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
        kind="修正"
        name={form.getValues("name")}
        onCancel={() => setIsOpen(false)}
        onOk={handleOk}
      />
    </main>
  );
};

export default GenreUpdate;
