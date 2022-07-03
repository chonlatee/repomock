//go:build wireinject
// +build wireinject

package main

import (
	"github.com/chonlatee/repomock/dbs"
	"github.com/chonlatee/repomock/repository"
	"github.com/chonlatee/repomock/routes"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type app struct {
	route *gin.Engine
}

func NewAPP(r *gin.Engine) *app {
	return &app{
		route: r,
	}
}

func initializeDI() (*app, error) {

	wire.Build(
		dbs.NewUserDB,
		wire.Bind(new(repository.Users), new(*dbs.User)),
		routes.Routes,
		NewAPP,
	)

	return &app{}, nil
}
