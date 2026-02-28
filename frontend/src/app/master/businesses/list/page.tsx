"use client";

import type { NextPage } from "next";
import Link from "next/link";

import { useState, useEffect } from "react";
import { BusinessMasterItem } from "@/types/master";

import ConfirmDialog from "@/components/modals/master";

const BusinessList: NextPage = () => {
  const [itemList, setItemList] = useState<BusinessMasterItem[]>([]);
  const [isOpen, setIsOpen] = useState(false);
  const [selectedId, setSelectedId] = useState<string>("");

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch("/api/master/businesses", {
        method: "GET",
      });
      const resJson = await response.json();
      setItemList(resJson);
    };
    fetchData();
  }, []);

  const deleteBusiness = async (id: string) => {
    setSelectedId(id);
    setIsOpen(true);
  };

  const handleOk = async () => {
    await fetch("/api/master/businesses/" + selectedId, {
      method: "DELETE",
    });
    setIsOpen(false);
    window.location.reload();
  };

  const list = itemList.map((item, i) => (
    <tr key={item.id} className="p-24 my-8">
      <td className="p-4 border border-slate-300 text-center">{i + 1}</td>
      <td className="p-4 border border-slate-300">{item.name}</td>
      <td className="p-2 border border-slate-300 text-center">
        <Link
          href={`/master/businesses/update/${item.id}`}
          className="edit-button"
        >
          編集
        </Link>
        <input
          type="button"
          name=""
          className="delete-button"
          onClick={() => deleteBusiness(item.id)}
          value="削除"
        />
      </td>
    </tr>
  ));

  return (
    <div className="z-12">
      <main className="main-container">
        <h1 className="heading1">職種一覧</h1>
        <div className="w-full xl:w-2/3 h-16 text-right">
          <Link href={`/master/businesses/create`} className="submit-button">
            作成
          </Link>
        </div>
        <div className="relative flex flex-col w-full xl:w-2/3 h-full overflow-scroll text-gray-700 bg-white shadow-md rounded-xl bg-clip-border">
          <table className="w-full text-left table-auto min-w-max">
            <thead>
              <tr>
                <th className="p-4 border border-slate-300 w-[10%]">ID</th>
                <th className="p-4 border border-slate-300 w-[60%]">職種名</th>
                <th className="p-4 border border-slate-300 w-[30%]">ボタン</th>
              </tr>
            </thead>
            <tbody>{list}</tbody>
          </table>
        </div>
        <ConfirmDialog
          open={isOpen}
          kind="削除"
          name={""}
          onCancel={() => setIsOpen(false)}
          onOk={handleOk}
        />
      </main>
    </div>
  );
};
export default BusinessList;
