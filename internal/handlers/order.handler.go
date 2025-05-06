package handlers

import (
	"fgo23-gin/internal/models"
	"fgo23-gin/internal/repositories"
	"fgo23-gin/pkg"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderRepo *repositories.OrderRepository
}

func NewOrderHandler(orderRepo *repositories.OrderRepository) *OrderHandler {
	return &OrderHandler{orderRepo: orderRepo}
}

func (o *OrderHandler) CreateTransaction(ctx *gin.Context) {
	// olah body
	var body models.Transaction
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Terjadi Kesalahan Server",
		})
		return
	}
	// ambil studentId
	payload, _ := ctx.Get("Payload")
	userClaims := payload.(*pkg.Claims)

	_, err := o.orderRepo.CreateTransaction(ctx.Request.Context(), userClaims.Id, body)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Terjadi Kesalahan Server",
		})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
