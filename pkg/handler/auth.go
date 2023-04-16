package handler

import (
	user "backend_ajax-people"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input user.User

	if err := c.BindJSON(&input); err != nil {
		fmt.Printf("Failed to create a user: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, "Failed to create a user")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		fmt.Printf("Failed data: %s\n", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Firstname, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
