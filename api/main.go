package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    //get the app configuration
    config := GetConfig()



    router := gin.Default()

    router.Use(cors.New(CORSConfig()))



    router.Run(config.port)
}