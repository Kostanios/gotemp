package repositories

import (
	"context"
	"github.com/kostanios/gotemp/models/usermodel"
	"github.com/kostanios/gotemp/test"
	"gorm.io/gorm"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var subEnding = int(time.Now().UTC().Add(time.Hour * 24).Unix())

var testUserSubscriptions = []usermodel.Subscription{
	{
		UserID:              1,
		Type:                usermodel.Yearly,
		Status:              usermodel.SubscriptionStatusActive,
		AutoRenew:           false,
		Ending:              &subEnding,
		AppleTransactionId:  nil,
		GoogleTransactionId: nil,
	},
}

func TestSubscriptions(t *testing.T) {
	ctx := context.Background()
	container, err := test.GetTestingDb()
	if err != nil {
		panic(err)
	}
	defer container.TerminateFunc(ctx)
	repo := NewSubscriptionRepository(container.DB)
	err = container.DB.AutoMigrate(&usermodel.Subscription{})
	if err != nil {
		panic(err)
	}

	t.Run("save subscriptions", func(t *testing.T) {
		err := repo.SaveUserSubscription(&testUserSubscriptions[0])
		assert.NoError(t, err)

		t.Run("get user's subscription successfully", func(t *testing.T) {
			sub, err := repo.GetUserSubscription(testUserSubscriptions[0].UserID)
			assert.NoError(t, err)
			assert.EqualValues(t, testUserSubscriptions[0], *sub)
		})

		t.Run("user subscription not found", func(t *testing.T) {
			sub, err := repo.GetUserSubscription(404)
			assert.EqualValues(t, 0, sub.UserID)
			assert.ErrorIs(t, gorm.ErrRecordNotFound, err)
		})
	})
}
