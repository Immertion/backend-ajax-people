package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) test(c *gin.Context) {

	token, err := c.Cookie("jwtToken")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userId, isAdmin, err := h.services.ParseToken(token)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"admin":  isAdmin,
		"userId": userId,
	})

}
