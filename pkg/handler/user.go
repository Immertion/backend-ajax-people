package handler

import (
	user "backend_ajax-people"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (h *Handler) createUser(c *gin.Context) {
	if checkAdmin(h, c) == false {
		c.JSON(http.StatusForbidden, "Forbidden")
		return
	}
	var input user.User

	if err := c.BindJSON(&input); err != nil {
		fmt.Printf("Failed to create a user: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, "Failed to create a user")
		return
	}

	id, err := h.services.UserAction.CreateUser(input)
	if err != nil {
		fmt.Printf("Failed data: %s\n", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllUsers(c *gin.Context) {
	userList, err := h.services.GetAllUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, userList)
}

func (h *Handler) getUserById(c *gin.Context) {
	userId, isAdmin, err := getJWT(h, c)
	getId, err := getUserId(c)
	if isAdmin == false && userId != getId {
		c.JSON(http.StatusForbidden, "Forbidden")
		return
	}

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.services.GetUserById(getId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) updateUser(c *gin.Context) {
	userId, isAdmin, err := getJWT(h, c)
	getId, err := getUserId(c)
	if isAdmin == false && userId != getId {
		c.JSON(http.StatusForbidden, "Forbidden")
		return
	}

	var input user.UpdateUserInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Printf("Failed to update a user: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, "Failed to update a user")
		return
	}

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.UpdateUser(getId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Ok")
}

func (h *Handler) deleteUser(c *gin.Context) {
	if checkAdmin(h, c) == false {
		c.JSON(http.StatusForbidden, "Forbidden")
		return
	}
	userId, err := getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.DeleteUser(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Ok")
}

type Message struct {
	Content string `json:"content" binding:"required"`
}

func (h *Handler) checkActivationUser(c *gin.Context) {
	token, err := c.Cookie("jwtToken")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userId, _, err := h.services.ParseToken(token)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var code Message

	if err := c.BindJSON(&code); err != nil {
		c.JSON(http.StatusBadRequest, "Fail")
		return
	}

	verified, err := h.services.CheckCodeActivation(userId, code.Content)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, verified)
}
