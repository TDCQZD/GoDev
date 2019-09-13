package model

import (
	"fmt"
	"go_project_code/APIServer/RESTful/APIServer_startup_script/src/pkg/auth"
	"go_project_code/APIServer/RESTful/APIServer_startup_script/src/pkg/constvar"

	validator "gopkg.in/go-playground/validator.v9"
)

type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (um *UserModel) TableName() string {
	return "tb_users"
}

func (um *UserModel) Create() error {
	return DB.Self.Create(&um).Error
}
func (um *UserModel) Update() error {
	return DB.Self.Save(um).Error
}
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id
	return DB.Self.Delete(&user).Error
}

func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("username = ?", username).First(&u)
	return u, d.Error
}

func ListUser(username string, offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}
	users := make([]*UserModel, 0)
	var count uint64

	where := fmt.Sprintf("username like '%%%s%%'", username)

	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}
	return users, count, nil
}

func (um *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(um.Password, pwd)
	return
}

func (um *UserModel) Encrypt() (err error) {
	um.Password, err = auth.Encrypt(um.Password)
	return
}

func (um *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(um)
}
