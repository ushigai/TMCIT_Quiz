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

type QuizDataStruct struct {
	RoomID string
	QuizID string
	QuizText string
	QuizAnswer string
	QuizTimeout string
	Choice1 string
	Choice2 string
	Choice3 string
	Choice4 string
	NextQuiz string
} 

type RoomDataStruct []struct {
	RoomID string
	QuizTitle string
	QuizSubTitle string
}

type QuizDiscription struct {
	RoomID string
	QuizTitle string
	QuizSubTitle string
	Author string
	Date string
	Comment string
}

func StartQuiz(c echo.Context) error {
	// TODO: ここらへんDBと連携する
	QuizData := QuizDataStruct{
		RoomID: c.Param("RoomID"),
		QuizID: c.Param("QuizID"),
		QuizText: "「家康の康ですよね？」は高専何年生のとき？",
		QuizAnswer: "高専2年",
		QuizTimeout: "6000",
		Choice1: "高専1年",
		Choice2: "高専2年",
		Choice3: "高専3年",
		Choice4: "高専4年",
		NextQuiz: "finish",
	}

	return c.Render(http.StatusOK, "quiz", QuizData)
}

func GetQuiz(c echo.Context) error {
	// TODO: ここらへんDBと連携する
	QuizDiscription := QuizDiscription{
		RoomID: c.Param("RoomID"),
		QuizTitle: "1300の歴史",
		QuizSubTitle: "じじじせいじ編",
		Author: "ushigai",
		Date: "2022/07/11",
		Comment: "Duoなんだよなぁ",
	}
	return c.Render(http.StatusOK, "room", QuizDiscription)
}

func GetRoom(c echo.Context) error {
	// TODO: ここらへんDBと連携する
	RoomData := RoomDataStruct {
		{
			RoomID: "12",
			QuizTitle: "1300の歴史",
			QuizSubTitle: "じじじせいじ編",
		},
		{
			RoomID: "23",
			QuizTitle: "1300の歴史",
			QuizSubTitle: "キッズ編",
		},
	}
	return c.Render(http.StatusOK, "lobby", RoomData)
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

    e.GET("/lobby", GetRoom)
	//e.GET("/lobby", func(c echo.Context) error{
		//// TODO: ここらへんDBと連携する
		//data := struct {
			//LobbyID string
			//LobbyName string
		//} {
			//"123456789",
			//"1300の歴史",
		//}
		//return c.Render(http.StatusOK, "lobby", data)
	//})
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
	//e.GET("/quiz", func(c echo.Context) error{
		//data := struct {
			//info string
		//} {
			//"quiz",
		//}
		//return c.Render(http.StatusOK, "quiz", data)
	//})
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
	
    e.GET("/quiz/:RoomID/:QuizID", StartQuiz)
    e.GET("/lobby/:QuizID", GetQuiz)
    e.GET("/room/:RoomID", GetQuiz)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Here is root :)")
	})

	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret1" {
			return true, nil
		}
		return false, nil
	}))

    e.Logger.Fatal(e.Start(":8080"))

}
