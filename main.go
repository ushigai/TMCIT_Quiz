package main

import (
    "net/http"
	"html/template"
	"io"
	"os"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
    "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
	Id    int    `json:"id" param:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

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

func getUsers(c echo.Context) error {
	db := sqlConnect()
	users := []User{}
	db.Find(&users)
	defer db.Close()
	return c.JSON(http.StatusOK, users)
}

func getUser(c echo.Context) error {
	db := sqlConnect()
	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	db.Take(&user)
	defer db.Close()

	return c.JSON(http.StatusOK, user)
}

func updateUser(c echo.Context) error {
	db := sqlConnect()
	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	db.Save(&user)
	defer db.Close()
	return c.JSON(http.StatusOK, user)
}

func createUser(c echo.Context) error {
	db := sqlConnect()
	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	db.Create(&user)
	defer db.Close()
	return c.JSON(http.StatusCreated, user)
}

func deleteUser(c echo.Context) error {
	db := sqlConnect()
	id := c.Param("id")
	db.Delete(&User{}, id)
	defer db.Close()
	return c.NoContent(http.StatusNoContent)
}


func main() {
	db := sqlConnect()
	db.AutoMigrate(&User{})
	defer db.Close()

	e := echo.New()
	tl, err := template.New("t").ParseGlob("views/*.html")
    tl.ParseGlob("views/common/*.html")
    tl.ParseGlob("views/quiz/*.html")
    t := &Template{
        templates: template.Must(tl, err),
    }
	e.Renderer = t

	e.GET("/login", func(c echo.Context) error{
		data := struct { info string } { "login", }
		return c.Render(http.StatusOK, "login", data)
	})
	e.GET("/home", func(c echo.Context) error{
		data := struct { info string } { "home", }
		return c.Render(http.StatusOK, "home", data)
	})

    e.GET("/lobby/:QuizID", GetQuiz)
    e.GET("/lobby", GetRoom)
    e.GET("/room/:RoomID", GetQuiz)
    e.GET("/quiz/:RoomID/:QuizID", StartQuiz)

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
	
	e.GET("/users", getUsers)

    e.Logger.Fatal(e.Start(":8080"))
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.POST("/users", createUser)
	e.DELETE("/users/:id", deleteUser)

}


func sqlConnect() (database *gorm.DB) {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_NAME")

	DBMS := "mysql"
	PROTOCOL := "tcp(db:3306)"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	fmt.Println("user : " + USER)
	fmt.Println("password : " + PASS)
	fmt.Println("DB name : " + DBNAME)
	fmt.Println("DBMS : " + DBMS)
	fmt.Println("protocol : " + PROTOCOL)

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic("DB接続失敗🤪🤪🤪")
	} else {
		fmt.Println("やったね😊")
	}
	//db.LogMode(true)
	return db
}

