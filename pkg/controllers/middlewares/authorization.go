package middlewares

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vsrecorder/vsr-apiserver/pkg/controllers/helpers"
)

const (
	TOKEN_LIFETIME_SECOND = time.Duration(15) * time.Second
)

type VSRClaims struct {
	jwt.RegisteredClaims
	UID string `json:"uid"`
}

func generateToken(uid string, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(TOKEN_LIFETIME_SECOND).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func parseToken(tokenString string, secretKey string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &VSRClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func RequiredAuthorization(ctx *gin.Context) {
	secretKey := os.Getenv("VSRECORDER_JWT_SECRET")

	header := http.Header{}
	header.Add("Authorization", ctx.GetHeader("Authorization"))
	tokenString := strings.TrimPrefix(header.Get("Authorization"), "Bearer ")

	token, err := parseToken(tokenString, secretKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	claims := token.Claims.(*VSRClaims)
	helpers.SetUID(ctx, claims.UID)
}

func OptionalAuthorization(ctx *gin.Context) {
	secretKey := os.Getenv("VSRECORDER_JWT_SECRET")

	header := http.Header{}
	header.Add("Authorization", ctx.GetHeader("Authorization"))

	if header.Get("Authorization") == "" {
		return
	}

	tokenString := strings.TrimPrefix(header.Get("Authorization"), "Bearer ")
	token, err := parseToken(tokenString, secretKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	claims := token.Claims.(*VSRClaims)
	helpers.SetUID(ctx, claims.UID)
}
