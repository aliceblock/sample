package repository

import (
	"github.com/aliceblock/sample/cmd/server/models"
	"github.com/aliceblock/sample/internal/mode"
)

// UserRepo binded with UserModel
type UserRepo struct {
	Mode  mode.DataMode
	Model models.UserModel
}

// NewUserRepo init new UserRepo
func NewUserRepo(dm mode.DataMode) *UserRepo {
	var model models.UserModel
	if dm == mode.MOCK {
		model = &models.MockUserModel{}
	} else {
		model = &models.ProdUserModel{}
	}
	return &UserRepo{
		Mode:  dm,
		Model: model,
	}
}

// SetMode set new mode for UserRepo
func (u *UserRepo) SetMode(dm mode.DataMode) {
	u.Mode = dm
	if dm == mode.MOCK {
		u.Model = &models.MockUserModel{}
	} else {
		u.Model = &models.ProdUserModel{}
	}
}
