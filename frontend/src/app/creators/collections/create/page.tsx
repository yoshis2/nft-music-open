"use client";

import { useState } from "react";
import type { NextPage } from "next";
import Image from "next/image";

const CreateCollectionPage: NextPage = () => {
  const [name, setName] = useState("");
  const [symbol, setSymbol] = useState("");
  const [description, setDescription] = useState("");
  const [coverImage, setCoverImage] = useState<File | null>(null);
  const [coverImagePreview, setCoverImagePreview] = useState<string | null>(null);

  const handleCoverImageChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      const file = e.target.files[0];
      setCoverImage(file);
      setCoverImagePreview(URL.createObjectURL(file));
    }
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    // TODO: APIへの送信処理を実装
    console.log({
      name,
      symbol,
      description,
      coverImage,
    });
    alert("コレクション作成データがコンソールに出力されました。");
  };

  return (
    <div className="container mx-auto max-w-3xl px-4 py-10">
      <h1 className="mb-8 text-3xl font-bold">Create New Collection</h1>
      <form onSubmit={handleSubmit} className="space-y-6">
        <div>
          <label htmlFor="name" className="block text-sm font-medium text-gray-700">
            Collection Name
          </label>
          <input
            type="text"
            id="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div>
          <label htmlFor="symbol" className="block text-sm font-medium text-gray-700">
            Symbol
          </label>
          <input
            type="text"
            id="symbol"
            value={symbol}
            onChange={(e) => setSymbol(e.target.value)}
            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div>
          <label htmlFor="description" className="block text-sm font-medium text-gray-700">
            Description
          </label>
          <textarea
            id="description"
            rows={4}
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
          />
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700">Cover Image</label>
          <div className="mt-1 flex items-center space-x-6">
            <div className="h-24 w-24 flex-shrink-0 overflow-hidden rounded-md bg-gray-100">
              {coverImagePreview ? (
                <Image
                  src={coverImagePreview}
                  alt="Cover preview"
                  width={96}
                  height={96}
                  className="h-full w-full object-cover"
                />
              ) : (
                <span className="flex h-full w-full items-center justify-center text-sm text-gray-500">Preview</span>
              )}
            </div>
            <input type="file" onChange={handleCoverImageChange} accept="image/*" />
          </div>
        </div>

        <div>
          <button
            type="submit"
            className="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
          >
            Create Collection
          </button>
        </div>
      </form>
    </div>
  );
};

export default CreateCollectionPage;
