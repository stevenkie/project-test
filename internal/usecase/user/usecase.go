package user

import (
	loginModel "github.com/stevenkie/project-test/internal/model/login"
	userModel "github.com/stevenkie/project-test/internal/model/user"
)

// Usecase contract for user account interface
type Usecase interface {
	//GetUserByID get single data of user
	GetUserByID(userID string) (userModel.User, error)
	//InsertUser insert single data to table 'users'
	InsertUser(data userModel.InsertUser) error
	//UpdateUser Update single data to table 'users'
	UpdateUser(data userModel.UpdateUser) error
	//DeleteUser non-activate single user to table 'users'
	DeleteUser(userID string) error
	//Login login process, check user from db, and save token
	Login(data loginModel.Login) (string, error)
	// ValidateSession for validation session
	ValidateSession(token string) (valid bool)
}
