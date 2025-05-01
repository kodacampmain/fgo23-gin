package middlewares

import (
	"fgo23-gin/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) AccessGateAdmin(ctx *gin.Context) {
	// 1. ambil payload/claims dari context gin
	claims, exist := ctx.Get("Payload")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Silahkan login terlebih dahulu",
		})
		return
	}
	// type assertion claims menjadi pkg.claims
	userClaims, ok := claims.(*pkg.Claims)
	// log.Println(userClaims)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Identitas login anda rusak, Silahkan login kembali",
		})
		return
	}
	// cek role yang ada di claims
	if userClaims.Role != "admin" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Anda tidak dapat mengakses sumber ini",
		})
		return
	}
	ctx.Next()
}
