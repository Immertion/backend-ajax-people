package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func checkAdmin(h *Handler, c *gin.Context) bool {
	token, err := c.Cookie("jwtToken")
	if err != nil {
		return false
	}

	_, isAdmin, err := h.services.ParseToken(token)
	if err != nil {
		return false
	}

	return isAdmin
}

func getJWT(h *Handler, c *gin.Context) (int, bool, error) {
	token, err := c.Cookie("jwtToken")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return 0, false, err
	}

	userId, isAdmin, err := h.services.ParseToken(token)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return 0, false, err
	}

	return userId, isAdmin, nil
}

func (h *Handler) userIdentify(c *gin.Context) {
	userId, isAdmin, err := getJWT(h, c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	getId, err := getUserId(c)
	if isAdmin == false && userId != getId {
		if getId == 0 {
			return
		}
		newErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

}

func (h *Handler) userIdentifyAdmin(c *gin.Context) {
	if checkAdmin(h, c) == false {
		newErrorResponse(c, http.StatusForbidden, "Forbidden")
		return
	}
}

func getUserId(c *gin.Context) (int, error) {

	str := c.Param("id")
	fmt.Println(str)

	if str == "" {
		return 0, nil
	}
	id, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return id, nil
}
