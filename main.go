package main

import (
	"github.com/gin-gonic/gin"
	"github.com/moverq1337/redisWeather/config"
	"github.com/moverq1337/redisWeather/controllers"
)

func main() {
	config.LoadEnv()
	r := gin.Default()
	r.GET("/weather/:city", controllers.GetWeather)
	r.Run()
}
