package controllers

import "github.com/gin-gonic/gin"

type Context interface {
	Param(key string) string
	Query(key string) string
	JSON(code int, obj interface{})
	Bind(interface{}) error
	Status(int)
	HTML(code int, name string, obj interface{})
}

var c gin.Context
