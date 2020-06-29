package main

import (
	"github.com/labstack/echo"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var routes *RoutesStruct

func main() {
	// read config
	err := parseJSONConfig()
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	e := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.html")),
	}
	e.Renderer = renderer

	e.GET("/*", index)
	e.GET("/open/:file", openFile)

	e.Logger.Fatal(e.Start(":5000"))
}

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", routes)
}

func openFile(c echo.Context) error {
	route, err := getRouteConfig(c.Param("file"))
	if err != nil {
		log.Printf("File not found: " + c.Param("file"))
		return c.String(http.StatusBadRequest, "File not found!")

	}
	cmd := exec.Command("open", route.FilePath)
	log.Printf("Opening '" + route.FilePath + "'...")
	err = cmd.Run()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Opening "+route.Name+"...")
}
