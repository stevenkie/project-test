package user

import "strings"

func (u *userUC) ValidateSession(token string) (valid bool) {
	if token != "" {
		bearerToken := strings.Split(token, " ")
		if len(bearerToken) == 2 {
			userID, err := u.sessionRedisRepo.GetToken(bearerToken[1])
			if err != nil || userID == "0" {
				return false
			}
			return true
		}
	}
	return
}
