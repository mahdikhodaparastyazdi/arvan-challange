ALTER TABLE `templates` DROP FOREIGN KEY `active_provider_id`;
ALTER TABLE  `templates` DROP `active_provider_id`;
ALTER TABLE `templates` DROP `use_provider_template`;
ALTER TABLE `templates` DROP `priority`;