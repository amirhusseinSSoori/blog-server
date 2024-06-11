package provider

import (
	articlesoutes "blog/internal/modules/articles/routes"
	homeRoutes "blog/internal/modules/home/routes"
	userRoutes "blog/internal/modules/user/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	articlesoutes.Routes(router)
	userRoutes.Routes(router)

}
