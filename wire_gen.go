// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/chonlatee/repomock/dbs"
	"github.com/chonlatee/repomock/routes"
	"github.com/gin-gonic/gin"
)

// Injectors from wiredi.go:

func initializeDI() (*app, error) {
	user := dbs.NewUserDB()
	engine := routes.Routes(user)
	mainApp := NewAPP(engine)
	return mainApp, nil
}

// wiredi.go:

type app struct {
	route *gin.Engine
}

func NewAPP(r *gin.Engine) *app {
	return &app{
		route: r,
	}
}
