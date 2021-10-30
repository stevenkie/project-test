package postgres

import (
	"github.com/pkg/errors"
	userModel "github.com/stevenkie/project-test/internal/model/user"
	"github.com/stevenkie/project-test/util"
	"golang.org/x/crypto/bcrypt"
)

//InsertUser insert single data to table 'users'
func (upg *userPostgresRepo) InsertUser(tx util.Transaction, data userModel.InsertUser) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "hash password failed")
	}
	_, err = tx.Exec(insertUser, data.Email, data.Address, hashedPassword)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//UpdateUser Update single data to table 'users'
func (upg *userPostgresRepo) UpdateUser(tx util.Transaction, data userModel.UpdateUser) error {
	_, err := tx.Exec(updateUser, data.ID, data.Email, data.Address)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//DeleteUser non-activate single user to table 'users'
func (upg *userPostgresRepo) DeleteUser(tx util.Transaction, userID string) error {
	_, err := tx.Exec(deleteUser, userID)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
