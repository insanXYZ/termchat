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

func (repo *UserRepository) FindGetUser(db *gorm.DB, req *entity.User) (*[]entity.User, error) {
	users := new([]entity.User)
	var err error
	if req.ID != "" {
		err = db.Where("id = ?", req.ID).Find(users).Error
	} else if req.Name != "" {
		err = db.Where("name like ?", "%"+req.Name+"%").Find(users).Error
	}
	return users, err
}
