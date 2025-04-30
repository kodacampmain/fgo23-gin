package handlers

import (
	"fgo23-gin/internal/models"
	"fgo23-gin/internal/repositories"
	"fgo23-gin/pkg"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthRepo *repositories.AuthRepo
}

func NewAuthHandler(AuthRepo *repositories.AuthRepo) *AuthHandler {
	return &AuthHandler{AuthRepo: AuthRepo}
}

func (a *AuthHandler) Register(ctx *gin.Context) {
	var body models.Student
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Terjadi kesalahan server",
		})
		return
	}

	hash := pkg.InitHashConfig()
	hash.UseDefaultConfig()
	hashedPass, err := hash.GenHashedPassword(body.Password)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Hash failed",
		})
		return
	}

	cmd, err := a.AuthRepo.AddNewUser(ctx.Request.Context(), body.Name, hashedPass)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Register failed",
		})
		return
	}
	if cmd.RowsAffected() == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Register failed",
		})
		return
	}
	// opsi kalau berhasil register
	// 1. otomatis login
	// buatkan token, lalu return token nya ke client
	// 2. disuruh login lagi
	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("Selamat Datang, %s! Silahkan login", body.Name),
	})
}
func (a *AuthHandler) Login(ctx *gin.Context) {
	var body models.Student
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Terjadi kesalahan server",
		})
		return
	}

	result, err := a.AuthRepo.GetUserData(ctx.Request.Context(), body.Name)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Terjadi kesalahan server",
		})
		return
	}

	hash := pkg.InitHashConfig()
	valid, err := hash.CompareHashAndPassword(result.Password, body.Password)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Terjadi kesalahan server",
		})
		return
	}
	if !valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid username/password",
		})
		return
	}
	// kalau sudah berhasil login, maka berikan identitas (jwt)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login Berhasil",
	})
}
