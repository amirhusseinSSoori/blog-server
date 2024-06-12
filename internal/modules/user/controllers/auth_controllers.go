package controllers

import (
	"blog/internal/modules/user/request/auth"
	UserService "blog/internal/modules/user/services"
	"blog/pkg/html"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService UserService.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: UserService.New(),
	}

}

func (controller *Controller) Register(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"message": "Register form"})
	html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "Register",
	})
}

func (controller *Controller) HandleRegister(c *gin.Context) {
	// validate the request
	var request auth.RegisterRequest
	// This will infer what binder to use depending on th content-type header.
	if err := c.ShouldBind(&request); err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	// create the user
	user, err := controller.userService.Create(request)

	// check if there is any errors on the user creation
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	// after creating th user > redirect  to home page
	log.Printf("the user Creat successfully with a name %s \n", user.Name)

	c.Redirect(http.StatusFound, "/")
}
