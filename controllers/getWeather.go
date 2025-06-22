package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/moverq1337/redisWeather/models"
	"net/http"
	"os"
)

func WeatherInfo(c *gin.Context) {
	var weather models.Weather
	city := c.Param("city")
	link := os.Getenv("url") + city + "&appid=" + os.Getenv("apiKey") + "&units=metric"
	fmt.Println(link)
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
