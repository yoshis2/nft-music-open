
-- +migrate Up
create TABLE `transactions`
(
  id                varchar(80) not null primary key comment 'ID',
  user_id           char(36) not null comment 'ユーザーID',
  chain_id          varchar(64) not null comment 'チェーンID',
  nonce             int not null comment 'Nonce',
  token_url         varchar(255) not null comment 'トークンURL',
  genre_id          char(36) not null comment 'ジャンルID',
  `to`              varchar(128) comment '送信先アドレス',
  price             decimal(40,10) not null comment '金額',
  insentive         int  comment 'インセンティブ',
  cost              bigint not null comment 'コスト',
  sale              boolean comment '販売可否',
  status            varchar(64) comment 'ステータス',
  created_at        datetime not null comment '作成日時',
  updated_at        datetime not null comment '更新日時',
  foreign key user_foreign_key (user_id) references users (id),
  foreign key genre_id_foregin_key (genre_id) references genre_masters (id)
) comment 'トランザクション処理';

-- +migrate Down
DROP TABLE `transactions`;