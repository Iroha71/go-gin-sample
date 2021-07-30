package controllers

import (
	"net/http"
	// "strconv"

	"gomodtest/repository"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (pc UserController) Index(c *gin.Context) {
	var u repository.UserRepository
	p, err := u.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}
