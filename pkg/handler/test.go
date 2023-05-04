package handler

import (
	user "backend_ajax-people"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Json struct {
	Interests []string `json:"interests"`
}

func (h *Handler) test(c *gin.Context) {
	//idUser, _, err := getJWT(h, c)
	var input user.User

	if err := c.BindJSON(&input); err != nil {
		fmt.Printf("Failed to selected a user: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, "Failed to selected users")
		return
	}

	//err = h.services.AddInterests(input.Interests, idUser)
	//if err != nil {
	//	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	//}

	c.JSON(http.StatusOK, input)
}
