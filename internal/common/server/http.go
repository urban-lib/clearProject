package server

import (
	"github.com/gin-gonic/gin"
	"github.com/urban-lib/envs"
	"github.com/urban-lib/logging/v2"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type RouterFunc func(router *gin.RouterGroup)

func RunHTTPServer(register RouterFunc) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	listenAddr := envs.Get("HOST") + ":" + envs.Get("PORT")
	go func() {
		if err := RunHttpServerOnAddr(listenAddr, register); err != nil && err != http.ErrServerClosed {
			logging.Fatalf("listen: %s\n", err)
		}
	}()

	<-quit
	logging.Info("Shutdown Server ...")
}

func RunHttpServerOnAddr(addr string, register RouterFunc) error {

	router := gin.New()
	setMiddlewares(router)

	apiGroup := router.Group("/api")

	register(apiGroup)

	httpServer := http.Server{
		Addr:           addr,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
		Handler:        router,
	}
	logging.Infof("App running at: http://%s", addr)
	return httpServer.ListenAndServe()
}

func setMiddlewares(router *gin.Engine) {
	router.Use(
		gin.Recovery(),
	)
}
