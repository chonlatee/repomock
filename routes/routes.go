package routes

import (
	"github.com/chonlatee/repomock/handlers"
	"github.com/chonlatee/repomock/repository"
	"github.com/gin-gonic/gin"
)

func Routes(dbs repository.Users) *gin.Engine {
	r := gin.Default()
	u := handlers.NewUserAPI(dbs)
	r.GET("/:userID", u.GetUserByID)

	return r
}
