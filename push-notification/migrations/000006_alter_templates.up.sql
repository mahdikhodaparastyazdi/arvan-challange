ALTER TABLE `templates` ADD `active_provider_id` bigint unsigned NOT NULL AFTER `params`;

ALTER TABLE `templates` ADD FOREIGN KEY (`active_provider_id`) REFERENCES `providers`(`id`);

ALTER TABLE `templates` ADD `use_provider_template` tinyint(1) NULL AFTER `active_provider_id`;

ALTER TABLE `templates` ADD `priority` enum('HIGH', 'MEDIUM', 'LOW') DEFAULT NULL AFTER `use_provider_template`;

ALTER TABLE `templates` ADD `type` enum('OTP', 'SMS') DEFAULT NULL AFTER `priority`;