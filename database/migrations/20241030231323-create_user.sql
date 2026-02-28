-- +migrate Up
create TABLE `users`
(
  id          char(36) not null primary key comment 'ID',
  wallet      char(42) not null comment 'ウォレットアドレス',
  name        nvarchar(64) not null comment '氏名',
  email       varchar(128) not null comment 'メールアドレス',
  address     varchar(255) comment '住所',
  business_id char(36) comment '職種ID',
  website     varchar(128) comment 'サイトURL',
  face_image  varchar(255) comment 'キャラ画像URL',
  eyecatch    varchar(255) comment 'アイキャッチ画像URL',
  profile     nvarchar(2048) comment 'プロフィール文',
  role        enum('member', 'creator', 'admin') comment '権限',
  created_at  datetime not null comment '作成日時',
  updated_at  datetime not null comment '更新日時',
	foreign key business_id_foregin_key (business_id) references business_masters (id)
);

-- +migrate Down
DROP TABLE `users`;
