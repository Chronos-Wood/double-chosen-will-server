package router

import(
	"github.com/gin-gonic/gin"
	"github.com/double-chosen-will-server/web/middleware"
	"github.com/double-chosen-will-server/config"
	"github.com/gomodule/redigo/redis"
	"fmt"
)


func Engine(conf *config.Config) *gin.Engine {
	defaultEngine := gin.Default()
	//TODO
	defaultEngine.Use(middleware.SessionValidation(&redis.Pool{Dial: func() (redis.Conn, error) {
		serverAddr := fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port)
		return redis.Dial("tcp", serverAddr)
	}}))
	defaultEngine.GET("/", func(context *gin.Context) {
		context.JSON(200, struct {
			Code int `json:"code"`
			Message string `json:"message"`
			Data interface{} `json:"data"`
		}{0, "success",""})
	})
	return defaultEngine
}

func init(){
}