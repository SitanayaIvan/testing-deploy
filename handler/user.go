package handler

import (
	"net/http"
	"strconv"

	"github.com/SitanayaIvan/testing-deploy/user"
	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	service user.Service
}

func NewHandlerUser(service user.Service) *HandlerUser {
	return &HandlerUser{service}
}

func (h HandlerUser) GetAllUser(c *gin.Context) {
	users, err := h.service.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   users,
	})
}

func (h HandlerUser) GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

func (h HandlerUser) CreateUser(c *gin.Context) {
	var user user.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "wrong input data!",
		})
		return
	}

	user.FirstName = "test"
	err = h.service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func (h HandlerUser) UpdateUser(c *gin.Context) {
	var user user.User
	err := c.ShouldBindJSON(&user)

	id, _ := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "wrong input data!",
		})
		return
	}

	err = h.service.UpdateUser(user, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})

}

func (h HandlerUser) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
