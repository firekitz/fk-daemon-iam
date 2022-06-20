package main

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

type FKClaims struct {
	ProjectId   int64 `json:"pi,omitempty"`
	DomainId    int64 `json:"di,omitempty"`
	AccountId   int64 `json:"ai,omitempty"`
	AccountType int64 `json:"at,omitempty"`
	jwt.StandardClaims
}

type AuthTokenClaims struct {
	UserID string   `json:"id"`
	Name   string   `json:"name"`
	Email  string   `json:"mail"`
	Role   []string `json:"role"`

	jwt.StandardClaims
}

func TestTokenBuild() {
	//at := AuthTokenClaims {
	//	StandardClaims: jwt.StandardClaims{
	//		IssuedAt: 1636680665,
	//		Subject: "0",
	//	},
	//}

	var JWT_SECRET = "ZHR4LWp3dC1zZWNyZXQ="
	//var JWT_SECRET = "fk-jwt-secret"

	//secretDecoded, _ := b64.StdEncoding.DecodeString(JWT_SECRET)
	//secret := b64.StdEncoding.EncodeToString(secretDecoded)
	//logrus.Infof("secret: %s", secret)
	//secret := b64.StdEncoding.EncodeToString([]byte(JWT_SECRET))

	//atoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &at)
	//signedAuthToken, err := atoken.SignedString([]byte(JWT_SECRET))

	//if err != nil {
	//	logrus.Infof("permissions: %s", err)
	//}
	var js = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIwIiwiaWF0IjoxNjM2Njk0NzYwLCJleHAiOjE2NjgyMzA3NjAsImRpIjoxMDAsInBpIjo1LCJhaSI6MiwiYXQiOjEsInR5cGUiOiJhY2Nlc3MifQ.XpFsc8jF6xyba_Dx2VVpT7Ho0blLBWwIbosnJaJW_yU"
	//var goToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2MzY2ODA2NjUsInN1YiI6IjAifQ.38_AAA4kdG1qwqGEpfJ1q76e3oO3qL6Tb-yiWBbEwKs"

	//var newToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjAsImlhdCI6MTYzNjYxNzI5NywiZXhwIjoyMjY3NzY5MjkwLCJkaSI6MTAwLCJwaSI6NSwiYWkiOjIsImF0IjoxLCJ0eXBlIjoiYWNjZXNzIn0.p4ax681nuiGE-vdEUpMOxD9Sn6peFbiy6HnUrtVp4-c"

	key := func(token *jwt.Token) (interface{}, error) {
		logrus.Infof("in key")
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			var ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
			logrus.Infof("in key 0")
			return nil, ErrUnexpectedSigningMethod
		}
		logrus.Infof("in key 1")
		return []byte(JWT_SECRET), nil
		//return []byte(secret), nil
		//return []byte("fk-jwt-secret"), nil
		//return []byte(secretDecoded), nil
	}

	//tok, err := jwt.ParseWithClaims(signedAuthToken, &AuthTokenClaims{}, key)
	tok, _ := jwt.ParseWithClaims(js, &FKClaims{}, key)
	println(tok.Valid)

	//logrus.Infof("permissions: %s", signedAuthToken)
}
func main() {
	TestTokenBuild()
}
