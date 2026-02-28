"use client";

import type { NextPage } from "next";
import { redirect } from "next/navigation";

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

const BusinessCreate: NextPage = () => {
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
    formState: { errors },
  } = form;

  const onSubmit = async () => {
    setIsOpen(true);
  };

  const handleOk = async () => {
    const data = form.getValues();
    await fetch("/api/master/businesses", {
      method: "POST",
      body: JSON.stringify(data),
    });
    setIsOpen(false);
    redirect("/master/businesses/list");
  };

  return (
    <main className="main-container">
      <h1 className="heading1">職種作成</h1>
      <FormProvider {...form}>
        <form onSubmit={handleSubmit(onSubmit)} className="form-contents">
          <label htmlFor="text" className="label-text">
            職種
          </label>
          <input type="text" {...register("name")} className="input-text" />
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

export default BusinessCreate;
