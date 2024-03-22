package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func NewRouter(db *gorm.DB, app *gin.Engine) {
	NewOrderRouter(db, app)
}
