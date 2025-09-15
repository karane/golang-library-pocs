package routes

import (
	"net/http"
	"strconv"

	"gin-poc/models"
	"gin-poc/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()
var users = []models.User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

func RegisterUserRoutes(rg *gin.RouterGroup) {
	rg.GET("/", getUsers)
	rg.GET("/:id", getUserByID)
	rg.POST("/", createUser)
}

func getUsers(c *gin.Context) {
	nameQuery := c.Query("name")
	if nameQuery != "" {
		filtered := []models.User{}
		for _, u := range users {
			if u.Name == nameQuery {
				filtered = append(filtered, u)
			}
		}
		c.JSON(http.StatusOK, filtered)
		return
	}
	c.JSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JSONError(c, http.StatusBadRequest, "invalid user id")
		return
	}
	for _, u := range users {
		if u.ID == id {
			c.JSON(http.StatusOK, u)
			return
		}
	}
	utils.JSONError(c, http.StatusNotFound, "user not found")
}

func createUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		utils.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Validate
	if err := validate.Struct(newUser); err != nil {
		utils.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	newUser.ID = len(users) + 1
	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}
