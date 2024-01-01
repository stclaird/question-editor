package questions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stclaird/question-editor/pkg/common/models"
)

func AddQuestion(ctx *gin.Context) {

    var question models.Question
    if err := ctx.ShouldBindJSON(&question); err != nil {
        ctx.Error(err)
        ctx.AbortWithStatus(http.StatusBadRequest)
        return
    }
    ctx.JSON(http.StatusOK, question)
}

func CheckJson()