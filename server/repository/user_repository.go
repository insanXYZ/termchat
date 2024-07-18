package repository

import (
	"backend/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) AppendChatUsers(db *gorm.DB, sender *entity.User, receiver *entity.User) error {
	return db.Model(sender).Association("ChatUsers").Append(receiver)
}

func (repo *UserRepository) FindWithName(db *gorm.DB, name string) ([]entity.User, error) {
	var users []entity.User
	err := db.Where("name like = ?", "%"+name+"%").Find(&users).Error
	return users, err
}
