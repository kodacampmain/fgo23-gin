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

// Register
// @summary					Register Student
// @router					/auth/new [post]
// @accept					json
// @param					body body models.AuthForm true "register information"
// @produce					json
// @failure					500 {object} models.ErrorResponse
// @success					201 {object} models.Response
func (a *AuthHandler) Register(ctx *gin.Context) {
	var body models.Student
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.InternalServerErrorCode,
				Status:  http.StatusInternalServerError,
				Message: "terjadi kesalahan server",
			},
		})
		return
	}

	hash := pkg.InitHashConfig()
	hash.UseDefaultConfig()
	hashedPass, err := hash.GenHashedPassword(body.Password)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.InternalServerErrorCode,
				Status:  http.StatusInternalServerError,
				Message: "terjadi kesalahan server",
			},
		})
		return
	}

	cmd, err := a.AuthRepo.AddNewUser(ctx.Request.Context(), body.Name, hashedPass)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.RegisterFailedCode,
				Status:  http.StatusInternalServerError,
				Message: "register failed",
			},
		})
		return
	}
	if cmd.RowsAffected() == 0 {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.RegisterFailedCode,
				Status:  http.StatusInternalServerError,
				Message: "register failed",
			},
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

// Login
// @summary					Login Student
// @router					/auth [post]
// @accept					json
// @param					body body models.AuthForm true "login information"
// @produce					json
// @failure					500 {object} models.ErrorResponse
// @failure					401 {object} models.ErrorResponse
// @success					200 {object} models.Response
func (a *AuthHandler) Login(ctx *gin.Context) {
	var body models.Student
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println("[DEBUG] Binding Error", err.Error())
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.InternalServerErrorCode,
				Status:  http.StatusInternalServerError,
				Message: "terjadi kesalahan server",
			},
		})
		return
	}

	log.Println("[DEBUG] body ", body)

	result, err := a.AuthRepo.GetUserData(ctx.Request.Context(), body.Name)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.InternalServerErrorCode,
				Status:  http.StatusInternalServerError,
				Message: "terjadi kesalahan server",
			},
		})
		return
	}

	hash := pkg.InitHashConfig()
	valid, err := hash.CompareHashAndPassword(result.Password, body.Password)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.InternalServerErrorCode,
				Status:  http.StatusInternalServerError,
				Message: "terjadi kesalahan server",
			},
		})
		return
	}
	if !valid {
		ctx.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.InvalidUsernamePasswordCode,
				Status:  http.StatusUnauthorized,
				Message: "invalid username/password",
			},
		})
		return
	}
	// kalau sudah berhasil login, maka berikan identitas (jwt)
	claims := pkg.NewClaims(result.Id, result.Role)
	token, err := claims.GenerateToken()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.InternalServerErrorCode,
				Status:  http.StatusInternalServerError,
				Message: "terjadi kesalahan server",
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Msg:  "Login Berhasil",
		Data: token,
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
