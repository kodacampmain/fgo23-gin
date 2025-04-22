package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin engine initialization
	router := gin.Default()
	// endpoint & resource
	// /ping => protocol://hostname/ping => http://localhost:port/ping
	router.GET("/ping", func(ctx *gin.Context) {
		// mengirimkan response suatu string berisikan pong
		// ctx.String(http.StatusOK, "pong")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// field datatype structTag
	type userStruct struct {
		Id   int    `json:"id" form:"identity"`
		Name string `json:"name" form:"nama"`
	}
	users := []userStruct{
		{Id: 1, Name: "Nana"},
		{Id: 2, Name: "Dudu"},
		{Id: 3, Name: "Nana"},
		{Id: 4, Name: "Dudul"},
	}
	// definisikan rute dengan params id
	router.GET("/users/:id", func(ctx *gin.Context) {
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
		var user []userStruct
		for _, u := range users { // pencarian data user berdasarkan id
			if u.Id == idInt {
				user = append(user, u)
				break
			}
		}

		if len(user) == 0 { // error handling jika user tidak ditemukan
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "User tidak ditemukan",
			})
			return
		}
		// pengiriman response jika user ditemukan
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "Success",
			"user": user[0],
		})
	})

	// /users?name=nana
	router.GET("/users", func(ctx *gin.Context) {
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
	})

	// /users
	router.POST("/users", func(ctx *gin.Context) {
		newUser := &userStruct{}
		if err := ctx.ShouldBind(newUser); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "Terjadi kesalahan sistem",
			})
			return
		}
		newUsers := append(users, *newUser)
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "Success",
			"data": newUsers,
		})
	})

	// jalankan service
	router.Run("127.0.0.1:8080")
	// router.Run(":8080")
}
