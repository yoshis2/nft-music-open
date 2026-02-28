import type { NextPage } from 'next';

export type ModalProps = {
  open: boolean;
  onCancel: () => void;
  onOk: () => void;
};

const Modal: NextPage<ModalProps> = (props) => {
  const end = () => {
    location.reload();
  };
  return props.open ? (
    <>
      <div className="bg-white  top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-120 h-96 p-5 flex flex-col items-start absolute z-20">
        <h1 className="text-xl font-bold mb-5">お問合せ完了のお知らせ</h1>
        <p className="text-lg mb-5 break-all">
          お問合せ完了しました。３営業日以内に返信しますので今しばらくお待ちください
        </p>
        <div className="flex mt-auto w-full">
          <button
            className="bg-slate-900 hover:bg-slate-700 text-white px-8 py-2 mx-auto"
            onClick={end}
          >
            OK
          </button>
        </div>
      </div>
      <div
        className="fixed bg-black bg-opacity-50 w-full h-full z-10"
        onClick={() => props.onCancel()}
      ></div>
    </>
  ) : (
    <></>
  );
};

export default Modal;
