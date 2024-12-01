package repositories

import (
	"github.com/kostanios/gotemp/models/usermodel"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(userID uint) (*usermodel.AppUser, error)
	GetUserByIDWithSubscription(userID uint) (*usermodel.AppUser, error)
	UpdateUser(user *usermodel.AppUser) error
	CreateUser(user *usermodel.AppUser) error
	GetInvitedUsersWithInviteCode(code string) ([]usermodel.AppUser, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) GetUserByID(userID uint) (*usermodel.AppUser, error) {
	user := new(usermodel.AppUser)
	err := repo.db.Table(user.TableName()).First(user, "id = ?", userID).Error
	return user, err
}

func (repo *userRepository) GetUserByIDWithSubscription(userID uint) (*usermodel.AppUser, error) {
	user := new(usermodel.AppUser)
	err := repo.db.Table(user.TableName()).Preload("Subscription").First(user, "id = ?", userID).Error
	return user, err
}

func (repo *userRepository) UpdateUser(user *usermodel.AppUser) error {
	return repo.db.Table(user.TableName()).Save(user).Error
}

func (repo *userRepository) CreateUser(user *usermodel.AppUser) error {
	return repo.db.Table(user.TableName()).Save(user).Error
}

func (repo *userRepository) GetInvitedUsersWithInviteCode(code string) ([]usermodel.AppUser, error) {
	var users []usermodel.AppUser
	err := repo.db.Table((*usermodel.AppUser)(nil).TableName()).Where("invite_code = ?", code).Scan(&users).Error
	return users, err
}
