package repositories

import (
	"github.com/kostanios/gotemp/models/usermodel"
	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	GetUserSubscription(userID uint) (*usermodel.Subscription, error)
	GetGoogleSubscription(userID uint, googleSubKey string) (*usermodel.Subscription, error)
	SaveUserSubscription(subscription *usermodel.Subscription) error
	DeleteUsersExistingSubscription(userID uint) error
}

type subscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &subscriptionRepository{
		db: db,
	}
}

func (repo *subscriptionRepository) GetUserSubscription(userID uint) (*usermodel.Subscription, error) {
	model := new(usermodel.Subscription)
	err := repo.db.Table(model.TableName()).First(model, "user_id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (repo *subscriptionRepository) GetGoogleSubscription(userID uint, googleSubKey string) (*usermodel.Subscription, error) {
	model := new(usermodel.Subscription)
	err := repo.db.Table(model.TableName()).First(model, "user_id = ? AND google_transaction_id = ?", userID, googleSubKey).Error
	if err != nil {
		return nil, err
	}
	return model, err
}

func (repo *subscriptionRepository) SaveUserSubscription(subscription *usermodel.Subscription) error {
	return repo.db.Table(subscription.TableName()).Save(subscription).Error
}

func (repo *subscriptionRepository) DeleteUsersExistingSubscription(userID uint) error {
	err := repo.db.Table((*usermodel.Subscription)(nil).TableName()).Delete(&usermodel.Subscription{}, userID).Error
	return err
}

func (repo *subscriptionRepository) UpdateSubscription(subscription *usermodel.Subscription) error {
	return repo.db.Table(subscription.TableName()).Where("user_id = ?", subscription.UserID).Updates(subscription).Error
}

func (repo *subscriptionRepository) SaveUserSubscriptionsBulk(subscriptions []*usermodel.Subscription) error {
	if len(subscriptions) == 0 {
		return nil
	}

	return repo.db.Transaction(func(tx *gorm.DB) error {
		for _, subscription := range subscriptions {
			if err := tx.Table(subscription.TableName()).Save(subscription).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
