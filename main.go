package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

type Response struct {
	Msg  string `json:"message"`
	Data any    `json:"data"`
}

func main() {
	// gin engine initialization
	router := gin.Default()

	// create database connection string
	dbEnv := []any{}
	dbEnv = append(dbEnv, os.Getenv("DBUSER"))
	dbEnv = append(dbEnv, os.Getenv("DBPASS"))
	dbEnv = append(dbEnv, os.Getenv("DBHOST"))
	dbEnv = append(dbEnv, os.Getenv("DBPORT"))
	dbEnv = append(dbEnv, os.Getenv("DBNAME"))
	// setup database connection
	dbString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbEnv...)
	dbClient, err := pgxpool.New(context.Background(), dbString)
	if err != nil {
		log.Printf("[ERROR] Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	// graceful shutdown
	// server jalan di goroutine
	// goroutine main handling shutdown
	defer func() {
		log.Println("Closing DB...")
		dbClient.Close()
	}()

	// endpoint & resource
	// /ping => protocol://hostname/ping => http://localhost:port/ping
	router.GET("/ping", func(ctx *gin.Context) {
		type Student struct {
			Id   int    `db:"id"`
			Name string `db:"name"`
		}
		query := "SELECT id, name FROM students"
		rows, err := dbClient.Query(ctx.Request.Context(), query)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "Terjadi kesalahan sistem",
			})
			return
		}
		defer rows.Close()
		var result []Student
		for rows.Next() {
			var student Student
			if err := rows.Scan(&student.Id, &student.Name); err != nil {
				log.Println(err.Error())
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "Terjadi kesalahan sistem",
				})
				return
			}
			// log.Printf("[DEBUG] data student per baris: %v\n", student)
			result = append(result, student)
		}

		// mengirimkan response suatu string berisikan pong
		// ctx.String(http.StatusOK, "pong")
		ctx.JSON(http.StatusOK, Response{
			Msg:  "pong",
			Data: result,
		})
	})
	// field datatype structTag
	type userStruct struct {
		Id         int       `json:"id" form:"identity"`
		Name       string    `json:"name" form:"nama"`
		Created_at time.Time `json:"-" form:"-" db:"created_at"`
	}
	users := []userStruct{
		{Id: 1, Name: "Nana"},
		{Id: 2, Name: "Dudu"},
		{Id: 3, Name: "Nana"},
		{Id: 4, Name: "Dudul"},
	}
	type Employee struct {
		Id     int    `db:"id" json:"id,omitempty"`
		Name   string `db:"name" json:"name"`
		Salary int    `db:"salary" json:"salary"`
		City   string `db:"city" json:"city,omitempty"`
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
		query := "SELECT id,name,salary FROM employee WHERE id=$1 AND name=$2"
		values := []any{idInt, name}
		var result Employee
		if err := dbClient.QueryRow(ctx.Request.Context(), query, values...).Scan(&result.Id, &result.Name, &result.Salary); err != nil {
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
		// newEmployee := &Employee{}
		var newEmployee Employee
		if err := ctx.ShouldBind(&newEmployee); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "Terjadi kesalahan sistem",
			})
			return
		}

		query := "INSERT INTO employee (name, salary, city) VALUES ($1, $2, $3)"
		values := []any{newEmployee.Name, newEmployee.Salary, newEmployee.City}
		cmd, err := dbClient.Exec(ctx.Request.Context(), query, values...)
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
	})

	// jalankan service
	router.Run("127.0.0.1:8080")
	// router.Run(":8080")
}
