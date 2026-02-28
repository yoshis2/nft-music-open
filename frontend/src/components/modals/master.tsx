import type { NextPage } from 'next';

export type ModalProps = {
  open: boolean;
  kind: string;
  name: string;
  onCancel: () => void;
  onOk: () => void;
};

const Modal: NextPage<ModalProps> = (props) => {
  return props.open ? (
    <>
      <div className="bg-white  top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-120 h-96 p-5 flex flex-col items-start absolute z-20">
        <h1 className="text-xl font-bold mb-5">{props.kind} 確認画面</h1>
        <p className="text-lg mb-5 break-all">
          {props.kind}後の内容は「{props.name}」です。
          <br />
          <br />
          OKボタンをクリックすると{props.kind}が完了します。
          <br />
          <br />
          もし、取消したい場合はキャンセルボタンをクリックしてください。
        </p>
        <div className="flex mt-auto w-full">
          <button
            className="bg-slate-900 hover:bg-slate-700 text-white px-8 py-2 mx-auto"
            onClick={() => props.onOk()}
          >
            OK
          </button>
          <button
            className="bg-slate-900 hover:bg-slate-700 text-white px-8 py-2 mx-auto"
            onClick={() => props.onCancel()}
          >
            キャンセル
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
