package handler

import (
	user "backend_ajax-people"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) test(c *gin.Context) {
	var input user.UpdateUserInput
	var userList []user.User

	if err := c.BindJSON(&input); err != nil {
		fmt.Printf("Failed to selected a user: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, "Failed to selected users")
		return
	}

	userList, err := h.services.SelectedDataUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, userList)
}
