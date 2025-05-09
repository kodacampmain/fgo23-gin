package handlers

import (
	"fgo23-gin/internal/models"
	"fgo23-gin/internal/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingHandler struct {
	pingRepo *repositories.PingRepository
}

// initialization
func NewPingHandler(pingRepo *repositories.PingRepository) *PingHandler {
	return &PingHandler{pingRepo: pingRepo}
}

// Handler

// PingGetStudentHandler
// @summary					Example of Ping then Get Students
// @produce					json
// @success					200 {object} models.Response
// @failure					500 {object} models.ErrorResponse
// @failure					404 {object} models.ErrorResponse
// @router					/ping [get]
func (p *PingHandler) GetStudents(ctx *gin.Context) {
	result, err := p.pingRepo.GetStudents(ctx.Request.Context())
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.InternalServerErrorCode,
				Status:  http.StatusInternalServerError,
				Message: "Terjadi Kesalahan Sistem",
			},
		})
		return
	}

	// Logika tambahan melihat isi dari result
	if len(result) == 0 {
		// error 404 not found
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.DataNotFoundCode,
				Status:  http.StatusNotFound,
				Message: "Data tidak ditemukan",
			},
		})
		return
	}

	// mengirimkan response suatu string berisikan pong
	// ctx.String(http.StatusOK, "pong")
	ctx.JSON(http.StatusOK, models.Response{
		Msg:  "pong",
		Data: result,
	})
}
