package handlers

import (
	"fgo23-gin/internal/models"
	"fgo23-gin/internal/repositories"
	"fgo23-gin/pkg"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	fp "path/filepath"
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

type UserHandler struct {
	userRepo *repositories.UserRepository
}

// Initialization
func NewUserHandler(userRepo *repositories.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
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
	result, err := u.userRepo.FindEmployeeById(ctx.Request.Context(), idInt, name)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Terjadi kesalahan sistem",
		})
		return
	}

	if result == (models.Employee{}) {
		// error 404 not found
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data tidak ditemukan",
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
		if status, msg := errorMsgBuilder(err); status != 0 {
			ctx.JSON(status, gin.H{
				"msg": msg,
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Terjadi kesalahan sistem",
		})
		return
	}

	// validasi dengan regex
	// regexp.Match()

	cmd, err := u.userRepo.CreateNewEmployee(ctx.Request.Context(), newEmployee)

	if err != nil { // error handling conversion
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Terjadi kesalahan server",
		})
		return
	}
	if cmd.RowsAffected() == 0 {
		log.Println("Query Gagal, Tidak merubah data di DB")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Data yang diberikan salah",
		})
		return
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

func (u *UserHandler) EditStudents(ctx *gin.Context) {
	// Handling File
	// file, err := ctx.FormFile("img")
	// if err != nil {
	// 	log.Println(err.Error())
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "Terjadi kesalahan server",
	// 	})
	// 	return
	// }
	var formBody models.StudentForm
	if err := ctx.ShouldBind(&formBody); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Terjadi kesalahan server",
		})
		return
	}
	file := formBody.Image
	var filename, filepath string
	if file != nil {
		var err error
		filename, filepath, err = fileHandling(ctx, file)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Terjadi kesalahan upload",
			})
		}
	}
	log.Println("[DEBUG] filename", filename)
	log.Println("[DEBUG] body", formBody)

	// Handling Non-File
	// ctx.PostForm
	// formBody
	// Send Response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update Success",
		"data": gin.H{
			"url": filepath,
		},
	})
}

func fileHandling(ctx *gin.Context, file *multipart.FileHeader) (filename, filepath string, err error) {
	claims, _ := ctx.Get("Payload")
	userClaims := claims.(*pkg.Claims)
	ext := fp.Ext(file.Filename)
	filename = fmt.Sprintf("%d_%d_students_image%s", time.Now().UnixNano(), userClaims.Id, ext)
	filepath = fp.Join("public", "img", filename)
	if err = ctx.SaveUploadedFile(file, filepath); err != nil {
		return "", "", err
	}
	return filename, filepath, nil
}

func errorMsgBuilder(err error) (status int, msg string) {
	if strings.Contains(err.Error(), "Field validation") {
		if strings.Contains(err.Error(), "gt") {
			return http.StatusBadRequest, "Salary harus lebih besar dari 10"
		}
		return http.StatusBadRequest, "Body harus berisikan name, salary (lebih besar dari 10) dan city"
	}
	return 0, ""
}
