package repositories

import (
	"context"
	"github.com/kostanios/gotemp/models/usermodel"
	"github.com/kostanios/gotemp/test"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testUsers = []usermodel.AppUser{
	{
		ID:                1,
		FirstName:         "foo",
		LastName:          "baz",
		Phone:             "+995 599 99 99 99",
		Email:             "some@email.com",
		IsEmailVerified:   false,
		Pronoun:           usermodel.He,
		Birthday:          time.Now(),
		PushToken:         nil,
		InviteCode:        nil,
		HasAccess:         true,
		IsProfileComplete: true,
		Password:          nil,
		Salt:              nil,
		TwoFactorEnabled:  false,
		Subscription:      nil,
		Created:           0,
		Updated:           0,
	},
}

func TestUsers(t *testing.T) {
	ctx := context.Background()
	container, err := test.GetTestingDb()
	if err != nil {
		panic(err)
	}
	defer container.TerminateFunc(ctx)
	repo := NewUserRepository(container.DB)
	err = container.DB.AutoMigrate(&usermodel.AppUser{})
	if err != nil {
		panic(err)
	}

	t.Run("save user", func(t *testing.T) {
		err := repo.CreateUser(&testUsers[0])
		assert.NoError(t, err)

		t.Run("retrieve saved user", func(t *testing.T) {
			user, err := repo.GetUserByID(testUsers[0].ID)
			assert.NoError(t, err)
			assert.EqualValues(t, testUsers[0].ID, user.ID)
			assert.EqualValues(t, testUsers[0].FirstName, user.FirstName)
			assert.EqualValues(t, testUsers[0].LastName, user.LastName)
		})

		t.Run("update saved user", func(t *testing.T) {
			user, _ := repo.GetUserByID(testUsers[0].ID)
			err := user.SetUserData("updatedName", "updatedLastname", "updated@email.com", "they", nil)
			assert.NoError(t, err)
			err = repo.UpdateUser(user)
			assert.NoError(t, err)

			user, err = repo.GetUserByID(testUsers[0].ID)
			assert.NoError(t, err)
			assert.EqualValues(t, "updatedName", user.FirstName)
			assert.EqualValues(t, "updatedLastname", user.LastName)
			assert.EqualValues(t, "updated@email.com", user.Email)
			assert.EqualValues(t, usermodel.They, user.Pronoun)
		})
	})
}
