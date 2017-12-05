package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
}

func Test(c *gin.Context) {
	user := c.Params.ByName("name")
	name := new(User)
	c.ShouldBindJSON(&name)

	fmt.Println("%v", name)

	c.JSON(200, gin.H{"user": user, "value": user})
}
