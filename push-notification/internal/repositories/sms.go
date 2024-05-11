package repositories

import (
	"context"
	"errors"
	"notification/internal/constants"
	"notification/internal/domain"
	"notification/internal/model"
	"time"

	"gorm.io/gorm"
)

type smsRepository struct {
	CommonBehaviourRepository[model.SMS, domain.SMS]
	db *gorm.DB
}

func NewSmsRepository(db *gorm.DB) SmsRepository {
	return smsRepository{
		CommonBehaviourRepository: NewCommonBehaviour[model.SMS, domain.SMS](db),
		db:                        db,
	}
}

func (s smsRepository) Create(ctx context.Context, sms domain.SMS) (domain.SMS, error) {
	mSms := model.SMS{
		Mobile:     sms.Mobile,
		Content:    sms.Content,
		Status:     sms.Status.String(),
		ProviderId: sms.ProviderId,
		TemplateID: sms.TemplateID,
		ExpiredAt:  sms.ExpiresAt,
	}

	if err := s.Save(ctx, &mSms); err != nil {
		return domain.SMS{}, err
	}
	ds := mSms.ToDomain().(domain.SMS)

	return ds, nil
}

func (s smsRepository) GetByID(ctx context.Context, id uint) (domain.SMS, error) {
	sms, err := s.ByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.SMS{}, constants.ErrSmsNotFound
		}
		return domain.SMS{}, err
	}

	return sms, nil
}

func (s smsRepository) UpdateStatus(c context.Context, id uint, status domain.NotificationStatus) error {
	return s.db.WithContext(c).Model(&model.SMS{}).Where("id = ?", id).Update("status", status.String()).Error
}

func (s smsRepository) CountSince(ctx context.Context, mobile string, since time.Time) (int, error) {
	var count int64

	err := s.db.WithContext(ctx).
		Model(&model.SMS{}).
		Where("mobile = ?", mobile).
		Where("updated_at >= ?", since).
		Where("status = ?", constants.SMS_STATUS_SENT).
		Count(&count).
		Error

	if err != nil {
		return 0, err
	}

	return int(count), nil
}
