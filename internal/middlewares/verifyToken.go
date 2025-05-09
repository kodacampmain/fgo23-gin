package middlewares

import (
	"fgo23-gin/pkg"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middleware struct{}

func InitMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) VerifyToken(ctx *gin.Context) {
	// 1. ambil token dari header
	bearerToken := ctx.GetHeader("Authorization")
	if bearerToken == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Silahkan login terlebih dahulu",
		})
		return
	}
	// verifikasi bearer token
	if !strings.Contains(bearerToken, "Bearer") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Silahkan login terlebih dahulu",
		})
		return
	}
	// 2. pisahkan token dari bearer
	token := strings.Split(bearerToken, " ")[1]
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Silahkan login terlebih dahulu",
		})
		return
	}
	// 3. verifikasi token
	claims := &pkg.Claims{}
	if err := claims.VerifyToken(token); err != nil {
		log.Println(err.Error())
		// if err.Error() == "expired token" || err.Error() == "token has invalid claims: token is expired" {
		if strings.Contains(err.Error(), "expired") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Sesi anda berakhir, Silahkan login kembali",
			})
			return
		}
		if strings.Contains(err.Error(), "malformed") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Identitas login anda rusak, Silahkan login kembali",
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Terjadi kesalahan server",
		})
		return
	}
	// kita masukkan claims/payload ke dalam gin context
	ctx.Set("Payload", claims)
	ctx.Next()
}
