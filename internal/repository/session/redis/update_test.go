package redis

import (
	"errors"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
)

func Test_sessionRedisRepo_SetUserToken(t *testing.T) {
	db, mock := redismock.NewClientMock()
	mock.ExpectSet("test", "test", DefaultCacheExpiration).SetVal("succes")

	dbFail, mockFail := redismock.NewClientMock()
	mockFail.ExpectSet("test", "test", DefaultCacheExpiration).SetErr(errors.New("test"))

	type fields struct {
		redis *redis.Client
	}
	type args struct {
		userID string
		token  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "positive test case",
			fields: fields{
				redis: db,
			},
			args: args{
				userID: "test",
				token:  "test",
			},
			wantErr: false,
		},
		{
			name: "negative test case",
			fields: fields{
				redis: dbFail,
			},
			args: args{
				userID: "test",
				token:  "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &sessionRedisRepo{
				redis: tt.fields.redis,
			}
			if err := sr.SetUserToken(tt.args.userID, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("sessionRedisRepo.SetUserToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
