package postgres

import (
	"testing"

	"github.com/jmoiron/sqlx"
	itemPGRepo "github.com/stevenkie/project-test/internal/repository/item"
)

func TestInitItemPGRepo(t *testing.T) {
	type args struct {
		db *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want itemPGRepo.Repository
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitItemPGRepo(tt.args.db)
		})
	}
}
