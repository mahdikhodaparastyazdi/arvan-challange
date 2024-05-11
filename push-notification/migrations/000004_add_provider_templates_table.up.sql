CREATE TABLE IF NOT EXISTS `provider_templates` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `provider_id` bigint unsigned NOT NULL,
    `template_id` bigint unsigned NOT NULL,
    `code` varchar(20)  NOT NULL,
    PRIMARY KEY(`id`),
    UNIQUE (`provider_id`, `template_id`, `code`),
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

ALTER TABLE `provider_templates` ADD FOREIGN KEY (`provider_id`) REFERENCES `providers`(`id`);
ALTER TABLE `provider_templates` ADD FOREIGN KEY (`template_id`) REFERENCES `templates`(`id`);