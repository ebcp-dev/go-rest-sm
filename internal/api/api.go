package api

import (
	"fmt"

	"github.com/ebcp-dev/go-rest-sm/internal/api/router"
	"github.com/ebcp-dev/go-rest-sm/internal/pkg/config"
	"github.com/ebcp-dev/go-rest-sm/internal/pkg/db"
	"github.com/gin-gonic/gin"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
	db.SetupDB()
	gin.SetMode(config.GetConfig().Server.Mode)
}

func Run(configPath string) {
	if configPath == "" {
		configPath = "data/config.yml"
	}
	setConfiguration(configPath)
	conf := config.GetConfig()
	web := router.Setup()
	fmt.Println("API Running on port " + conf.Server.Port)
	fmt.Println("<==================>")
	_ = web.Run(":" + conf.Server.Port)
}
