
-- +migrate Up
CREATE TABLE `collections`
(
  id                char(36) not null primary key comment 'ID',
  user_id           char(36) not null comment 'ユーザーID',
  chain_id          varchar(64) not null comment 'チェーンID',
  name              varchar(255) not null comment 'コレクション名',
  contract_address  varchar(64) not null comment 'コントラクトアドレス',
  description       varchar(4000) comment '説明文',
  image_url         varchar(255) comment '画像URL',
  banner_image_url  varchar(255) comment 'バナー画像URL',
  external_url      varchar(255) comment '外部URL',
  royalty           int not null comment 'ロイヤリティ',
  royalty_receiver  char(42) not null comment 'ロイヤリティアドレス',
  created_at        datetime not null comment '作成日時',
  updated_at        datetime not null comment '更新日時',
  foreign key user_foreign_key (user_id) references users (id)
) comment 'コレクション';

-- +migrate Down
DROP TABLE `collections`;