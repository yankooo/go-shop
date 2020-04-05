package m_token

import (
	"github.com/yankooo/school-eco/be/constant"
	"github.com/yankooo/school-eco/be/model"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(constant.JwtSecret)

func GenerateToken(accountId uint64, openId string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := model.Claims{
		AccountId:   accountId,
		OpenId:      openId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "michael.liu",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*model.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {

		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*model.Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
