package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"redisWeather/config"
)

type Weather struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

const apiKey = "19051382c62dbbf9bcc65e3661f26287"
const url = "https://api.openweathermap.org/data/2.5/weather?q="

func getWeather(c *gin.Context) {
	var weather Weather
	city := c.Param("city")
	link := os.Getenv("weatherUrl") + city + "&appid=" + os.Getenv("apiKey") + "&units=metric"
	r, err := http.Get(link)
	if err != nil {
		fmt.Println("Ошибка доступа к сервису")
		return
	}
	err = json.NewDecoder(r.Body).Decode(&weather)
	if err != nil {
		fmt.Println(err, "Хуево что-то")
		return
	}
	c.JSON(200, gin.H{
		"message": weather.Main.Temp,
	})
}

func main() {
	config.LoadEnv()
	r := gin.Default()
	r.GET("/weather/:city", getWeather)
	r.Run()
}
