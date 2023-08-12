package services

import (
	"net/http"
	"strings"

	"repo/components/redis"
	"repo/crypto"
	"repo/flags"
	"repo/repositories"
	"repo/util"
)

type AuthService interface {
	// ExtractToken takes value of Authorization Header, validate expected token type and extract token value
	ExtractToken(authString string, expectedType string) (string, error)
	ValidateAccessToken(r *http.Request, purpose string) (string, string, error)
}

type authService struct {
	player repositories.PlayerRepository
}

func (s *authService) ExtractToken(authString string, expectedType string) (string, error) {
	if authString == "" {
		log.Error("Authorization Header is empty")
		return "", util.NewError("401")
	}
	// Extract token
	splitToken := strings.Split(authString, " ")
	if len(splitToken) != 2 {
		log.Error("Bearer token is malformed")
		return "", util.NewError("401")
	}

	tokenType := splitToken[0]
	if tokenType != expectedType {
		log.Errorf("Unmatch token type. Type: %s", tokenType)
		return "", util.NewError("401")
	}
	return splitToken[1], nil
}

func (s *authService) ValidateAccessToken(r *http.Request, purpose string) (string, string, error) {
	// Get authentication string
	authString := r.Header.Get(flags.HeaderKeyCOBRAAuthorization)
	// If client secret

	// Validate jwt token and get claim
	claim, err := s.ValidateJWT(authString)
	if err != nil {
		return "", claim, util.NewError("500")
	}

	data, err := redis.Get(claim)
	if err != nil {
		return "", claim, util.NewError("500")
	}
	return data, claim, nil
}

func (s *authService) ValidateJWT(authString string) (string, error) {
	// Extract token
	token, err := s.ExtractToken(authString, crypto.TokenTypeBearer)
	if err != nil {
		log.Error(err)
		return "", err
	}
	return token, nil
}

// func (s *authService) NewAnonymousToken() (*response.Success, error) {

// 	// Initiate new jwt with purpose
// 	token, expiredAt, err := crypto.NewJWTAnonymous(flags.ACLAuthenticatedAnonymous, 20160)
// 	if err != nil {
// 		return nil, util.NewError("401")
// 	}
// 	// Compose success message
// 	success := response.Success{
// 		Result: "Success",
// 		Header: response.NewToken(token, expiredAt),
// 	}
// 	return &success, nil
// }
