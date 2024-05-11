ALTER TABLE `sms` ADD FOREIGN KEY (`template_id`) REFERENCES `templates`(`id`);
ALTER TABLE `sms` DROP `type`;