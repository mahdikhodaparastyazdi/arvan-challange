CREATE TABLE IF NOT EXISTS `providers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `active` tinyint(1) DEFAULT NULL,
  `otp_template` text,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `sms` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `mobile` varchar(20) DEFAULT NULL,
  `type` enum('otp','sms') DEFAULT NULL,
  `content` text,
  `status` enum('pending','sent','success','fail') DEFAULT NULL,
  `provider_id` bigint unsigned DEFAULT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `templates` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(100) DEFAULT NULL,
  `content` text,
  `params` text,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_templates_code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `clients` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `balance` BIGINT UNSIGNED NOT NULL,
    `sms_rate` enum('low','high') DEFAULT NULL,
    `name` VARCHAR(255) NOT NULL UNIQUE,
    `api_key` VARCHAR(255) NOT NULL UNIQUE,
    `is_active` tinyint(1) DEFAULT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

ALTER TABLE `sms` ADD FOREIGN KEY (`provider_id`) REFERENCES `providers`(`id`);
