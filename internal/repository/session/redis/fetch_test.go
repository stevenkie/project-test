package redis

import (
	"errors"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
)

func Test_sessionRedisRepo_GetToken(t *testing.T) {
	db, mock := redismock.NewClientMock()
	mock.ExpectGet("test").SetVal("test")

	dbFail, mockFail := redismock.NewClientMock()
	mockFail.ExpectGet("test").SetErr(errors.New("test"))

	type fields struct {
		redis *redis.Client
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "positive test case",
			fields: fields{
				redis: db,
			},
			args: args{
				token: "test",
			},
			want: "test",
		},
		{
			name: "negative test case",
			fields: fields{
				redis: dbFail,
			},
			args: args{
				token: "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &sessionRedisRepo{
				redis: tt.fields.redis,
			}
			got, err := sr.GetToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("sessionRedisRepo.GetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sessionRedisRepo.GetToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
