

-- +migrate Up
ALTER TABLE `transactions`
  ADD COLUMN `contract_address` varchar(64) COMMENT 'コントラクトアドレス' AFTER `chain_id`,
  ADD COLUMN `creator_address` char(42) COMMENT 'クリエイターアドレス' AFTER `status`;

-- +migrate Down
ALTER TABLE `transactions`
  DROP COLUMN `contract_address`,
  DROP COLUMN `creator_address`;