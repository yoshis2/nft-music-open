
-- +migrate Up
create TABLE `wallets`
(
  id          char(36) not null primary key comment 'ID',
  address     char(42) not null comment 'ウォレットアドレス',
  created_at  datetime not null comment '作成日時'
);

-- +migrate Down
DROP TABLE `wallets`;
