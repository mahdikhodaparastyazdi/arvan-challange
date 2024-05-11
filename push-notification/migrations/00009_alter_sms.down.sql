ALTER TABLE `sms` DROP FOREIGN KEY (`template_id`);
ALTER TABLE `sms` ADD COLUMN `type` enum('otp','sms') DEFAULT NULL;