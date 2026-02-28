"use client";

import type { NextPage } from "next";
import { redirect, useParams } from "next/navigation";

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

const BusinessUpdate: NextPage = () => {
  const pathname = useParams();
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
      const response = await fetch("/api/master/businesses/" + id, {
        method: "GET",
      });
      const resJson = await response.json();
      setValue("name", resJson.name);
    };
    fetchData();
  }, [id, setValue]);

  const onSubmit = async () => {
    setIsOpen(true);
  };

  const handleOk = async () => {
    const data = form.getValues();
    await fetch("/api/master/businesses/" + id, {
      method: "PUT",
      body: JSON.stringify(data),
    });
    setIsOpen(false);
    redirect("/master/businesses/list");
  };

  return (
    <main className="main-container">
      <h1 className="heading1">職種更新</h1>
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
        kind="修正"
        name={form.getValues("name")}
        onCancel={() => setIsOpen(false)}
        onOk={handleOk}
      />
    </main>
  );
};

export default BusinessUpdate;
