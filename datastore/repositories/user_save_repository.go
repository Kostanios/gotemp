package repositories

import (
	"github.com/kostanios/gotemp/models/savedmodel"
	"gorm.io/gorm"
)

type UserSaveRepository interface {
	GetUserSave(userID uint, saveType string, id string) (*savedmodel.Saved, error)
	GetSaves(userID uint, offset int, limit int) ([]savedmodel.Saved, error)
	GetSavesByType(userID uint, saveType string, page int) ([]savedmodel.Saved, error)
	NewSave(model *savedmodel.Saved) error
	DeleteSave(model *savedmodel.Saved) error
}

type userSaveRepository struct {
	db *gorm.DB
}

func NewUserSavesRepository(db *gorm.DB) UserSaveRepository {
	return &userSaveRepository{db: db}
}

func (repo *userSaveRepository) GetUserSave(userID uint, saveType string, id string) (*savedmodel.Saved, error) {
	save := new(savedmodel.Saved)
	err := repo.db.First(save, "user_id = ? AND type = ? AND id = ?", userID, saveType, id).Error
	return save, err
}

func (repo *userSaveRepository) GetSaves(userID uint, offset int, limit int) ([]savedmodel.Saved, error) {
	saves := make([]savedmodel.Saved, 0, 20)
	err := repo.db.Offset(offset).Limit(limit).Find(&saves, "user_id = ?", userID).Order("created desc").Error
	return saves, err
}

func (repo *userSaveRepository) GetSavesByType(userID uint, saveType string, page int) ([]savedmodel.Saved, error) {
	saves := make([]savedmodel.Saved, 0, 20)
	err := repo.db.Offset(20*page).Limit(20).Find(&saves, "user_id = ? AND type = ?", userID, saveType).Order("created desc").Error
	return saves, err
}

func (repo *userSaveRepository) NewSave(model *savedmodel.Saved) error {
	return repo.db.Table(model.TableName()).Save(model).Error
}

func (repo *userSaveRepository) DeleteSave(model *savedmodel.Saved) error {
	return repo.db.Delete(model).Error
}
