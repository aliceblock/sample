package main

import (
	"flag"
	"os"

	"github.com/aliceblock/sample/cmd/server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	dataMode := flag.String("mode", "debug", "Data mode {debug|release|dev|prod}")
	flag.Parse()

	switch *dataMode {
	case "dev":
		*dataMode = "debug"
	case "prod":
		*dataMode = "release"
	}

	os.Setenv("DATAMODE", *dataMode)
	if *dataMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	routes.APIRoutes(r)

	r.Run()
}
