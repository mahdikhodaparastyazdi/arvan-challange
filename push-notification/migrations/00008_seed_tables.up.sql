INSERT INTO `templates`
(`code`, `content`, `params`, `active_provider_id`, `use_provider_template`, `priority`, `type`, `created_at`, `updated_at`)
VALUES
    ('auth.otp', "آروان کلاد\nرمز یکبار‌ مصرف ورود🤫\ncode: {code}", 'code', 1, 0, 'HIGH', 'OTP', now(), now()),
    ('arvan.delete_account.otp', '⚠️با حذف حساب آروان کلاد همه اطلاعات شما پاک می‌شود!\ncode:{code}', 'code',1, 0, 'HIGH', 'OTP', now(), now()),
    ('reward.kyc_done', 'کد تخفیف یک میلیون تومانی محصولات ابری فعال شد🤩', NULL,1, 0, 'MEDIUM', 'SMS', now(), now()),
    ('otp', '\nورود به آروان آکادمی\n کد ورود: {token}\n{link}', 'token,link',2, 0, 'HIGH', 'SMS', now(), now()),
    ('online_payment_at_store', 'آروان آکادمی\nبرای ادامه‌ی فرایند خرید از طریق لینک زیر اقدام کنید:\n{link}', 'link',2, 0, 'HIGH', 'SMS', now(), now()),
    ('change_password_otp', 'کد تغییر رمز عبور شما در آروان کلاد :{token}\n{link}', 'token,link',2, 0, 'HIGH', 'SMS', now(), now()),
    ('forget_password_otp', 'کد تغییر رمز عبور شما در آروان کلاد :{token}\n{link}', 'token,link',2, 0, 'HIGH', 'SMS', now(), now()),
    ('change_card_password', 'کد تغییر رمز عبور شما در آروان کلاد :{token}\n{link}', 'token,link',2, 0, 'HIGH', 'SMS', now(), now()),
    ('register_otp', 'کد ثبت نام شما در آروان کلاد :{token}\n{link}', 'token,link',2, 0, 'HIGH', 'SMS', now(), now())
ON DUPLICATE KEY UPDATE `updated_at` = now();

INSERT INTO `provider_templates`
(`provider_id`, `template_id`, `code`, `created_at`, `updated_at`)
VALUES
    (1, 1, '168', now(), now()),
    (2, 2, 'verify', now(), now())
ON DUPLICATE KEY UPDATE `updated_at` = now();
