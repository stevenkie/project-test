package user

import (
	"github.com/pkg/errors"
	userModel "github.com/stevenkie/project-test/internal/model/user"
	"github.com/stevenkie/project-test/util"
)

//GetUserByID get single data of user
func (u *userUC) GetUserByID(userID string) (userModel.User, error) {
	user, err := u.userPGRepo.GetUserByID(userID)
	if err != nil {
		return user, errors.WithStack(err)
	}
	return user, nil
}

//InsertUser insert single data to table 'users'
func (u *userUC) InsertUser(data userModel.InsertUser) error {
	err := data.Validate()
	if err != nil {
		return errors.WithStack(err)
	}
	err = util.WithTransaction(u.userPGRepo.GetDB(), func(tx util.Transaction) error {
		err := u.userPGRepo.InsertUser(tx, data)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//UpdateUser Update single data to table 'users'
func (u *userUC) UpdateUser(data userModel.UpdateUser) error {
	err := data.Validate()
	if err != nil {
		return errors.WithStack(err)
	}
	err = util.WithTransaction(u.userPGRepo.GetDB(), func(tx util.Transaction) error {
		err := u.userPGRepo.UpdateUser(tx, data)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//DeleteUser non-activate single user to table 'users'
func (u *userUC) DeleteUser(userID string) error {
	return util.WithTransaction(u.userPGRepo.GetDB(), func(tx util.Transaction) error {
		err := u.userPGRepo.DeleteUser(tx, userID)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
}
