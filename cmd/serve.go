package cmd

import (
	"blog/pkg/config"
	"fmt"
	"net/http"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve App on the dev server",
	Long:  "Application will be served on host and port defind in config.yml file",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	config.Set()
	configs := config.Get()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"app name": viper.Get("App.Name"),
		})

	})
	r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
}
