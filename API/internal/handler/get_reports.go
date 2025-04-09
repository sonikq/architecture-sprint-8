package handler

import (
	"backend-api/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetReports(ctx *gin.Context) {
	reports := make([]models.Report, 0, 10)
	for i := 0; i <= 10; i++ {
		reports = append(reports, models.Report{
			ID:          i,
			Title:       fmt.Sprintf("Название отчета: %d", i),
			Description: fmt.Sprintf("Описание отчета: %d", i),
			CreatedAt:   time.Now(),
		})
	}
	ctx.JSON(http.StatusOK, reports)
}
