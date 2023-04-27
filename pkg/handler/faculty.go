package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getAllFaculties(c *gin.Context) {
	facultiesList, err := h.services.Faculty.GetAll()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, facultiesList)
}
