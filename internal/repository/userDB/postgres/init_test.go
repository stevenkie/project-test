package postgres

import (
	"testing"

	"github.com/jmoiron/sqlx"
	userPGRepo "github.com/stevenkie/project-test/internal/repository/userdb"
)

func TestInitUserPGRepo(t *testing.T) {
	type args struct {
		db *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want userPGRepo.Repository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitUserPGRepo(tt.args.db)
		})
	}
}
