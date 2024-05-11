Begin;

ALTER TABLE `sms` ADD COLUMN `expires_at` TIMESTAMP DEFAULT now() AFTER `template_id`;

UPDATE `sms`
    JOIN `templates` ON sms.template_id = templates.id
    SET sms.expires_at = DATE_ADD(sms.created_at,INTERVAL 2 MINUTE)
WHERE templates.type = 'OTP';

ALTER TABLE `sms` ALTER COLUMN `expires_at` DROP DEFAULT;

UPDATE `sms`
    JOIN `templates` ON sms.template_id = templates.id
    SET sms.expires_at = DATE_ADD(sms.created_at, INTERVAL 1 DAY)
WHERE templates.type = 'SMS';

ALTER TABLE `sms` MODIFY COLUMN `status` ENUM('PENDING', 'SENT',  'RETRY', 'FAILED', 'EXPIRED');

Commit;