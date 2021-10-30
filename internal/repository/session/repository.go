package session

// Repository repository interface for interacting with Session Redis
type Repository interface {
	//SetUserToken to redis, this cache is used for token auth
	SetUserToken(userID string, token string) error
	//GetToken from redis and get user_id
	GetToken(token string) (string, error)
}
