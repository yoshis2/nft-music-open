
-- +migrate Up
ALTER TABLE users ADD CONSTRAINT unique_wallet UNIQUE (wallet);

-- +migrate Down
ALTER TABLE users DROP INDEX unique_wallet;