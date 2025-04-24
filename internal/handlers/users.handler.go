package handlers

import (
	"fgo23-gin/internal/models"
	"fgo23-gin/internal/repositories"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// field datatype structTag
type userStruct struct {
	Id         int       `json:"id" form:"identity"`
	Name       string    `json:"name" form:"nama"`
	Created_at time.Time `json:"-" form:"-" db:"created_at"`
}

type UserHandler struct{}

// Initialization
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetEmployeeById(ctx *gin.Context) {
	idStr, ok := ctx.Params.Get("id")
	// idStr := ctx.Param("id")
	// ambil params id
	if !ok { // error handling pengambilan params id
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Param id is needed",
		})
		return
	}
	idInt, err := strconv.Atoi(idStr)
	// conversion dari string menjadi int
	if err != nil { // error handling conversion
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Terjadi kesalahan server",
		})
		return
	}
	// var user []userStruct
	// for _, u := range users { // pencarian data user berdasarkan id
	// 	if u.Id == idInt {
	// 		user = append(user, u)
	// 		break
	// 	}
	// }

	// if len(user) == 0 { // error handling jika user tidak ditemukan
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"msg": "User tidak ditemukan",
	// 	})
	// 	return
	// }
	name := ctx.Query("name")
	result, err := repositories.UserRepo.FindEmployeeById(ctx.Request.Context(), idInt, name)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Terjadi kesalahan sistem",
		})
		return
	}

	// pengiriman response jika user ditemukan
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Success",
		"user": result,
	})
}

func (u *UserHandler) GetUsers(ctx *gin.Context) {
	users := []userStruct{
		{Id: 1, Name: "Nana"},
		{Id: 2, Name: "Dudu"},
		{Id: 3, Name: "Nana"},
		{Id: 4, Name: "Dudul"},
	}
	// ambil query
	nameQ := ctx.Query("name")
	if nameQ == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "Success",
			"data": users,
		})
		return
	}
	result := []userStruct{}
	for _, user := range users {
		condition := strings.EqualFold(user.Name, nameQ)
		// condition := user.Name == nameQ
		if condition {
			result = append(result, user)
		}
	}
	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "User tidak ditemukan",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "User ditemukan",
		"data": result,
	})
}

func (u *UserHandler) AddEmployee(ctx *gin.Context) {
	// newEmployee := &Employee{}
	var newEmployee models.Employee
	if err := ctx.ShouldBind(&newEmployee); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Terjadi kesalahan sistem",
		})
		return
	}

	cmd, err := repositories.UserRepo.CreateNewEmployee(ctx.Request.Context(), newEmployee)

	if err != nil { // error handling conversion
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Terjadi kesalahan server",
		})
		return
	}
	if cmd.RowsAffected() == 0 {
		log.Println("Query Gagal, Tidak merubah data di DB")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
	// newUsers := append(users, *newUser)
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"msg":  "Success",
	// 	"data": newUsers,
	// })
}
