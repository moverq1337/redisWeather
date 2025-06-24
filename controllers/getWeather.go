package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/moverq1337/redisWeather/config"
	"github.com/moverq1337/redisWeather/models"
	"net/http"
	"os"
	"time"
)

func GetWeather(c *gin.Context) {
	var weather models.Weather
	city := c.Param("city")
	val, err := config.Client.Get(config.Сtx, city).Result()
	if err != nil {
		fmt.Println("Запись не найдена", err)

	} else {
		fmt.Println("Запись нашлась, сейчас отдам")
		fmt.Println(val)
		c.JSON(200, gin.H{
			"message": val,
			"Success": "from redis done",
		})
		return
	}
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
	err = config.Client.Set(config.Сtx, city, weather.Main.Temp, 100*time.Minute).Err()
	if err != nil {
		c.JSON(400, gin.H{
			"message": weather.Main.Temp,
			"error":   "redis error"})
		return
	}
	c.JSON(200, gin.H{
		"message": weather.Main.Temp,
	})
}
