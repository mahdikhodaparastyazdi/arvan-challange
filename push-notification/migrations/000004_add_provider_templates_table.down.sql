ALTER TABLE `provider_template` DROP FOREIGN KEY `provider_id`;
ALTER TABLE `provider_template` DROP FOREIGN KEY `template_id`;
DROP TABLE `provider_templates`;