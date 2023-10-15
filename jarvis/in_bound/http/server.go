package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"html/template"
	"io"
	"jarvis/out_bound/grpc"
	"net/http"
)

func Run(host string, port string) {

	e := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("template/index.html")),
	}
	e.Renderer = renderer

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"Chat": "Hello!",
		})
	})

	e.POST("/", GenerateTextHandler)

	e.Logger.Fatal(e.Start(host + ":" + port))
}

func GenerateTextHandler(c echo.Context) error {
	log.Info("request: ", c)
	res := grpc.GetGeneratedText(c.FormValue("text"))
	log.Info("response: ", res)
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"Chat": res,
	})
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
