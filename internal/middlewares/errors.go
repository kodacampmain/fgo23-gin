package middlewares

import (
	"fgo23-gin/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) Error(ctx *gin.Context) {
	ctx.Next()

	if len(ctx.Errors) > 0 {
		// for _, err := range ctx.Errors {}
		err := ctx.Errors.Last().Err
		log.Printf("\nInternal Error: %v\n", err)
		if !ctx.Writer.Written() {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: &models.ErrorResponseDetail{
					Code:    "INTERNAL_ERROR",
					Message: "Internal Server Error",
					Status:  http.StatusInternalServerError,
				},
			})
		}
	}
}
