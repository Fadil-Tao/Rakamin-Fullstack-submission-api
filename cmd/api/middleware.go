package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rakamins-pbi/final-task-pbi-rakamin-fullstack-HadadFadilah/Internals/models"
)

func CheckAuth(ctx *gin.Context) {
	// Mendapatkan token dari http cookie
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	// decode dan validasi
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// cek ekspayer
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		//  nyari menggunakan GORM 
		var user models.User
		models.DB.First(&user, "id = ?", claims["sub"])

		// kalau ga nemu
		if user.ID == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		// attach ke request ctx gin
		ctx.Set("user", user.ID)

		fmt.Println("Success on middleware")
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
