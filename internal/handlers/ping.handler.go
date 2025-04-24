package handlers

import (
	"fgo23-gin/internal/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Msg  string `json:"message"`
	Data any    `json:"data"`
}

type PingHandler struct{}

// initialization
func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

// Handler
func (p *PingHandler) GetStudents(ctx *gin.Context) {
	result, err := repositories.PingRepo.GetStudents(ctx.Request.Context())
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Terjadi kesalahan sistem",
		})
		return
	}

	// mengirimkan response suatu string berisikan pong
	// ctx.String(http.StatusOK, "pong")
	ctx.JSON(http.StatusOK, Response{
		Msg:  "pong",
		Data: result,
	})
}
