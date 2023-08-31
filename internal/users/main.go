package main

import (
	"clearProject/internal/common/server"
	"clearProject/internal/users/controller"
	"github.com/gin-gonic/gin"
	"github.com/urban-lib/envs"
	"github.com/urban-lib/logging/v2"
	"log"
	"net/http"
)

func init() {
	envs.NewEnv("HOST", true, "0.0.0.0")
	envs.NewEnv("PORT", true, "8000")
	envs.NewEnv("LOG_LEVEL_CONSOLE", false, "debug")
	envs.NewEnv("LOG_FILE_ENABLE", false, "false")
	envs.NewEnv("LOG_LEVEL_FILE", false, "error")
	envs.NewEnv("LOG_FILE_PATH", false, "logs/users.log")
	envs.NewEnv("LOG_FILE_MAX_SIZE", false, "50")
	envs.NewEnv("LOG_FILE_MAX_BACKUPS", false, "2")
	envs.NewEnv("LOG_FILE_MAX_AGE", false, "1")
}

func HandlerRegister(serv controller.IServer, router *gin.RouterGroup) {
	router.Handle(http.MethodGet, "/", serv.GetAllUser)
	router.Handle(http.MethodPost, "/", serv.CreateUser)
	router.Handle(http.MethodGet, "/:userID", serv.GetUser)
	router.Handle(http.MethodPut, "/:userID", serv.UpdateUser)
	router.Handle(http.MethodDelete, "/:userID", serv.DeleteUser)
}

func main() {
	envs.CheckEnvironments()
	if _, _, err := logging.GetLogger(); err != nil {
		log.Fatal(err)
	}
	server.RunHTTPServer(func(router *gin.RouterGroup) {
		HandlerRegister(controller.HttpServer{}, router.Group("/users"))
	})

}
