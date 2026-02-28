
-- +migrate Up
create TABLE `genre_masters`
(
  id          char(36) not null primary key comment 'ID',
  name        nvarchar(64) not null comment '名称',
  created_at  datetime not null comment '作成日時',
  updated_at  datetime not null comment '更新日時'
);

-- +migrate Down
DROP TABLE `genre_masters`;
