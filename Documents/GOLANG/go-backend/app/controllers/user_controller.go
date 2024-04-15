package controllers

import (

	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)


type UserController struct {
	repo repositories.UserRepository
}

func NewUserController(repo repositories.UserRepository) *UserController {
	return &UserController{
		repo: repo,
	}
}

// THIS IS WHERE THE CRUD REQUESTS BELONG

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.repo.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := c.repo.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) getUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, u := range users {
		if strconv.Itoa(u.ID) == id {
			ctx.IndentedJSON(http.StatusOK, u)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func (c *UserController) updateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	for i, u := range users {
		if strconv.Itoa(u.ID) == id {
			var updatedUser User

			if err := ctx.BindJSON(&updatedUser); err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			updatedUser.ID = u.ID
			users[i] = updatedUser

			ctx.JSON(http.StatusOK, updatedUser)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func (c *UserController) deleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, u := range users {
		if strconv.Itoa(u.ID) == id {
			users = append(users[:i], users[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}



// var userFields = getUserFields()

// func getUserFields() []string {
//     var field []string
//     v := reflect.ValueOf(User{})
//     for i := 0; i < v.Type().NumField(); i++ {
//         field = append(field, v.Type().Field(i).Tag.Get("json"))
//     }
//     return field
// }
// func stringInSlice(strSlice []string, s string) bool {
//     for _, v := range strSlice {
//         if v == s {
//             return true
//         }
//     }
//     return false
// }

// func validateAndReturnSort(sortBy string) (string, error) {
// 	splits := strings.Split(sortBy, ".")

// 	if len(splits) != 2 {
// 		return "", errors.New("malformed sortBy query parameter, should be field.orderdirection")
	
// 	}

// 	field, order := splits[0], splits[1]

// 	if order != "desc" && order != "asc" {
// 		return "", errors.New("malformed orderdirection in sortBy query parameter, should be asc or desc")
// 	}

// 	if !stringInSlice(userFields, field) {
//         return "", errors.New("unknown field in sortBy query parameter")
//     }
//     return fmt.Sprintf("%s %s", field, strings.ToUpper(order)), nil

// }