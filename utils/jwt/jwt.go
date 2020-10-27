package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"/utils/responses"
	"io/ioutil"
	"net/http"
	"time"
)

func GenerateJWT(id int64) (string, responses.ApiResponseList) {
	sign := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"aud": 6,
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 365).Unix(),
		"jti": "8986520cc8f9f6ea4cce9fa82cbf7d10e2c9d0af97cdf9e717c0b04773493fd01e66a8d18b1158fe",
		"sub": id,
	})

	privateKey, readErr := ioutil.ReadFile("./configs/rsa-key/oauth-private.key")
	if readErr != nil {
		return "", responses.ErrorApiResponse(http.StatusInternalServerError, readErr.Error())
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", responses.ErrorApiResponse(http.StatusInternalServerError, err.Error())
	}

	token, err := sign.SignedString(key)

	if err != nil {
		return "", responses.ErrorApiResponse(http.StatusInternalServerError, err.Error())
	}

	return token, nil
}
