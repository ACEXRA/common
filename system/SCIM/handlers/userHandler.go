package handlers

import (
	"net/http"

	"github.com/acexra/scim/config"
	"github.com/acexra/scim/model"
	"github.com/gin-gonic/gin"
)

var user model.User

func Create(c *gin.Context) {

	if err:= c.ShouldBindBodyWithJSON(&user); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := user.Create(config.DB); 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user": createdUser,
	})
}

func GetById(c *gin.Context){
	
}