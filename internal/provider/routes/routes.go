package provider

import (
	articlesoutes "blog/internal/modules/articles/routes"
	homeRoutes "blog/internal/modules/home/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	articlesoutes.Routes(router)
}
