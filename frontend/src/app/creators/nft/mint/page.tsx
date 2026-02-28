"use client";

import type { NextPage } from "next";
import { useRouter } from "next/navigation";
import Image from "next/image";

import { useEffect, useRef, useState } from "react";
import { FormProvider, useForm } from "react-hook-form";

import { zodResolver } from "@hookform/resolvers/zod";
import { ethers } from "ethers";

import { IPFS_GATEWAY_PORT, IPFS_HOST, LOCAL_HOST } from "@/components/lib/constants";
import ErrorDialog from "@/components/modals/error";
import SuccessDialog from "@/components/modals/success";
import { GenreMasterItem } from "@/types/master";
import { defaultFormValues, formSchema, type NftMintFormValues } from "./form";

const NftMint: NextPage = () => {
  const router = useRouter();
  const imageInputFile = useRef<HTMLInputElement>(null);
  const audioInputFile = useRef<HTMLInputElement>(null);
  const videoInputFile = useRef<HTMLInputElement>(null);

  const [genreList, setGenreList] = useState<GenreMasterItem[]>([]);

  const [fileType, setFileType] = useState<"audio" | "video">("audio");
  const [imageCid, setImageCid] = useState("");
  const [audioCid, setAudioCid] = useState(""); // For audio or video
  const [videoCid, setVideoCid] = useState(""); // For audio or video

  const [wallet, setWallet] = useState("");
  const [ethersChainId, setEthersChainId] = useState<bigint | null>(null);

  const [uploading, setUploading] = useState({ image: false, audio: false, video: false });
  const [unregistered, setUnregistered] = useState(false);
  const [isSuccessOpen, setIsSuccessOpen] = useState(false);

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch("/api/master/genres/", {
        method: "GET",
      });
      const genre = await response.json();
      setGenreList(genre);
    };
    fetchData();
  }, []);

  const handleFileChange = async (e: React.ChangeEvent<HTMLInputElement>, type: "image" | "audio" | "video") => {
    if (!e.target.files) {
      return;
    }

    if (!window.ethereum) {
      return;
    }

    const provider = new ethers.BrowserProvider(window.ethereum);
    const network = await provider.getNetwork();
    const signer = await provider.getSigner();
    const walletAddress = await signer.getAddress();

    setEthersChainId(network.chainId);
    setWallet(walletAddress);

    try {
      setUploading((prev) => ({ ...prev, [type]: true }));

      const data = new FormData();
      data.set("file", e.target.files[0]);
      data.set("wallet", walletAddress);

      const response = await fetch("/api/ipfs/file", {
        method: "POST",
        body: data,
      });

      const imageAddResponse = await response.json();
      if (!imageAddResponse.user_id) {
        setUnregistered(true);
        return;
      }

      if (type === "image") {
        setImageCid(imageAddResponse.cid);
      } else if (type === "audio") {
        setAudioCid(imageAddResponse.cid);
      } else if (type === "video") {
        setVideoCid(imageAddResponse.cid);
      }
      setUploading((prev) => ({ ...prev, [type]: false }));
    } catch (err) {
      setUploading((prev) => ({ ...prev, [type]: false }));
      console.error("ファイルアップロードエラー:", err);
      alert("ファイルのアップロード中にエラーが発生しました");
      return;
    }
  };

  const form = useForm<NftMintFormValues>({
    resolver: zodResolver(formSchema),
    defaultValues: defaultFormValues,
  });

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = form;

  const createAccount = async () => {
    setUnregistered(false);
    setUploading({ image: false, audio: false, video: false });
    router.push("/creators/users/create");
  };

  const genres = genreList.map((genre) => (
    <option key={genre.id} value={genre.id}>
      {genre.name}
    </option>
  ));

  const onSubmit = async () => {
    if ((fileType === "audio" && (!imageCid || !audioCid)) || (fileType === "video" && !videoCid)) {
      alert("必要なファイルを全てアップロードしてください。");
      return;
    }
    // setActive(true); // This state is not defined, but was in the original code.
    setIsSuccessOpen(true);
  };

  const onHandler = async () => {
    const data = form.getValues(); // <- varをconstに、不要なJSON処理を削除
    const meta = JSON.stringify({
      // meta はこの時点でJSON文字列です
      chain_id: ethersChainId ? Number(ethersChainId) : null,
      wallet: wallet,
      genre_id: String(data.genre_id),
      name: data.name,
      description: data.description,
      file_type: fileType,
      image_cid: imageCid,
      audio_cid: audioCid,
      video_cid: videoCid,
      status: "mint",
      price: String(data.price),
      insentive: String(data.insentive),
      sale: data.sale,
    });

    try {
      const res = await fetch("/api/nft", {
        method: "POST",
        body: meta,
      });
      const resData = await res.json();
      console.log(resData);
    } catch (e) {
      alert(`NFTミント中にエラーが発生しました: ${e}`);
    }
    setIsSuccessOpen(false);
    router.push("/creators/nft/listing");
  };

  return (
    <main className="main-container max-w-4xl mx-auto">
      <h1 className="heading1">NFT登録</h1>
      <section className="p-6 mb-6 bg-white border border-gray-200 rounded-lg shadow w-full">
        <h2 className="text-xl font-semibold mb-2">1. ファイルタイプを選択</h2>
        <div className="flex items-center space-x-4">
          <div className="flex items-center">
            <input
              type="radio"
              id="audio"
              name="fileType"
              value="audio"
              checked={fileType === "audio"}
              onChange={() => setFileType("audio")}
              className="h-4 w-4 text-indigo-600 border-gray-300 focus:ring-indigo-500"
            />
            <label htmlFor="audio" className="ml-2 block text-sm font-medium text-gray-700">
              MP3と画像
            </label>
          </div>
          <div className="flex items-center">
            <input
              type="radio"
              id="video"
              name="fileType"
              value="video"
              checked={fileType === "video"}
              onChange={() => setFileType("video")}
              className="h-4 w-4 text-indigo-600 border-gray-300 focus:ring-indigo-500"
            />
            <label htmlFor="video" className="ml-2 block text-sm font-medium text-gray-700">
              MP4動画
            </label>
          </div>
        </div>
      </section>

      {fileType === "audio" && (
        <>
          <section className="p-6 mb-6 bg-white border border-gray-200 rounded-lg shadow w-full">
            <h2 className="text-xl font-semibold mb-2">2. 画像をアップロード</h2>
            <p className="text-sm text-gray-600 mb-4">NFTのジャケット画像（JPG, PNG, GIFなど）を選択してください。</p>
            <input
              type="file"
              className="hidden"
              ref={imageInputFile}
              onChange={(e) => handleFileChange(e, "image")}
              accept="image/*"
            />
            <button
              type="button"
              className="submit-button"
              disabled={uploading.image}
              onClick={() => imageInputFile.current?.click()}
            >
              {uploading.image ? "アップロード中..." : "画像を選択"}
            </button>
            {imageCid && (
              <div className="mt-6 flex justify-center">
                <Image
                  src={`${IPFS_HOST}${IPFS_GATEWAY_PORT}/ipfs/${imageCid}`}
                  alt="アップロードされた画像"
                  className="rounded-lg border border-gray-200"
                  width={400}
                  height={300}
                />
              </div>
            )}
          </section>
          <section className="p-6 mb-6 bg-white border border-gray-200 rounded-lg shadow w-full">
            <h2 className="text-xl font-semibold mb-2">3. 音声ファイルをアップロード</h2>
            <p className="text-sm text-gray-600 mb-4">NFTにする音声ファイル（MP3）を選択してください。</p>
            <input
              type="file"
              className="hidden"
              ref={audioInputFile}
              onChange={(e) => handleFileChange(e, "audio")}
              accept="audio/mpeg"
            />
            <button
              type="button"
              className="submit-button"
              disabled={uploading.audio}
              onClick={() => audioInputFile.current?.click()}
            >
              {uploading.audio ? "アップロード中..." : "音声ファイルを選択"}
            </button>
            {audioCid && <p className="mt-4 text-green-600">音声ファイルのアップロードが完了しました。</p>}
          </section>
        </>
      )}

      {fileType === "video" && (
        <section className="p-6 mb-6 bg-white border border-gray-200 rounded-lg shadow w-full">
          <h2 className="text-xl font-semibold mb-2">2. 動画ファイルをアップロード</h2>
          <p className="text-sm text-gray-600 mb-4">NFTにする動画ファイル（MP4）を選択してください。</p>
          <input
            type="file"
            className="hidden"
            ref={videoInputFile}
            onChange={(e) => handleFileChange(e, "video")}
            accept="video/mp4"
          />
          <button
            type="button"
            className="submit-button"
            disabled={uploading.video}
            onClick={() => videoInputFile.current?.click()}
          >
            {uploading.video ? "アップロード中..." : "動画ファイルを選択"}
          </button>
          {videoCid && (
            <div className="mt-6 flex justify-center">
              <video
                src={`${LOCAL_HOST}${IPFS_GATEWAY_PORT}/ipfs/${videoCid}`}
                controls
                className="rounded-lg border border-gray-200"
                width="400"
              >
                お使いのブラウザはビデオタグをサポートしていません。
              </video>
            </div>
          )}
        </section>
      )}

      <section className="p-6 bg-white border border-gray-200 rounded-lg shadow">
        <h2 className="text-xl font-semibold mb-4">{fileType === "audio" ? "4" : "3"}. NFT情報を入力</h2>
        <FormProvider {...form}>
          <form onSubmit={handleSubmit(onSubmit)}>
            <label htmlFor="genre_id">ジャンル</label>
            <select className="input-text" id="genre_id" {...register("genre_id")}>
              <option value="">選択してください</option>
              {genres}
            </select>
            {errors.genre_id?.message && <p className="error-text">{errors.genre_id?.message}</p>}
            <hr />
            <label htmlFor="name">NFT名</label>
            <input type="text" id="name" {...register("name")} className="input-text" />
            {errors.name && <p className="error-text">{errors.name.message}</p>} <hr />
            <label htmlFor="description">説明文</label>
            <textarea id="description" className="input-text" {...register("description")} />
            {errors.description && <p className="error-text">{errors.description.message}</p>} <hr />
            <label htmlFor="price">価格</label>
            <div>
              <div className="inline-flex justify-between bg-gray-100 rounded w-full">
                <input
                  type="number"
                  id="price"
                  step="any"
                  placeholder="半角数字をご入力ください"
                  className="input-text"
                  {...register("price")}
                />
                <div className="inline bg-gray-200 py-2 px-4 text-gray-600 select-none whitespace-nowrap">ETH</div>
              </div>
              イーサリアムやポリゴン、Flareの価格が0.001から購入が可能です。
            </div>
            {errors.price && <p className="error-text">{errors.price.message}</p>} <hr />
            <label htmlFor="insentive">製作者のインセンティブ割合</label>
            <div>
              <div className="inline-flex justify-between bg-gray-100 rounded w-full">
                <input type="number" id="insentive" step="any" {...register("insentive")} className="input-text" />
                <div className="inline bg-gray-200 py-2 px-4 text-gray-600 select-none whitespace-nowrap">
                  パーセント
                </div>
              </div>
            </div>
            {errors.insentive && <p className="error-text">{errors.insentive.message}</p>} <hr />
            <label htmlFor="sale">販売する</label>
            <input type="checkbox" id="sale" {...register("sale")} className="input-text" />
            {errors.sale && <p className="error-text">{errors.sale.message}</p>} <hr />
            <div className="text-center">
              <button
                type="submit"
                className="submit-button"
                disabled={uploading.image || uploading.audio || uploading.video}
              >
                送信
              </button>
            </div>
          </form>
        </FormProvider>
      </section>
      <ErrorDialog
        open={unregistered}
        kind="NFTデータアップロード"
        wallet={wallet}
        onCancel={() => setUnregistered(false)}
        createAccount={createAccount}
      />
      <SuccessDialog
        open={isSuccessOpen}
        kind="NFT ミント(作成)"
        wallet={wallet}
        onCancel={() => setIsSuccessOpen(false)}
        onHandler={onHandler}
      />
    </main>
  );
};
export default NftMint;
