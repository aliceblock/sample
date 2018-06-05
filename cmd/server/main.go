package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aliceblock/sample/cmd/server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	dataMode := flag.String("mode", "debug", "Data mode {debug|release}")
	flag.Parse()

	os.Setenv("DATAMODE", *dataMode)
	if *dataMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)

	routes.APIRoutes(r)

	r.Run()
}

// loadTemplate loads templates embedded by go-assets-builder
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		fmt.Println(name)
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
