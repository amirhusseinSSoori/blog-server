package html

import (
	"github.com/gin-gonic/gin"
)

func LoadHtml(router *gin.Engine) {
	router.LoadHTMLGlob("internal/**/**/**/*html")
}
