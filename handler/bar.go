package handler

import (
	"github.com/evsamsonov/trading-history-server/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func CreateBars(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var json model.Bars
		err := context.BindJSON(&json)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var symbol model.Symbol
		var resolution model.Resolution

		db.FirstOrCreate(&symbol, model.Symbol{Code: json.Symbol})
		db.FirstOrCreate(&resolution, model.Resolution{Code: json.Resolution})

		for _, barItem := range json.Items {
			var bar model.Bar
			bar.BarItem = barItem
			bar.SymbolID = symbol.ID
			bar.ResolutionID = resolution.ID

			db.FirstOrCreate(
				&bar,
				map[string]interface{}{"symbol_id": symbol.ID, "resolution_id": resolution.ID, "time": barItem.Time})
		}

		context.JSON(http.StatusOK, gin.H{})
	}
}
