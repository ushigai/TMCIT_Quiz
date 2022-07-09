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

func getUserName(c echo.Context) error {
    id := c.Param("id")
    return c.String(http.StatusOK, id)
    return c.Render(http.StatusOK, "lobby", id)
}

func main() {
	e := echo.New()
	tl, err := template.New("t").ParseGlob("views/*.html")
    tl.ParseGlob("views/common/*.html")
    tl.ParseGlob("views/quiz/*.html")
    t := &Template{
        templates: template.Must(tl, err),
    }
	e.Renderer = t

	e.GET("/lobby", func(c echo.Context) error{
		data := struct {
			LobbyID string
			LobbyName string
		} {
			"123456789",
			"1300の歴史",
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
	e.GET("/quiz", func(c echo.Context) error{
		data := struct {
			info string
		} {
			"quiz",
		}
		return c.Render(http.StatusOK, "quiz", data)
	})
	e.GET("/collect", func(c echo.Context) error{
		data := struct {
			info string
		} {
			"collect",
		}
		return c.Render(http.StatusOK, "collect", data)
	})
	e.GET("/wrong", func(c echo.Context) error{
		data := struct {
			info string
		} {
			"wrong",
		}
		return c.Render(http.StatusOK, "wrong", data)
	})
	e.GET("/timeout", func(c echo.Context) error{
		data := struct {
			info string
		} {
			"timeout",
		}
		return c.Render(http.StatusOK, "timeout", data)
	})
    e.GET("/lobby/:id", getUserName)
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

    e.Logger.Fatal(e.Start(":8080"))

}
