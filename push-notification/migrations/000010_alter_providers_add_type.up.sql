Begin;

ALTER TABLE `providers` ADD COLUMN `type` ENUM('SMS','PUSH') NOT NULL DEFAULT 'SMS' AFTER `name`;

Commit;