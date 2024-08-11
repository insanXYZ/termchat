package service

import (
	"backend/entity"
	"backend/model"
	"backend/repository"
	"backend/utils/valid"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	DB        *gorm.DB
	Viper     *viper.Viper
	Validator *validator.Validate
	UserRepo  *repository.UserRepository
}

func NewUserService(DB *gorm.DB, viper *viper.Viper, validator *validator.Validate, userRepo *repository.UserRepository) *UserService {
	return &UserService{DB: DB, Viper: viper, Validator: validator, UserRepo: userRepo}
}

func (service *UserService) Register(req *model.RegisterUser) error {
	err := service.Validator.Struct(req)
	if err != nil {
		return valid.HandleValidatorStruct(err)
	}

	if count := service.UserRepo.CountWhere(service.DB, "email = ?", req.Email); count == 0 {
		passwordByte, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
		if err != nil {
			return err
		}

		user := entity.User{
			ID:       uuid.New().String(),
			Name:     req.Name,
			Email:    req.Email,
			Bio:      "Hi, im use termchat",
			Password: string(passwordByte),
		}

		err = service.UserRepo.Create(service.DB, &user)
		if err != nil {
			return err
		}

		return nil
	}

	return errors.New("email already taken")

}

func (service *UserService) Login(req *model.LoginUser) (*entity.User, *string, error) {
	err := service.Validator.Struct(req)
	if err != nil {
		return nil, nil, valid.HandleValidatorStruct(err)
	}

	user := &entity.User{
		Email: req.Email,
	}

	err = service.UserRepo.Take(service.DB, user)
	if err != nil {
		return nil, nil, errors.New("username or password wrong")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, nil, errors.New("username or password wrong")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"name": user.Name,
		"exp":  time.Now().Add(time.Duration(service.Viper.GetInt("JWT_EXP")) * time.Minute).Unix(),
	})

	signedString, err := claims.SignedString([]byte(service.Viper.GetString("JWT_SECRET_KEY")))
	if err != nil {
		return nil, nil, errors.New("username or password wrong")
	}

	return user, &signedString, nil
}

func (service *UserService) Refresh(claims jwt.MapClaims) (*string, error) {
	newClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  claims["sub"],
		"name": claims["name"],
		"exp":  time.Now().Add(time.Duration(service.Viper.GetInt("JWT_EXP")) * time.Minute).Unix(),
	})

	signedString, err := newClaims.SignedString([]byte(service.Viper.GetString("JWT_SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	return &signedString, nil
}

func (service *UserService) GetUser(req *model.GetUser) (*[]entity.User, error) {

	err := service.Validator.Struct(req)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		ID:   req.ID,
		Name: req.Name,
	}

	users, err := service.UserRepo.FindGetUser(service.DB, user)
	if err != nil {
		return nil, err
	}

	if len(*users) == 0 {
		return nil, errors.New("result not found")
	}

	return users, nil
}

func (service *UserService) UpdateUser(claims jwt.MapClaims, req *model.UpdateUser) (*entity.User, error) {
	err := service.Validator.Struct(req)
	if err != nil {
		return nil, valid.HandleValidatorStruct(err)
	}

	user := &entity.User{
		ID: claims["sub"].(string),
	}

	err = service.UserRepo.Take(service.DB, user)
	if err != nil {
		return nil, err
	}

	update := &entity.User{
		Name:  req.Name,
		Email: req.Email,
		Bio:   req.Bio,
	}

	if req.Password != "" {

		password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
		if err != nil {
			return nil, err
		}
		update.Password = string(password)
	}

	err = service.UserRepo.Updates(service.DB, user, update)
	if err != nil {
		return nil, err
	}

	return user, nil

}
