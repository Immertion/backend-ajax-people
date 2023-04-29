package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getPostById(c *gin.Context) {
	var postId int

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	post, err := h.services.GetPostById(postId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) getAllPosts(c *gin.Context) {
	postsList, err := h.services.GetAllPosts()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, postsList)
}

type PostInput struct {
	Text string `json:"text" binding:"required"`
}

func (h *Handler) createPost(c *gin.Context) {
	var input PostInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Failed to create a post")
		return
	}

	userId, _, err := getJWT(h, c)
	if err != nil {
		return
	}

	id, err := h.services.CreatePost(input.Text, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed data: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type postUpdateInput struct {
	IsModerated bool `json:"isModerated" binding:"required"`
}

func (h *Handler) updatePost(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input postUpdateInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Failed to modify post")
		return
	}

	if err := h.services.UpdatePost(postId, input.IsModerated); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) deletePost(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.DeletePost(postId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "OK")
}
