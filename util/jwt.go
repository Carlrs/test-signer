package util

import "github.com/golang-jwt/jwt/v5"
import "github.com/google/uuid"

func ExtractUserUUID(jwtString string) (uuid.UUID, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(jwtString, jwt.MapClaims{})
	if err != nil {
		return [16]byte{}, err
	}
	userString, err := token.Claims.GetSubject()
	if err != nil {
		return [16]byte{}, err
	}
	return uuid.Parse(userString)
}
