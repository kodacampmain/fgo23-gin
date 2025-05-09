package middlewares

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) CORSMiddleware(ctx *gin.Context) {
	// setup whitelist origin
	whitelistOrigin := []string{"http://localhost:5173"}
	origin := ctx.GetHeader("Origin")
	// log.Println("[DEBUG] Origin: ", origin)
	if slices.Contains(whitelistOrigin, origin) {
		// log.Println("[DEBUG] whitelisted")
		ctx.Header("Access-Control-Allow-Origin", origin)
	}
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, HEAD, PATCH, PUT, DELETE, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// handle preflight
	if ctx.Request.Method == http.MethodOptions {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}

	ctx.Next()
}
