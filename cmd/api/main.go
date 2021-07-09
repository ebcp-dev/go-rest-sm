package main

import (
	_ "github.com/ebcp-dev/go-rest-sm/docs"
	"github.com/ebcp-dev/go-rest-sm/internal/api"
)

// @Golang API REST
// @version 1.0
// @description API REST in Golang with Gin Framework

// @contact.name Earl Perez
// @contact.url http://earlperez.me
// @contact.email ebcp.dev@gmail.com

// @license.name MIT
// @license.url https://github.com/ebcp-dev/go-rest-sm/blob/master/LICENSE

// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	api.Run("")
}
