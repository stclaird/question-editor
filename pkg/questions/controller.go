// Package questions
// Provide api functionailty for question objects
package questions

import (
	"github.com/gin-gonic/gin"
)


func RegisterRoutes(router *gin.Engine) {
    questionRoutes := router.Group("/questions")
    questionRoutes.GET("/", AddQuestion())

}
