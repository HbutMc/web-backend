package main

import (
	"encoding/json"
	"errors"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Resp(message string, err error, data any) gin.H {
	return gin.H{
		"message": message,
		"error":   err.Error(),
		"data":    data,
	}
}

func SetBodyToCtx(c *gin.Context) error {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	c.Set("body", body)
	return nil
}

func GetBodyByCtx(c *gin.Context, dest any) error {

	dataAny, exists := c.Get("body")
	if !exists {
		return errors.New("键body不存在")
	}

	data, ok := dataAny.([]byte)
	if !ok {
		return errors.New("键body值的类型不是[]byte")
	}

	return json.Unmarshal(data, dest)
}

func GetJwt(userId uint) (string, error) {
	token, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userId,
			"exp":    time.Now().Add(24 * time.Hour).Unix(),
		}).SignedString([]byte(Config.JWTKey))
	return "Bearer " + token, err
}
