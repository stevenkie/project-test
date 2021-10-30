package user

import (
	"github.com/pkg/errors"
	loginModel "github.com/stevenkie/project-test/internal/model/login"
	"github.com/stevenkie/project-test/util"
	"golang.org/x/crypto/bcrypt"
)

//Login login process, check user from db, and save token
func (u *userUC) Login(data loginModel.Login) (string, error) {
	var token string
	user, err := u.userPGRepo.GetUserByEmail(data.Email)
	if err != nil {
		return token, errors.WithStack(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return token, errors.WithStack(err)
	}
	token, err = util.GenerateToken(u.config.Server.Secret)
	if err != nil {
		return token, errors.WithStack(err)
	}
	err = u.sessionRedisRepo.SetUserToken(user.ID, token)
	if err != nil {
		return token, errors.WithStack(err)
	}
	return token, nil
}
