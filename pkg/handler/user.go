package handler

import (
	user "backend_ajax-people"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createUser(c *gin.Context) {
	var reqUser user.User

	if err := c.BindJSON(&reqUser); err != nil {
		fmt.Printf("Failed to create a user: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, "Failed to create a user")
		return
	}

	c.JSON(http.StatusOK, reqUser)
}

func (h *Handler) getAllUsers(c *gin.Context) {

}

func (h *Handler) getUserById(c *gin.Context) {

}

func (h *Handler) updateUser(c *gin.Context) {

}

func (h *Handler) deleteUser(c *gin.Context) {

}
