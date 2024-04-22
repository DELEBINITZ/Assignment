package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/akshay/libraryAssign/internal/library/exceptions"
	"github.com/akshay/libraryAssign/internal/library/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(email, role string) (string, error) {
	signingKey := []byte("hashKeyMaker")

	claims := &Claims{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Auth(c *gin.Context) {
	var token string
	if values := c.Request.Header["Authorization"]; len(values) > 0 {
		token = values[0]
	} else {
		err := errors.New("authorisation header missing")
		err = exceptions.GenerateNewServerError(exceptions.JWTTokenInvalid, err, "token missing", http.StatusUnauthorized)
		utils.RespondError(c, err, utils.ConstructErrorAPIResp(err))
		c.Abort()
		return
	}

	claims, err := VerifyToken(token)
	if err != nil {
		err = exceptions.GenerateNewServerError(exceptions.JWTTokenInvalid, err, "token unverified", http.StatusUnauthorized)
		utils.RespondError(c, err, utils.ConstructErrorAPIResp(err))
		c.Abort()
		return
	}

	c.Set("claims", claims)
	c.Next()
}

func VerifyToken(tokenString string) (*Claims, error) {
	signingKey := []byte("hashKeyMaker")
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func GetClaimsFromContext(ctx context.Context) (*Claims, error) {
	claimsInterface := ctx.Value("claims")
	claims, ok := claimsInterface.(*Claims)
	if !ok {
		serverErr := exceptions.GenerateNewServerError(exceptions.NoClaimsFound, nil, "claims not found", http.StatusBadRequest)
		return nil, serverErr
	}
	return claims, nil

}
