package questions

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stclaird/question-editor/pkg/common/models"
)

func AddQuestion(ctx *gin.Context) {

    var question models.Question
    if err := ctx.ShouldBindJSON(&question); err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }

    questionJsonBytes, _ := json.Marshal(question)

    questionName := generateQuestionFileName(question)

    os.WriteFile(questionName, questionJsonBytes, os.ModePerm)

    ctx.JSON(http.StatusOK, question)
}

