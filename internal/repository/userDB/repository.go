package userdb

import (
	"github.com/jmoiron/sqlx"
	userModel "github.com/stevenkie/project-test/internal/model/user"
	"github.com/stevenkie/project-test/util"
)

// Repository repository interface for interacting with user DB
type Repository interface {
	//GetUserByID from db and return single data of user
	GetUserByID(userID string) (userModel.User, error)
	//InsertUser insert single data to table 'users'
	InsertUser(tx util.Transaction, data userModel.InsertUser) error
	//UpdateUser Update single data to table 'users'
	UpdateUser(tx util.Transaction, data userModel.UpdateUser) error
	//DeleteUser non-activate single user to table 'users'
	DeleteUser(tx util.Transaction, userID string) error
	//GetUserByEmail used for login process, get user by email
	GetUserByEmail(email string) (userModel.User, error)
	//GetDB for used withing util.WithTransaction
	GetDB() *sqlx.DB
}
