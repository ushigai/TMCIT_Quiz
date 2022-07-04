package main

import (
    "net/http"
	"html/template"
	"io"
    "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()

	e.Renderer = t

	e.GET("/lobby", func(c echo.Context) error{
		data := struct {
			info string
		} {
			"lobby",
		}
		return c.Render(http.StatusOK, "lobby", data)
	})

	e.GET("/login", func(c echo.Context) error{
		data := struct {
			info string
		} {
			"login",
		}
		return c.Render(http.StatusOK, "login", data)
	})
	e.GET("/home", func(c echo.Context) error{
		data := struct {
			info string
		} {
			"home",
		}
		return c.Render(http.StatusOK, "home", data)
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Here is root :)")
	})

	//e.GET("/users/:id", func(c echo.Context) error {
		//return c.Render(http.StatusOK, "/users/:id")
	//})
	//e.GET("/rooms/:RoomNumber", func(c echo.Context) error {
		//return c.Render(http.StatusOK, "/rooms/:RoomNumber")
	//})
	
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret1" {
			//data := struct {
				//info string
			//} {
				//"admin",
			//}
			return true, nil
		}
		return false, nil
	}))

    e.Logger.Fatal(e.Start(":1323"))
}
