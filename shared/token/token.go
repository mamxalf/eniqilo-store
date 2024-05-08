package token

import (
	"github.com/mamxalf/eniqilo-store/config"
	"github.com/mamxalf/eniqilo-store/shared/constant"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/rs/zerolog/log"
)

// JWTSigningMethod is JWT's signing method
var jwtSigningMethod = jwt.SigningMethodHS256

// GenerateToken will generate both access and refresh token
// for current user.
// Access Token will be expired in 15 Minutes
// Refresh Token will be expired in 6 Months
func GenerateToken(user *UserData, params *GenerateTokenParams) (token Token, err error) {
	jwtToken, err := GenerateJWT(user, params.AccessTokenSecret, params.AccessTokenExpiry)
	if err != nil {
		return
	}

	token = Token{
		AccessToken: jwtToken,
	}

	return
}

// GenerateJWT is
func GenerateJWT(user *UserData, tokenSecret string, tokenExpiry time.Duration) (signedToken string, err error) {
	exp := time.Now().UTC().Add(tokenExpiry)
	claims := JWTToken{
		StandardClaims: jwt.StandardClaims{
			Issuer:    config.Get().AppName,
			ExpiresAt: exp.Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
			Subject:   user.ID,
		},
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		OwnerID:     user.ID,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err = token.SignedString([]byte(tokenSecret))
	if err != nil {
		return signedToken, err
	}

	return signedToken, err
}

func VerifyJwtToken(token, tokenSecret string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Err(constant.ErrInvalidAuthorization).Msg("VerifyJwtToken")
			return nil, constant.ErrInvalidAuthorization
		}
		return []byte(tokenSecret), nil
	})
	if err != nil {
		log.Err(constant.ErrInvalidAuthorization).Msg("VerifyJwtToken")
		return nil, err
	}
	return jwtToken, nil
}
