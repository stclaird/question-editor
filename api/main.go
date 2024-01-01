package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stclaird/question-editor/pkg/questions"
)

func main() {
    //get the app configuration
    config := GetConfig()

    router := gin.Default()

    router.Use(cors.New(CORSConfig()))
    questions.RegisterRoutes(router)
    router.Run(config.port)
}