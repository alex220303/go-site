package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    r := gin.Default()

    // Настройка CORS
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"*"} // Разрешаем запросы с любого источника
    config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
    config.AllowHeaders = []string{"*"}
    config.AllowCredentials = true

    r.Use(cors.New(config))

    r.GET("/ping", pong)
    r.Run(":9090")
}

func pong(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "pon",
    })
}
