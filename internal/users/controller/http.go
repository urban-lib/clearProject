package controller

import "github.com/gin-gonic/gin"

type IServer interface {
	GetUser(ctx *gin.Context)
	GetAllUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
}

type HttpServer struct {
}

func (this HttpServer) GetUser(ctx *gin.Context)    {}
func (this HttpServer) GetAllUser(ctx *gin.Context) {}
func (this HttpServer) CreateUser(ctx *gin.Context) {}
func (this HttpServer) DeleteUser(ctx *gin.Context) {}
func (this HttpServer) UpdateUser(ctx *gin.Context) {}
