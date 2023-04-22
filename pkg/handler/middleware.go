package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func getUserId(c *gin.Context) (int, error) {

	str := c.Param("id")
	fmt.Println(str)

	id, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return id, nil
}
