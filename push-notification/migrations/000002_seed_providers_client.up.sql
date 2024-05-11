INSERT INTO `providers` 
  (`id`, `name`, `active`, `otp_template`, `created_at`, `updated_at`)
VALUES
  (1, 'provider1', true, 'otp', now(), now()),
  (2, 'provider2', false, 'verify', now(), now())
ON DUPLICATE KEY UPDATE
  `updated_at` = now();

INSERT INTO `clients` (`name`, `balance`, `sms_rate`,`api_key`, `is_active`, `created_at`, `updated_at`)
VALUES
    ('client1', 60, 'low', 'XgKD7VSvd7UAMmNf2LcsXz13mcE0dFS7UYlE87ld39TcmsQCabMo377jhL4YCPzk', 1, now(), now()),
    ('client2', 80, 'high', 'n59e0xosxQDnJUto4FJDbXFThgsqh8CXprE9E4qATsniqhlqs98fPRtxK70Ax9bd', 1, now(), now()),
    ('client3', 90, 'high', 'FqoBfPO4Ngrx7mCWAicKzt1SsyH3hkeHY2ghKkVI0VPZzaMX85SpXV66SbCilFgH', 1, now(), now()),
    ('activity', 600, 'low', 'hyDqdTZp5DGeMOLYQL7C8CeUDQ4XmAoiPnaXwf0bbBfKA4pRm7TvupeicM5cjhGy', 1, now(), now())
ON DUPLICATE KEY UPDATE `updated_at` = now();

