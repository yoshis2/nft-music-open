import type { NextPage } from "next";

export type ModalProps = {
  open: boolean;
  kind: string;
  wallet: string;
  onCancel: () => void;
  onHandler: () => void;
};

const Modal: NextPage<ModalProps> = (props) => {
  return props.open ? (
    <>
      <div className="modal">
        <div className="out" onClick={() => props.onCancel()}>
          <div className="flame">
            <div className="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
              <h3>{props.kind} 確認画面</h3>
              <hr />
              <div className="mt-2">
                <p>以下のアドレスに{props.kind}します。</p>
                <p>{props.wallet}</p>
                <p>「ミント」ボタンクリックで{props.kind}を実行します。</p>
                <p>取消す場合は「キャンセル」ボタンをクリックしてください。</p>
              </div>
            </div>
            <div className="background">
              <button className="ok" onClick={() => props.onHandler()}>
                ミント
              </button>
              <button className="cancel" onClick={() => props.onCancel()}>
                キャンセル
              </button>
            </div>
          </div>
        </div>
      </div>
      <div className="fixed inset-0 bg-gray-500/75 transition-opacity"></div>
    </>
  ) : (
    <></>
  );
};

export default Modal;
