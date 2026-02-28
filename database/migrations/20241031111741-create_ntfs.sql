
-- +migrate Up
create TABLE `nfts`
(
  id                char(36) not null primary key comment 'ID',
  user_id           char(36) not null comment 'ユーザーID',
  contract_address  varchar(64) not null comment 'コントラクトアドレス',
  chain_id          varchar(64) comment 'チェーンID',
  transaction_id    varchar(64) not null comment 'トランザクションID',
  token_url         varchar(255) not null comment 'トークンURL',
  genre_id          char(36) comment 'ジャンルID',
  status            enum('mint','buy','gift') comment 'ステータス',
  sale              boolean comment '販売可否',
  price             int comment '金額',
  insentive         int  comment 'インセンティブ',
  created_at        datetime not null comment '作成日時',
  updated_at        datetime not null comment '更新日時',
  foreign key user_foreign_key (user_id) references users (id),
  foreign key genre_id_foregin_key (genre_id) references genre_masters (id)
);


-- +migrate Down
DROP TABLE `nfts`;
