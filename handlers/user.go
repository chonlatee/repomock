package handlers

import (
	"net/http"

	"github.com/chonlatee/repomock/repository"
	"github.com/gin-gonic/gin"
)

type user struct {
	uRepo repository.Users
}

func (u *user) GetUserByID(c *gin.Context) {
	userID := c.Param("userID")
	res, err := u.uRepo.GetUserByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	ur := UserRes{}
	ur.UserID = res.UserID
	ur.FirstName = res.FirstName
	ur.LastName = res.LastName

	c.JSON(http.StatusOK, ur)
}

func NewUserAPI(ur repository.Users) *user {
	return &user{
		uRepo: ur,
	}
}
