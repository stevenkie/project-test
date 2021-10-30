package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	userModel "github.com/stevenkie/project-test/internal/model/user"
)

// GetDB function for getting database for util WithTransactionFunction
func (upg *userPostgresRepo) GetDB() *sqlx.DB {
	return upg.db
}

func (upg *userPostgresRepo) GetUserByID(userID string) (userModel.User, error) {
	var result userModel.User
	err := upg.db.Get(&result, getUserByID, userID)
	if err != nil {
		return result, errors.WithStack(err)
	}
	return result, nil
}

//GetUserByEmail used for login process, get user by email
func (upg *userPostgresRepo) GetUserByEmail(email string) (userModel.User, error) {
	var result userModel.User
	err := upg.db.Get(&result, getUserByEmail, email)
	if err != nil {
		return result, errors.WithStack(err)
	}
	return result, nil
}
