package controller

import (
	"event-app/config"
	"event-app/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthInputRegister struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type AuthInputLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

func RegisterUser(context *gin.Context) {
	var input AuthInputRegister

	// validation
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// HashPassword
	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if errHash != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Gagal Encripsi password",
		})
		return
	}

	// Simpan ke DB
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	userCreated := config.DB.Create(&user).Error

	if userCreated != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Email Mungkin sudah terdaftar",
		})
		return
	}

	// Response

	context.JSON(http.StatusCreated, gin.H{
		"Message": "Berhasil daftar",
		"user": gin.H{
			"id":     user.ID,
			"name":   user.Name,
			"email":  user.Email,
			"events": user.Events,
		},
	})
}

func LoginUser(context *gin.Context) {
	var input AuthInputLogin

	// validation
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	userData := config.DB.Where("emial = ?", input.Email).First(&user)
	if userData != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "email belum terdaftar",
		})
		return
	}

	errMatchPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if errMatchPassword != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "email belum terdaftar",
		})
		return
	}

	// Token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal Membuat Token",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login Berhasil",
		"token":   tokenString,
		"user": gin.H{
			"id":     user.ID,
			"name":   user.Name,
			"email":  user.Email,
			"events": user.Events,
		},
	})
}
