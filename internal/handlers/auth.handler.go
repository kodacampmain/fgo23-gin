package handlers

import (
	"fgo23-gin/internal/models"
	"fgo23-gin/internal/repositories"
	"fgo23-gin/pkg"
	"fmt"
	"log"
	"net/http"
	"strings"

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
		log.Println("[DEBUG] Binding Error", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Terjadi kesalahan server",
		})
		return
	}

	log.Println("[DEBUG] body ", body)

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
	claims := pkg.NewClaims(result.Id, result.Role)
	token, err := claims.GenerateToken()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Terjadi kesalahan server",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login Berhasil",
		"token":   token,
	})
}

func (a *AuthHandler) VerifyToken(ctx *gin.Context) {
	// 1. ambil token dari header
	bearerToken := ctx.GetHeader("Authorization")
	if bearerToken == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Silahkan login terlebih dahulu",
		})
		return
	}
	// 2. pisahkan token dari bearer
	token := strings.Split(bearerToken, " ")[1]
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Silahkan login terlebih dahulu",
		})
		return
	}
	// 3. verifikasi token
	claims := &pkg.Claims{}
	if err := claims.VerifyToken(token); err != nil {
		log.Println(err.Error())
		if err.Error() == "expired token" || err.Error() == "token has invalid claims: token is expired" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Silahkan login kembali",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Terjadi kesalahan server",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    claims,
	})
}
